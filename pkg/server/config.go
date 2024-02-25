package server

import (
	"flag"
)

const (
	defaultPort    = "8080"
	defaultWriteTO = 10
	defaultReadTO  = 10
)

func ConfigFromFlags() *Config {
	config := &Config{}

	flag.StringVar(&config.Port, "port", defaultPort, "Server port")
	flag.IntVar(&config.ReadTimeout, "read", defaultReadTO, "Server read timeout in seconds")
	flag.IntVar(&config.WriteTimeout, "write", defaultWriteTO, "Server write timeout in seconds")
	flag.Parse()

	return config
}
