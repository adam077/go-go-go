FROM golang:1.11.1 AS build
#FROM alpine:latest
ENV TZ=Asia/Shanghai LANG=C.UTF-8 TERM=xterm
WORKDIR /go/src/go-go-go
COPY . .
RUN go build -tags=jsoniter -o main src/main.go
#暴露端口
EXPOSE 8080
#最终运行docker的命令
ENTRYPOINT ["./main"]