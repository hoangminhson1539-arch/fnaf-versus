FROM golang:1.23-alpine AS build
WORKDIR /src
COPY server.go .
RUN CGO_ENABLED=0 go build -o /server server.go
FROM alpine:3.21
WORKDIR /app
COPY --from=build /server /app/server
COPY index.html /app/
ENV FNAF_PUBLIC=1 PORT=8080
EXPOSE 8080
USER 65534:65534
CMD ["/app/server"]
