# Nhật ký cập nhật FNAF 1 VERSUS

Mã nguồn chính: `main`. [Chơi game](https://hoangminhson1539-arch.github.io/fnaf-versus/).

**Quy trình bắt buộc:** Mỗi lần cập nhật game, phải cập nhật file `CHANGELOG.md` trên GitHub cùng phiên bản mã nguồn. Ghi rõ phiên bản, ngày, tính năng mới, sửa lỗi, thay đổi cân bằng, giới hạn và kết quả kiểm thử thực tế; không ghi tính năng chưa triển khai là đã hoàn thành.

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
