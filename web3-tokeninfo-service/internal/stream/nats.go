package stream

type nats struct{}

func (f *nats) init() error {
	return nil
}

func (f *nats) Ping() error {
	return nil
}

func (f *nats) Subscribe(string, func(msg string)) {

}
