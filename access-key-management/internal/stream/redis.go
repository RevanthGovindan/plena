package stream

import (
	"access-key-management/internal/models"
	"context"
	"encoding/json"

	goRedis "github.com/redis/go-redis/v9"
)

type redis struct {
	db *goRedis.Client
}

func (f *redis) Init() error {
	f.db = goRedis.NewClient(&goRedis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping to test connection
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

func (f *redis) Subscribe(string) error {
	return nil
}
