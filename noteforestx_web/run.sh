#!/bin/bash
# 临时调试 Nuxt SSR 错误

# 创建日志目录
mkdir -p logs

echo "Starting NoteForestX in debug mode..."
echo "Logs: logs/server-out.log & logs/server-error.log"

# 开发模式，打印更多错误
export NODE_ENV=development
export DEBUG=nuxt:*

# 启动 Nuxt 3 SSR
node --enable-source-maps ./.output/server/index.mjs \
  > logs/server-out.log \
  2> logs/server-error.log