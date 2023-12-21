package helpers

import (
	"encoding/json"
	"errors"
	"os"
)

type Unit struct {
	Version           int
	Game_id           string
	Last_modification string
}

func GetVersion(config Config, game_id string) (int, error) {
	var payload Unit

	content, err := os.ReadFile(config.Save_path + "/" + game_id + "/metadata.json")

	if err != nil {
		return 0, errors.New("Error opening metadata file.")
	}

	err = json.Unmarshal(content, &payload)

	if err != nil {
		return 0, errors.New("Error during Unmarshal.")
	}

	return payload.Version, nil
}

func GetUnits(config Config) ([]Unit, error) {
	var units []Unit

	dir, err := os.Open(config.Save_path)
	if err != nil {
		return units, errors.New("Error getting units.")
	}
	defer dir.Close()

	dirEntries, _ := dir.ReadDir(0)

	for _, entry := range dirEntries {
		if entry.IsDir() {
			var payload Unit

			content, err := os.ReadFile(config.Save_path + "/" + entry.Name() + "/metadata.json")
			if err != nil {
				return units, errors.New("Error opening metadata file of: " + entry.Name())
			}

			err = json.Unmarshal(content, &payload)

			if err != nil {
				return units, errors.New("Error during Unmarshal.")
			}

			units = append(units, payload)
		}
	}

	return units, nil
}
