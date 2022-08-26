package main

import (
	"fmt"

	"github.com/tayron/golang-estudos/redis/cache/redis"
)

func main() {
	redis.Ping()
	redis.Set("Nome", "Tayron Miranda")
	fmt.Println(redis.Get("Nome"))
}
