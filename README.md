# FNAF 1 Versus · V49.0 Reconnect & Squad Intelligence

**[Chơi ngay trên trình duyệt](https://hoangminhson1539-arch.github.io/fnaf-versus/)**

**[Nhật ký cập nhật chi tiết — CHANGELOG.md](CHANGELOG.md)**

## Nhật ký cập nhật trên README

> Quy trình bắt buộc: mỗi lần cập nhật game phải đồng bộ phiên bản, mục changelog mới nhất trong README.md và lịch sử đầy đủ trong CHANGELOG.md; nêu tính năng, sửa lỗi, cân bằng, kiểm thử và giới hạn thực tế.

### V49.0 — Reconnect & Squad Intelligence (26/11/2026)

- Tự nối lại kênh dữ liệu P2P của khách trong cùng tab tối đa 30 giây; giữ nguyên vị trí, vai trò và mã ca, tạm dừng cả đội khi đang chơi. Xác thực yêu cầu nối lại bằng Peer ID và token ngẫu nhiên 128 bit riêng của chỗ; không lưu token vào bộ nhớ lâu dài. Nếu hết hạn, kết thúc ca và thông báo cho các thành viên còn lại trước khi đóng kết nối. Không hỗ trợ khôi phục sau khi đóng/tải lại tab hoặc mất chủ phòng.
- Xác nhận sẵn sàng theo phiên bản cấu hình lobby, loại bỏ xác nhận cũ sau khi đổi thời lượng/độ khó/đội hình. Sửa lỗi thông báo kết thúc ca bị mất do đóng kênh quá sớm; lưu kết quả ca hiện tại để gửi lại cho khách vừa khôi phục.
- Co-op, Thực địa và Foxy Rush P2P 3–4 người: quản lý trạng thái và lựa chọn camera theo từng khách, giữ camera chung hoạt động nếu còn người xem, xử lý kỹ năng phụ thuộc camera theo lựa chọn của người gửi. Chia sẻ điện, cửa và hồi chiêu theo luật hiện tại.
- Giữ đường đi và tiến độ thao tác thực địa khi tạm dừng mạng P2P; không tiến hành nhiệm vụ trong thời gian tạm dừng. Giữ hệ thống điện, cân bằng Foxy Rush và trợ giúp đêm 1 đã có.
- HUD đội hình nhỏ hiển thị P1–P4 và chỗ đang khôi phục; cập nhật thông báo trong lobby, cài đặt, và trận đấu cho nhóm nhiều người. Sửa cập nhật trạng thái phát voice từ kết nối cũ và thử phát lại audio bị tạm dừng.
- Kiểm tra cú pháp toàn bộ JavaScript đạt. 27 kiểm tra tích hợp giả lập 4 người đạt cho phiên bản trước tinh chỉnh HUD/kết quả/voice: xác nhận cấu hình, camera độc lập, ngắt kênh, tạm dừng, nối lại đúng vị trí, đồng bộ sau nối lại, hết hạn và xác thực token. Cú pháp bản cuối đạt. Chưa thử WebRTC, mic, đồ họa và gameplay trên thiết bị thật. P2P yêu cầu cả nhóm V49.0; relay vẫn hai người, không áp dụng cơ chế giữ chỗ mới.

### V48.0 — Party & Tactical Overhaul (26/11/2026)

- Phòng P2P 2–4 người có xác nhận sẵn sàng từng thành viên; chủ phòng chỉ bắt đầu khi mọi khách đã xác nhận. Đổi chế độ, độ khó, thời lượng, sức chứa hoặc đội hình sẽ yêu cầu xác nhận lại. Hiển thị vai trò dự kiến của phòng 3–4 người và thêm preset Chiến thuật 9 phút.
- Tách gói dữ liệu theo mã ca và số thứ tự, loại bỏ gói trận cũ, trùng hoặc sai cấu trúc; kiểm tra trạng thái điện, thời gian, cửa, hiệu ứng và vị trí quái. Đồng bộ chơi lại và nhận diện P1 trong chat. Kiểm tra kết nối còn mở khi đánh giá sức khỏe phòng.
- Voice nhóm có nút tắt nghe từng thành viên, giữ mute khi nối lại voice, xóa mute khi thành viên rời phòng; kiểm tra mức mic qua WebRTC stats nếu trình duyệt cung cấp. Khôi phục Audio Session khi quay lại tab; bổ sung nút khôi phục nhạc/SFX và đổi đầu ra ngay trong cài đặt phòng.
- Thực địa: thu gọn/ẩn bản đồ cho cả bảo vệ và camera, giữ đường đi cũ nếu yêu cầu đường mới thất bại, tránh hủy đường do chạm lặp, thử tìm đường lại tối đa một lần nếu kẹt. Điều chỉnh phí điện quét/sửa/dụ quái theo thời lượng và đồng bộ điều kiện nút.
- Tối ưu HUD ngang, đồng hồ phút:giây và lượng điện tiếp tế mốc kế tiếp, bảng voice nhỏ gọn, lịch sử 12 trận có tổng thắng/thua và tỷ lệ thắng cục bộ. Giữ tutorial đêm 1 và cân bằng Foxy Rush.
- Kiểm tra cú pháp toàn bộ JavaScript đạt; 91 kiểm tra hàm readiness/validation và 13 kiểm tra vòng đời P2P giả lập 4 người đạt. Chưa kiểm thử giao diện, mic/loa thật, kết nối qua NAT hoặc cân bằng thực chiến; không khẳng định đã hết mọi lỗi. Giao thức P2P V48.0 mới yêu cầu cùng bản; relay vẫn 2 người và không áp dụng ready-check mới.

### V47.5 — Night 1 Tutorial & Landscape UI (26/11/2026)

- Chỉ hiển thị hỗ trợ tự động tên quái, phía cửa, đếm ngược tấn công và nhãn nhận diện CCTV trong Đêm 1 cốt truyện (Baby/Easy/Normal). Từ Đêm 2, Custom, Challenge, Boss và multiplayer không có trợ giúp tự động này; giữ cảm biến/radar được người chơi chủ động kích hoạt, âm thanh và hình ảnh thật của quái.
- Bỏ hộp thoại chọn hỗ trợ lúc 2 AM gây gián đoạn trận; chuẩn hóa lựa chọn hỗ trợ từ dữ liệu lưu cũ.
- Thu gọn thông báo, nhãn camera, bảng tên đêm 1 và hộp thoại; thông báo thông thường tự ẩn trong tối đa 2,5 giây, không chặn thao tác. Nội dung cuộc gọi đầy đủ vẫn có trong nhật ký.
- Bố cục landscape ưu tiên sân chơi, đặt HUD multiplayer bên phải có cuộn riêng, thanh điện/thời gian ở trên; tối ưu các nút cửa, camera, menu, chat và khu vực an toàn trên màn hình thấp. Cập nhật canvas khi xoay màn hình/đổi kích thước viewport; không bắt buộc khóa hướng màn hình.
- Kiểm tra cú pháp JavaScript và 450 tổ hợp logic tutorial theo đêm, chế độ, độ khó, online/offline đạt. Chưa kiểm thử trực quan trên điện thoại/PC thật. Giữ cân bằng điện V47.4, voice nhóm và giao thức P2P V47.4.

### V47.4 — Mobile Voice Audio & Power Balance (26/11/2026)

- Khắc phục xung đột giữa micro và âm thanh game: yêu cầu chế độ Audio Session play-and-record khi thu mic và playback khi ngừng thu trên trình duyệt hỗ trợ; kiểm tra lại đầu ra sau khi cấp quyền, đổi thiết bị hoặc ngắt mic. Không tự giảm nhạc 35% khi voice kết nối.
- Dùng bus âm thanh game riêng, giữ voice độc lập. Thêm nút ĐỔI ĐẦU RA GAME trong chat voice và cài đặt: có thể chuyển Web Audio trực tiếp sang MediaStream audio thay thế; nếu bị chặn, lỗi hoặc dừng phát, tự quay lại đầu ra trực tiếp. Nút KHÔI PHỤC NHẠC/SFX vẫn dùng được trong phòng.
- Cân bằng điện cho multiplayer 3/6/9/12/15 phút, giữ 100% lúc bắt đầu. Hệ số tiêu hao thời gian 180/d; giảm nhẹ hao điện nền theo thời lượng; tiếp tế 5 mốc tăng dần, tổng tiềm năng Foxy Rush 22,5/26,25/30/33,75/37,5%, Versus và Co-op 20/23,75/27,5/31,25/35%, Thực địa 25/28,75/32,5/36,25/40%. Chỉ nạp đến 100% và không cấp khi đã hết điện.
- Giữ áp lực đóng hai cửa liên tục bằng mức hao phụ cửa; giảm chi phí tối thiểu của đuổi quái, gia cố, sửa nhanh, radar hỗ trợ trong trận dài. Hell multiplayer giảm hệ số hao điện 2,2 xuống 1,75; không thay đổi chơi đơn. Giữ hệ số Foxy Rush V47.2.
- Kiểm tra cú pháp JavaScript đạt; 16 kiểm tra giả lập định tuyến Audio Session/đầu ra, lỗi autoplay, timeout và mute đạt; 30 kiểm tra ngân sách điện danh nghĩa (4 chế độ × 5 thời lượng và áp lực hai cửa ở 2 chế độ đối đầu) đạt. Chưa kiểm tra âm thanh thật trên iPhone/Android/PC hay cân bằng qua trận thật. Trình duyệt có thể cần chạm nút khôi phục; không bảo đảm ép loa ngoài hoặc hoạt động trong mọi trình duyệt. P2P dùng mã phòng V47.4 mới, relay vẫn hai người.

### V47.3 — Audio Recovery (26/11/2026)

- Khởi tạo Web Audio từ thao tác chạm/bấm/phím của người chơi; phục hồi AudioContext bị suspend/interrupted, đóng hoặc đồng hồ âm thanh bị treo. Không khởi động lại nhạc liên tục khi bấm nhiều lần; loại bỏ kết quả resume thuộc context đã thay thế.
- Khôi phục đồ thị nhạc nền, hiệu ứng, Foley và âm thanh môi trường; xếp hàng ngắn hiệu ứng khi audio đang mở khóa, tách lỗi bộ lập lịch và phục hồi khi trở lại tab.
- Nút KHÔI PHỤC ÂM THANH xuất hiện khi âm thanh bị chặn/treo và luôn có trong cài đặt audio. Khi chủ động bấm nút, bật master sound, khôi phục mức nhạc/hiệu ứng mặc định nếu đã về 0 và bật chế độ nhạc thích ứng nếu đã tắt. Không tự ghi đè lựa chọn tắt âm thông thường.
- Sửa xung đột phím U với thao tác mở khóa; đồng bộ thanh âm lượng, giữ tùy chọn audio khi vào/rời multiplayer, áp dụng âm lượng voice cho tất cả luồng nhóm.
- Kiểm tra cú pháp JavaScript đạt; 24 kiểm tra giả lập AudioContext đạt trước tinh chỉnh giao diện/phím tắt, cú pháp cuối đạt. Chưa thử đầu ra âm thanh thật trên Safari/iPhone, Android hoặc PC. Trình duyệt có thể yêu cầu thao tác của người dùng sau khi chuyển nền. Không thay đổi cân bằng gameplay hoặc giao thức P2P V47.2.

### V47.2 — Foxy Rush Balance & Stability (26/11/2026)

- Foxy Rush tự cân bằng theo số người trong phòng (2/3/4): hệ số tốc độ 0,94/1,00/1,06; sát thương cửa cơ bản 14/15,4/16,8 điểm (trước đây 16). Giữ tăng áp lực theo giờ và thời lượng.
- Siêu tốc x1,75/x1,875/x2 trong 4/4,5/5 giây, hồi cơ bản 10/9,5/9 giây; kỹ năng phá rối hồi 15/14/13 giây, nhiễu camera 4 giây, rút 2/2,25/2,5% điện cơ bản gần cửa. Hao điện mỗi cú đập 0,65/0,75/0,85% cơ bản. Chi phí điện và hồi chiêu vẫn chịu hệ số thời lượng/giai đoạn.
- Bảo vệ 2 người được tăng nhẹ thời gian phản ứng và tốc độ sửa cửa; hiệu chỉnh theo số người để Foxy không quá yếu trước nhiều người dùng chung bàn trực. Không đổi cân bằng Versus thường hoặc chơi đơn.
- Sửa điều kiện bắt đầu khi kết nối đang mở hoặc người chơi ở nền; đồng bộ tạm dừng nhanh hơn, ngăn lệnh từ kết nối cũ và đặt lại biểu quyết chơi lại nếu nhóm chưa sẵn sàng.
- Hiển thị đúng quyền chọn quái và chỉ huy chiến thuật; sửa khôi phục voice khi một đầu mất kết nối và lỗi hủy cấp quyền micro.
- Kiểm tra cú pháp JavaScript đạt, 60 kiểm tra hệ số cân bằng/không ảnh hưởng Versus và 35 kiểm tra voice giả lập đạt. Chưa thử trận thật, âm thanh thật trên bốn máy hoặc xác nhận cân bằng bằng tỷ lệ thắng. P2P cần cùng V47.2; relay vẫn 2 người.

### V47.1 — Party Voice (26/11/2026)

- Voice chat WebRTC mesh cho 2–4 người qua P2P, tối đa ba kết nối âm thanh mỗi máy; sáu cặp trong phòng bốn người.
- Thêm định tuyến tín hiệu voice theo từng người, trạng thái kết nối từng thành viên, chỉnh âm lượng, mute, tự thử nối lại và dọn phiên âm thanh khi rời lobby.
- Đồng bộ tạm dừng toàn nhóm khi mất tín hiệu/ở nền; bỏ giới hạn lệnh bảo vệ trái/phải không nhất quán trong Foxy Rush.
- Kiểm tra cú pháp JavaScript và 29 kiểm tra mô phỏng đạt; chưa thử voice trên bốn thiết bị thật. Relay Go vẫn giới hạn hai người. Micro cần HTTPS/localhost và quyền truy cập; một số mạng cần TURN.
- P2P dùng mã phòng V47.1 riêng; tất cả người chơi cần tải lại cùng phiên bản.

### V47.0 — Party (26/11/2026)

- Phòng P2P 2–4 người, chọn sức chứa, lobby có vị trí người chơi, chia quyền Animatronics trong Versus, chat nhóm và chơi lại khi cả đội đồng ý.
- Co-op/Thực địa hỗ trợ nhiều camera chung; Foxy Rush hỗ trợ nhiều người chia sẻ bàn trực. Chưa có nhiều nhân vật bảo vệ độc lập.
- Thêm preset Khẩn cấp 3 phút, Tiêu chuẩn 6 phút, Cầm cự 12 phút, Ác mộng 15 phút.

### Lịch sử các bản trước

V41.9: Sửa lệnh cửa online khi camera đang mở, xác nhận kết quả từ host, thêm nút cửa nhanh trong camera và xử lý quái đứng chồng ngưỡng phía ngoài. Giữ cơ chế cửa hỏng, sửa cửa và giao thức 43. Cả hai thiết bị nên tải lại bản mới.

V41.8: Chủ phòng chọn 3/6/9/12/15 phút cho Versus và Foxy Rush (mặc định 6 phút). Mỗi giờ trong game tăng tốc độ, sát thương cửa và giảm hồi kỹ năng của quái; tiêu hao điện được điều chỉnh theo thời lượng, nhiệm vụ chia đều 3 giai đoạn. Co-op vẫn 3 phút. Giao thức 43: cả hai người và relay riêng phải cập nhật cùng phiên bản.

V41.7: Versus và Foxy Rush có 3 giai đoạn 60 giây với nhiệm vụ phụ, tiến độ và thưởng tự động cho hai phe. Sửa relay Go từ chối trạng thái AI tổng công kích; mồi nhử tác động cả quái điều khiển thủ công gần cửa. Giữ giao thức 42; nên dùng hai client cùng bản V41.7 và cập nhật relay nếu sử dụng máy chủ riêng.

V41.6: Hành lang hai cửa tối sâu khi tắt đèn. Chuẩn/Khó/Hell không thấy mắt quái khi tắt đèn. Đập cửa có sát thương vừa phải, hồi chung 5s/cửa và hao điện nhẹ; sửa lỗi điều khiển, AI và khôi phục kết nối. Giữ radar tên phòng/CAM, bản đồ di động và sửa cửa chậm. Dùng hai client cùng phiên bản; relay giao thức 42.

Game fan-made chạy trên trình duyệt. Có chế độ một người, Co-op, Versus, Foxy Rush và Thực địa.

## Chơi trên web

- Mở trang GitHub Pages của kho trên Chrome, Safari hoặc Firefox. Game nằm ngay ở trang chủ.
- Để chơi 2–4 người, chọn **Chơi cùng bạn → Khác mạng Internet → Tạo phòng** bằng P2P. Chọn sức chứa, gửi link/QR hoặc mã phòng 6 ký tự, chọn chế độ và bắt đầu khi có ít nhất hai người.
- Chơi online trên GitHub Pages dùng PeerJS và WebRTC. Nếu mạng chặn kết nối trực tiếp, có thể dùng TURN đã cấu hình hoặc máy chủ HTTPS chạy relay Go (relay chỉ hỗ trợ hai người).

## Voice chat nhóm

Trong lobby hoặc trận, mở **CHAT / VOICE → BẬT VOICE** trên từng thiết bị và cấp quyền micro. Nếu không nghe được, chạm **BẬT NGHE** hoặc **NỐI LẠI VOICE**. Dùng link HTTPS trực tiếp trên Safari/Chrome. Không bảo đảm hoạt động ở mọi mạng hoặc khi ứng dụng chạy nền; TURN có thể cần cho mạng NAT nghiêm ngặt.

## Chạy relay

Cần Go 1.23 trở lên. Chạy lệnh trong thư mục dự án, cạnh `index.html`.

```sh
go run server.go
```

Máy chủ hiển thị địa chỉ để các thiết bị cùng Wi-Fi mở game. Với dịch vụ hỗ trợ Docker và HTTPS, có thể dùng `Dockerfile` và `render.yaml`. GitHub Pages chỉ phục vụ trang tĩnh.

Xem [hướng dẫn phím điều khiển, kỹ năng và kết nối](HUONG_DAN.txt).
