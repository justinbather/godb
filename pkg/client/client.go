package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type requestItem struct {
	Value   string `json:"value"`
	Key     string `json:"key"`
	TTL     int    `json:"ttl"`
	Sliding bool   `json:"sliding"`
}

func Get(cfg *Config, key string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s?key=%s", cfg.Address, key))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	// Need to figure out how we can give back the correct types as they are stored instead of a string
	// Maybe we take an interface as an argument then marshal the json to that interface?

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))

	return string(body), nil
}

func Set(cfg *Config, key string, val string, ttl int, sliding bool) error {
	payload := requestItem{
		Value:   val,
		Key:     key,
		TTL:     ttl,
		Sliding: sliding,
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", cfg.Address, bytes.NewBuffer(payloadJson))
	if err != nil {
		return err
	}
	log.Println("Request created")

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	log.Println("Made Request to server")

	if resp.StatusCode != http.StatusCreated {
		log.Println("Client did not get StatusOK")
	}

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
