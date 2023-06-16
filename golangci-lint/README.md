# Rodando lint com biblioteca golang CI Lint
[golangci-lint](https://github.com/golangci/golangci-lint)

Para rodar execute o comando com Makefile
```sh
make build-image
```

Neste exemplo a verificação foi feita na geração da image, onde é executado o link, logo depois o teste unitário para no fim gerar o binário da aplicação.

Caso o lint quebre o processo é encerrado pela metade, até que o problema seja resolvido, como no exemplo abaixo:

```sh
.....
 ---> 775268663460
Step 6/14 : COPY . ./
 ---> 9b8088be7926
Step 7/14 : RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.53.3
 ---> Running in 9ffc5ad269d4
golangci/golangci-lint info checking GitHub for tag 'v1.53.3'
golangci/golangci-lint info found version: 1.53.3 for v1.53.3/linux/amd64
golangci/golangci-lint info installed ./bin/golangci-lint
Removing intermediate container 9ffc5ad269d4
 ---> f59c7854959a
Step 8/14 : RUN ./bin/golangci-lint --version
 ---> Running in 7a7c2c30d84e
golangci-lint has version 1.53.3 built with go1.20.5 from 2dcd82f3 on 2023-06-15T10:50:11Z
Removing intermediate container 7a7c2c30d84e
 ---> ac6a30001e42
Step 9/14 : RUN ./bin/golangci-lint run ./...
 ---> Running in 864e88f4e78b
main.go:11:3: S1023: redundant break statement (gosimple)
                break
                ^
The command '/bin/sh -c ./bin/golangci-lint run ./...' returned a non-zero code: 1
make: *** [Makefile:13: build-image] Erro 1
```

Neste exemplo o erro está na instrução switch onde não se deve usar o comando break.
```sh
....
	switch letraAlfabeto {
	case "a":
		fmt.Println("Letra A encontrada")
		break  <----- Deve ser removido para passar no lint ---
	}
```