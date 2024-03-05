package client

import (
	"fmt"
	"net/http"
)

func Get(cfg *Config, key string) (interface{}, error) {
	_, err := http.Get(fmt.Sprintf("%s/key=%s", cfg.Address, key))
	if err != nil {
		return nil, err
	}
	// Need to figure out how we can give back the correct types as they are stored instead of a string

	return nil, nil
}

func Set(cfg *Config, key string, val interface{}) error {

	return nil
}

func Delete(cfg *Config, key string) error {

	return nil
}

// User should run this to ensure the config is correct and docker server is running, ready to be used
func Connect(cfg *Config) error {

	// Just need a simple ping to the server to ensure its running and the users config is correct
	_, err := http.Get(cfg.Address)
	if err != nil {
		return err
	}

	return nil
}
