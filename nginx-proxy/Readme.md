## Nginx Proxy
基于ubuntu:bionic制作的nginx反向代理
```docker
#打包
docker build -t nginx-proxy .
#使用
docker run -itd --name nginx_proxy -e REMOTE_ADDRESS="http://www.baidu.com/" -p 8080:80 nginx-proxy
```