# FNAF 1 Versus · V41.3

**[Chơi ngay trên trình duyệt](https://hoangminhson1539-arch.github.io/fnaf-versus/)**

**[Nhật ký cập nhật chi tiết — CHANGELOG.md](CHANGELOG.md)**

V41.3: Đập cửa ưu tiên gây tiếng động đánh lạc hướng; giảm mạnh sát thương, hồi chung 8s/cửa, không trực tiếp rút điện. Giữ radar tên phòng/CAM, bản đồ di động và sửa cửa chậm. Dùng hai client cùng phiên bản; relay giao thức 42.

Game fan-made chạy trên trình duyệt. Có chế độ một người, Co-op, đối đầu đội quái và Foxy Rush.

## Chơi trên web

- Mở trang GitHub Pages của kho trên Chrome, Safari hoặc Firefox. Game nằm ngay ở trang chủ.
- Để chơi hai người, chọn **Chơi cùng bạn → Khác mạng Internet → Tạo phòng**. Gửi link mời hoặc mã phòng 6 ký tự cho người còn lại. Chọn vai và bắt đầu khi phòng đủ hai người.
- Chơi online trên GitHub Pages dùng PeerJS và WebRTC. Nếu mạng chặn kết nối trực tiếp, cả hai có thể vào cùng máy chủ HTTPS chạy relay Go.

## Chạy relay

Cần Go 1.23 trở lên. Chạy lệnh trong thư mục dự án, cạnh `index.html`.

```sh
go run server.go
```

Máy chủ hiển thị địa chỉ để các thiết bị cùng Wi-Fi mở game. Với dịch vụ hỗ trợ Docker và HTTPS, có thể dùng `Dockerfile` và `render.yaml`. GitHub Pages chỉ phục vụ trang tĩnh.

Xem [hướng dẫn phím điều khiển, kỹ năng và kết nối](HUONG_DAN.txt).
