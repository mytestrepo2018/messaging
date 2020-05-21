package messaging

import (
	"errors"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type natsReceiver struct {
	nc *nats.Conn
	rx chan []byte
}

func (n *natsReceiver) Name() string {
	return "nats"
}

func (n *natsReceiver) Init(topic string, options ...Option) error {
	c := &Config{}
	for _, o := range options {
		o(c)
	}

	opts := []nats.Option{}
	opts = convertToNatsConfig(*c, opts)

	nc, err := nats.Connect(c.ServerURL, opts...)
	if err != nil {
		return err
	}
	n.nc = nc
	n.rx = make(chan []byte)
	n.nc.QueueSubscribe(topic, "queue", func(m *nats.Msg) {
		n.rx <- m.Data
	})
	n.nc.Flush()

	if err := nc.LastError(); err != nil {
		return err
	}

	log.Printf("Listening on [%s], queue group [%s]", topic, "queue")

	return nil
}

func (n *natsReceiver) Receive() ([]byte, error) {
	if n == nil || n.nc == nil {
		return nil, errors.New("need a initialised receiver")
	}
	data := <-n.rx
	return data, nil
}

func (n *natsReceiver) Close() {
	if n != nil && n.nc != nil {
		n.nc.Drain()
	}
}

func convertToNatsConfig(c Config, opts []nats.Option) []nats.Option {
	reconnectDelay := time.Second * time.Duration(c.ReconnectWait)

	opts = append(opts, nats.Name(c.Name))
	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(c.Retry))

	return opts
}

func init() {
	AddReceiver(&natsReceiver{})
}
