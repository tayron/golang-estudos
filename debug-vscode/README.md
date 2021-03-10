# Exemplo de configuração para debug no VS Code

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/main.go",
            "env": {
                "AMBIENTE": "desenvolvimento",
                "DB_PORTA": 80,
                "DB_USUARIO": "root",
                "DB_SENHA": "123456",
                "DB_BANCO": "sistema"
            },
            "args": []
        }
    ]
}
```

## Observações:
*  **{workspaceFolder}** - Pega caminho do diretório raiz o projeto, bastando informar nome do arquivo de encontrada do programa, exemplo: **main.go**
* **env** - Seta todas as variaveis de ambiente do programa
* **args** - Seta todos os parametros passados via linha de comando (Não testado)

Após feita essas configurações basta apertar **F5** para inicializar o debug.
