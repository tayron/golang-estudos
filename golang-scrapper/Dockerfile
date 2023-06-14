FROM debian:latest

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install chromium -y

WORKDIR /app
COPY . /app
RUN chmod +x scrapper

RUN rm -rf /var/cache/apk/*
RUN rm -rf /tmp/*
RUN rm -rf /var/log/*