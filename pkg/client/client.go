package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type requestItem struct {
	Value   string `json:"value"`
	Key     string `json:"key"`
	TTL     int    `json:"ttl"`
	Sliding bool   `json:"sliding"`
}

func Get(cfg *Config, key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.Timeout))
	defer cancel()

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s?key=%s", cfg.Address, key), nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
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

	return string(body), nil
}

func Set(cfg *Config, key string, val string, ttl int, sliding bool) error {
	payload := requestItem{
		Value:   val,
		Key:     key,
		TTL:     ttl,
		Sliding: sliding,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.Timeout))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cfg.Address, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return err
	}

	log.Println("Request created")

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	log.Println("Made Request to server")

	if resp.StatusCode != http.StatusCreated {
		log.Println("Client did not get StatusOK")
	}

	return nil
}

func Delete(_ *Config, _ string) error {
	return nil
}

func Connect(cfg *Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.Timeout))
	defer cancel()

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.Address, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
