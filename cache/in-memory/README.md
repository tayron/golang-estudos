# Cache em Memoria

Este exemplo mostra um cache simples mantido dentro do proprio processo usando `map` e `sync.RWMutex`.

## Beneficios

- Muito rapido, sem rede.
- Simples de implementar.
- Excelente para aplicacoes com apenas uma instancia.

## Quando usar

- APIs pequenas rodando em um unico processo.
- Dados pequenos e muito acessados.
- Configuracoes, flags e respostas quentes de leitura.

## Limitacoes

- Cada instancia tem seu proprio cache.
- O conteudo e perdido quando a aplicacao reinicia.
- Nao serve bem para estado compartilhado entre varias replicas.

## Executando

```sh
go run .
```
