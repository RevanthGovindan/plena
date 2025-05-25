package stream

import (
	"access-key-management/internal/global"
	"access-key-management/internal/models"
	"strings"
	"sync"
)

type Stream interface {
	Init() error
	Publish(string, models.EventMessage) error
	Ping() error
}

var (
	streamer Stream
	once     sync.Once
)

func GetStreamer() Stream {
	once.Do(func() {
		if strings.EqualFold(global.Config.StreamType, "local") {
			streamer = &redis{}
		} else {
			streamer = &nats{}
		}
		streamer.Init()
	})
	return streamer
}
