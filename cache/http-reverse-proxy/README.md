# Cache HTTP / Reverse Proxy

Este exemplo usa Nginx como reverse proxy na frente de um backend Go. O proxy faz cache de requisicoes GET em `/products`.

## Beneficios

- Excelente para APIs com muito GET.
- Reduz carga no backend.
- Pode ser aplicado sem mudar grande parte da logica da aplicacao.

## Quando usar

- APIs REST muito lidas.
- Endpoints publicos e relativamente estaveis.
- Conteudo que pode ser servido por alguns segundos ou minutos sem problema.

## Subindo o ambiente

```sh
docker-compose up --build
```

## Testando

Acesse varias vezes:

```sh
curl -i http://localhost:8081/products
```

Veja o header `X-Proxy-Cache`:

- `MISS`: proxy buscou do backend.
- `HIT`: proxy respondeu do proprio cache.

## Exemplo de uso real

- APIs de catalogo.
- Conteudo institucional.
- Endpoints de busca ou listagem com alta taxa de leitura.
