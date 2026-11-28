# Nhật ký cập nhật FNAF 1 VERSUS

Mã nguồn chính: `main`. [Chơi game](https://hoangminhson1539-arch.github.io/fnaf-versus/).

**Quy trình bắt buộc:** Mỗi lần cập nhật game, phải cập nhật file `CHANGELOG.md` trên GitHub cùng phiên bản mã nguồn. Ghi rõ phiên bản, ngày, tính năng mới, sửa lỗi, thay đổi cân bằng, giới hạn và kết quả kiểm thử thực tế; không ghi tính năng chưa triển khai là đã hoàn thành.

## V62.1 — Cuộc gọi Đêm 2 (28/11/2026)

- Bổ sung 4 câu thoại của Minh ở cuối Đêm 2: các khu bí mật đóng khi nhà hàng đón khách ban ngày, mở về đêm; có camera riêng nhưng cả bảo vệ lẫn Minh không có quyền xem; Minh không biết mục đích sử dụng. Không tiết lộ thêm cốt truyện.
- Phát tại các mốc 170, 188, 206, 224 giây của ca 240 giây, tự co giãn theo thời lượng ca; mỗi đoạn hiện 9 giây. Giữ 4 đoạn cũ, giữ mã lưu nhật ký đã nghe (n2:0–7), không dừng gameplay. Đọc lại qua nhật ký cuộc gọi.
- Chỉ cập nhật nội dung cốt truyện, giữ gameplay V62.0 và giao thức P2P v620. Kiểm tra cú pháp JavaScript và 24 điều kiện lịch phát/nhật ký/tương thích đạt; chưa kiểm tra hiển thị thực tế trên điện thoại hoặc xác nhận Pages triển khai.

## V62.0 — Fair Repair & Power (27/11/2026)

- Versus / Foxy Rush: giảm hệ số tốc độ sửa chung .88 → .68; hệ số sửa Foxy .92 → .85 ở 2 người (vẫn giảm .04 mỗi người thêm); bonus giờ bảo trì 15% → 8%. Mỗi lượt sửa tối đa 30 điểm độ bền Versus / 25 Foxy; hồi chung hai cửa 8s / 10s sau khi kết thúc hoặc bị ngắt. Không tự tiếp tục khi quái rời đi; đổi cửa giữa lượt sẽ hủy lượt và kích hoạt hồi. Chờ 4s / 6s sau va đập, cấm sửa khi cửa còn trong thời gian hỏng, camera bật hoặc quái cách điểm cửa dưới 100 / 115 đơn vị. Kiểm tra host và camera nhiều người. HUD hiện tiến độ và thời gian hồi; chặn nút khi chưa đủ điều kiện.
- Chi phí sửa: 0,17% / 0,20% điện cho mỗi điểm độ bền (trận 3 phút, Versus / Foxy), nhân max(0,55; sqrt(180/thời lượng giây)). Không vượt 100% độ bền, không âm điện. Vá nhanh giảm 20 → 12 / 8 điểm, chi phí max(2%; 4,5% / 5% × sqrt(180/thời lượng)), hồi 55s / 65s nhân hệ số thời lượng; dùng chung khóa sửa, không vá khi đang sửa, cửa chưa hết thời gian hỏng, quái ở gần hoặc camera bật.
- Điện multiplayer: phụ phí mỗi cửa Versus .29 → .25, Foxy .35 → .30 trước hệ số thời lượng; hệ số hao nền 3 phút Versus 1.00, Foxy 1.02, Co-op/Thực địa 1.00, Tiếp tế khẩn cấp .96; giảm .006 mỗi bậc thời lượng. Tiếp tế Foxy giảm bonus theo bậc .75 → .60, giữ tiếp tế theo giờ và pin dự phòng. Giữ độ khó, kỹ năng phe quái và cân bằng chơi đơn. Không giảm tốc độ mô phỏng.
- Sửa khai thác vật thể thực địa: thiết bị nhiệm vụ của mode tương ứng chỉ cấp thưởng qua nhiệm vụ, không nhận thưởng phụ trước/sau khi hoàn thành. Giữ tương tác với vật thể khác và đi khẽ V61.0.
- Kiểm thử: cú pháp toàn bộ JavaScript đạt; 317 kiểm tra hàm/điều kiện giả lập trên cân bằng, chi phí, giới hạn sửa, hủy khi mất an toàn, hồi chung, 5 thời lượng, 5 mode, 2–4 người ở đối đầu; 16 kiểm tra phân loại thiết bị nhiệm vụ đạt. Chưa thử gameplay thật, tỷ lệ thắng, FPS hoặc xác nhận Pages triển khai. P2P cần cùng V62.0; server.go không đổi từ V61.0.

## V61.0 — Field Interaction & Stealth (27/11/2026)

- Thực địa và Tiếp tế khẩn cấp: tương tác với 43 vật thể trong 10 khu vực multiplayer hiện có. Chạm vật thể ở xa để đi đến; khi đủ gần và không bị tường/cửa chắn, chạm lại để dùng. Nút Tương tác hoặc F/Space chọn nhiệm vụ gần nhất đang mở, nếu không thì dùng vật thể gần nhất; joystick và các nút điều khiển cũ giữ nguyên.
- Đi khẽ: nút riêng hoặc phím C; tốc độ di chuyển tay và tự động 62% bình thường, tiếng bước chân giảm (noiseLevel 40→12), AI dùng cơ chế crouch/hearing đã có. Đồng bộ riêng cho từng bảo vệ; không còn đặt lại crouch mỗi tick. Không bất tử hoặc tàng hình tuyệt đối.
- Vật thể: cửa/bảng cửa, máy phát, pin, mạch điện, tủ cầu chì, tủ đồ/khói/giáp, máy nước/thiết bị giảm tiếng, loa/chuông đánh lạc hướng, tiếp sóng/đài liên lạc, bàn trực và hồ sơ. Vật tư có giới hạn dùng chung trong ca; host xác thực khoảng cách, đường nhìn, trạng thái dừng, hồi chiêu và quyền vai trò. Hồ sơ/thẻ/băng/cửa thoát cho phép kiểm tra nhưng không cấp tiến trình chơi đơn hoặc kết thúc ca sớm. Chưa mở các khu bí mật và mở rộng ngoài 10 khu hiện tại.
- Versus/Foxy Rush: Animatronics có thao tác phá thiết bị theo ngữ cảnh khi cách thiết bị tối đa 73 đơn vị và có đường nhìn. Phá máy phát/cầu chì/trạm sạc làm hỏng đèn 5s; phá bàn giám sát/tiếp sóng gây nhiễu CCTV 5s. Dùng 28 AP, chia sẻ hồi chiêu phá tín hiệu 36s nhân hệ số thời lượng, không thêm tài nguyên. Kiểm tra quyền nhân vật, giữ cơ chế di chuyển/chiến đấu cũ. Co-op và các mode bàn trực giữ vai trò cố định, mở rộng lại kỹ năng chính (gia cố, khởi động cầu chì, theo dõi hành lang, đổi hướng/phá tín hiệu); không chuyển thành thực địa tự do.
- Giữ HUD và tối ưu Auto V60.0, thêm nút đi khẽ và tương tác theo ngữ cảnh; tránh tính lặp vật thể gần trong danh sách hành động. Không thêm animation nặng hoặc bản đồ mới. Giữ giới hạn nhiệm vụ bắt buộc và luật thắng ca.
- Relay Go hai người: bổ sung danh sách cho phép lệnh vật thể, đi khẽ, phá thiết bị, tích áp và khởi động cầu chì ở nhánh đối đầu. Phải triển khai lại server.go nếu dùng relay riêng; P2P tất cả máy cần V61.0. Chưa biên dịch/chạy relay thật.
- Kiểm thử: cú pháp JavaScript toàn bộ đạt trên bản cuối; 453 kiểm tra hàm/điều kiện/cấu trúc, gồm 43 vật thể, quyền guard/camera, whitelist relay, tọa độ, chống lệnh sai, khoảng cách/đường nhìn, vật tư giới hạn, cooldown và phá thiết bị. Chưa kiểm thử trên thiết bị thật, đường đi đầy đủ tới từng vật thể, gameplay nhiều người, cân bằng thực chiến, FPS hoặc xác nhận triển khai Pages.

## V60.0 — Focused Gameplay & Adaptive Performance (26/11/2026)

- Nội dung chơi đơn: thêm Mạch điện luân phiên (mỗi giờ đổi cửa quá tải trái/phải, đóng cửa bên đó hao thêm 0,20% điện/giây; thưởng 170 xu) và Camera giới hạn (ngân sách CCTV bằng 25% thời lượng ca, tự tắt và khóa khi hết ngân sách; thưởng 180 xu). Dùng đồng hồ, cửa, âm thanh và camera hiện có, không thêm tài nguyên hay nút mới. Giữ cơ chế gian lận và lưu camWatched hiện có. Tổng 25 thử thách; menu đề xuất 6 thử thách, 19 thử thách cũ còn lại trong mục mở rộng.
- Menu: ưu tiên chơi đơn, multiplayer, tiếp tục, tài khoản và cài đặt; chuyển độ khó vào mục thu gọn, chuyển phím tắt và nhật ký vào mục tiến trình/hướng dẫn. Sửa số lượng thử thách hiển thị và mô tả mở boss. Không xóa tiến trình, thử thách hay vật phẩm đã có.
- Multiplayer: chế độ chiến thuật gọn chỉ hiện hai kỹ năng chính theo vai trò (bảo vệ: tiết kiệm/pin khẩn; camera: quét đội/khôi phục hệ thống; quái: cường hóa/tích áp), pin đã dùng sẽ ẩn. Nút Thêm kỹ năng mở toàn bộ kỹ năng cũ; giữ nút cửa, sửa cửa, điều khiển quái và điều phối thực địa. Nút kỹ năng giữ DOM ổn định khi thời gian hồi thay đổi, giảm mất thao tác do dựng lại nút. Không đổi cân bằng Foxy Rush V58.0 hay luật multiplayer.
- Hiệu năng: HUD thông thường trên cảm ứng cập nhật 150 ms thay vì 100 ms, HUD multiplayer 200 ms (vẫn cập nhật ngay khi thao tác yêu cầu); chỉ dựng tiêu đề/danh sách chọn camera khi CAM thay đổi; bỏ tính toán mô tả chiến thuật đang ẩn. Auto đánh giá mỗi 1,5 giây thay vì 2,5 giây, giảm FPS/DPR khi chi phí render cao hoặc tỷ lệ khung lấy mẫu chậm vượt 25%; cần 3 cửa sổ ổn định để tăng chất lượng. Giữ giới hạn FPS và thời gian mô phỏng cũ.
- Kiểm thử: cú pháp toàn bộ script đạt; 68 kiểm tra hàm/điều kiện/cấu trúc ở commit nội dung, 36 kiểm tra hàm Auto ở commit tối ưu đạt. Chưa đo FPS thật, chơi thử trình duyệt/đa thiết bị, kiểm thử luồng OTP/OAuth hoặc xác nhận triển khai GitHub Pages. P2P yêu cầu cùng V60.0; relay Go không thay đổi. Không bảo đảm mượt mọi lúc trên mọi máy.

## V59.0 — Stability & Mobile Recovery (26/11/2026)

- Multiplayer: tách bước gửi snapshot của host, render và HUD thành các vùng bắt lỗi độc lập; gửi trạng thái trước render để lỗi đồ họa không làm mất bước đồng bộ. Giới hạn log lỗi riêng từng vùng 2 giây. Không sửa giao thức relay Go; P2P sử dụng mã phòng V59.0 riêng, tất cả máy phải cập nhật cùng phiên bản.
- Mobile: không dựng lại dải 11 camera bị ẩn; không cập nhật danh sách nút đội/nhiệm vụ khi bảng đội đang đóng, vẫn cập nhật nút thao tác nhanh; khi mở bảng sẽ cập nhật lại. Sửa nhãn chất lượng Auto thành 18–60 FPS. Không thay đổi tốc độ mô phỏng.
- Tài khoản: giữ OTP đang nhập và tiêu điểm khi form được vẽ lại nếu vẫn cùng email yêu cầu mã; không lưu mã vào localStorage/sessionStorage. Giữ trạng thái lỗi kiểm tra dịch vụ khi chưa có kết quả, giới hạn kiểm tra lại sau lỗi 10 giây, kiểm tra thành công được cache 60 giây. Không thay cấu hình SMTP/OAuth/RLS.
- Giữ nguyên cân bằng V58.0, bao gồm sức mạnh Foxy Rush, hao điện và tốc độ sửa cửa.
- Kiểm tra trước commit: cú pháp toàn bộ script JavaScript đạt; 7 kiểm tra thực thi hàm cô lập lỗi/log và 5 kiểm tra cấu trúc mã đạt. Chưa kiểm thử trên trình duyệt/điện thoại thật, FPS, gameplay đa thiết bị, OTP/OAuth hoặc xác nhận triển khai Pages. Không khẳng định toàn bộ lỗi đã hết.

## V58.0 — Foxy Rush Rebalance & Mobile Lite (26/11/2026)

- Tài khoản: kết nối URL và publishable key của dự án Supabase thật của chủ game; xác minh bảng fnaf_saves đã có RLS, quyền chỉ authenticated, policy theo auth.uid và trigger revision. Không dùng secret. Game kiểm tra public auth settings, chỉ hiển thị OAuth đã bật, thông báo rõ lỗi SMTP/email/provider; giữ đăng ký email OTP/link, mật khẩu tùy chọn và hướng dẫn trong game. Chưa kiểm thử đăng nhập thật hoặc cấu hình SMTP/redirect/provider. SMTP mặc định giới hạn thành viên dự án và 2 thư/giờ; cần quản trị viên thiết lập SMTP để mở đăng ký công khai. Không tắt xác minh email.
- Foxy Rush: hệ số tốc độ cơ sở .84→.94, sát thương .75→.87, tăng tốc kỹ năng 1.55→1.72, thời gian 3.5→3.8s, hồi 13→11.5s, gây nhiễu hồi 19→17s, hao điện gây nhiễu 1.5→1.8, va cửa .50→.65, thời gian cửa hỏng 8→9s. Giữ hệ số theo số người và chống camp cửa. Hệ số điện nền .88→1.02, phụ phí cửa .31→.35; giữ tiếp tế.
- Nerf sửa cửa multiplayer Versus/Foxy Rush: hệ số chung 1.10→.88 (giảm 20% tốc độ), hệ số Foxy Rush 1.06→.92 ở 2 người (tổng giảm khoảng 30,6%); phí điện sửa .25→.30 Versus và .36 Foxy Rush theo hệ số thời lượng. Chờ 2s sau va đập (4s Foxy Rush) mới sửa, giữ yêu cầu mở cửa, hạ camera, tránh quái. Không đổi điện các mode còn lại.
- Camera mobile: một vùng chơi chính trong viewport, nút ĐỘI/SOS mở bảng nhiệm vụ, kỹ năng, chat và điều khiển cửa riêng; đóng để trở lại camera. Thay 11 nút CAM bằng danh sách chọn gọn khi xem CCTV, giữ Bản đồ/CCTV, chỉ đường/chọn CAM theo phòng; chia màn chỉ PC, ba nút nhanh Báo quái/Cứu viện/Tập hợp, bỏ đồng hồ/điện trùng. Giữ panel đầy đủ trên PC.
- Hiệu năng: Auto cảm ứng giảm DPR xuống .75/.9/1 theo mức 18/24/30 FPS, mục tiêu tối đa 650k pixel; bỏ lớp trang trí khí quyển ở Auto cảm ứng, tránh dựng map phụ khi CCTV và bảng đội đóng, không dựng danh sách đội ẩn; giảm tần suất HUD thực địa trên cảm ứng và lúc mất mạng. Không đổi logic thời gian mô phỏng.
- Sửa lỗi hiển thị dấu cộng/trừ hao điện, lỗi bố cục camera lặp/tràn trên mobile, tránh cập nhật layout đã đóng; tăng rõ ràng thông báo lỗi đăng nhập. Kiểm tra cú pháp toàn bộ JS và 61 kiểm tra logic trích hàm trên bản cuối đạt (5 thời lượng thực tế, 2–4 người, sửa cửa, DPR, provider). Chưa thử thực tế iPhone/PC, đo FPS, gameplay đa thiết bị, OTP/OAuth hay xác nhận GitHub Pages. P2P cần cùng V58.0; server.go không đổi. Nhật ký trong game, README và CHANGELOG đã cập nhật.

## V57.0 — In-game Accounts & Performance (26/11/2026)

- Tài khoản: nhập email → gửi mã → nhập OTP 6 số trực tiếp trong game; lần đầu tự tạo tài khoản qua Supabase, giữ link email và OAuth làm phương án thay thế. Kiểm tra mã, giới hạn thử lại 3 giây và gửi lại 60 giây, xác minh đúng email đã yêu cầu, hỗ trợ bàn phím OTP/Enter, nhớ email đang xác minh trong sessionStorage tối đa 10 phút khi tải lại. Không lưu mã hoặc mật khẩu trong bộ nhớ lưu trữ. Đăng xuất chỉ phiên thiết bị hiện tại.
- Hướng dẫn tài khoản, lưu/tải tiến trình, kích hoạt dịch vụ và tối ưu hiệu năng có ngay trong game. Chưa có cấu hình Supabase thật; cần quản trị viên kích hoạt một lần, đặt mẫu Magic Link có Token 6 số và cấu hình SMTP/provider khi cần. ACCOUNT_SETUP.md cập nhật. Người chơi không tự cấu hình backend; không có tài khoản giả.
- Hiệu năng thực địa: cache lớp phòng tĩnh của SVG, cập nhật lớp dấu động riêng, không dựng lại phòng mỗi lần vị trí thay đổi. Giữ map khi đang nhấn; dấu không tương tác không chặn chọn phòng. Culling cửa, tường, nhãn phòng và vật thể nằm ngoài khung nhìn có khoảng đệm; không đổi collision, AI hay luật chơi.
- Chất lượng Auto tính thêm chi phí mô phỏng trước lúc render, giảm hiệu ứng nhiễu ở mức tải thấp; joystick cache kích thước khi kéo, không đo layout mỗi lần di chuyển. Giữ tối ưu V56.
- Giao diện: trường nhập 16px hạn chế zoom tự động Safari, OTP riêng, nút tài khoản lớn; camera gọn trên màn hình dọc và ngang thấp, ẩn thanh hành động trùng trên ngang dưới 420px (vẫn có ở sidebar), giới hạn tràn nội dung.
- Kiểm tra cú pháp toàn bộ JavaScript đạt; 22 kiểm tra logic trích hàm (OTP, vùng culling, cache và cập nhật lớp bản đồ) đạt. Chưa kiểm thử trình duyệt thực tế, FPS, OAuth/OTP thật hoặc triển khai Pages; P2P yêu cầu cùng V57.0. Không đổi server.go. Cập nhật changelog trong game, README và CHANGELOG.

## V56.0 — One-tap Accounts & Smooth Play (26/11/2026)

- Safari iPhone/Android: chặn chọn chữ, menu nhấn giữ và kéo văn bản trong khu chơi; xóa vùng chọn cũ khi chạm, giữ khả năng nhập/chọn chữ trong chat, cài đặt và tài khoản. Thu gọn joystick/HUD, giảm chiều cao cảnh dọc và cho phép cuộn bảng bên dưới. Nút khôi phục âm thanh nổi không che game multiplayer (vẫn có nút trong bảng điều khiển).
- Camera thực địa: dùng grid hàng rõ ràng, cố định thanh thao tác nhanh và camera sát đáy, giữ khoảng trống giữa màn CCTV; chia màn giữ bản đồ trong vùng riêng, bỏ dòng trạng thái trùng, giảm kích thước thông báo tạm dừng.
- Hiệu năng: bỏ render canvas khi đang xem bản đồ lớn; giới hạn render nền/tạm dừng 8 FPS; DPR tự động tối đa 1,25 trên cảm ứng, 1,5 trên PC ở mức cao nhất của Auto, giảm theo tải và giới hạn pixel canvas mục tiêu 1,6 triệu; vẫn cho phép chọn chất lượng thủ công. Không đổi tốc độ mô phỏng. Điều tiết HUD thực địa 125 ms (350 ms khi tạm dừng), giảm cập nhật UI từ snapshot thường, giữ cập nhật cửa/pha ngay; giảm tính lặp tầm nhìn và không thay SVG khi đang chạm. Giữ nút hành động phản hồi ngay.
- Multiplayer: tăng ngưỡng heartbeat khỏe từ 5,5 lên 8 giây và dùng pong cập nhật thời điểm kết nối, giữ ngưỡng ngắt 12 giây; bỏ snapshot trạng thái khi hàng đợi P2P quá lớn, không bỏ lệnh điều khiển/kết quả; loại route cũ khi tạo snapshot. Khi người chơi thực sự ở nền vẫn tạm dừng để giữ đồng bộ, không hứa chạy nền iOS.
- Sửa lỗi biến s chưa khai báo trong updateFoxy ở nhánh va cửa và đi thực địa; chọn camera khi đang tắt gửi một lệnh bật đúng camera thay vì hai lệnh. Giữ thông tin nhập email khi UI cập nhật, giới hạn chờ tải SDK 15 giây, xử lý lỗi callback, xóa trạng thái recovery khi đổi người dùng, không vẽ lại form theo token refresh.
- Tài khoản: ưu tiên đăng ký/đăng nhập chung bằng link email không mật khẩu (tự tạo tài khoản khi lần đầu), OAuth một nút, mật khẩu tùy chọn thu gọn; chờ 60 giây giữa lần yêu cầu link thành công. Cấu hình Supabase vẫn bắt buộc, chưa kích hoạt; ACCOUNT_SETUP.md rút gọn còn ba bước, email trước và OAuth tùy chọn. Giữ nguyên bảo mật RLS và xác nhận lưu/tải.
- Kiểm tra cú pháp JavaScript toàn bộ đạt; 27 kiểm tra logic (DPR, backpressure, heartbeat, snapshot, khóa chọn chữ) và 2 kiểm tra nhánh Foxy đạt trên các commit. Chưa thử trình duyệt thật, đo FPS thực tế, P2P đa thiết bị, OAuth hay xác nhận triển khai Pages. Cần tất cả người chơi P2P dùng V56.0. Không sửa server.go vì không thay lệnh relay. Cập nhật nhật ký trong game, README và CHANGELOG.

## V55.0 — Cloud Accounts & Tactical Camera Overhaul (26/11/2026)

- Menu chính có Tài khoản & lưu online, Nhật ký cập nhật trong game (V51–V55, link lịch sử đầy đủ). Chơi khách và lưu offline vẫn độc lập; không yêu cầu đăng nhập để chơi.
- Tích hợp Supabase Auth cho email/mật khẩu, xác nhận email, quên/đổi mật khẩu, Google, Discord, Apple và GitHub. Dùng PKCE và SDK chỉ tải khi mở tài khoản hoặc nhận callback. **Chưa kích hoạt dịch vụ**: cần project Supabase, publishable key, OAuth provider credentials và redirect URL do quản trị viên thiết lập. Không có tài khoản giả, không lưu mật khẩu hoặc khóa bí mật trong repo.
- Giao diện kiểm tra/lưu/tải cloud theo tài khoản; xác nhận ghi đè, lưu phiên bản và kiểm tra revision để tránh hai thiết bị ghi đồng thời. Lưu đêm đã mở, xu, vật phẩm, nâng cấp, thành tích, hồ sơ, nhiệm vụ và tối đa 12 trận; không lưu ca đang dở, cấu hình máy, chat/voice. Không tự gộp/tự đồng bộ; đăng xuất giữ tiến trình offline, cần chú ý thiết bị dùng chung. Tải cloud có sao lưu cục bộ trước khi thay tiến trình.
- Thêm supabase/cloud_saves.sql với RLS chỉ cho chủ tài khoản đọc/ghi, giới hạn payload và trigger revision; ACCOUNT_SETUP.md hướng dẫn thiết lập và kiểm thử. Chưa chạy SQL trên dịch vụ thật hoặc thử OAuth.
- Camera thực địa/Tiếp tế khẩn cấp: bản đồ, CCTV và chia màn hình (CCTV với bản đồ nhỏ), thanh cảnh báo tín hiệu/quái/điện/giờ, thanh thao tác nhanh báo quái, xung cứu viện, tập hợp, cửa trái/phải. Nhóm nút Tất cả/Điều phối/Hỗ trợ/Cửa; giao diện ngang gọn hơn. Không đổi cân bằng kỹ năng V54.
- Sửa chọn camera theo quái khi chạm marker, giữ nút điều khiển/camera ổn định khi cooldown và trạng thái quái cập nhật để tránh mất thao tác, xử lý trạng thái account khi chuyển màn, kiểm tra session trước khi cập nhật trạng thái lưu, phục hồi tiến trình và lịch sử cục bộ nếu tải cloud ghi thất bại.
- Kiểm tra cú pháp toàn bộ JavaScript đạt trên hai commit; 23 kiểm tra logic trích hàm trên bản cuối đạt (lọc tiến trình, giới hạn, schema, lọc lịch sử, phân nhóm kỹ năng). Chưa kiểm thử giao diện thực tế, mạng P2P, SQL/RLS hoặc đăng nhập thật. GitHub Pages chưa xác nhận triển khai; P2P cần cùng V55.0. Không thay server.go vì không thêm lệnh mạng, account tách riêng Supabase.

## V54.0 — Field Command Center & Emergency Response (26/11/2026)

- Thiết kế trung tâm điều phối mới cho người xem camera trong Thực địa và Tiếp tế khẩn cấp: bản đồ chiến thuật lớn trên sân chơi, chuyển giữa bản đồ và CCTV, thanh chọn nhanh camera, hiển thị tín hiệu/mất kết nối, công tắc bật/tắt camera để tiết kiệm điện. Hai công cụ tương tác bản đồ riêng: chỉ đường và chọn camera theo phòng; tránh đặt điểm ngoài ý muốn. Bố cục PC và điện thoại ngang, có điều chỉnh cho màn hình dọc.
- Bảng theo dõi đội: từng bảo vệ có vị trí phòng, trạng thái an toàn/cảnh giác/truy đuổi, giáp, khói, trạng thái di chuyển hoặc tiến độ thao tác, mốc đường đã tới. Danh sách ba nhiệm vụ và mốc giờ mở; nhật ký điều phối 8 sự kiện gần nhất. Camera hiển thị tiến độ của bảo vệ đang thao tác, kể cả BV2.
- Theo dõi tuyến chỉ đường riêng cho BV1/BV2; tự ghi nhận khi tới cách mốc tối đa 35 đơn vị. Nút Đi mốc tiếp để tự tìm đường theo tuyến đã nhận; bảo vệ tự quyết định di chuyển. Đổi tuyến đặt lại tiến độ. HUD hiện mốc kế tiếp và khoảng cách.
- SOS/Báo an toàn: từng bảo vệ có hồi chung 10 giây, tín hiệu vị trí phòng tồn tại 15 giây, camera xác nhận riêng từng người. SOS đánh dấu trên bản đồ và phát âm thanh cho người xem camera; nút SOS/Mốc tiếp ngay trên cảnh chơi. Không ép di chuyển hoặc cấp vật phẩm.
- Tập hợp: chỉ đường về phòng trực, hiệu lực 45 giây, hồi 25 giây; đồng bộ cho cả hai bảo vệ. Xung cứu viện: làm choáng một quái trong CAM đang chọn, cách ít nhất một bảo vệ không quá 180 đơn vị, trong 2,5 giây; tiêu hao 3% điện cơ sở theo thời lượng (tối thiểu 1,2%), hồi 42 giây nhân hệ số thời lượng. Không áp dụng khi camera tắt/nhiễu, thiếu điện, quái đang choáng/rút lui hoặc ngoài phạm vi. Giữ hệ thống cảnh báo và hỗ trợ thao tác V53.
- Sửa lỗi snapshot lấy nhầm vị trí chính khi BV2 đang thực hiện hành động gây gửi trạng thái ngay; chặn vẽ HUD giữa lúc hoán đổi trạng thái BV2. Đồng bộ hướng khi tự tìm đường, trạng thái AI và thời gian choáng của quái; giới hạn di chuyển thực địa trong các phòng hợp lệ; giữ mục tiêu chạm bản đồ khi SVG cập nhật và hủy thao tác nếu mất pointer capture. Tương thích cấu trúc checkpoint cũ của hệ thống điều phối.
- Relay hai người: thêm whitelist SOS, xác nhận, đi mốc, tập hợp, xung cứu viện; phải triển khai lại server.go để sử dụng. Tiếp tế khẩn cấp vẫn P2P; các máy P2P cần cùng V54.0. Giữ cân bằng điện và Foxy Rush V52.0.
- Kiểm tra cú pháp toàn bộ JavaScript đạt ở hai commit; 27 kiểm tra logic trích hàm trong môi trường giả lập trên blob cuối đạt, gồm quyền/cooldown SOS, xác nhận hai người, tuyến/mốc, tập hợp, điều kiện xung cứu viện, giới hạn nhật ký, lọc snapshot và lỗi hoán đổi BV2. Chưa chạy thử trình duyệt, gameplay đa thiết bị, biên dịch Go hoặc xác nhận triển khai GitHub Pages. Không khẳng định mọi lỗi đã hết.

## V53.0 — Field Command & Tactical Intelligence (26/11/2026)

- Thực địa và Tiếp tế khẩn cấp: người xem camera chạm phòng trên bản đồ để đặt tuyến tối đa 3 điểm theo thứ tự, có dấu xanh đánh số trên bản đồ và cảnh chính của cả hai bảo vệ. Dấu tồn tại 35 giây, hồi đặt/xóa 1,5 giây; nút chỉ mục tiêu đặt nhiệm vụ hiện tại hoặc phòng trực. Bảo vệ tự quyết định đường di chuyển, có xác nhận riêng BV1/BV2, không bị điều khiển từ xa.
- Cảnh báo quái: chạm quái đang hiện trên camera/bản đồ hoặc bấm Báo quái để chọn quái nhìn thấy gần đội nhất. Host xác thực camera đang hoạt động, không bị nhiễu, chỉ chấp nhận quái trong camera được chọn hoặc trong đợt quét. Đánh dấu vị trí ghi nhận (không bám theo quái) trong 9 giây, hồi chung 14 giây, đồng bộ HUD và âm thanh cho bảo vệ.
- Hỗ trợ thao tác từ xa: tăng tốc độ thực hiện nhiệm vụ 40% trong 6 giây khi có ít nhất một bảo vệ đang thao tác; tiêu hao 1,5% điện cơ sở theo thời lượng (tối thiểu 0,65%), hồi 28 giây nhân hệ số thời lượng. Không bỏ qua yêu cầu vị trí, mốc giờ hoặc nhiệm vụ.
- Giao diện: bảng điều phối hiển thị tuyến, thời gian còn lại, xác nhận từng bảo vệ, cảnh báo và hỗ trợ; HUD cảnh chính hiển thị thông tin cô đọng. Màu xanh cho chỉ đường, cam cho cảnh báo, xanh lá cho hỗ trợ. Sửa hướng dẫn phím thực địa và thao tác chạm trên bản đồ SVG thường xuyên cập nhật; không tính vuốt dài là chọn mục tiêu.
- Sửa lỗi Thực địa/Tiếp tế: mồi âm thanh không được đặt sát bất kỳ bảo vệ nào và đồng bộ mồi cho cả hai; đi thẳng nếu đường không bị chắn (tránh lỗi đường rỗng cùng ô điều hướng); vô hiệu nút thao tác khi đang di chuyển; ẩn HUD/joystick ở màn kết quả. Giữ cơ chế điện và nerf Foxy V52.0.
- Snapshot chỉ đường/cảnh báo/hồi chiêu được chuẩn hóa khi nhận, đặt lại ở ca mới. Phân quyền host và thời gian hồi chung giữa các camera. Relay hai người có whitelist lệnh mới trong server.go; cần triển khai lại relay. Tiếp tế khẩn cấp vẫn P2P.
- Kiểm tra cú pháp JavaScript toàn bộ đạt ở cả hai commit; 25 kiểm tra logic trích hàm trong môi trường giả lập đạt: phân quyền, giới hạn tuyến và thời gian hồi, xác nhận hai bảo vệ, cảnh báo có xác thực tầm nhìn, giữ nguyên vị trí ghi nhận, hỗ trợ và lọc snapshot. Chưa thử trình duyệt/thiết bị thật, chưa biên dịch Go hay xác nhận triển khai GitHub Pages; P2P yêu cầu cùng V53.0.

## V52.0 — Direct Field Control, Anti-Camp & Emergency Supply (26/11/2026)

- Thực địa: WASD/phím mũi tên và joystick cảm ứng 8 hướng; tốc độ 145 đơn vị/s, chuẩn hóa đường chéo, dùng va chạm hiện hữu. Lệnh điều khiển làm hủy đường tự động và thao tác đang thực hiện, tự hết hạn sau 0,5 giây; gửi lại mỗi 125ms khi giữ. Dừng khi thả, mất tiêu điểm hoặc tạm dừng mạng. F/Space thao tác, X dùng khói, R về phòng. Bổ sung whitelist relay hai người cho điều khiển bảo vệ và túi tiếp tế. Relay cần triển khai lại server.go; chưa xác nhận đã triển khai.
- Thực địa: HUD khu vực, số giáp/khói, cảnh giác, nhiệm vụ; joystick và bốn nút nhanh trên cảnh chính; bảng nhiệm vụ mới, viền trạng thái, thanh bên thu gọn. Desktop tối đa 1280px, thanh bên 245–290px trên màn hình đủ lớn; có bố cục riêng cho điện thoại ngang. Đồng bộ hướng, trạng thái di chuyển, giáp và khói; ngăn thao tác nhiệm vụ tiếp tục nhầm mục tiêu khi đồng đội hoàn thành trước.
- Mode multiplayer P2P mới **Tiếp tế khẩn cấp**: 2–4 người, 75% điện xuất phát, 1 khói/bảo vệ; ba nhiệm vụ lần lượt ở vệ sinh, sân khấu, Pirate Cove, mở tại 12 AM, 1 AM, 3 AM, thưởng tối đa 6/7/9% điện; về phòng trước 6 AM. Dùng chung cơ chế một hoặc hai bảo vệ và camera điều phối. Không có mode này trên relay.
- Điện: tăng hệ số tiêu hao multiplayer Co-op 0,93→1,04; Versus 0,90→1,01; Thực địa 0,90→1,02, trước hệ số thời lượng; giảm tiếp tế giờ và bonus thời lượng ngoài Foxy Rush. Solo tăng tiêu hao cơ sở 5%, không thay đổi cheat noDrain. Giữ nguyên công thức điện và tiếp tế giờ Foxy Rush; mức sát thương điện từ Foxy giảm theo nerf kỹ năng.
- Anti-camp: trong bán kính 85 đơn vị quanh cửa, quái người chơi tích thời gian khi cửa đóng hoặc đứng chờ/di chuyển khi cửa mở; buộc rút sau 5,5s (Foxy Rush 3,8s), không reset chỉ vì đổi lệnh. Tấn công thực sự qua cửa mở vẫn hợp lệ. AI có giới hạn đứng gần cửa đóng. Phe quái Versus/Foxy Rush phải quay tới gần điểm xuất phát và chờ hồi trước khi ra lệnh tiếp.
- Foxy Rush: giảm hệ số tốc độ 0,94→0,84 (phòng 2 người), sát thương 0,875→0,75; giảm tốc độ/thời lượng tăng tốc và sát thương điện; tăng hồi chiêu, thời gian rút lui và cửa sổ phản ứng bảo vệ, áp dụng hệ số riêng cho phòng 3–4 người.
- Kiểm tra cú pháp toàn bộ JavaScript đạt trước commit 746a48e5037a26acefc0ef7479ce1776e8636d4d; đọc lại blob và chạy 68 kiểm tra logic trích hàm với môi trường giả lập: đạt (điện theo 5 thời lượng, giữ nguyên công thức điện Foxy Rush, hệ số nerf, di chuyển và hết hạn, phân quyền, mode tiếp tế, anti-camp). Chưa chạy thử gameplay P2P/relay, chưa biên dịch Go, chưa kiểm tra thực tế Safari/iPhone và các kích thước màn hình. P2P yêu cầu cùng V52.0.

## V51.0 — Field Patrol Duo & Compact Desktop (26/11/2026)

- Thực địa P2P: phòng 2 người có 1 bảo vệ và 1 camera; phòng 3–4 người tự phân 2 bảo vệ và 1–2 camera. Bảo vệ thứ hai có vị trí, tuyến đường, thao tác nhiệm vụ, trạng thái nguy hiểm, 1 khói và 1 giáp riêng. Nhiệm vụ, điện, cửa và tiếp tế phụ dùng chung. Chủ phòng xử lý lệnh theo seat của kết nối; quái chọn bảo vệ gần hơn. Cả hai bảo vệ phải về phòng trực để thắng.
- Tái sử dụng cảnh khám phá tự do của chế độ thường cho góc nhìn bảo vệ thực địa: nhân vật, vật cản, ánh sáng, tuyến đi, đồng đội và mục tiêu. Chạm vào phòng trên cảnh chính hoặc bản đồ để đi; điều phối viên giữ góc camera. Không bổ sung điều khiển di chuyển tự do WASD trong bản này.
- Chuyển bảng bản đồ/nhiệm vụ khỏi lớp phủ cảnh sang thanh điều khiển; thu gọn nút trùng, bản đồ tỷ lệ cố định, giới hạn chiều rộng desktop 1440px và thanh bên 260–320px trên màn hình đủ lớn. Giữ khả năng cuộn thanh bên khi chơi ngang trên điện thoại.
- Sửa kiểm tra vùng cửa để tránh đóng cửa khi một trong hai bảo vệ ở ngưỡng cửa; khói và tiến độ đường đi của bảo vệ thứ hai được đưa vào snapshot.
- Đã đọc lại index.html từ GitHub sau commit 9b4a7bf91fcf45704b6b41c67f3582b40a67db91 và kiểm tra cú pháp JavaScript toàn bộ script: đạt. 19 kiểm tra logic với các hàm trích từ mã đã commit và môi trường giả lập: đạt (phân vai phòng 2–4 người, vị trí riêng, chia sẻ nhiệm vụ/tiếp tế, snapshot, khôi phục ngữ cảnh khi lỗi, điều kiện thắng, kiểm tra ngưỡng cửa, relay). Chưa chạy thử gameplay multiplayer, WebRTC hoặc đo bố cục trên thiết bị thật; không xác nhận đã sửa toàn bộ lỗi giao diện. Phân đội tự động theo số người, chưa có tùy chọn 1/2 độc lập ở lobby. P2P dùng mã phiên bản V51.0; relay vẫn hai người.

## V50.0 — Command Center & Tactical Expansion (26/11/2026)

- Xác nhận lệnh cửa riêng theo người gửi, mã lệnh và ca trực; bỏ lệnh trùng, giới hạn thao tác quá nhanh; hỗ trợ cửa thường và cửa Thực địa P2P. Không dùng thông báo cửa của người khác để xác nhận lệnh đang chờ. Sửa nút mở cửa bị vô hiệu hóa khi cửa hỏng nhưng còn đóng; giảm âm thanh cửa lặp.
- Phát hiện kênh P2P còn báo mở nhưng không truyền heartbeat trong 12 giây khi hai phía ở foreground, thử tạo lại kênh bằng hệ thống giữ chỗ 30 giây. Không tự tạo nhiều kết nối khôi phục cùng lúc. Không đảm bảo khôi phục nếu đóng tab hoặc mất chủ phòng.
- Bảng TÍN HIỆU ĐỘI thu gọn cho Co-op và Thực địa P2P: trái, phải, an toàn, tiết kiệm điện, sửa hệ thống, về phòng; hiện P1–P4 gửi tín hiệu, hạn chế spam. Relay giữ ba tín hiệu cũ.
- Ba kỹ năng chiến thuật multiplayer: khởi động cầu chì (2,5% điện cơ bản, hồi 36s), phục hồi camera (3%, hồi 34s), tích áp +12 AP khi quái không tấn công/điều khiển tay/rút lui và áp lực không quá 55 (hồi 40s). Chi phí/hồi chiêu áp dụng hệ số thời lượng; quyền quái và cooldown vẫn kiểm tra ở host.
- Thực địa: túi tiếp tế phụ một lần tại khu vệ sinh từ 3 AM, tối đa +4% điện và +1 khói (tối đa 2 khói), cần đứng gần, không đang di chuyển/làm nhiệm vụ. Không tăng số nhiệm vụ bắt buộc. Chỉnh thống kê pin khẩn cấp theo lượng thực nhận.
- Giải phóng track đầu ra audio thay thế khi đổi đường phát. Giữ bố cục ngang, voice 4 người, tutorial đêm 1, cân bằng pin/Foxy Rush trước đó.
- Kiểm tra JavaScript cuối đạt; 38 kiểm tra mô phỏng P2P 4 người/biên heartbeat và mã lệnh, 23 kiểm tra logic kỹ năng/túi tiếp tế đạt trước lần tinh chỉnh cuối về chống lặp và điều kiện quái. Chưa thử trên thiết bị thật. P2P yêu cầu đồng phiên bản V50.0; relay vẫn hai người.

## V49.0 — Reconnect & Squad Intelligence (26/11/2026)

- Tự nối lại kênh dữ liệu P2P của khách trong cùng tab tối đa 30 giây; giữ nguyên vị trí, vai trò và mã ca, tạm dừng cả đội khi đang chơi. Xác thực yêu cầu nối lại bằng Peer ID và token ngẫu nhiên 128 bit riêng của chỗ; không lưu token vào bộ nhớ lâu dài. Nếu hết hạn, kết thúc ca và thông báo cho các thành viên còn lại trước khi đóng kết nối. Không hỗ trợ khôi phục sau khi đóng/tải lại tab hoặc mất chủ phòng.
- Xác nhận sẵn sàng theo phiên bản cấu hình lobby, loại bỏ xác nhận cũ sau khi đổi thời lượng/độ khó/đội hình. Sửa lỗi thông báo kết thúc ca bị mất do đóng kênh quá sớm; lưu kết quả ca hiện tại để gửi lại cho khách vừa khôi phục.
- Co-op, Thực địa và Foxy Rush P2P 3–4 người: quản lý trạng thái và lựa chọn camera theo từng khách, giữ camera chung hoạt động nếu còn người xem, xử lý kỹ năng phụ thuộc camera theo lựa chọn của người gửi. Chia sẻ điện, cửa và hồi chiêu theo luật hiện tại.
- Giữ đường đi và tiến độ thao tác thực địa khi tạm dừng mạng P2P; không tiến hành nhiệm vụ trong thời gian tạm dừng. Giữ hệ thống điện, cân bằng Foxy Rush và trợ giúp đêm 1 đã có.
- HUD đội hình nhỏ hiển thị P1–P4 và chỗ đang khôi phục; cập nhật thông báo trong lobby, cài đặt, và trận đấu cho nhóm nhiều người. Sửa cập nhật trạng thái phát voice từ kết nối cũ và thử phát lại audio bị tạm dừng.
- Kiểm tra cú pháp toàn bộ JavaScript đạt. 27 kiểm tra tích hợp giả lập 4 người đạt cho phiên bản trước tinh chỉnh HUD/kết quả/voice: xác nhận cấu hình, camera độc lập, ngắt kênh, tạm dừng, nối lại đúng vị trí, đồng bộ sau nối lại, hết hạn và xác thực token. Cú pháp bản cuối đạt. Chưa thử WebRTC, mic, đồ họa và gameplay trên thiết bị thật. P2P yêu cầu cả nhóm V49.0; relay vẫn hai người, không áp dụng cơ chế giữ chỗ mới.

## V48.0 — Party & Tactical Overhaul (26/11/2026)

- Phòng P2P 2–4 người có xác nhận sẵn sàng từng thành viên; chủ phòng chỉ bắt đầu khi mọi khách đã xác nhận. Đổi chế độ, độ khó, thời lượng, sức chứa hoặc đội hình sẽ yêu cầu xác nhận lại. Hiển thị vai trò dự kiến của phòng 3–4 người và thêm preset Chiến thuật 9 phút.
- Tách gói dữ liệu theo mã ca và số thứ tự, loại bỏ gói trận cũ, trùng hoặc sai cấu trúc; kiểm tra trạng thái điện, thời gian, cửa, hiệu ứng và vị trí quái. Đồng bộ chơi lại và nhận diện P1 trong chat. Kiểm tra kết nối còn mở khi đánh giá sức khỏe phòng.
- Voice nhóm có nút tắt nghe từng thành viên, giữ mute khi nối lại voice, xóa mute khi thành viên rời phòng; kiểm tra mức mic qua WebRTC stats nếu trình duyệt cung cấp. Khôi phục Audio Session khi quay lại tab; bổ sung nút khôi phục nhạc/SFX và đổi đầu ra ngay trong cài đặt phòng.
- Thực địa: thu gọn/ẩn bản đồ cho cả bảo vệ và camera, giữ đường đi cũ nếu yêu cầu đường mới thất bại, tránh hủy đường do chạm lặp, thử tìm đường lại tối đa một lần nếu kẹt. Điều chỉnh phí điện quét/sửa/dụ quái theo thời lượng và đồng bộ điều kiện nút.
- Tối ưu HUD ngang, đồng hồ phút:giây và lượng điện tiếp tế mốc kế tiếp, bảng voice nhỏ gọn, lịch sử 12 trận có tổng thắng/thua và tỷ lệ thắng cục bộ. Giữ tutorial đêm 1 và cân bằng Foxy Rush.
- Kiểm tra cú pháp toàn bộ JavaScript đạt; 91 kiểm tra hàm readiness/validation và 13 kiểm tra vòng đời P2P giả lập 4 người đạt. Chưa kiểm thử giao diện, mic/loa thật, kết nối qua NAT hoặc cân bằng thực chiến; không khẳng định đã hết mọi lỗi. Giao thức P2P V48.0 mới yêu cầu cùng bản; relay vẫn 2 người và không áp dụng ready-check mới.

## V47.5 — Night 1 Tutorial & Landscape UI (26/11/2026)

- Chỉ hiển thị hỗ trợ tự động tên quái, phía cửa, đếm ngược tấn công và nhãn nhận diện CCTV trong Đêm 1 cốt truyện (Baby/Easy/Normal). Từ Đêm 2, Custom, Challenge, Boss và multiplayer không có trợ giúp tự động này; giữ cảm biến/radar được người chơi chủ động kích hoạt, âm thanh và hình ảnh thật của quái.
- Bỏ hộp thoại chọn hỗ trợ lúc 2 AM gây gián đoạn trận; chuẩn hóa lựa chọn hỗ trợ từ dữ liệu lưu cũ.
- Thu gọn thông báo, nhãn camera, bảng tên đêm 1 và hộp thoại; thông báo thông thường tự ẩn trong tối đa 2,5 giây, không chặn thao tác. Nội dung cuộc gọi đầy đủ vẫn có trong nhật ký.
- Bố cục landscape ưu tiên sân chơi, đặt HUD multiplayer bên phải có cuộn riêng, thanh điện/thời gian ở trên; tối ưu các nút cửa, camera, menu, chat và khu vực an toàn trên màn hình thấp. Cập nhật canvas khi xoay màn hình/đổi kích thước viewport; không bắt buộc khóa hướng màn hình.
- Kiểm tra cú pháp JavaScript và 450 tổ hợp logic tutorial theo đêm, chế độ, độ khó, online/offline đạt. Chưa kiểm thử trực quan trên điện thoại/PC thật. Giữ cân bằng điện V47.4, voice nhóm và giao thức P2P V47.4.

## V47.4 — Mobile Voice Audio & Power Balance (26/11/2026)

- Khắc phục xung đột giữa micro và âm thanh game: yêu cầu chế độ Audio Session play-and-record khi thu mic và playback khi ngừng thu trên trình duyệt hỗ trợ; kiểm tra lại đầu ra sau khi cấp quyền, đổi thiết bị hoặc ngắt mic. Không tự giảm nhạc 35% khi voice kết nối.
- Dùng bus âm thanh game riêng, giữ voice độc lập. Thêm nút ĐỔI ĐẦU RA GAME trong chat voice và cài đặt: có thể chuyển Web Audio trực tiếp sang MediaStream audio thay thế; nếu bị chặn, lỗi hoặc dừng phát, tự quay lại đầu ra trực tiếp. Nút KHÔI PHỤC NHẠC/SFX vẫn dùng được trong phòng.
- Cân bằng điện cho multiplayer 3/6/9/12/15 phút, giữ 100% lúc bắt đầu. Hệ số tiêu hao thời gian 180/d; giảm nhẹ hao điện nền theo thời lượng; tiếp tế 5 mốc tăng dần, tổng tiềm năng Foxy Rush 22,5/26,25/30/33,75/37,5%, Versus và Co-op 20/23,75/27,5/31,25/35%, Thực địa 25/28,75/32,5/36,25/40%. Chỉ nạp đến 100% và không cấp khi đã hết điện.
- Giữ áp lực đóng hai cửa liên tục bằng mức hao phụ cửa; giảm chi phí tối thiểu của đuổi quái, gia cố, sửa nhanh, radar hỗ trợ trong trận dài. Hell multiplayer giảm hệ số hao điện 2,2 xuống 1,75; không thay đổi chơi đơn. Giữ hệ số Foxy Rush V47.2.
- Kiểm tra cú pháp JavaScript đạt; 16 kiểm tra giả lập định tuyến Audio Session/đầu ra, lỗi autoplay, timeout và mute đạt; 30 kiểm tra ngân sách điện danh nghĩa (4 chế độ × 5 thời lượng và áp lực hai cửa ở 2 chế độ đối đầu) đạt. Chưa kiểm tra âm thanh thật trên iPhone/Android/PC hay cân bằng qua trận thật. Trình duyệt có thể cần chạm nút khôi phục; không bảo đảm ép loa ngoài hoặc hoạt động trong mọi trình duyệt. P2P dùng mã phòng V47.4 mới, relay vẫn hai người.

## V47.3 — Audio Recovery (26/11/2026)

- Khởi tạo Web Audio từ thao tác chạm/bấm/phím của người chơi; phục hồi AudioContext bị suspend/interrupted, đóng hoặc đồng hồ âm thanh bị treo. Không khởi động lại nhạc liên tục khi bấm nhiều lần; loại bỏ kết quả resume thuộc context đã thay thế.
- Khôi phục đồ thị nhạc nền, hiệu ứng, Foley và âm thanh môi trường; xếp hàng ngắn hiệu ứng khi audio đang mở khóa, tách lỗi bộ lập lịch và phục hồi khi trở lại tab.
- Nút KHÔI PHỤC ÂM THANH xuất hiện khi âm thanh bị chặn/treo và luôn có trong cài đặt audio. Khi chủ động bấm nút, bật master sound, khôi phục mức nhạc/hiệu ứng mặc định nếu đã về 0 và bật chế độ nhạc thích ứng nếu đã tắt. Không tự ghi đè lựa chọn tắt âm thông thường.
- Sửa xung đột phím U với thao tác mở khóa; đồng bộ thanh âm lượng, giữ tùy chọn audio khi vào/rời multiplayer, áp dụng âm lượng voice cho tất cả luồng nhóm.
- Kiểm tra cú pháp JavaScript đạt; 24 kiểm tra giả lập AudioContext đạt trước tinh chỉnh giao diện/phím tắt, cú pháp cuối đạt. Chưa thử đầu ra âm thanh thật trên Safari/iPhone, Android hoặc PC. Trình duyệt có thể yêu cầu thao tác của người dùng sau khi chuyển nền. Không thay đổi cân bằng gameplay hoặc giao thức P2P V47.2.

## V47.2 — Foxy Rush Balance & Stability (26/11/2026)

- Foxy Rush tự cân bằng theo số người trong phòng (2/3/4): hệ số tốc độ 0,94/1,00/1,06; sát thương cửa cơ bản 14/15,4/16,8 điểm (trước đây 16). Giữ tăng áp lực theo giờ và thời lượng.
- Siêu tốc x1,75/x1,875/x2 trong 4/4,5/5 giây, hồi cơ bản 10/9,5/9 giây; kỹ năng phá rối hồi 15/14/13 giây, nhiễu camera 4 giây, rút 2/2,25/2,5% điện cơ bản gần cửa. Hao điện mỗi cú đập 0,65/0,75/0,85% cơ bản. Chi phí điện và hồi chiêu vẫn chịu hệ số thời lượng/giai đoạn.
- Bảo vệ 2 người được tăng nhẹ thời gian phản ứng và tốc độ sửa cửa; hiệu chỉnh theo số người để Foxy không quá yếu trước nhiều người dùng chung bàn trực. Không đổi cân bằng Versus thường hoặc chơi đơn.
- Sửa điều kiện bắt đầu khi kết nối đang mở hoặc người chơi ở nền; đồng bộ tạm dừng nhanh hơn, ngăn lệnh từ kết nối cũ và đặt lại biểu quyết chơi lại nếu nhóm chưa sẵn sàng.
- Hiển thị đúng quyền chọn quái và chỉ huy chiến thuật; sửa khôi phục voice khi một đầu mất kết nối và lỗi hủy cấp quyền micro.
- Kiểm tra cú pháp JavaScript đạt, 60 kiểm tra hệ số cân bằng/không ảnh hưởng Versus và 35 kiểm tra voice giả lập đạt. Chưa thử trận thật, âm thanh thật trên bốn máy hoặc xác nhận cân bằng bằng tỷ lệ thắng. P2P cần cùng V47.2; relay vẫn 2 người.

## V47.1 Party Voice — Preview (26/11/2026)

### Tính năng mới
- Voice P2P mở rộng từ 2 lên 2–4 người bằng WebRTC mesh: mỗi máy dùng một luồng micro được người dùng cấp quyền và tối đa ba kết nối âm thanh riêng. Phòng 4 người có sáu cặp kết nối.
- Host định tuyến tín hiệu voice theo vị trí người gửi/nhận, xác thực vị trí và giới hạn SDP, không phát tín hiệu đàm phán của một cặp sang các thành viên khác. Âm thanh truyền qua WebRTC giữa các cặp, không qua kênh dữ liệu gameplay.
- Hiển thị trạng thái kết nối voice từng người, chỉnh âm lượng chung, tắt tiếng micro, nút bật nghe khi autoplay bị chặn và hỗ trợ cấu hình STUN/TURN. Theo dõi ICE và thử kết nối lại khi kết nối lỗi.

### Sửa lỗi
- Không còn tự ngắt voice khi thành viên thứ ba hoặc thứ tư vào phòng. Dọn kết nối âm thanh của thành viên rời phòng, bảo toàn các kết nối còn lại trong lobby.
- Heartbeat của host báo trạng thái toàn nhóm; khách cũng tạm dừng gameplay khi một thành viên ở nền hoặc mất tín hiệu. Foxy Rush cho các bảo vệ chia sẻ bàn trực sử dụng các lệnh hợp lệ, bỏ giới hạn trái/phải không nhất quán.
- Mã phòng P2P V47.1 riêng để tránh trộn giao thức voice V47.0. Relay giữ voice 2 người và giao thức cũ.
- Cập nhật nhãn phiên bản và hướng dẫn trong lobby.

### Kiểm thử và giới hạn
- Kiểm tra cú pháp toàn bộ JavaScript: đạt. Mô phỏng bốn máy với SDP/ICE giả lập: 29 kiểm tra đạt, gồm 6 cặp kết nối/12 đầu WebRTC, ngắt voice, dọn thành viên rời phòng và dọn toàn bộ kết nối. Đây không phải thử nghiệm âm thanh thực tế.
- Chưa kiểm thử WebRTC thật, âm thanh thực, Safari/iPhone, NAT nghiêm ngặt hoặc TURN thực trên bốn thiết bị. Micro cần HTTPS hoặc localhost và quyền truy cập; trình duyệt, mạng và hệ điều hành có thể chặn voice hoặc ngắt khi chạy nền. Chưa merge main/triển khai Pages. Không xác nhận đã sửa hết lỗi toàn bộ game.

## V47.0 Party — Preview (26/11/2026)

- P2P: mở rộng phòng từ 2 lên 2–4 người; chủ phòng chọn sức chứa, lobby hiển thị từng vị trí, trạng thái kết nối và số người thực tế. Mã phòng V47 riêng, không ghép với V46.
- Versus 3–4 người: chủ phòng là bảo vệ; các khách được chia quyền điều khiển Bonnie, Chica, Freddy và Foxy. Kiểm tra quyền gửi lệnh trên host, hạn chế spam lệnh và tránh AI tự điều động quái đã phân công.
- Co-op/Thực địa: một bảo vệ và nhiều người hỗ trợ camera chung. Foxy Rush: một Foxy và nhiều người chia sẻ bàn trực, không phải nhiều nhân vật bảo vệ độc lập.
- Thêm bốn preset ca trực: Khẩn cấp 3 phút, Tiêu chuẩn 6 phút, Cầm cự 12 phút, Ác mộng 15 phút; giữ chọn thời lượng/độ khó riêng.
- Chat chữ chuyển tiếp trong phòng, nhãn người gửi, chơi lại khi tất cả thành viên đồng ý, kiểm tra phòng đầy và từ chối vào giữa trận.
- Host duy trì trạng thái trận; heartbeat theo từng kết nối, tạm dừng khi thành viên mất tín hiệu/ở nền. Nếu kênh P2P thực sự đóng giữa trận thì kết thúc phòng.
- Giới hạn: voice và relay vẫn chỉ hỗ trợ hai người; phòng nhóm dùng chat chữ. Chưa có voice nhóm, nhân vật bảo vệ độc lập, hay vào lại trận đã mất kênh.
- Kiểm tra: bốn tổ hợp phân quyền vị trí (2, 3, 4 người và vị trí không liên tục) được kiểm tra logic; chưa chạy trình duyệt, kiểm thử WebRTC thật trên 4 thiết bị hoặc xác nhận cân bằng. Nhánh preview `v47-party`, chưa merge main hoặc triển khai Pages.

## V41.9 — Online Door Control Hotfix

- Đóng/mở cửa online qua đường xử lý riêng trên host, tự hạ camera nếu cần; không bỏ lệnh vì camera chưa đồng bộ.
- Khách gửi lệnh rồi đợi trạng thái xác nhận từ host; chống gửi trùng trong lúc chờ, có báo thất bại và quá hạn. Nút nhanh cửa trái/phải khi đang mở camera.
- Trong đối đầu, quái chồng ngưỡng cửa nhưng tâm vẫn ngoài hành lang được đưa lùi về điểm tiếp cận hợp lệ để cửa đóng; không đẩy quái đã qua ngưỡng về ngoài. Giữ nguyên luật hỏng cửa, hết điện, khóa khẩn.
- Giao thức 43 giữ nguyên; kiểm tra cú pháp JavaScript và 12 tình huống mô phỏng với đối tượng giả lập đạt. Chưa kiểm thử trực tiếp trên hai thiết bị hoặc xác nhận Pages đã triển khai.

## V41.8 — Dynamic Match Length & Hourly Pressure

- Chủ phòng chọn 3/6/9/12/15 phút trước khi tạo hoặc trong phòng chờ; mặc định 6 phút. Versus và Foxy Rush áp dụng thời lượng mới, Co-op giữ 3 phút.
- Đồng bộ thời lượng qua P2P, relay, rematch và checkpoint; relay chấp nhận thời lượng tối đa 900 giây. Giao thức 43; không ghép phòng khác phiên bản.
- Tốc độ quái tăng 3 điểm % theo mỗi bậc thời lượng và 7 điểm % mỗi giờ game; sát thương cửa tăng 2 và 5 điểm % tương ứng. Hồi kỹ năng giảm 1,5 và 4 điểm % tương ứng. Thời gian phản ứng cửa giảm 2,5 điểm % mỗi giờ.
- Hệ số tiêu hao điện liên tục và phần lớn chi phí kỹ năng = 180 / số giây của ván. Tốc độ sửa cửa từ 5 đến 7 điểm/giây. HUD hiển thị áp lực 1/6–6/6, phí kỹ năng và nhiệm vụ theo 1/3 thời lượng trận.
- Kiểm tra cú pháp JavaScript, 30 tổ hợp thời lượng/giờ và 30 kiểm tra biên giai đoạn đạt; chưa biên dịch relay Go, chưa kiểm thử P2P/relay đầu cuối, chưa cân bằng qua trận chơi thật.

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
