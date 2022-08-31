package pubsub

import (
	"gopkg.in/redis.v2"
	"log"
	"net"
	"reflect"
	"time"
)

type Subscriber struct {
	pubsub   *redis.PubSub
	channel  string
	callback processFunc
}

type processFunc func(string, string)

func NewSubscriber(channel string, fn processFunc) (*Subscriber, error) {
	var err error
	// TODO Timeout param?

	s := Subscriber{
		pubsub:   Service.client.PubSub(),
		channel:  channel,
		callback: fn,
	}

	// Subscribe to the channel
	err = s.subscribe()
	if err != nil {
		return nil, err
	}

	// Listen for messages
	go s.listen()

	return &s, nil
}

func (s *Subscriber) subscribe() error {
	var err error

	err = s.pubsub.Subscribe(s.channel)
	if err != nil {
		log.Println("Error subscribing to channel.")
		return err
	}
	return nil
}

func (s *Subscriber) listen() error {
	var channel string
	var payload string

	for {
		msg, err := s.pubsub.ReceiveTimeout(time.Second)
		if err != nil {
			if reflect.TypeOf(err) == reflect.TypeOf(&net.OpError{}) && reflect.TypeOf(err.(*net.OpError).Err).String() == "*net.timeoutError" {
				// Timeout, ignore
				continue
			}
			// Actual error
			log.Print("Error in ReceiveTimeout()", err)
		}

		channel = ""
		payload = ""

		switch m := msg.(type) {
		case *redis.Subscription:
			log.Printf("Subscription Message: %v to channel '%v'. %v total subscriptions.", m.Kind, m.Channel, m.Count)
			continue
		case *redis.Message:
			channel = m.Channel
			payload = m.Payload
		case *redis.PMessage:
			channel = m.Channel
			payload = m.Payload
		}

		// Process the message
		go s.callback(channel, payload)
	}
}
