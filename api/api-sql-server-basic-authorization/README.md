## API Data Warehouse

## Execução
```env AMBIENTE=desenvolvimento go run *.go```

## Serviços
### Clientes
* GET 
    1. http://127.0.0.1:3002/api/v1/clientes?pagina=1&limit=5
    1. http://127.0.0.1:3002/api/v1/cliente?cpf=11111111124

## Autenticação
Autenticação do tipo Basic Auth, deve-se informar usuário e senha.
Os dados de usuário e senha deve ser configurado no arquivo de configuração .env que está na raiz do projeto

## configuração
Porta e os dados de conexão com banco de dados deve ser configurado no arquivo de configuração *.env* que está na raiz do projeto