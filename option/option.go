package messaging

/*
// pattern taken from github.com/febytanzil/gobroker
*/

// Option configures Sender & Receiver
type Option func(c *config)

type config struct {
	// serverURL saves server address to specify messaging server
	serverURL string

	// port to use to connect to messaging server
	port int

	// retry counts maximum retry attempts to reconnect to server
	// 0 means unlimited retry
	retry int

	// wait number of seconds
	reconnectWait int

	// a descriptive name
	name string

	// ... other config values as required by technology
}

// NoopConfig configures Sender and Receiver for noop connection
func NoopConfig() Option {
	return nil
}

// NatsIOConfig configures Sender and Receiver for NATS.io connection
func NatsIOConfig(server string, name string, retry int, reconnectWait int) Option {
	return func(c *config) {
		c.serverURL = server
		c.name = name
		c.retry = retry
		c.reconnectWait = reconnectWait
	}
}
