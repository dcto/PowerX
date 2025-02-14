# 第一阶段
FROM golang:1.21-alpine AS builder

# 配置go proxy为中国国内proxy
ENV GOPROXY=https://goproxy.cn,direct

# 拷贝当前目录到docker内
WORKDIR /app
RUN ls -l
COPY ./ /app/

# 默认使用 linux 架构
ARG TARGETARCH=amd64
ARG TARGETOS=linux

# 编译不同架构的文件
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /app/powerx_$TARGETARCH $(BUILD_EXE_PATH)
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /app/powerxctl_$TARGETARCH $(BUILD_CTL_PATH)




# 第二阶段
FROM alpine:latest

# 安装编译工具和依赖项
RUN apk update && apk add build-base

# 拷贝文件
COPY --from=builder /app/powerxctl /app/powerxctl
COPY --from=builder /app/Makefile /app/Makefile
COPY --from=builder /app/powerx /app/powerx
COPY --from=builder /app/etc/ /app/etc/

RUN chmod +x /app/powerxctl
RUN chmod +x /app/Makefile
RUN chmod +x /app/powerx

WORKDIR /app

EXPOSE 8888

# 运行可执行文件
CMD ["make", "-f", "/app/Makefile","-C", "/app", "app-init"]