package stream

type nats struct{}

func (f nats) Init() error {
	return nil
}

func (f nats) Publish(string, string) error {
	return nil
}

func (f nats) Subscribe(string) error {
	return nil
}
