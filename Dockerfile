FROM golang:latest AS builder

ENV GO111MODULE on
ENV GOPATH=/go
#ENV TZ=Asia/Shanghai
WORKDIR $GOPATH/src/qiniu
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o qiniu main.go

FROM alpine:latest
WORKDIR /server/
COPY --from=builder /go/src/qiniu/qiniu .
COPY --from=builder /go/src/qiniu/wait-for-it.sh .
COPY --from=builder /go/src/qiniu/conf/. conf/
COPY --from=builder /go/src/qiniu/logfile/. logfile/
EXPOSE 8080
# 设置时区为上海
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata \
    && apk add --no-cache bash
CMD ["./wait-for-it.sh", "mysql:3306", "--", "./qiniu"]