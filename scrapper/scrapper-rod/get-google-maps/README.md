# Scrapper utilizando biblioteca Go Rod
Scrapper é uma técnica de programação que automatiza a extração de dados de websites de forma estruturada.

Este exemplo acessa o site [www.hospeda.app](https://www.hospeda.app) e exibe o html da página no terminal utilizando a bilioteca [go-rod]( https://github.com/go-rod/rod).


## Configuração
Para executar o exemplo feito com a biblioteca go-rod, você deve previamente instalar o programa [upx](https://upx.github.io) em seu comptutado, ele será utilizado na otimização 
do binário da aplicação gerada.

```sh
apt-get install upx
```

Para geração do binário, build da imagem e execução, há os comandos no arquivo Makefile, portando caso não queira utilizar o upx, basta abrir o arquivo e comentar a linha que faz a utilização dele.

## Gerando binário do scrapper
Para gerar o binário do scrapper execute:
```sh
make build-scrapper
```

## Gerando imagem docker
Para gerar a imagem docker, deve-se antes gerar o binário do scrapper, logo depois execute:
```sh
make build-image
```

## Executando imagem gerada
Execute o comando abaixo para subir o container e acessar o terminal da imagem gerada:
```sh
make run-image
```

Dentro do container execute o comando para rodar o scrapper:
```sh
./scrapper
```

Logo em seguida será impresso o html da página do site que foi acessado

## Removendo imagem
Para excluir a imagem gerada execute:
```sh
make remove-image
```

## Para obter o tamanho da imagem
Para obter o tamanho da imagem gerada execute:
```sh
make get-size-image
```

