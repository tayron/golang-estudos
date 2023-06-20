# Controlando timeout utilizando biblioteca Go Rod
Scrapper é uma técnica de programação que automatiza a extração de dados de websites de forma estruturada.

Este exemplo acessa o site [www.hospeda.app](https://www.hospeda.app) e exibe o html da página no terminal utilizando a bilioteca [go-rod]( https://github.com/go-rod/rod).
Portando foi adicionado um timeout de 0 segundo para que não dê tempo da tela carregar e se possa controlar o erro evitando panic.

O exemplo de como controlar o timeout foi retirado da documentação oficial sobre [context e timeout](https://github.com/go-rod/go-rod.github.io/blob/main/context-and-timeout.md).
## Execução
```sh
go run *.go
```

O retorno deverá ser
```sh
timeout error
```

