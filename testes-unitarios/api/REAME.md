# Criando teste unitário para API

## Testando endpoint GET
```sh
curl http://localhost:8181/hello-world
```

## Testando endpoint POST
```sh
curl -X POST http://localhost:8181/hello-world -d '{"name": "Tayron"}'
```

## Testando endpoint UPLOAD 
De dentro da pasta api execute
```sh
curl -F  file=@./testdata/bobs-burgers.jpg http://localhost:8181/upload
```