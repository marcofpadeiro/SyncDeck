package main

import (
	"fmt"
	"log"

	"github.com/marcofpadeiro/SyncDeck/helpers"
)

func HandleAdd(config Config, unit_id string, path string) {
	remote_units, err := helpers.GetRemoteUnits(config.IP, config.Port)
	if err != nil {
		log.Panic(err)
	}
	local_units, err := helpers.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}

	exists := helpers.CheckExists(local_units, unit_id)
	if exists != -1 {
		fmt.Println(unit_id + " already exists in your units!")
		return
	}

	exists = helpers.CheckExists(remote_units, unit_id)
	if exists == -1 {
		fmt.Println(unit_id + " does not exist in remote units!")
		return
	}

	//api call to get files

	remote_units[exists].Path = path

	err = helpers.AddUnit(config.Units_metadata, remote_units[exists])
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Unit " + unit_id + " added successfully!")
}

func HandleDel(config Config, unit_id string) {
	local_units, err := helpers.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}

	exists := helpers.CheckExists(local_units, unit_id)
	if exists == -1 {
		fmt.Println(unit_id + " does not exist in local units!")
		return
	}

	err = helpers.DeleteUnit(config.Units_metadata, unit_id)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Unit " + unit_id + " deleted successfully!")
}

func HandleAddRemote(config Config, unit_id string, folder_path string) {
	remote_units, err := helpers.GetRemoteUnits(config.IP, config.Port)
	if err != nil {
		log.Panic(err)
	}

	if helpers.CheckExists(remote_units, unit_id) != -1 {
		fmt.Println("Already exists")
		return
	}

	URL := "http://" + config.IP + ":" + config.Port + "/upload"

	zipData, err := helpers.ZipFolder(folder_path)
	if err != nil {
		log.Panic(err)
	}

	err = helpers.SendZipFile(zipData, URL, unit_id)
	if err != nil {
		log.Panic(err)
	}

	helpers.AddUnit(config.Units_metadata, helpers.Unit{ID: unit_id, Version: 1, Path: folder_path})
	fmt.Println("Successfully added " + unit_id + " to remote")
}

func HandleList(config Config) {
	units, err := helpers.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}
	remote_units, err := helpers.GetRemoteUnits(config.IP, config.Port)
	if err != nil {
		log.Panic(err)
	}

	for _, unit := range remote_units {
		exists := helpers.CheckExists(units, unit.ID)
		var local helpers.Unit
		if exists != -1 {
			local = units[exists]
		}
		fmt.Printf("%s v%d|v%d \t-> %s\n", unit.ID, local.Version, unit.Version, local.Path)
	}
}

func HandleFetch(config Config, unit_id string) {

}

func HandleUpload(config Config, unit_id string) {
	local_units, err := helpers.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}

	exists := helpers.CheckExists(local_units, unit_id)
	if exists == -1 {
		fmt.Println("You are not subscribed to that unit!")
		return
	}

	URL := "http://" + config.IP + ":" + config.Port + "/upload"

	zipData, err := helpers.ZipFolder(local_units[exists].Path)
	if err != nil {
		log.Panic(err)
	}

	err = helpers.SendZipFile(zipData, URL, unit_id)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Successfully uploaded " + unit_id + " to remote")
}
