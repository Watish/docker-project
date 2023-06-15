FROM ubuntu:bionic

RUN apt update
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt install tzdata -y
RUN apt-get install -y --no-install-recommends ca-certificates curl

CMD ["/bin/bash"]