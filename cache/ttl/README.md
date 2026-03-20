# Cache com TTL

Este exemplo expira automaticamente itens apos um tempo definido no momento do `Set`.

## Beneficios

- Evita dados velhos por tempo indefinido.
- Simples de controlar frescor dos dados.
- Muito usado com token, sessao, consulta e resposta de API.

## Quando usar

- Dados que perdem valor com o tempo.
- Integracoes externas.
- Sessao, autenticacao e consultas temporarias.

## Limitacoes

- Se o TTL for curto demais, o cache perde eficiencia.
- Se o TTL for longo demais, aumenta risco de dado desatualizado.

## Executando

```sh
go run .
```
