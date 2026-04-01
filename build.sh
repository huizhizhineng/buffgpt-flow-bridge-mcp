#!/bin/bash
# 构建镜像（推送到私有仓库时请自行修改镜像名并先 docker login）
docker build -t buffgpt-flow-bridge-mcp:latest .

# docker push <your-registry>/buffgpt-flow-bridge-mcp:latest
