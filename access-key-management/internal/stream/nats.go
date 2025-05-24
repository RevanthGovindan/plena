package stream

import "access-key-management/internal/models"

type nats struct{}

func (f *nats) init() error {
	return nil
}

func (f *nats) Publish(string, models.EventMessage) error {
	return nil
}

func (f *nats) Ping() error {
	return nil
}
