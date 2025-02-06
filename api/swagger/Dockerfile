# Etapa de build
FROM golang:1.22 as build

# Adiciona o código-fonte
ADD . /server-app

WORKDIR /server-app

# Inicializa e organiza o módulo Go
RUN go mod tidy

# Instala o Go Swagger
RUN go install github.com/go-swagger/go-swagger/cmd/swagger@latest

# Garante que o diretório do Go binário esteja no PATH
ENV PATH=$PATH:/go/bin

# Gera o spec Swagger
RUN swagger generate spec -o ./swagger-ui/swagger.json

# Compila o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o server .

RUN ls -la
RUN ls -la ./swagger-ui

# Etapa de produção: cria uma imagem mínima com `scratch` novamente
FROM scratch

WORKDIR /app

# Copia os binários e o arquivo swagger.json para a imagem final
COPY --from=build /server-app/server /app
COPY --from=build /server-app/swagger-ui /app/swagger-ui

# Define a variável de ambiente e expõe a porta
ENV PORT 8080
EXPOSE 8080

# Comando de entrada do container
ENTRYPOINT [ "./server" ]
