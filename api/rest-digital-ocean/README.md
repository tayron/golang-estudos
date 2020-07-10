## Exemplo Digital Ocean
https://kenyaappexperts.com/blog/how-to-deploy-golang-to-production-step-by-step/

## Execução
```go run main.go```

## Execução em produção

```sudo vim /etc/systemd/system/jobs.service```

Conteudo do arquivo:

```
Description= instance to serve jobs api
After=network.target

[Service]
User=root
Group=www-data
# /var/www/html/jobs (jobs é o binario da aplicação)
ExecStart=/var/www/html/jobs
[Install]
WantedBy=multi-user.target
```
Depois executar 
```service jobs start```
