package main

import (
	"fmt"

	"github.com/tayron/go-lang-estudos/redis/cache/redis"
)

func main() {
	redis.Ping()
	redis.Set("Nome", "Tayron Miranda")
	fmt.Println(redis.Get("Nome"))
}
