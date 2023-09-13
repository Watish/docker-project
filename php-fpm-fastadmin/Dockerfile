FROM ubuntu:bionic


RUN apt update
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt install tzdata -y
RUN apt update
RUN apt install software-properties-common -y
RUN apt install -y curl

# 创建目录
RUN mkdir -p /data/www/html
RUN chmod 777 /data/www/html -R
RUN mkdir -p /run/php/

# Nginx
RUN echo "deb http://nginx.org/packages/mainline/ubuntu `lsb_release -cs` nginx" \
     | tee /etc/apt/sources.list.d/nginx.list
RUN curl -o /tmp/nginx_signing.key https://nginx.org/keys/nginx_signing.key
RUN mv /tmp/nginx_signing.key /etc/apt/trusted.gpg.d/nginx_signing.asc
RUN apt update
RUN apt install -y nginx
RUN apt search nginx


# PHP 7.4
RUN add-apt-repository ppa:ondrej/php -y
RUN apt update
RUN apt search php7
RUN apt install -y php7.2-fpm
# PHP 拓展
RUN apt install -y php7.2-yaml php7.2-mbstring php7.2-soap php7.2-odbc php7.2-mysql php7.2-sqlite3 php7.2-redis php7.2-gmp php7.2-amqp php7.2-readline php7.2-snmp php7.2-xml php7.2-zip 
RUN apt update && apt install -y php7.2-gd php7.2-curl  --fix-missing
RUN mkdir -p /var/lib/php/sessions/
RUN chmod 777 /var/lib/php/sessions/ -R

# 删除原来的 nginx配置文件
RUN rm /etc/nginx/nginx.conf
# 替换新的
COPY ./nginx.conf /etc/nginx/
# 删除配置目录
# RUN rm /etc/nginx/sites-enabled/ -r
RUN rm /etc/nginx/conf.d/ -r
# 重新创建
RUN mkdir -p /etc/nginx/sites-enabled/
RUN mkdir -p /etc/nginx/conf.d/
# 放入fastadmin专用配置
COPY ./fastadmin.conf /etc/nginx/sites-enabled/

# 放入入口文件
COPY entrypoint.sh /
#RUN chmod 777 /entrypoint.sh
RUN apt install -y bash
RUN apt install -y dos2unix

RUN dos2unix /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 80

ENTRYPOINT [ "/entrypoint.sh" ]