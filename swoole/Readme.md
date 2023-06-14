## swoole
基于ubuntu:bionic制作的swoole运行环境
```docker
#打包
docker build -t swoole .
#使用
docker run -itd --name swoole_project_dev -v /path/to/project:/opt/app swoole php /opt/app/xxx.php 
```