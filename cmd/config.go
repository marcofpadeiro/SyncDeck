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
	Backup_Path    string
	Backup_Size    int
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

	return payload, nil
}
