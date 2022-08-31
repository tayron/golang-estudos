package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/tayron/golang-estudos/redis/pub-sub/driver/pubsub"
	"gopkg.in/redis.v2"
)

type food struct {
	Name     string
	Calories float64
}

type car struct {
	Make  string
	Model string
}

func main() {
	var pub *redis.IntCmd
	var err error

	// Create a subscriber
	_, err = pubsub.NewSubscriber("food", eat)
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}

	// Create another subscriber
	_, err = pubsub.NewSubscriber("cars", drive)
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}
	log.Print("Subscriptions done. Publishing...")

	// -- Publish some stuf --
	pub = pubsub.Service.Publish("food", food{"Pizza", 50.1})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}

	pub = pubsub.Service.Publish("food", food{"Big Mac", 200})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}

	pub = pubsub.Service.Publish("cars", car{"Subaru", "Impreza"})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}

	pub = pubsub.Service.Publish("cars", car{"Tesla", "Model S"})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}

	log.Print("Publishing done. Sleeping...")

	for {
		time.Sleep(time.Second)
	}
}

func eat(channel, payload string) {
	var f food

	err := json.Unmarshal([]byte(payload), &f)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
	}

	log.Printf("Eating a %v.", f.Name)
}

func drive(channel, payload string) {
	var c car

	err := json.Unmarshal([]byte(payload), &c)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
	}

	log.Printf("Driving a %v.", c.Make)
}
