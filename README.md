# FNAF 1 Versus · V54.0 Field Command Center & Emergency Response

**[Chơi ngay trên trình duyệt](https://hoangminhson1539-arch.github.io/fnaf-versus/)**

**[Nhật ký cập nhật chi tiết — CHANGELOG.md](CHANGELOG.md)**

## Nhật ký cập nhật trên README

> Quy trình bắt buộc: mỗi lần cập nhật game phải đồng bộ phiên bản, mục changelog mới nhất trong README.md và lịch sử đầy đủ trong CHANGELOG.md; nêu tính năng, sửa lỗi, cân bằng, kiểm thử và giới hạn thực tế.

### V54.0 — Field Command Center & Emergency Response (26/11/2026)

- Thiết kế trung tâm điều phối mới cho người xem camera trong Thực địa và Tiếp tế khẩn cấp: bản đồ chiến thuật lớn trên sân chơi, chuyển giữa bản đồ và CCTV, thanh chọn nhanh camera, hiển thị tín hiệu/mất kết nối, công tắc bật/tắt camera để tiết kiệm điện. Hai công cụ tương tác bản đồ riêng: chỉ đường và chọn camera theo phòng; tránh đặt điểm ngoài ý muốn. Bố cục PC và điện thoại ngang, có điều chỉnh cho màn hình dọc.
- Bảng theo dõi đội: từng bảo vệ có vị trí phòng, trạng thái an toàn/cảnh giác/truy đuổi, giáp, khói, trạng thái di chuyển hoặc tiến độ thao tác, mốc đường đã tới. Danh sách ba nhiệm vụ và mốc giờ mở; nhật ký điều phối 8 sự kiện gần nhất. Camera hiển thị tiến độ của bảo vệ đang thao tác, kể cả BV2.
- Theo dõi tuyến chỉ đường riêng cho BV1/BV2; tự ghi nhận khi tới cách mốc tối đa 35 đơn vị. Nút Đi mốc tiếp để tự tìm đường theo tuyến đã nhận; bảo vệ tự quyết định di chuyển. Đổi tuyến đặt lại tiến độ. HUD hiện mốc kế tiếp và khoảng cách.
- SOS/Báo an toàn: từng bảo vệ có hồi chung 10 giây, tín hiệu vị trí phòng tồn tại 15 giây, camera xác nhận riêng từng người. SOS đánh dấu trên bản đồ và phát âm thanh cho người xem camera; nút SOS/Mốc tiếp ngay trên cảnh chơi. Không ép di chuyển hoặc cấp vật phẩm.
- Tập hợp: chỉ đường về phòng trực, hiệu lực 45 giây, hồi 25 giây; đồng bộ cho cả hai bảo vệ. Xung cứu viện: làm choáng một quái trong CAM đang chọn, cách ít nhất một bảo vệ không quá 180 đơn vị, trong 2,5 giây; tiêu hao 3% điện cơ sở theo thời lượng (tối thiểu 1,2%), hồi 42 giây nhân hệ số thời lượng. Không áp dụng khi camera tắt/nhiễu, thiếu điện, quái đang choáng/rút lui hoặc ngoài phạm vi. Giữ hệ thống cảnh báo và hỗ trợ thao tác V53.
- Sửa lỗi snapshot lấy nhầm vị trí chính khi BV2 đang thực hiện hành động gây gửi trạng thái ngay; chặn vẽ HUD giữa lúc hoán đổi trạng thái BV2. Đồng bộ hướng khi tự tìm đường, trạng thái AI và thời gian choáng của quái; giới hạn di chuyển thực địa trong các phòng hợp lệ; giữ mục tiêu chạm bản đồ khi SVG cập nhật và hủy thao tác nếu mất pointer capture. Tương thích cấu trúc checkpoint cũ của hệ thống điều phối.
- Relay hai người: thêm whitelist SOS, xác nhận, đi mốc, tập hợp, xung cứu viện; phải triển khai lại server.go để sử dụng. Tiếp tế khẩn cấp vẫn P2P; các máy P2P cần cùng V54.0. Giữ cân bằng điện và Foxy Rush V52.0.
- Kiểm tra cú pháp toàn bộ JavaScript đạt ở hai commit; 27 kiểm tra logic trích hàm trong môi trường giả lập trên blob cuối đạt, gồm quyền/cooldown SOS, xác nhận hai người, tuyến/mốc, tập hợp, điều kiện xung cứu viện, giới hạn nhật ký, lọc snapshot và lỗi hoán đổi BV2. Chưa chạy thử trình duyệt, gameplay đa thiết bị, biên dịch Go hoặc xác nhận triển khai GitHub Pages. Không khẳng định mọi lỗi đã hết.

### V53.0 — Field Command & Tactical Intelligence (26/11/2026)

- Thực địa và Tiếp tế khẩn cấp: người xem camera chạm phòng trên bản đồ để đặt tuyến tối đa 3 điểm theo thứ tự, có dấu xanh đánh số trên bản đồ và cảnh chính của cả hai bảo vệ. Dấu tồn tại 35 giây, hồi đặt/xóa 1,5 giây; nút chỉ mục tiêu đặt nhiệm vụ hiện tại hoặc phòng trực. Bảo vệ tự quyết định đường di chuyển, có xác nhận riêng BV1/BV2, không bị điều khiển từ xa.
- Cảnh báo quái: chạm quái đang hiện trên camera/bản đồ hoặc bấm Báo quái để chọn quái nhìn thấy gần đội nhất. Host xác thực camera đang hoạt động, không bị nhiễu, chỉ chấp nhận quái trong camera được chọn hoặc trong đợt quét. Đánh dấu vị trí ghi nhận (không bám theo quái) trong 9 giây, hồi chung 14 giây, đồng bộ HUD và âm thanh cho bảo vệ.
- Hỗ trợ thao tác từ xa: tăng tốc độ thực hiện nhiệm vụ 40% trong 6 giây khi có ít nhất một bảo vệ đang thao tác; tiêu hao 1,5% điện cơ sở theo thời lượng (tối thiểu 0,65%), hồi 28 giây nhân hệ số thời lượng. Không bỏ qua yêu cầu vị trí, mốc giờ hoặc nhiệm vụ.
- Giao diện: bảng điều phối hiển thị tuyến, thời gian còn lại, xác nhận từng bảo vệ, cảnh báo và hỗ trợ; HUD cảnh chính hiển thị thông tin cô đọng. Màu xanh cho chỉ đường, cam cho cảnh báo, xanh lá cho hỗ trợ. Sửa hướng dẫn phím thực địa và thao tác chạm trên bản đồ SVG thường xuyên cập nhật; không tính vuốt dài là chọn mục tiêu.
- Sửa lỗi Thực địa/Tiếp tế: mồi âm thanh không được đặt sát bất kỳ bảo vệ nào và đồng bộ mồi cho cả hai; đi thẳng nếu đường không bị chắn (tránh lỗi đường rỗng cùng ô điều hướng); vô hiệu nút thao tác khi đang di chuyển; ẩn HUD/joystick ở màn kết quả. Giữ cơ chế điện và nerf Foxy V52.0.
- Snapshot chỉ đường/cảnh báo/hồi chiêu được chuẩn hóa khi nhận, đặt lại ở ca mới. Phân quyền host và thời gian hồi chung giữa các camera. Relay hai người có whitelist lệnh mới trong server.go; cần triển khai lại relay. Tiếp tế khẩn cấp vẫn P2P.
- Kiểm tra cú pháp JavaScript toàn bộ đạt ở cả hai commit; 25 kiểm tra logic trích hàm trong môi trường giả lập đạt: phân quyền, giới hạn tuyến và thời gian hồi, xác nhận hai bảo vệ, cảnh báo có xác thực tầm nhìn, giữ nguyên vị trí ghi nhận, hỗ trợ và lọc snapshot. Chưa thử trình duyệt/thiết bị thật, chưa biên dịch Go hay xác nhận triển khai GitHub Pages; P2P yêu cầu cùng V53.0.

### V52.0 — Direct Field Control, Anti-Camp & Emergency Supply (26/11/2026)

- Thực địa: WASD/phím mũi tên và joystick cảm ứng 8 hướng; tốc độ 145 đơn vị/s, chuẩn hóa đường chéo, dùng va chạm hiện hữu. Lệnh điều khiển làm hủy đường tự động và thao tác đang thực hiện, tự hết hạn sau 0,5 giây; gửi lại mỗi 125ms khi giữ. Dừng khi thả, mất tiêu điểm hoặc tạm dừng mạng. F/Space thao tác, X dùng khói, R về phòng. Bổ sung whitelist relay hai người cho điều khiển bảo vệ và túi tiếp tế. Relay cần triển khai lại server.go; chưa xác nhận đã triển khai.
- Thực địa: HUD khu vực, số giáp/khói, cảnh giác, nhiệm vụ; joystick và bốn nút nhanh trên cảnh chính; bảng nhiệm vụ mới, viền trạng thái, thanh bên thu gọn. Desktop tối đa 1280px, thanh bên 245–290px trên màn hình đủ lớn; có bố cục riêng cho điện thoại ngang. Đồng bộ hướng, trạng thái di chuyển, giáp và khói; ngăn thao tác nhiệm vụ tiếp tục nhầm mục tiêu khi đồng đội hoàn thành trước.
- Mode multiplayer P2P mới **Tiếp tế khẩn cấp**: 2–4 người, 75% điện xuất phát, 1 khói/bảo vệ; ba nhiệm vụ lần lượt ở vệ sinh, sân khấu, Pirate Cove, mở tại 12 AM, 1 AM, 3 AM, thưởng tối đa 6/7/9% điện; về phòng trước 6 AM. Dùng chung cơ chế một hoặc hai bảo vệ và camera điều phối. Không có mode này trên relay.
- Điện: tăng hệ số tiêu hao multiplayer Co-op 0,93→1,04; Versus 0,90→1,01; Thực địa 0,90→1,02, trước hệ số thời lượng; giảm tiếp tế giờ và bonus thời lượng ngoài Foxy Rush. Solo tăng tiêu hao cơ sở 5%, không thay đổi cheat noDrain. Giữ nguyên công thức điện và tiếp tế giờ Foxy Rush; mức sát thương điện từ Foxy giảm theo nerf kỹ năng.
- Anti-camp: trong bán kính 85 đơn vị quanh cửa, quái người chơi tích thời gian khi cửa đóng hoặc đứng chờ/di chuyển khi cửa mở; buộc rút sau 5,5s (Foxy Rush 3,8s), không reset chỉ vì đổi lệnh. Tấn công thực sự qua cửa mở vẫn hợp lệ. AI có giới hạn đứng gần cửa đóng. Phe quái Versus/Foxy Rush phải quay tới gần điểm xuất phát và chờ hồi trước khi ra lệnh tiếp.
- Foxy Rush: giảm hệ số tốc độ 0,94→0,84 (phòng 2 người), sát thương 0,875→0,75; giảm tốc độ/thời lượng tăng tốc và sát thương điện; tăng hồi chiêu, thời gian rút lui và cửa sổ phản ứng bảo vệ, áp dụng hệ số riêng cho phòng 3–4 người.
- Kiểm tra cú pháp toàn bộ JavaScript đạt trước commit 746a48e5037a26acefc0ef7479ce1776e8636d4d; đọc lại blob và chạy 68 kiểm tra logic trích hàm với môi trường giả lập: đạt (điện theo 5 thời lượng, giữ nguyên công thức điện Foxy Rush, hệ số nerf, di chuyển và hết hạn, phân quyền, mode tiếp tế, anti-camp). Chưa chạy thử gameplay P2P/relay, chưa biên dịch Go, chưa kiểm tra thực tế Safari/iPhone và các kích thước màn hình. P2P yêu cầu cùng V52.0.

### V51.0 — Field Patrol Duo & Compact Desktop (26/11/2026)

- Thực địa P2P: phòng 2 người có 1 bảo vệ và 1 camera; phòng 3–4 người tự phân 2 bảo vệ và 1–2 camera. Bảo vệ thứ hai có vị trí, tuyến đường, thao tác nhiệm vụ, trạng thái nguy hiểm, 1 khói và 1 giáp riêng. Nhiệm vụ, điện, cửa và tiếp tế phụ dùng chung. Chủ phòng xử lý lệnh theo seat của kết nối; quái chọn bảo vệ gần hơn. Cả hai bảo vệ phải về phòng trực để thắng.
- Tái sử dụng cảnh khám phá tự do của chế độ thường cho góc nhìn bảo vệ thực địa: nhân vật, vật cản, ánh sáng, tuyến đi, đồng đội và mục tiêu. Chạm vào phòng trên cảnh chính hoặc bản đồ để đi; điều phối viên giữ góc camera. Không bổ sung điều khiển di chuyển tự do WASD trong bản này.
- Chuyển bảng bản đồ/nhiệm vụ khỏi lớp phủ cảnh sang thanh điều khiển; thu gọn nút trùng, bản đồ tỷ lệ cố định, giới hạn chiều rộng desktop 1440px và thanh bên 260–320px trên màn hình đủ lớn. Giữ khả năng cuộn thanh bên khi chơi ngang trên điện thoại.
- Sửa kiểm tra vùng cửa để tránh đóng cửa khi một trong hai bảo vệ ở ngưỡng cửa; khói và tiến độ đường đi của bảo vệ thứ hai được đưa vào snapshot.
- Đã đọc lại index.html từ GitHub sau commit 9b4a7bf91fcf45704b6b41c67f3582b40a67db91 và kiểm tra cú pháp JavaScript toàn bộ script: đạt. 19 kiểm tra logic với các hàm trích từ mã đã commit và môi trường giả lập: đạt (phân vai phòng 2–4 người, vị trí riêng, chia sẻ nhiệm vụ/tiếp tế, snapshot, khôi phục ngữ cảnh khi lỗi, điều kiện thắng, kiểm tra ngưỡng cửa, relay). Chưa chạy thử gameplay multiplayer, WebRTC hoặc đo bố cục trên thiết bị thật; không xác nhận đã sửa toàn bộ lỗi giao diện. Phân đội tự động theo số người, chưa có tùy chọn 1/2 độc lập ở lobby. P2P dùng mã phiên bản V51.0; relay vẫn hai người.

### V50.0 — Command Center & Tactical Expansion (26/11/2026)

- Xác nhận lệnh cửa riêng theo người gửi, mã lệnh và ca trực; bỏ lệnh trùng, giới hạn thao tác quá nhanh; hỗ trợ cửa thường và cửa Thực địa P2P. Không dùng thông báo cửa của người khác để xác nhận lệnh đang chờ. Sửa nút mở cửa bị vô hiệu hóa khi cửa hỏng nhưng còn đóng; giảm âm thanh cửa lặp.
- Phát hiện kênh P2P còn báo mở nhưng không truyền heartbeat trong 12 giây khi hai phía ở foreground, thử tạo lại kênh bằng hệ thống giữ chỗ 30 giây. Không tự tạo nhiều kết nối khôi phục cùng lúc. Không đảm bảo khôi phục nếu đóng tab hoặc mất chủ phòng.
- Bảng TÍN HIỆU ĐỘI thu gọn cho Co-op và Thực địa P2P: trái, phải, an toàn, tiết kiệm điện, sửa hệ thống, về phòng; hiện P1–P4 gửi tín hiệu, hạn chế spam. Relay giữ ba tín hiệu cũ.
- Ba kỹ năng chiến thuật multiplayer: khởi động cầu chì (2,5% điện cơ bản, hồi 36s), phục hồi camera (3%, hồi 34s), tích áp +12 AP khi quái không tấn công/điều khiển tay/rút lui và áp lực không quá 55 (hồi 40s). Chi phí/hồi chiêu áp dụng hệ số thời lượng; quyền quái và cooldown vẫn kiểm tra ở host.
- Thực địa: túi tiếp tế phụ một lần tại khu vệ sinh từ 3 AM, tối đa +4% điện và +1 khói (tối đa 2 khói), cần đứng gần, không đang di chuyển/làm nhiệm vụ. Không tăng số nhiệm vụ bắt buộc. Chỉnh thống kê pin khẩn cấp theo lượng thực nhận.
- Giải phóng track đầu ra audio thay thế khi đổi đường phát. Giữ bố cục ngang, voice 4 người, tutorial đêm 1, cân bằng pin/Foxy Rush trước đó.
- Kiểm tra JavaScript cuối đạt; 38 kiểm tra mô phỏng P2P 4 người/biên heartbeat và mã lệnh, 23 kiểm tra logic kỹ năng/túi tiếp tế đạt trước lần tinh chỉnh cuối về chống lặp và điều kiện quái. Chưa thử trên thiết bị thật. P2P yêu cầu đồng phiên bản V50.0; relay vẫn hai người.

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

Game fan-made chạy trên trình duyệt. Có chế độ một người và 5 chế độ multiplayer: Co-op, Versus, Foxy Rush, Thực địa và Tiếp tế khẩn cấp (P2P).

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
