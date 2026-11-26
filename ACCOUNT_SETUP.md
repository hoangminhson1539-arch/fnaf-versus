# Tài khoản FNAF Versus V58.0

**Đã kết nối dự án Supabase của chủ game.** Bảng public.fnaf_saves đã tồn tại, có RLS chỉ cho chủ tài khoản đọc/ghi, quyền authenticated và trigger revision; anon không có quyền đọc. Game dùng URL và publishable key công khai, không dùng service role/secret.

## Người chơi
Menu → Tài khoản & lưu online → nhập email, gửi mã, xác minh mã 6 số trong game (hoặc mở link bằng cùng trình duyệt nếu thư chỉ có link). Có lựa chọn email/mật khẩu. Các nút Google/Discord/Apple/GitHub chỉ xuất hiện nếu máy chủ xác thực báo provider đã bật. Sau khi đăng nhập, chọn Lưu lên online hoặc Tải tiến trình; xác nhận trước khi ghi đè. Chơi khách không cần tài khoản.

## Quản trị viên: phần còn lại để mở đăng ký công khai
- Dashboard dự án: https://supabase.com/dashboard/project/thsrdwldepzoprvpcaym
- Authentication → URL Configuration: đặt Site URL và Redirect URL đúng https://hoangminhson1539-arch.github.io/fnaf-versus/ .
- Authentication → Email/SMTP: cấu hình máy chủ SMTP riêng nếu muốn gửi thư cho tất cả người chơi; SMTP mặc định của Supabase giới hạn địa chỉ thành viên dự án và 2 thư/giờ. Không tắt xác minh email để né giới hạn. Tài liệu: https://supabase.com/docs/guides/auth/auth-smtp
- Email Templates → Magic Link: thêm mã {{ .Token }} khi SMTP/gói dịch vụ cho phép chỉnh mẫu; có thể giữ {{ .ConfirmationURL }}. Nếu chưa chỉnh mẫu, dùng link mặc định.
- OAuth: bật từng provider trong Authentication và tạo OAuth app, đặt callback theo Supabase, nhập client ID/secret tại dashboard. Không commit secret. Không cần sửa code để nút hiện sau khi provider bật; game đọc /auth/v1/settings. Tài liệu: https://supabase.com/docs/guides/auth/social-sign-in

## Lưu trữ và kiểm thử
Lưu đêm, xu, vật phẩm, thành tích, nhiệm vụ, hồ sơ, 12 trận gần nhất; không lưu ca đang dở, chat/voice hay thiết lập máy. Đăng xuất giữ tiến trình offline. Tải cloud sao lưu cục bộ trước khi thay thế. RLS đã kiểm tra trên DB thật; chưa kiểm thử email/OTP/OAuth, gửi email SMTP, callback hoặc lưu/tải qua hai thiết bị. Kiểm tra toàn bộ trước khi công bố đăng ký công khai. Một cảnh báo advisor hiện hữu về quyền gọi hàm rls_auto_enable nằm ngoài bảng lưu game; cần rà soát riêng: https://supabase.com/docs/guides/database/database-linter?lint=0028_anon_security_definer_function_executable .

