package helpers

import (
	"encoding/json"
	"errors"
	"os"
)

func UnmarshallJson(filepath string) ([]Unit, error) {
	var units []Unit

	content, err := os.ReadFile(filepath)
	if err != nil {
		return units, errors.New("Error opening json file." + err.Error())
	}

	err = json.Unmarshal(content, &units)
	if err != nil {
		return units, errors.New("Error during Unmarshal." + err.Error())
	}
	return units, nil
}

func MarshallJson(filepath string, units []Unit) error {
	content, err := json.Marshal(units)
	if err != nil {
		return errors.New("Error during Marshal." + err.Error())
	}

	err = os.WriteFile(filepath, content, 0644)
	if err != nil {
		return errors.New("Error writing JSON file:" + err.Error())
	}
	return nil
}
