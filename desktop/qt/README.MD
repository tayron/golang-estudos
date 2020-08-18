# Sistema altera valor de pedido

Programa feito para estudo em GO e QT que permite visualizar valor do pedido e o valor pago pelo cliente.
Permite alteração do valor pago pelo cliente através da interface gráfica

![Alt text](./imagens/1.png?raw=true "Página inicial do programa")
![Alt text](./imagens/2.png?raw=true "Página visualização pedido")
![Alt text](./imagens/3.png?raw=true "Mensagem sucesso ao alterar valor pago")

## Para executar (Somente Linux)
```./pedido-painel-franqueado``` 

## Configuração
Altere arquivo *.env* na raiz do projeto com:
1. Dados de conexão com banco de dados

Exemplo
```
# Dados de conexão do banco de dados
# -------------------------------------------------
DB_LOCALHOST=endereco-servidor
DB_PORTA=3306
DB_USUARIO=nome-usuario
DB_SENHA=senha-banco
DB_BANCO=nome-banco
```
