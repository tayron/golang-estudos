# Uso de publish e subscriber com Redis

[Referencia do código de exempĺo usado](https://github.com/dabfleming/go-redis-pubsub-example)

## Subindo container Docker com Redis
```sh
docker
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