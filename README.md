# FNAF 1 Versus · V47.1 Party Voice

**[Chơi ngay trên trình duyệt](https://hoangminhson1539-arch.github.io/fnaf-versus/)**

**[Nhật ký cập nhật chi tiết — CHANGELOG.md](CHANGELOG.md)**

## Nhật ký cập nhật trên README

> Quy trình bắt buộc: mỗi lần cập nhật game phải đồng bộ phiên bản, mục changelog mới nhất trong README.md và lịch sử đầy đủ trong CHANGELOG.md; nêu tính năng, sửa lỗi, cân bằng, kiểm thử và giới hạn thực tế.

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
