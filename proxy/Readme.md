## proxy
一个用Go编写的简易端口转发工具

```docker
# 构建
docker build -t proxy .
# 运行
docker run -itd --name proxy_example -e SERVER_ADDR="baidu.com" -e SERVER_PORT=80 -e LISTEN_PORT=80 -p 8080:80 proxy
```