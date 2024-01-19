# Rabbit MQ com Golang
Para saber o que é Rabbit MQ assista o vídeo [RabbitMQ (Mensageria Robusta pros seus Softwares) // Dicionário do Programador
](https://www.youtube.com/watch?v=_Uo14nxB_iA)
Referência: [https://x-team.com/blog/set-up-rabbitmq-with-docker-compose](https://x-team.com/blog/set-up-rabbitmq-with-docker-compose)


## Subindo container Docker com Rabbit MQ
```sh
docker-composer up --build -d
```

## Acessando Rabbit MQ pelo navegador
Acesse usando IP do container na porta 15672, exemplo: ```http://localhost:15672```

As credenciais para acessar são:

* Usuário: guest
* Senha: guest

## Consumindo mensagem
Para consumir mensagem da fila criada anteriormente utilize o comando:
```sh
go run consumer.go
```

O consumer vai ficar esperando uma mensagem ser enviada e irá exibir uma mensagem semelhante abaixo:

```sh
2022/08/17 14:14:52  [*] Waiting for messages. To exit press CTRL+C
```

Não feche o terminal, deixe ele aberto de lado e abra um novo terminal para fazer o envio de uma mensagem, que se segue no passo seguinte.

## Publicando mensagem
Para publicar mensagem utilize o comando:
```sh
go run producer.go
```

O comando acima irá perguntar qual mensagem enviar
```sh
What message do you want to send?
```

Digite a mensagem em seguida
```sh
Hello World
```
O resultado total será semelhante ao exibido abaixo:
```sh
What message do you want to send?
Hello Worl
2022/08/17 14:14:07  [x] Congrats, sending message: Hello Worl
```

Voltando ao passo anterior (Consumindo mensagem), ao visualizar o terminal você verá a mensagem enviada:
```sh
tayron@tayron:.../rabbit-mq$ go run consumer.go
2022/08/17 14:14:52  [*] Waiting for messages. To exit press CTRL+C
2022/08/17 14:16:40 Received a message: Hello World
```