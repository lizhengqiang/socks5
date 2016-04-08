FROM golang:latest
COPY ./ $GOPATH/src/socks5
WORKDIR $GOPATH/src/socks5

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && go get \
    && go build

ENV LISTEN 0.0.0.0:1080

EXPOSE 1080

CMD ["./socks5"]