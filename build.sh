#!/bin/bash
# 构建镜像
#docker build -t flow-bridge-mcp:latest .
docker build -t crpi-puijznmazn81i1ou.cn-hangzhou.personal.cr.aliyuncs.com/test_huizhi/flow-bridge-mcp:latest .


#docker login crpi-puijznmazn81i1ou.cn-hangzhou.personal.cr.aliyuncs.com

 # 推送镜像到阿里云容器镜像仓库
docker push crpi-puijznmazn81i1ou.cn-hangzhou.personal.cr.aliyuncs.com/test_huizhi/flow-bridge-mcp:latest