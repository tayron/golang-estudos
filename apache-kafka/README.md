# Apache Kafka e Zookeeper

## O que é Apache Kafka
O Apache Kafka(chamado só Kafka neste post) é uma plataforma de streaming distribuído capaz de manipular trilhões de eventos por dia. Criado e disponibilizado Open source pelo LinkedIn em 2011, o Kafka foi inicialmente concebido como uma fila de mensagens, porém rapidamente evoluiu para um plataforma de streaming completa, possuindo não somente funcionalidades de message broker, mas também a capacidade de armazenar e processar os dados em um fluxo.
[https://medium.com/trainingcenter/apache-kafka-838882261e83](https://medium.com/trainingcenter/apache-kafka-838882261e83)

## O que é Apache Zookeeper
o Apache Zookeeper é um serviço centralizado para manter informações de configurações e nomenclaturas entre serviços distribuídos. O Kafka utiliza o Zookeeper para sincronizar as configurações entre diferentes clusters.
[https://medium.com/trainingcenter/apache-kafka-codifica%C3%A7%C3%A3o-na-pratica-9c6a4142a08f](https://medium.com/trainingcenter/apache-kafka-codifica%C3%A7%C3%A3o-na-pratica-9c6a4142a08f)

## Startando Apache Kafka e Zookeeper

## Iniciando kafka usando docker-compose
```sh
docker-compose up --build -d
```

Após rodar o comando acima execute:
```sh
docker ps
```

Após rodar o comando poderemos ver o kafka rodando na porta 29092
```sh
CONTAINER ID   IMAGE                              COMMAND                  CREATED          STATUS         PORTS                                                     NAMES
a13f41cf8b47   confluentinc/cp-kafka:latest       "/etc/confluent/dock…"   9 seconds ago    Up 8 seconds   9092/tcp, 0.0.0.0:29092->29092/tcp, :::29092->29092/tcp   apache-kafka_kafka_1
0dd52b809ebc   confluentinc/cp-zookeeper:latest   "/etc/confluent/dock…"   10 seconds ago   Up 8 seconds   2181/tcp, 2888/tcp, 3888/tcp                              apache-kafka_zookeeper_1
```

## Criando um topico
```sh
docker-compose exec kafka kafka-topics --create --topic processar-despesa --partitions 1 --replication-factor 1 --if-not-exists --bootstrap-server localhost:29092
```

Retorno do comando acima deverá ser:
```sh
Created topic processar-despesa.
```

## Confirmando se o tópico foi criado
```
docker-compose exec kafka kafka-topics --describe --topic processar-despesa --bootstrap-server localhost:29092
```

Retorno do comando acima deverá ser:
```sh
Topic: processar-despesa        TopicId: 4VZRSkjfTwibTWYzd28Fpw PartitionCount: 1       ReplicationFactor: 1    Configs: 
        Topic: processar-despesa        Partition: 0    Leader: 1       Replicas: 1     Isr: 1
```

# Criando cliente
A implementação foi criado com base neste vídeo [https://www.youtube.com/watch?v=P-VHtR_GkgM] (https://www.youtube.com/watch?v=P-VHtR_GkgM)
Repositório do projeto criado no vídeo acima: [https://github.com/hyperyuri/kafka-with-go](https://github.com/hyperyuri/kafka-with-go)

## Enviando e recebendo mensagem
Execute o comando abaixo:

```sh
go run *.go
```

Retorno do comando acima deverá ser:
```sh
Sessão 123 encontrada!
Mensagem recebida:  Hello World
```
