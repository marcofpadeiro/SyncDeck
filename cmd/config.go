package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	IP             string `json:"server_ip"`
	Port           string `json:"server_port"`
	Units_metadata string `json:"unit_metadata"`
	Api_key        string
}

func ReadConfig() (Config, error) {
	var payload Config

	home, _ := os.UserHomeDir()
	config_file := filepath.Join(home, ".config/syncdeck/config.json")
	content, err := os.ReadFile(config_file)

	if err != nil {
		return payload, errors.New("Error opening config file.")
	}

	err = json.Unmarshal(content, &payload)
	if err != nil {
		return payload, errors.New("Error during Unmarshal.")
	}

	if payload.Units_metadata == "" {
		payload.Units_metadata = filepath.Join(home, ".config/syncdeck/units.json")
		json.Marshal(payload)
		os.WriteFile(config_file, content, 0644)
	}

	return payload, nil
}
