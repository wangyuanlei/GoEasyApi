#!/bin/bash
 
# 假设你的Go应用名为myapp
APP_NAME=GoEasyApi
 
# 检查应用是否已经在运行
if pgrep -x "$APP_NAME" >/dev/null
then
    echo "$APP_NAME is already running"
else
    # 如果没有运行，则启动应用
    echo "Starting $APP_NAME..."
    ./$APP_NAME &
fi