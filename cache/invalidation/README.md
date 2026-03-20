# Cache com Invalidacao

Este exemplo mostra a invalidacao manual do cache quando o dado principal muda no banco.

## Beneficios

- Evita servir dado antigo por muito tempo.
- Mantem mais controle sobre consistencia.
- E uma abordagem comum em CRUDs e sistemas transacionais.

## Quando usar

- Cadastro de usuario, plano, permissao e preco.
- Sempre que uma escrita deve refletir rapido nas leituras.
- Sistemas em que o dado nao pode ficar stale por muito tempo.

## Limitacoes

- Exige disciplina para invalidar em todos os pontos de escrita.
- Se algum fluxo esquecer de invalidar, o cache fica incorreto.

## Executando

```sh
go run .
```
