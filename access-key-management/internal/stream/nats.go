package stream

import "access-key-management/internal/models"

type nats struct{}

func (f *nats) Init() error {
	return nil
}

func (f *nats) Publish(string, models.EventMessage) error {
	return nil
}

func (f *nats) Subscribe(string) error {
	return nil
}
