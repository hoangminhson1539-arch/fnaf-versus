# Nhật ký cập nhật FNAF 1 VERSUS

Mã nguồn chính: `main`. [Chơi game](https://hoangminhson1539-arch.github.io/fnaf-versus/).

## V41.3 — Door Knock / Distraction Balance

- Đập cửa ưu tiên nghi binh, tiếng va đập định hướng trái/phải và viền cảnh báo ngắn 1,6s trên HUD/nút cửa. Không tự khóa điều khiển, không trực tiếp làm hao điện hoặc gây nhiễu camera.
- Giảm sát thương độ bền mỗi lần đủ điều kiện: Bonnie 5, Chica 5, Freddy 6, Foxy Versus 8, Foxy Rush 10 (từ 15/15/18/20/34).
- Mỗi cửa có hồi sát thương chung 8s, kể cả nhiều quái cùng tấn công; tiếng động có hồi riêng 3s/cửa. Hai cửa độc lập. Foxy Rush cần đứng trước cửa đóng 1,5s trước khi đập.
- Cửa nguyên vẹn cần tối thiểu 20/20/17/13/10 cú gây sát thương tương ứng nếu không sửa, không tính thời gian tiếp cận và rút lui.
- Cửa hỏng khóa tối thiểu 4s (Versus) / 6s (Foxy Rush), cần sửa tới 20% để đóng lại. Giữ nguyên sửa 5 điểm/giây, phí 0,25% điện/giây, radar và tốc độ/kỹ năng Foxy Rush.
- Solo/Co-op không thay đổi cơ chế đập cửa. Giao thức 42 giữ nguyên, nên cả hai người chơi tải lại V41.3 để dùng cùng luật. Kiểm tra cú pháp và mô phỏng logic đạt; chưa kiểm thử thiết bị thật, multiplayer đầu cuối, GitHub Pages.

## V41.2 — Radar & Door Repair

### Radar và bản đồ
- Radar của bảo vệ trong Versus/Foxy Rush: 12 giây, 2% điện, hồi 18 giây. Hiển thị tên từng quái, tên phòng và mã CAM từ tọa độ thật; Foxy Rush chỉ hiện Foxy.
- Chạm kết quả radar để mở/chuyển tới camera tương ứng; camera không có góc nhìn trực tiếp không được gán mã giả.
- Đánh dấu CAM có quái bằng viền vàng và số lượng; hiển thị tên quái trên camera đang xem. Radar hoạt động độc lập khi camera nhiễu.
- Bản đồ chiến thuật luôn vẽ tên phòng, rút gọn và xuống dòng khi cần. Camera map có tên tiếng Việt; nút CAM nhanh có tên phòng. Hiển thị phòng của quái đang chọn.

### Phá cửa và sửa cửa
- Hai cửa có 100% độ bền riêng. Versus: Bonnie/Chica 15, Freddy 18, Foxy 20 điểm độ bền mỗi lần đập cửa đang đóng. Hao 1% điện (Foxy 1,5%).
- Foxy Rush: 34 điểm độ bền và 4% điện mỗi lần đập, giữ nguyên buff tốc độ và kỹ năng V41.1.
- Khi về 0%, cửa mở và bị khóa tối thiểu 7 giây (Versus) / 12 giây (Foxy Rush); phải sửa lên ít nhất 35% mới đóng lại được.
- Bảo vệ: Y/O hoặc nút SỬA TRÁI/PHẢI để bắt đầu/dừng sửa một cửa. Hồi 5 điểm độ bền/giây; tốn 0,25% điện/giây. Tạm ngừng khi cửa đóng, camera mở, hết điện hoặc quái cách cửa dưới 90 đơn vị. Hết điện hoặc đủ 100% thì kết thúc sửa.
- HUD hai phe hiển thị độ bền; bảo vệ có trạng thái sửa. Chủ phòng xử lý sát thương và sửa cửa, đồng bộ qua snapshot.
- Chơi đơn/Co-op không đổi cơ chế cửa.

### Kết nối và kiểm thử
- Giao thức mạng nội bộ 42, phiên bản hiển thị V41.2. Dùng cả hai client V41.2; relay Go cũ phải cập nhật và khởi động lại.
- Đã kiểm tra cú pháp JavaScript và mô phỏng logic sát thương, ngắt sửa khi có quái/camera mở, khóa cửa tối thiểu, khôi phục, tiêu hao điện, lọc radar Foxy Rush. Chưa kiểm thử trình duyệt/thiết bị thật, multiplayer đầu cuối hoặc biên dịch Go.

## V41.1 — Foxy Door Breaker
- Foxy Rush: tốc độ cơ bản 155, kỹ năng x2,2; X hồi 8s, Z hồi 12s, rút lui tối đa 3,5s.
- Foxy phá cửa sau 3 cú đập, khóa cửa 12s (quy tắc phục hồi được thay đổi ở V41.2).
- Bảo vệ: xung hồi 16s, pin +18%, radar 12s.

## V41 — Direct Control & Total Assault
- Điều khiển animatronics bằng WASD/mũi tên/joystick 8 hướng, va chạm, lệnh tự hết hạn.
- Versus cho 3 quái tấn công đồng thời; AI tổng công kích, nghi binh, tuần tra, giữ vị trí.
- Tăng tốc quái, giảm hồi kỹ năng, cải thiện hiển thị bản đồ, emoji và hotkey.
- Bảo vệ: pin +15%, radar 10s, xung 8s.

## V40.3 — Đồng bộ điều khiển và nhãn phiên bản
- Sửa nhãn V24 còn sót; dùng số phiên bản trong tiêu đề HTML làm nguồn hiển thị.
- Cải thiện phản hồi lệnh, kiểm tra mất kết nối và cho phép đổi quái khi tạm dừng.

## V40.2 / V40.1 / V40 — Squad Tactics
- Lệnh điều quân, kỹ năng riêng, điều khiển nhiều quái, AI hỗ trợ, công cụ phòng thủ và mô tả kỹ năng.

Các bản cũ được tóm tắt từ tính năng và lịch sử đã phát triển, không phải nhật ký kiểm thử đầy đủ.
