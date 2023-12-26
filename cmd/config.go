package main

import (
	"encoding/json"
	"errors"
	"os"
)

const CONFIG_PATH = "../configs/client.json"

type Config struct {
	IP             string `json:"server_ip"`
	Port           string `json:"server_port"`
	Units_metadata string `json:"unit_metadata"`
}

func ReadConfig() (Config, error) {
	var payload Config

	content, err := os.ReadFile(CONFIG_PATH)

	if err != nil {
		return payload, errors.New("Error opening config file.")
	}

	err = json.Unmarshal(content, &payload)
	if err != nil {
		return payload, errors.New("Error during Unmarshal.")
	}

	return payload, nil
}
