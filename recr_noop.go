package messaging

type noopReceiver struct{}

func (noopReceiver) Name() string {
	return "noop"
}

func (noopReceiver) Init(string, ...Option) error {
	return nil
}

func (noopReceiver) Receive() ([]byte, error) {
	return nil, nil
}

func (noopReceiver) Close() {}

func init() {
	AddReceiver(&noopReceiver{})
}
