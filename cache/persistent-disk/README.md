# Cache Persistente em Disco

Este exemplo salva cada item em arquivo JSON no disco. Ele e util quando o volume de dados e maior, ou quando voce quer preservar o cache entre reinicios.

## Beneficios

- Sobrevive ao restart da aplicacao.
- Pode armazenar volumes maiores que a memoria principal.
- Simples para estudos e caches locais.

## Quando usar

- Dados grandes.
- Processamentos em lote.
- Aplicacoes desktop ou jobs que se beneficiam de reuso local em disco.

## Limitacoes

- Mais lento que memoria.
- Requer limpeza e organizacao dos arquivos.
- Pode nao ser a melhor opcao para altissima concorrencia.

## Executando

```sh
go run .
```
