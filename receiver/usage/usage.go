package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/mytestrepo2018/messaging/receiver"
)

var (
	argMessaging       string
	argMessagingConfig string
	argMessagingTopic  string

	rcvr receiver.Receiver
)

// Command line
func init() {
	flag.StringVar(&argMessaging, "msgtech", "nats", "Messaging technology to use: default 'nats'")
	flag.StringVar(&argMessagingConfig, "msgconf", "demo.nats.io", "Messaging technology config to use: disabled by default")
	flag.StringVar(&argMessagingTopic, "msgtopic", "snort", "Messaging topic to use: default 'snort'")

	flag.Parse()
}

// Logging
func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	log.Log().
		Dict("args", zerolog.Dict().
			Str("msgconf", argMessagingConfig).
			Str("msgtopic", argMessagingTopic),
		).Msg("Command args")
}

// Receiver
func init() {
	if argMessagingConfig != "" {
		var err error
		rcvr, err = receiver.GetReceiver(argMessaging)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to get a receiver")
		}

		err = rcvr.Init(argMessagingTopic, receiver.NatsIOConfig(argMessagingConfig, "A test NATS receiver", 5, 10))
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Messaging Init error")
		}
	}
}

func main() {

	// Now handle signal to terminate so we cam drain on exit.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		// Wait for signal
		<-c
		log.Printf("<caught signal - draining>")
		rcvr.Close()
		os.Exit(0)
	}()

	// Send Alerts to the messaging provider
	for rcvr != nil {
		out, err := rcvr.Receive()
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Messaging Receive error")
		}
		fmt.Printf("%s\n", string(out))
	}
}
