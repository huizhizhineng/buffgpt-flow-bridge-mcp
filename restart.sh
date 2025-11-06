#!/bin/bash

docker pull crpi-puijznmazn81i1ou.cn-hangzhou.personal.cr.aliyuncs.com/test_huizhi/flow-bridge-mcp:latest

docker-compose down

docker-compose up -d