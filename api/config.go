package main

import (
	"os"
	"strconv"
)

type Config struct {
	Save_path    string
	History_size int
	IP           string `json:"server_ip"`
	Port         string `json:"server_port"`
    Api_Key      string
}

func ReadConfig() (Config, error) {
    backup_size, _ := strconv.Atoi(os.Getenv("BACKUP_SIZE"))
    payload := Config {
        Save_path: "/syncdeck_data",
        IP: "0.0.0.0",
        Port: "8080",
        History_size: backup_size,
        Api_Key: os.Getenv("API_KEY"),
    }

	return payload, nil
}
