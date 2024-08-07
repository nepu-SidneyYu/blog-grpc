# 使用官方的 golang 基础镜像
FROM golang:1.17-alpine AS build

# 设置工作目录
WORKDIR /app

# 拷贝项目文件到工作目录
COPY . .

# 构建应用
RUN go build -o myapp .

# 运行编译好的应用
CMD ["./myapp"]
