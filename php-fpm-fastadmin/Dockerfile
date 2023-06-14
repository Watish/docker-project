FROM ubuntu:bionic

RUN apt update
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt install tzdata -y
RUN apt install software-properties-common -y
RUN add-apt-repository ppa:ondrej/php -y
RUN apt update

# 创建目录
RUN mkdir -p /data/www/html
RUN chmod 777 /data/www/html -R
RUN mkdir -p /run/php/

# Nginx
RUN apt install -y nginx

# PHP 7.4
RUN apt install -y php7.4-fpm
# PHP 拓展
RUN apt install -y php7.4-yaml php7.4-mbstring php7.4-soap php7.4-odbc php7.4-mysql php7.4-sqlite3 php7.4-redis php7.4-gmp php7.4-amqp php7.4-readline php7.4-snmp php7.4-xml php7.4-zip 
RUN apt update && apt install -y php7.4-gd php7.4-curl  --fix-missing

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
# 放入fastadmin专用配置
COPY ./fastadmin.conf /etc/nginx/sites-enabled/

# 放入入口文件
COPY ./entrypoint.sh /
RUN chmod 777 /entrypoint.sh

EXPOSE 80

ENTRYPOINT [ "/entrypoint.sh" ]




