package stream

import "access-key-management/pkg/utils"

type Stream interface {
	Init() error
	Publish(string, string) error
	Subscribe(string) error
}

func NewStreamer() Stream {
	if utils.STREAM_TYPE == "redis" {
		return redis{}
	}
	return nats{}
}
