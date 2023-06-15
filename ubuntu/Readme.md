## Ubuntu
基于ubuntu:bionic制作，默认时区为Asia/Shanghai
```docker
# 构建
docker build -t ubuntu-test .
# 运行
docker run -itd --name ubuntu ubuntu-test
```