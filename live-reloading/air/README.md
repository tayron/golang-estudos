Livre Reload com AIR 

Require go1.25.5
# install
```sh
go install github.com/air-verse/air@latest
```

## rodar 
```sh
air --build.cmd "go build -o bin/api cmd/main.go" --build.entrypoint "./bin/api"
```