package messaging

import (
	"fmt"
        option "github.com/mytestrepo2018/messaging/option" 
)

// Receiver is a message receiver!
type Receiver interface {
	Name() string
	Init(topic string, options ...option.Option) error
	Receive() ([]byte, error)
	Close()
}

var receivers []Receiver

// AddReceiver used to register a receiver
func AddReceiver(s Receiver) {
	receivers = append(receivers, s)
}

// GetReceiver finds a receiver matching name
func GetReceiver(name string) (Receiver, error) {
	for _, r := range receivers {
		if r.Name() == name {
			return r, nil
		}
	}
	return nil, fmt.Errorf("No receiver found with name %s", name)
}
