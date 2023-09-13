#!/bin/bash
php-fpm7.2 -c /etc/php/7.2/fpm/php-fpm.conf
nginx -c /etc/nginx/nginx.conf
chmod 777 /data/www/html -R
echo "Started..."
tail -f /var/log/nginx/access.log