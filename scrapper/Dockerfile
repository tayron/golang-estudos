FROM debian:latest

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install chromium -y

WORKDIR /app
COPY scrapper /app
RUN chmod +x scrapper

RUN rm -rf /var/lib/{apt,dpkg,cache,log}/ /tmp/* /var/tmp/*