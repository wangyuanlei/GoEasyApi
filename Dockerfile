#开始编译
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o GoEasyApi_linux .
#构建docker
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/GoEasyApi_linux /app/GoEasyApi_linux
COPY --from=builder /app/db.sql /app/db.sql
COPY --from=builder /app/config.yaml /app/config.yaml
COPY --from=builder /app/start.sh /app/start.sh

ENV PORT 8008
ENTRYPOINT ["sh","/app/start.sh"]

EXPOSE 8008