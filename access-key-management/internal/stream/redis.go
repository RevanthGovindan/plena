package stream

import (
	"context"

	goRedis "github.com/redis/go-redis/v9"
)

var (
	Rdb *goRedis.Client
	Ctx = context.Background()
)

type redis struct{}

func (f redis) Init() error {
	Rdb = goRedis.NewClient(&goRedis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping to test connection
	_, err := Rdb.Ping(Ctx).Result()
	return err
}

func (f redis) Publish(string, string) error {
	return nil
}

func (f redis) Subscribe(string) error {
	return nil
}
