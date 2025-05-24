package stream

import (
	"strings"
	"sync"
	"web3-tokeninfo/pkg/utils"
)

type Stream interface {
	init() error
	Ping() error
	Subscribe(string, func(msg string))
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
