# API gRPC

Configuração da gRPC
- https://grpc.io/docs/languages/go/quickstart/


## Instalação copilador do protocol buffer

```sh
apt-get install -y protobuf-compiler
```
Verificando se copilador está instalado
```sh
protoc --version
```

## Instalando plugin do protocol buffer para trabalhar com golang

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```
```sh
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Atualizando patch
```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```
## Criando entidades golang definida no diretorio /propo
```sh
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```