# Estrategias de Cache em Go

Esta pasta contem exemplos prontos de diferentes tecnicas de cache em Go.

Cada estrategia fica em sua propria pasta e possui um `README.md` com beneficios, cenarios ideais de uso e instrucoes para executar.

## Estrategias implementadas

1. `in-memory`
2. `distributed-redis`
3. `persistent-disk`
4. `http-reverse-proxy`
5. `function-cache`
6. `ttl`
7. `invalidation`

## Tabela rapida de decisao

| Cenario | Melhor opcao |
| --- | --- |
| App simples com 1 instancia | In-memory |
| Microservicos e escala horizontal | Redis |
| Dados criticos compartilhados | Redis |
| Alta performance extrema | In-memory + fallback Redis |
| Dados grandes | Disco |
| APIs com muito GET | HTTP cache |

## Como executar

### 1. In-memory
```sh
cd in-memory
go run .
```

### 2. Cache distribuido com Redis
```sh
cd distributed-redis
docker-compose up -d
go run .
```

### 3. Cache persistente em disco
```sh
cd persistent-disk
go run .
```

### 4. Cache HTTP com reverse proxy
```sh
cd http-reverse-proxy
docker-compose up --build
```

Depois acesse `http://localhost:8081/products` varias vezes para observar o cache do proxy.

### 5. Cache de aplicacao
```sh
cd function-cache
go run .
```

### 6. Cache com TTL
```sh
cd ttl
go run .
```

### 7. Cache com invalidacao
```sh
cd invalidation
go run .
```
