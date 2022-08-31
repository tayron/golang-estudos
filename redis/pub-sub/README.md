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
2022/08/31 18:28:13 Subscription Message: subscribe to channel 'food'. 1 total subscriptions.
2022/08/31 18:28:13 Subscriptions done. Publishing...
2022/08/31 18:28:13 Subscription Message: subscribe to channel 'cars'. 1 total subscriptions.
2022/08/31 18:28:13 Eating a Pizza.
2022/08/31 18:28:13 Eating a Big Mac.
2022/08/31 18:28:13 Driving a Subaru.
2022/08/31 18:28:13 Publishing done. Sleeping...
2022/08/31 18:28:13 Driving a Tesla.
2022/08/31 18:28:14 Error in ReceiveTimeout()read tcp 127.0.0.1:57526->127.0.0.1:6379: i/o timeout
2022/08/31 18:28:14 Error in ReceiveTimeout()read tcp 127.0.0.1:57522->127.0.0.1:6379: i/o timeout
2022/08/31 18:28:14 Unmarshal error: unexpected end of JSON input
2022/08/31 18:28:14 Driving a .

```