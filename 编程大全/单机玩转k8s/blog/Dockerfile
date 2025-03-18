# 构建阶段
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp

# 最终生成镜像
FROM alpine:latest
WORKDIR /
COPY --from=builder /myapp /myapp
EXPOSE 5678
CMD ["/myapp"]
