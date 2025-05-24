package stream

import (
	"context"
	"encoding/json"
	"fmt"
	"web3-tokeninfo/internal/models"

	goRedis "github.com/redis/go-redis/v9"
)

type redis struct {
	db         *goRedis.Client
	subscriber *goRedis.PubSub
}

func (f *redis) init() error {
	f.db = goRedis.NewClient(&goRedis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return nil
}

func (f *redis) Ping() error {
	_, err := f.db.Ping(context.Background()).Result()
	return err
}

func (f *redis) Publish(topic string, message models.EventMessage) error {
	strMsg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	cmd := f.db.Publish(context.Background(), topic, strMsg)
	return cmd.Err()
}

func (f *redis) Subscribe(topic string, callback func(msg string)) {
	f.subscriber = f.db.Subscribe(context.Background(), topic)
	for {
		msg, err := f.subscriber.ReceiveMessage(context.Background())
		if err != nil {
			fmt.Printf("Error receiving message: %s\n", err.Error())
			continue
		}
		callback(msg.Payload)
	}
}
