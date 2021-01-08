
# SQLite

Exemple how to use SQLite with Golang using the [tutorial](https://www.bogotobogo.com/GoLang/GoLang_SQLite.php).

## Installing dependency
```
go get github.com/mattn/go-sqlite3
```

## Enable CGO
```
export CGO_ENABLED=1
```
Because it is a CGO enabled package we are required to set the environment variable CGO_ENABLED=1 and have a gcc compile present within our path. Note CGO enables the creation of Go packages that call C code.

## To execute
```
go run *.go
```

Output:
```
1: Rob Gronkowski
```

The file **bogo.db** will be created into directory.