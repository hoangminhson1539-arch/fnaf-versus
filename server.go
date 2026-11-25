package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// Portable, dependency-free, local network game relay.
type M = map[string]interface{}
type Client struct {
	token, role    string
	room           *Room
	queue          []M
	seen           time.Time
	closed         bool
	foreground     bool
	lastCommand    time.Time
	commandCredit  float64
	notify         chan struct{}
	lastState      time.Time
	lastChat       time.Time
	pending        []M
	batch, lastSeq int
}
type Room struct {
	code, diff, hostJob string
	host, guest         *Client
	started, finished   bool
	rematch             map[string]bool
	latest              M
	won                 bool
}

var mu sync.Mutex
var clients = map[string]*Client{}
var rooms = map[string]*Room{}
var lan = "127.0.0.1"
var lans []string
var game []byte
var publicServer = os.Getenv("FNAF_PUBLIC") == "1"

func random(n int) string {
	b := make([]byte, n)
	if _, e := rand.Read(b); e != nil {
		panic(e)
	}
	return hex.EncodeToString(b)
}
func code() string {
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	b := make([]byte, 6)
	rand.Read(b)
	for i := range b {
		b[i] = chars[int(b[i])%len(chars)]
	}
	return string(b)
}
func str(m M, k string) string { s, _ := m[k].(string); return s }
func send(c *Client, m M) {
	if c == nil || c.closed {
		return
	}
	if contains(str(m, "type"), "command", "signal") {
		m["relayQueuedAt"] = time.Now().UnixMilli()
	}
	if str(m, "type") == "state" {
		for i := len(c.queue) - 1; i >= 0; i-- {
			if str(c.queue[i], "type") == "state" {
				c.queue = append(c.queue[:i], c.queue[i+1:]...)
				break
			}
		}
	}
	if len(c.queue) >= 64 {
		c.queue = c.queue[1:]
	}
	c.queue = append(c.queue, m)
	select {
	case c.notify <- struct{}{}:
	default:
	}
}
func other(c *Client) *Client {
	if c.room == nil {
		return nil
	}
	if c.role == "host" {
		return c.room.guest
	}
	return c.room.host
}
func drop(c *Client) {
	if c == nil || c.closed {
		return
	}
	c.closed = true
	delete(clients, c.token)
	if r := c.room; r != nil {
		delete(rooms, r.code)
		p := other(c)
		if p != nil && !p.closed {
			p.room = nil
			send(p, M{"type": "closed"})
		}
		c.room = nil
	}
}
func validNum(m M, k string, lo, hi float64) bool {
	v, ok := m[k].(float64)
	return ok && v >= lo && v <= hi
}
func validFlags(v interface{}) bool {
	m, ok := v.(map[string]interface{})
	if !ok {
		return false
	}
	_, a := m["left"].(bool)
	_, b := m["right"].(bool)
	return a && b
}
func contains(s string, values ...string) bool {
	for _, v := range values {
		if s == v {
			return true
		}
	}
	return false
}
func validState(m M) bool {
	if !contains(str(m, "phase"), "playing", "blackout", "jumpscare", "deathAftermath", "winAnim", "won", "lost") {
		return false
	}
	for k, v := range map[string]float64{"secs": 360, "duration": 360, "power": 100, "foxy": 3, "jumpTimer": 100, "deathTimer": 100, "blackoutClock": 200} {
		if !validNum(m, k, 0, v) {
			return false
		}
	}
	if d, ok := m["duel"].(map[string]interface{}); ok {
		if d["order"] != nil && !contains(str(d, "order"), "Bonnie", "Chica", "Freddy", "Foxy") {
			return false
		}
		if !contains(str(d, "tactic"), "patrol", "decoy", "assault", "hold") {
			return false
		}
		for _, k := range []string{"readyAt", "skillAt", "skillUntil", "guardAt", "guardUntil"} {
			if !validNum(d, k, 0, 1000) {
				return false
			}
		}
		if _, ok := d["returning"].(bool); !ok {
			return false
		}
	}
	if !validFlags(m["doors"]) || !validFlags(m["lights"]) {
		return false
	}
	eff, ok := m["effects"].(map[string]interface{})
	if !ok {
		return false
	}
	for _, k := range []string{"cameraJamUntil", "ghostUntil", "flickerUntil"} {
		if !validNum(eff, k, 0, 1000) {
			return false
		}
	}
	hs, ok := m["hunters"].([]interface{})
	if !ok || len(hs) != 4 {
		return false
	}
	seen := map[string]bool{}
	for _, h := range hs {
		v, ok := h.(map[string]interface{})
		if !ok {
			return false
		}
		n := str(v, "name")
		if seen[n] || !contains(n, "Freddy", "Bonnie", "Chica", "Foxy") || !validNum(v, "x", 0, 2050) || !validNum(v, "y", 0, 1090) {
			return false
		}
		seen[n] = true
	}
	return true
}
func handle(c *Client, m M) {
	t := str(m, "type")
	if c.room == nil {
		switch t {
		case "create":
			if len(rooms) >= 64 {
				send(c, M{"type": "error", "message": "Máy chủ đầy."})
				return
			}
			k := code()
			for rooms[k] != nil {
				k = code()
			}
			d := str(m, "difficulty")
			if !contains(d, "easy", "normal", "hard", "hell") {
				d = "easy"
			}
			r := &Room{code: k, diff: d, hostJob: "guard", host: c, rematch: map[string]bool{}}
			rooms[k] = r
			c.room = r
			c.role = "host"
			send(c, M{"type": "room", "role": "host", "code": k, "ready": false, "lan": lan, "lans": lans})
		case "join":
			r := rooms[str(m, "code")]
			if r == nil || r.guest != nil || r.started {
				send(c, M{"type": "error", "message": "Mã phòng sai, phòng đầy hoặc đã bắt đầu."})
				return
			}
			r.guest = c
			c.room = r
			c.role = "guest"
			send(c, M{"type": "room", "role": "guest", "code": r.code, "ready": true, "hostJob": r.hostJob})
			send(r.host, M{"type": "peer", "ready": true})
		}
		return
	}
	r := c.room
	p := other(c)
	if t == "roles" && c.role == "host" && !r.started {
		job := str(m, "hostJob")
		if contains(job, "guard", "camera", "duelguard", "monster", "foxguard", "foxmonster") {
			r.hostJob = job
			send(p, M{"type": "roles", "hostJob": job})
		}
		return
	}
	if p == nil {
		return
	}
	if t == "start" && c.role == "host" && !r.started {
		if time.Since(p.seen) > 4*time.Second || !p.foreground {
			send(c, M{"type": "notice", "message": "Bạn chơi chưa sẵn sàng. Hãy mở lại game."})
			return
		}
		r.started = true
		send(c, M{"type": "start", "difficulty": r.diff, "hostJob": r.hostJob})
		send(p, M{"type": "start", "difficulty": r.diff, "hostJob": r.hostJob})
		return
	}
	if t == "rematch" && r.finished {
		r.rematch[c.role] = true
		n := len(r.rematch)
		send(c, M{"type": "rematch", "count": n})
		send(p, M{"type": "rematch", "count": n})
		if n == 2 {
			r.finished = false
			r.latest = nil
			r.won = false
			r.rematch = map[string]bool{}
			send(c, M{"type": "start", "difficulty": r.diff, "hostJob": r.hostJob})
			send(p, M{"type": "start", "difficulty": r.diff, "hostJob": r.hostJob})
		}
		return
	}
	if t == "chat" {
		text := strings.TrimSpace(str(m, "text"))
		if text != "" && utf8.RuneCountInString(text) <= 240 && time.Since(c.lastChat) >= 300*time.Millisecond {
			c.lastChat = time.Now()
			send(p, M{"type": "chat", "text": text})
		}
		return
	}
	if t == "voice" {
		kind := str(m, "kind")
		sdp := str(m, "sdp")
		if contains(kind, "ready", "bye", "retry") {
			send(p, M{"type": "voice", "kind": kind})
		} else if contains(kind, "offer", "answer") && len(sdp) > 0 && len(sdp) <= 20000 && ((kind == "offer" && c.role == "host") || (kind == "answer" && c.role == "guest")) {
			send(p, M{"type": "voice", "kind": kind, "sdp": sdp})
		}
		return
	}
	if !r.started || r.finished {
		return
	}
	switch t {
	case "signal":
		if !c.foreground || !p.foreground || time.Since(p.seen) > 4*time.Second {
			return
		}
		if contains(str(m, "signal"), "left", "right", "ok") {
			send(p, M{"type": "signal", "signal": str(m, "signal")})
		}
	case "command":
		if c.role != "guest" || !c.foreground || !p.foreground || time.Since(p.seen) > 4*time.Second {
			return
		}
		now := time.Now()
		if c.lastCommand.IsZero() {
			c.commandCredit = 12
		} else {
			c.commandCredit += now.Sub(c.lastCommand).Seconds() * 20
			if c.commandCredit > 12 {
				c.commandCredit = 12
			}
		}
		c.lastCommand = now
		if c.commandCredit < 1 {
			return
		}
		c.commandCredit--
		a := str(m, "action")
		if contains(r.hostJob, "duelguard", "foxguard") {
			if a == "tactic" && r.hostJob == "duelguard" {
				send(p, M{"type": "command", "action": a})
			} else if contains(a, "select", "attack", "recall", "skill", "skill2", "moveDining", "moveWest", "moveEast", "moveStage", "moveCove", "moveBack", "moveKitchen", "moveRest", "moveSupply", "manualN", "manualNE", "manualE", "manualSE", "manualS", "manualSW", "manualW", "manualNW", "manualStop") && contains(str(m, "bot"), "Bonnie", "Chica", "Freddy", "Foxy") && (r.hostJob != "foxguard" || str(m, "bot") == "Foxy") {
				send(p, M{"type": "command", "action": a, "bot": str(m, "bot")})
			}
			return
		}
		if contains(r.hostJob, "monster", "foxmonster") && contains(a, "leftDoor", "rightDoor", "leftLight", "rightLight", "fuseRepair", "guardSkill", "guardReserve", "guardLure", "guardScan", "guardRepairLeft", "guardRepairRight") {
			send(p, M{"type": "command", "action": a})
			return
		}
		if r.hostJob == "camera" {
			if contains(a, "leftDoor", "rightDoor", "leftLight", "rightLight", "fuseRepair") {
				send(p, M{"type": "command", "action": a})
			}
			return
		}
		if a == "repair" {
			send(p, M{"type": "command", "action": "repair"})
		} else if a == "watch" && contains(str(m, "cam"), "1A", "1B", "1C", "2A", "2B", "3", "4A", "4B", "5", "6", "7") {
			if enabled, ok := m["enabled"].(bool); ok {
				send(p, M{"type": "command", "action": "watch", "cam": str(m, "cam"), "enabled": enabled})
			}
		}
	case "state":
		if c.role == "host" && time.Since(c.lastState) > 70*time.Millisecond && validState(m) {
			c.lastState = time.Now()
			r.latest = m
			send(p, m)
		}
	case "result":
		if c.role == "host" {
			if won, ok := m["won"].(bool); ok {
				r.finished = true
				r.won = won
				r.rematch = map[string]bool{}
				send(p, M{"type": "result", "won": won})
			}
		}
	}
}
func write(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	json.NewEncoder(w).Encode(v)
}
func api(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "POST required", 405)
		return
	}
	if origin := r.Header.Get("Origin"); origin != "" && origin != "http://"+r.Host && origin != "https://"+r.Host {
		http.Error(w, "Invalid origin", 403)
		return
	}
	var m M
	if e := json.NewDecoder(http.MaxBytesReader(w, r.Body, 65536)).Decode(&m); e != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	action := strings.TrimPrefix(r.URL.Path, "/api/")
	if action == "open" {
		if v, ok := m["version"].(float64); !ok || v != 42 {
			http.Error(w, "Use FNAF V41.2", 409)
			return
		}
		if len(clients) >= 128 {
			http.Error(w, "Server full", 503)
			return
		}
		id := random(24)
		clients[id] = &Client{token: id, seen: time.Now(), foreground: true, notify: make(chan struct{}, 1)}
		write(w, M{"token": id})
		return
	}
	c := clients[str(m, "token")]
	if c == nil {
		http.Error(w, "Session closed", 410)
		return
	}
	c.seen = time.Now()
	switch action {
	case "presence":
		if v, ok := m["foreground"].(bool); ok {
			if c.foreground != v {
				c.foreground = v
				if p := other(c); p != nil {
					select {
					case p.notify <- struct{}{}:
					default:
					}
				}
			}
		}
		write(w, M{"ok": true})
	case "send":
		seq, _ := m["seq"].(float64)
		if seq == 0 || int(seq) > c.lastSeq {
			if data, ok := m["data"].(map[string]interface{}); ok {
				handle(c, data)
			}
			if seq > 0 {
				c.lastSeq = int(seq)
			}
		}
		write(w, M{"ok": true})
	case "poll":
		if v, ok := m["foreground"].(bool); ok {
			if c.foreground != v {
				c.foreground = v
				if p := other(c); p != nil {
					select {
					case p.notify <- struct{}{}:
					default:
					}
				}
			}
		}
		ack, _ := m["ack"].(float64)
		if int(ack) == c.batch {
			c.pending = nil
		}
		waited := int64(0)
		if c.pending == nil && len(c.queue) == 0 {
			wait, _ := m["wait"].(float64)
			if wait > 1000 {
				wait = 1000
			}
			if wait > 0 {
				select {
				case <-c.notify:
				default:
				}
				began := time.Now()
				mu.Unlock()
				select {
				case <-c.notify:
				case <-time.After(time.Duration(wait) * time.Millisecond):
				case <-r.Context().Done():
				}
				mu.Lock()
				waited = time.Since(began).Milliseconds()
				if c.closed {
					http.Error(w, "Session closed", 410)
					return
				}
			}
		}
		if c.pending == nil && len(c.queue) > 0 {
			c.batch++
			c.pending = c.queue
			c.queue = nil
		}
		peer := other(c)
		now := time.Now()
		healthy := peer != nil && !peer.closed && now.Sub(peer.seen) < 4*time.Second
		q := []M{}
		for _, msg := range c.pending {
			if contains(str(msg, "type"), "command", "signal") {
				at, _ := msg["relayQueuedAt"].(int64)
				if !c.foreground || !healthy || !peer.foreground || now.UnixMilli()-at > 2500 {
					continue
				}
			}
			q = append(q, msg)
		}
		response := M{"messages": q, "batch": c.batch, "waited": waited, "peerOnline": healthy, "peerForeground": peer == nil || peer.foreground}
		if want, _ := m["resync"].(bool); want && c.room != nil {
			if c.role == "guest" {
				response["snapshot"] = c.room.latest
			}
			response["finished"] = c.room.finished
			response["won"] = c.room.won
		}
		write(w, response)
	case "close":
		drop(c)
		write(w, M{"ok": true})
	default:
		http.NotFound(w, r)
	}
}

