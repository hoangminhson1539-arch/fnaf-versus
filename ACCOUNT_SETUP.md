# Kích hoạt tài khoản FNAF Versus V56.0 — 3 bước

Chỉ quản trị viên cần làm một lần. Người chơi không cần tự tạo project hay cấu hình gì. Sau kích hoạt, một nút Google/Discord/Apple/GitHub hoặc một địa chỉ email là đủ để đăng ký hoặc đăng nhập. Hiện chưa cấu hình Supabase nên tính năng online chưa hoạt động; chơi khách vẫn bình thường.

## Bước 1 — Tạo Supabase
1. Mở https://supabase.com/dashboard và tạo project.
2. SQL Editor → New query → dán nội dung [cloud_saves.sql](supabase/cloud_saves.sql) → Run.
3. Authentication → URL Configuration: đặt Site URL và thêm Redirect URL chính xác: https://hoangminhson1539-arch.github.io/fnaf-versus/ . Nếu dùng domain khác, thêm URL tương ứng.

## Bước 2 — Kích hoạt email trước (nhanh nhất)
- Authentication → Providers → Email: bật đăng ký và đăng nhập email. Giữ xác minh email. Dùng mẫu email Magic Link có ConfirmationURL mặc định, cấu hình SMTP khi cần.
- Trong index.html tìm V55_AUTH_CONFIG. Điền Project URL và khóa công khai sb_publishable_... từ dashboard. Chỉ để danh sách providers là [] nếu chưa bật đăng nhập xã hội. Không dùng khóa service_role, sb_secret_ hoặc OAuth secret.
- Commit lên GitHub. Người chơi vào Tài khoản → nhập email → Gửi link đăng nhập → mở link trên cùng trình duyệt. Tài khoản mới tự được tạo, không cần mật khẩu. Có tùy chọn mật khẩu cũ.

## Bước 3 — Tùy chọn Google/Discord/Apple/GitHub
Khi cần từng nút, tạo OAuth application ở nhà cung cấp, dùng callback URL do Supabase hiển thị, nhập client ID và secret vào Supabase; bật provider và thêm tên tương ứng vào providers trong V55_AUTH_CONFIG. Apple có thể cần Apple Developer/Services ID và gia hạn secret. Không đưa secret lên GitHub.
Tài liệu: https://supabase.com/docs/guides/auth/social-sign-in và https://supabase.com/docs/guides/auth/redirect-urls.

## Sử dụng và kiểm tra
- Bản lưu online: Lưu/Tải tiến trình có xác nhận, không tự ghi đè. Lưu đêm, xu, vật phẩm, thành tích, nhiệm vụ, hồ sơ, 12 trận gần nhất; không lưu ca dở, chat/voice hay thiết lập máy. Đăng xuất giữ dữ liệu offline. Tải cloud sao lưu cục bộ trước khi thay tiến trình.
- Kiểm tra email magic link trên cùng trình duyệt, callback OAuth, đăng xuất, hai tài khoản không đọc/ghi dữ liệu nhau, hai thiết bị cạnh tranh revision, mất mạng/CDN và hết dung lượng localStorage.
- SQL RLS chỉ cho chủ tài khoản đọc/ghi. Không dùng cloud save client-side làm chứng cứ thành tích chống gian lận.
- Chưa thử đăng nhập hoặc cơ sở dữ liệu thật. Không công bố hoạt động trước khi hoàn tất cấu hình và kiểm thử.
