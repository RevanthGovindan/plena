package stream

import (
	"access-key-management/internal/models"
	"access-key-management/pkg/utils"
	"strings"
	"sync"
)

type Stream interface {
	init() error
	Publish(string, models.EventMessage) error
	Ping() error
}

var (
	streamer Stream
	once     sync.Once
)

func GetStreamer() Stream {
	once.Do(func() {
		if strings.EqualFold(utils.DB_TYPE, "local") {
			streamer = &redis{}
		} else {
			streamer = &nats{}
		}
		streamer.init()
	})
	return streamer
}
