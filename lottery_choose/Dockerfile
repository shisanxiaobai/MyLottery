
FROM golang:1.19-alpine
#为镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

#工作目录
WORKDIR /project/go-docker/mylottery

ADD lottery-user /project/go-docker/mylottery/lottery-choose
RUN go mod download

ENTRYPOINT [ "/project/go-docker/mylottery/lottery-choose" ]
