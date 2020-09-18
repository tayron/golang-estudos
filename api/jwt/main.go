package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

var (
	router = gin.Default()
	client *redis.Client
)

func init() {
	//this should be in an env file
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	os.Setenv("ACCESS_PORT", "8181")
	os.Setenv("REDIS_DSN", "172.22.0.2:6379")

	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr:     dsn, //redis port
		Password: "Redis2019!",
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

}

func main() {
	router.POST("/login", Login)
	router.POST("/todo", CreateTodo)
	router.DELETE("/logout", Logout)

	fmt.Printf("Service started at http://127.0.0.1:%s \n", os.Getenv("ACCESS_PORT"))
	log.Fatal(router.Run(":" + os.Getenv("ACCESS_PORT")))
}
