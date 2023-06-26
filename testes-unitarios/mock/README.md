# Exemplo de teste unitiário usando Mock
Biblioteca usada [https://github.com/golang/mock](https://github.com/golang/mock)
 
Instalação 
```sh
go install github.com/golang/mock/mockgen@v1.6.0
```

Gerando mock
```sh
mockgen -destination=mocks/mock_foo.go -package=mocks . Foo
```