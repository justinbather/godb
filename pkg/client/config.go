package client

type Config struct {
	Address string
}

func New(addr string) *Config {
	return &Config{Address: addr}
}