// Prefer physical RFC1918 interfaces over CGNAT/VPN/virtual adapters.
type candidate struct {
	ip, name string
	score    int
}

func ipScore(ip net.IP, name string, route net.IP) int {
	v := ip.To4()
	if v == nil || ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsUnspecified() || ip.IsMulticast() {
		return -10000
	}
	n := strings.ToLower(name)
	score := 0
	if v[0] == 192 && v[1] == 168 {
		score = 120
	} else if v[0] == 172 && v[1] >= 16 && v[1] <= 31 {
		score = 100
	} else if v[0] == 10 {
		score = 100
	} else if v[0] == 100 && v[1] >= 64 && v[1] <= 127 {
		score = -100
	} else {
		score = -60
	}
	if strings.Contains(n, "ethernet") || strings.Contains(n, "wi-fi") || strings.Contains(n, "wifi") || strings.HasPrefix(n, "eth") || strings.HasPrefix(n, "en") || strings.HasPrefix(n, "wl") {
		score += 25
	}
	for _, x := range []string{"vpn", "virtual", "vethernet", "vmware", "virtualbox", "tailscale", "wireguard", "wintun", "tun", "tap", "hamachi", "zerotier", "docker", "wsl", "bluetooth", "clash"} {
		if strings.Contains(n, x) {
			score -= 180
			break
		}
	}
	if route != nil && ip.Equal(route) && score >= 80 {
		score += 15
	}
	return score
}
func lanIPs() []string {
	var route net.IP
	c, e := net.Dial("udp4", "192.0.2.1:9")
	if e == nil {
		if a, ok := c.LocalAddr().(*net.UDPAddr); ok {
			route = a.IP
		}
		c.Close()
	}
	var cs []candidate
	ifs, _ := net.Interfaces()
	for _, i := range ifs {
		if i.Flags&net.FlagUp == 0 || i.Flags&net.FlagLoopback != 0 {
			continue
		}
		as, _ := i.Addrs()
		for _, a := range as {
			ip, _, _ := net.ParseCIDR(a.String())
			score := ipScore(ip, i.Name, route)
			if score > -10000 {
				cs = append(cs, candidate{ip.String(), i.Name, score})
			}
		}
	}
	sort.SliceStable(cs, func(i, j int) bool { return cs[i].score > cs[j].score })
	out := []string{}
	seen := map[string]bool{}
	for _, c := range cs {
		if !seen[c.ip] {
			out = append(out, c.ip)
			seen[c.ip] = true
		}
	}
	if len(out) == 0 {
		out = append(out, "127.0.0.1")
	}
	return out
}
func openBrowser(url string) {
	var c *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		c = exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", url)
	case "darwin":
		c = exec.Command("open", url)
	default:
		c = exec.Command("xdg-open", url)
	}
	if e := c.Start(); e != nil {
		fmt.Println("Open this address in your browser:", url)
	}
}
func main() {
	exe, _ := os.Executable()
	root := filepath.Dir(exe)
	var e error
	for _, path := range []string{
		filepath.Join(root, "index.html"),
		filepath.Join(root, "FNAF1_VERSUS_V40.html"),
		"index.html",
		"FNAF1_VERSUS_V40.html",
	} {
		game, e = os.ReadFile(path)
		if !os.IsNotExist(e) {
			break
		}
	}
	if e != nil {
		fmt.Println("Keep the EXE and HTML together. Extract the whole ZIP first.")
		fmt.Println("Press Enter to close.")
		fmt.Scanln()
		return
	}
	lans = lanIPs()
	lan = lans[0]
	if publicServer {
		lan = ""
		lans = nil
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/config", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "GET required", 405)
			return
		}
		servers := []M{{"urls": "stun:stun.l.google.com:19302"}, {"urls": "stun:stun.cloudflare.com:3478"}}
		if raw := os.Getenv("FNAF_ICE_SERVERS"); raw != "" {
			var custom []M
			if json.Unmarshal([]byte(raw), &custom) == nil && len(custom) > 0 && len(custom) <= 12 {
				servers = custom
			}
		}
		write(w, M{"iceServers": servers})
	})
	mux.HandleFunc("/api/", api)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		meta := "<meta name=\"fnaf-portable-server\" content=\"1\">"
		if publicServer {
			meta += "<meta name=\"fnaf-public-server\" content=\"1\">"
		}
		io.WriteString(w, strings.Replace(string(game), "<!--PORTABLE_SERVER-->", meta, 1))
	})
	var l net.Listener
	if publicServer {
		p := os.Getenv("PORT")
		if p == "" {
			p = "8080"
		}
		l, e = net.Listen("tcp4", "0.0.0.0:"+p)
	} else {
		for port := 8765; port < 8776; port++ {
			l, e = net.Listen("tcp4", fmt.Sprintf("0.0.0.0:%d", port))
			if e == nil {
				break
			}
		}
	}
	if e != nil {
		log.Fatal(e)
	}
	port := l.Addr().(*net.TCPAddr).Port
	local := fmt.Sprintf("http://127.0.0.1:%d/?host=1", port)
	fmt.Println("FNAF V41.2 - SERVER READY\nKeep this window open. Allow private-network access if prompted.\nHost:", local, "\nSame Wi-Fi:", fmt.Sprintf("http://%s:%d/", lan, port), "\nOther network addresses:", strings.Join(lans, ", "), "\nScan QR or share the invite link.\nCtrl+C to stop.")
	if !publicServer && os.Getenv("FNAF_NO_BROWSER") != "1" {
		openBrowser(local)
	}
	go func() {
		for range time.Tick(10 * time.Second) {
			mu.Lock()
			for _, c := range clients {
				if time.Since(c.seen) > 900*time.Second {
					drop(c)
				}
			}
			mu.Unlock()
		}
	}()
	srv := &http.Server{Handler: mux, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 30 * time.Second, MaxHeaderBytes: 16384}
	log.Fatal(srv.Serve(l))
}
