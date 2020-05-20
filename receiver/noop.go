package messaging

import option "github.com/mytestrepo2018/messaging/option"

type noopReceiver struct{}

func (noopReceiver) Name() string {
	return "noop"
}

func (noopReceiver) Init(string, ...option.Option) error {
	return nil
}

func (noopReceiver) Receive() ([]byte, error) {
	return nil, nil
}

func (noopReceiver) Close() {}

func init() {
	AddReceiver(&noopReceiver{})
}
