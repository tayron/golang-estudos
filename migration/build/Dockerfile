
FROM golang:1.24 

WORKDIR /app

RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -

# Instala o migrate diretamente do repositório oficial
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate

# Comando para mandar o container de pé
CMD ["tail", "-f", "/dev/null"]