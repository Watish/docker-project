FROM golang:1.20.2-alpine

ENV SERVER_ADDR ""
ENV SERVER_PORT 8080
ENV LISTEN_PORT 80

RUN mkdir -p /opt/app/
COPY . /opt/app

RUN cd /opt/app/ && go build main.go
RUN mv /opt/app/main /usr/bin/proxy && chmod 777 /usr/bin/proxy

EXPOSE 80

ENTRYPOINT '/opt/app/entrypoint.sh'

