package client

type Config struct {
	Address string
	Timeout int
}

func New(addr string, timeout int) *Config {
	return &Config{Address: addr, Timeout: timeout}
}
