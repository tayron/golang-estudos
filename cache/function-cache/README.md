# Cache de Aplicacao

Este exemplo mostra memoization, ou seja, cacheando o resultado de uma funcao cara para chamadas futuras com os mesmos parametros.

## Beneficios

- Remove recomputacoes desnecessarias.
- Muito util para funcoes puras ou com custo alto.
- Pode reduzir bastante tempo de CPU.

## Quando usar

- Calculos repetidos.
- Regras deterministicas.
- Funcoes que recebem os mesmos parametros com frequencia.

## Limitacoes

- Precisa de cuidado com crescimento de memoria.
- Nao e ideal para funcoes com efeitos colaterais.
- Parametros precisam ser bons candidatos para chave de cache.

## Executando

```sh
go run .
```
