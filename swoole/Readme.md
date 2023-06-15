## swoole
基于ubuntu:bionic制作的swoole运行环境
```docker
# 构建
docker build -t swoole .
# 运行
docker run -itd --name swoole_project_dev -v /path/to/project:/opt/app swoole php /opt/app/xxx.php 
```