# Armazenando cache com Redis

[Golang e Redis tutorial](https://tutorialedge.net/golang/go-redis-tutorial)

## Subindo container Docker com Redis
```sh
docker-compose up --build -d
```

## Acessando Rabbit MQ pelo navegador
Acesse usando IP do container na porta 6379, exemplo: ```http://localhost:6379```

As credenciais para acessar são:

* localhost: guest
* porta: 6379
* Senha: Redis2019!

## Executando o teste
```sh
 go run *.go
```

O retorno deverá ser semelhante a:
```sh
PONG <nil>
Tayron Miranda <nil>

```