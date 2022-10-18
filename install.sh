#!/bin/sh
git pull
docker rm -f novel_server
docker rmi novel_server:v1
docker build -t novel_server:v1 .
docker run -d -p 8093:8093 --name crawler_server novel_server:v1