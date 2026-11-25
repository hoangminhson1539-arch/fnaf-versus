# Nhật ký cập nhật FNAF 1 VERSUS

Mã nguồn chính: `main`. [Chơi game](https://hoangminhson1539-arch.github.io/fnaf-versus/).

## V41.7 — Tactical Phase Contracts

- Thêm 3 giai đoạn 60 giây trong Versus và Foxy Rush. HUD hiển thị mục tiêu, phần thưởng, tiến độ và thời gian còn lại.
- Bảo vệ: radar, sửa 15 điểm độ bền, đẩy lùi quái bằng mồi nhử; thưởng 3/4/5% điện theo giai đoạn.
- Animatronics: 2 cú gây sát thương cửa, 2 quái khác nhau gây sát thương (Foxy Rush: 2 cú), 3 cú gây sát thương; thưởng tăng tốc và/hoặc hao điện của bảo vệ.
- Sửa lỗi validator relay không chấp nhận AI assault. Mồi nhử tác động cả quái đang điều khiển thủ công đủ gần cửa.
- Thưởng chỉ do host xử lý và tối đa một lần mỗi pha. Không đổi giao thức 42. Giữ chế độ solo/co-op.
- Kiểm tra cú pháp JavaScript và 16 kiểm tra mô phỏng nhiệm vụ đạt. Chưa kiểm tra trình duyệt, thiết bị thật, build Go hoặc multiplayer đầu cuối.

## V41.6 — Pitch-dark Hallways

- Hai cửa phòng trực: khi tắt đèn, hành lang ngoài cửa mở gần như đen hoàn toàn; chỉ giữ viền và ngưỡng cửa rất mờ. Không vẽ tường, sàn hay đèn hành lang khi không có ánh sáng.
- Kính quan sát trên cửa đóng tối sâu, hạn chế phản chiếu giả. Bật đèn bên nào thì chỉ bên đó hiện hành lang/quái như bình thường.
- Giữ quy tắc mắt V41.5 (Chuẩn/Khó/Hell 0%, Dễ 50%, Baby 100%). Không thay đổi gameplay, radar, sát thương, giao thức 42 hay tính năng nhìn gần cửa sổ.
- Kiểm tra cú pháp và mô phỏng lệnh vẽ 9 tình huống đạt; chưa kiểm tra ảnh thực tế trên trình duyệt.

## V41.5 — Dark Door Eyes

- Chuẩn (Normal): xác suất thấy mắt animatronic khi tắt đèn giảm từ 25% xuống 0%, kể cả lượt ngẫu nhiên đã lưu trước đó. Khó/Hell giữ 0%, Dễ 50%, Baby 100%.
- Bật đèn vẫn hiển thị quái tại cửa; không thay đổi AI, âm thanh, cảnh báo, sát thương cửa, kỹ năng hay giao thức 42.
- Cập nhật mô tả độ khó và nhãn phiên bản. Kiểm tra cú pháp và xác suất cả năm mức đạt; chưa kiểm thử giao diện thực tế.

## V41.4 — Balanced Door Pressure & Control Fixes

- Tăng sát thương so với V41.3 nhưng thấp hơn V41.2: Bonnie/Chica 9, Freddy 11, Foxy Versus 13, Foxy Rush 18 điểm độ bền mỗi cú hợp lệ. Hao điện trực tiếp chỉ 0,4% (Versus) hoặc 0,8% (Rush) khi gây sát thương.
- Hồi sát thương chung 5s/cửa cho cả đội; tiếng va đập hồi riêng 2,5s/cửa. HUD báo sát thương vừa nhận và thời gian hồi đập cửa cho phe quái. Không khóa thao tác hay gây nhiễu camera do đập cửa.
- Cửa nguyên vẹn cần tối thiểu 12/12/10/8/6 cú hợp lệ nếu không sửa. Khóa cửa vỡ tối thiểu 5s (Versus) / 8s (Rush); cần sửa tới 25%. Tốc độ sửa vẫn 5 điểm/giây.
- AI tổng công kích và nghi binh không tự điều động quái đang được chọn nếu quái đang rảnh; chọn quái được đồng bộ tới chủ phòng. Lệnh bị từ chối không xóa phím di chuyển đang giữ. Thời hạn lệnh hướng 0,55s và thử lại khi gửi lệnh tạm thất bại.
- Sửa trạng thái cửa thiếu dữ liệu, ngừng sửa khi mất kết nối, chặn sửa camera từ xa lúc trận tạm dừng, xóa lệnh di chuyển cũ khi khôi phục phiên. Xung phòng thủ đặt lại cả tiến độ đập cửa và tấn công.
- Giữ giao thức 42, radar, buff tốc độ và kỹ năng Foxy Rush. Hai người dùng cùng V41.4. Đã kiểm tra cú pháp và mô phỏng logic cửa/di chuyển; chưa kiểm thử trình duyệt, thiết bị thật hoặc multiplayer đầu cuối.

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
