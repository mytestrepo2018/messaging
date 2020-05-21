package messaging

/*
// pattern taken from github.com/febytanzil/gobroker
*/

// Option configures Sender & Receiver
type Option func(c *Config)

type Config struct {
	// serverURL saves server address to specify messaging server
	ServerURL string

	// port to use to connect to messaging server
	Port int

	// retry counts maximum retry attempts to reconnect to server
	// 0 means unlimited retry
	Retry int

	// wait number of seconds
	ReconnectWait int

	// a descriptive name
	Name string

	// ... other config values as required by technology
}

// NoopConfig configures Sender and Receiver for noop connection
func NoopConfig() Option {
	return nil
}

// NatsIOConfig configures Sender and Receiver for NATS.io connection
func NatsIOConfig(server string, name string, retry int, reconnectWait int) Option {
	return func(c *Config) {
		c.ServerURL = server
		c.Name = name
		c.Retry = retry
		c.ReconnectWait = reconnectWait
	}
}
