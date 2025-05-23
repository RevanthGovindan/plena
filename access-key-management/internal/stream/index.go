package stream

import (
	"access-key-management/internal/models"
	"access-key-management/pkg/utils"
)

type Stream interface {
	Init() error
	Publish(string, models.EventMessage) error
	Subscribe(string) error
}

var streamer Stream

func InitalizeStreamer() error {
	var err error
	if utils.STREAM_TYPE == "redis" {
		streamer = &redis{}
	} else {
		streamer = &nats{}
	}
	err = streamer.Init()
	return err
}

func GetStreamer() Stream {
	return streamer
}
