package redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

const (
	localhost = "localhost"
	port      = 6379
	pass      = "Redis2019!"
)

func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", localhost, port),
		Password: pass,
		DB:       0,
	})

	return client
}

func Ping() {
	client := RedisConnect()
	defer client.Close()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func Set(key string, value string) error {

	client := RedisConnect()
	defer client.Close()

	err := client.Set(key, value, 0).Err()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return err
}

func Get(key string) (string, error) {

	client := RedisConnect()
	defer client.Close()

	value, err := client.Get(key).Result()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return value, err
}
