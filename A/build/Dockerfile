FROM golang:alpine AS builder
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
COPY .. .
RUN go build -ldflags="-s -w" -o /app/main main.go

FROM alpine:3.17
ENV TZ Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 40001
CMD ["./main"]

