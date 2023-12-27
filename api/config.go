package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	Save_path    string
	History_size int
	IP           string `json:"server_ip"`
	Port         string `json:"server_port"`
    Api_Key      string
}

func ReadConfig() (Config, error) {
	var payload Config

	home, _ := os.UserHomeDir()
	content, err := os.ReadFile(filepath.Join(home, ".config/syncdeck/server.json"))

	if err != nil {
		return payload, errors.New("Error opening config file.")
	}

	os.MkdirAll(payload.Save_path, os.ModePerm)

	err = json.Unmarshal(content, &payload)
	if err != nil {
		return payload, errors.New("Error during Unmarshal.")
	}

	return payload, nil
}
