package utils

import (
	"encoding/json"
	"errors"
	"os"
)

type Unit struct {
	Version int    `json:"version"`
	ID      string `json:"id"`
	Path    string `json:"path"`
}

func GetVersion(path string, id string) (int, error) {
	var payload []Unit

	content, err := os.ReadFile(path)
	if err != nil {
		return 0, errors.New("Error opening metadata file.")
	}

	err = json.Unmarshal(content, &payload)
	if err != nil {
		return 0, errors.New("Error during Unmarshal.")
	}

	for _, item := range payload {
		if item.ID == id {
			return item.Version, nil
		}
	}

	return 0, errors.New("Does not exist")
}

func GetUnits(path string) ([]Unit, error) {
	var units []Unit

	content, err := os.ReadFile(path)
	if err != nil {
		return units, errors.New("Error opening metadata file.")
	}

	err = json.Unmarshal(content, &units)
	if err != nil {
		return units, errors.New("Error during Unmarshal.")
	}

	return units, nil
}

func AddUnit(json_path string, unit Unit) error {
	units, err := UnmarshallJson(json_path)
	if err != nil {
		return err
	}

	for _, c := range units {
		if c.ID == unit.ID {
			return nil
		}
	}

	units = append(units, unit)

	err = MarshallJson(json_path, units)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUnit(json_path string, unit_id string) error {
	units, err := UnmarshallJson(json_path)
	if err != nil {
		return err
	}

	index := -1

	for i, unit := range units {
		if unit_id == unit.ID {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("Unit " + unit_id + " does not exist")
	}

	units = append(units[:index], units[index+1:]...)

	err = MarshallJson(json_path, units)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUnit(json_path string, unit Unit, version int) error {
	units, err := UnmarshallJson(json_path)
	if err != nil {
		return err
	}

	for i, u := range units {
		if u.ID == unit.ID {
			units[i].Version = version
			break
		}
	}

	err = MarshallJson(json_path, units)
	if err != nil {
		return err
	}

	return nil
}

func EditUnit(json_path string, unit Unit, path string) error {
	units, err := UnmarshallJson(json_path)
	if err != nil {
		return err
	}

	for i, u := range units {
		if u.ID == unit.ID {
			units[i].Path = path
			break
		}
	}

	err = MarshallJson(json_path, units)
	if err != nil {
		return err
	}

	return nil
}

func CheckExists(units []Unit, unit_id string) int {
	exists := -1
	for i, unit := range units {
		if unit.ID == unit_id {
			exists = i
		}
	}
	return exists
}
