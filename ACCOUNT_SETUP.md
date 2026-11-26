# Kích hoạt tài khoản FNAF Versus V55.0

Hiện tích hợp đã có trong game nhưng CHƯA có dự án Supabase được cấp quyền. Không thể đăng nhập thật hoặc lưu cloud cho tới khi quản trị viên cấu hình. Game offline và multiplayer không phụ thuộc vào dịch vụ này.

## 1. Supabase
- Tạo project tại https://supabase.com/dashboard.
- SQL Editor: chạy `supabase/cloud_saves.sql`. Bảng fnaf_saves dùng khóa user_id tham chiếu auth.users, chỉ chủ tài khoản đọc/ghi bằng RLS. Không dùng cho leaderboard hoặc xác minh thành tích chống gian lận.
- Authentication → URL Configuration: Site URL và Redirect URLs thêm chính xác `https://hoangminhson1539-arch.github.io/fnaf-versus/`. Nếu triển khai domain khác thì thêm đúng URL đó.
- Authentication → Providers: bật Email; giữ xác nhận email, cấu hình SMTP và giới hạn gửi email theo nhu cầu.
- Mở từng nhà cung cấp muốn hỗ trợ: Google, Discord, Apple, GitHub. Tạo OAuth application trong console của nhà cung cấp, đặt callback theo URL hiện trong dashboard Supabase (thường là `https://<project-ref>.supabase.co/auth/v1/callback`), nhập client ID và secret vào Supabase, không đưa secret lên GitHub. Apple Web cần cấu hình Apple Developer/Services ID và secret hợp lệ, có thể cần tài khoản Developer và gia hạn secret.
- Tài liệu: https://supabase.com/docs/guides/auth/social-sign-in ; https://supabase.com/docs/guides/auth/redirect-urls

## 2. Kết nối game
Trong index.html tìm V55_AUTH_CONFIG, thay url bằng Project URL HTTPS và publishableKey bằng khóa công khai bắt đầu sb_publishable_. Chỉ giữ tên các provider đã bật trong providers. KHÔNG bao giờ dùng service_role, sb_secret_, OAuth secret hoặc mật khẩu trong GitHub. Commit thay đổi. SDK Supabase JS 2.49.4 tải từ CDN khi mở tài khoản hoặc có callback; nếu mạng chặn thì game khách vẫn hoạt động.

## 3. Sử dụng
Menu → Tài khoản & lưu online → đăng ký email/mật khẩu hoặc đăng nhập nhà cung cấp. Xác nhận email nếu được yêu cầu. Khi đã đăng nhập, bấm Kiểm tra bản lưu / Lưu lên online / Tải tiến trình. Mỗi thao tác lưu/tải yêu cầu xác nhận; không tự tải/ghi đè. Tải tiến trình bỏ ca đang dở, giữ thiết lập cục bộ, sao lưu tiến trình trước khi tải vào localStorage (fnaf_v55_before_cloud). Tài khoản chỉ lưu tiến trình, xu, nâng cấp, thành tích, hồ sơ cốt truyện, nhiệm vụ và tối đa 12 trận gần đây; không lưu phiên multiplayer, ca đang dở, mic/chat hay cài đặt cá nhân. Đăng xuất không xóa tiến trình offline; thiết bị dùng chung cần chú ý không tải tiến trình của người khác lên tài khoản của mình. Không có tự đồng bộ hay gộp tiến trình. Game client-side nên dữ liệu cloud không chứng minh thành tích hợp lệ.

## 4. Kiểm thử trước khi công bố
- Email: đăng ký, xác nhận, đăng nhập, quên/đổi mật khẩu, đăng xuất.
- Từng OAuth: callback trên Chrome/Safari, cả máy tính và điện thoại; provider chưa bật không được quảng bá hoạt động.
- Tài khoản A và B: không thể đọc hoặc ghi bản lưu của nhau; anon không đọc được. Hai tab cạnh tranh revision, chỉ một ghi thành công.
- Hai thiết bị cùng tài khoản: lưu ở A, tải ở B; kiểm tra xu/đêm/thành tích, xác nhận hủy không thay đổi dữ liệu.
- Thử lỗi CDN, offline, mất kết nối, token hết hạn, hết dung lượng localStorage; chơi khách vẫn chạy.
- Kiểm tra lại camera thực địa, split view, thao tác chạm và P2P sau khi GitHub Pages triển khai.

Không công bố đăng nhập hoạt động thực tế nếu chưa hoàn thành cấu hình và kiểm thử trên dịch vụ thật.
