FROM ubuntu:bionic

ENV REMOTE_ADDRESS "http://yourdomain:port/"

RUN apt update
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt install tzdata -y

RUN apt install -y nginx

# 删除原来的 nginx配置文件
RUN rm /etc/nginx/nginx.conf
# 替换新的
COPY ./nginx.conf /etc/nginx/
# 删除配置目录
RUN rm /etc/nginx/sites-enabled/ -r
RUN rm /etc/nginx/conf.d/ -r
# 重新创建
RUN mkdir -p /etc/nginx/sites-enabled/
RUN mkdir -p /etc/nginx/conf.d/
COPY ./proxy.conf /etc/nginx/sites-enabled/
RUN sed -i "s|HTTP_ADDRESS|${REMOTE_ADDRESS}|g" /etc/nginx/sites-enabled/proxy.conf

# 放入入口文件
COPY ./entrypoint.sh /
RUN chmod 777 /entrypoint.sh

ENTRYPOINT "/entrypoint.sh"

