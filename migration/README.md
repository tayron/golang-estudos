# Golang Migrate 

* [Link biblioteca golang-migrate](https://github.com/golang-migrate/migrate)



Comando para criar arquivo de migrate
```sh
migrate create -ext sql -dir migrations -seq create_table_clients
migrate create -ext sql -dir migrations -seq add_data_nascimento_table_clients
```

Comando para rodar migrate para adicionar campos ou criar tabelas no banco de dados


```sh
migrate -path ./migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${DATABASE_NAME}" up
```

Comando para rodar migrate para desfazer/remover campos ou tabelas no banco de dados


```sh
migrate -path ./migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${DATABASE_NAME}" down
```

Comando para rodar migrate para desfazer/remover campos ou tabelas no banco de dados, onde o down [2] representa o número de versões que se deve reverter


```sh
migrate -path ./migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${DATABASE_NAME}" down 2
```