## php-fpm-fastadmin
基于ubuntu:bionic为fastadmin量身定制的php7.4(fpm) + nginx 运行环境
```docker
# 构建
docker build -t . php-fpm-fastadmin
# 运行
docker run -itd --name -p 80:80 -v /path/to/project:/data/www/html php-fpm-fastadmin
```