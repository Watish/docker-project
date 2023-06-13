#!/bin/bash
php-fpm7.4 -c /etc/php/7.4/fpm/php-fpm.conf
nginx -c /etc/nginx/nginx.conf
chmod 777 /data/www/html -R
echo "Started..."
tail -f /var/log/access.log