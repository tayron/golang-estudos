# Cache Distribuido com Redis

Este exemplo usa Redis como cache compartilhado entre varias instancias. O codigo usa protocolo RESP via TCP puro, sem dependencia externa, para demonstrar `PING`, `SET` e `GET`.

## Beneficios

- Compartilhado entre varias aplicacoes e replicas.
- Muito usado em microservicos.
- Bom para dados quentes que precisam ser vistos por todo o cluster.

## Quando usar

- Sistemas escalando horizontalmente.
- Dados criticos compartilhados.
- Sessao, rate limit, token e cache de consulta entre varias instancias.

## Limitacoes

- Tem custo de rede.
- Precisa de infraestrutura extra.
- Exige observabilidade e politicas de expiracao.

## Subindo o Redis

```sh
docker-compose up -d
```

## Executando o exemplo

```sh
go run .
```

## Exemplo de uso real

- Microservicos acessando o mesmo catalogo de produtos.
- API Gateway compartilhando rate limits.
- Aplicacao com alta performance usando in-memory local e fallback Redis.
