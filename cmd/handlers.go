package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/marcofpadeiro/SyncDeck/utils"
)

func HandleAdd(config Config, unit_id string, path string) {
	remote_units, err := utils.GetRemoteUnits(config.IP, config.Port)
	if err != nil {
		log.Panic(err)
	}
	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}

	exists := utils.CheckExists(local_units, unit_id)
	if exists != -1 {
		fmt.Println(unit_id + " already exists in your units!")
		return
	}

	exists = utils.CheckExists(remote_units, unit_id)
	if exists == -1 {
		fmt.Println(unit_id + " does not exist in remote units!")
		return
	}

	URL := "http://" + config.IP + ":" + config.Port + "/download/" + unit_id

	zipPath := filepath.Join("/tmp", unit_id+".zip")
	err = utils.Download(URL, zipPath)
	if err != nil {
		log.Panic(err.Error())
	}

	// Unzip the downloaded file
	err = utils.Extract(zipPath, path)
	if err != nil {
		fmt.Println("Error extracting file:", err)
		return
	}

	remote_units[exists].Path = path

	err = utils.AddUnit(config.Units_metadata, remote_units[exists])
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Unit " + unit_id + " added successfully!")
}

func HandleDel(config Config, unit_id string) {
	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}

	exists := utils.CheckExists(local_units, unit_id)
	if exists == -1 {
		fmt.Println(unit_id + " does not exist in local units!")
		return
	}

	err = utils.DeleteUnit(config.Units_metadata, unit_id)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Unit " + unit_id + " deleted successfully!")
}

func HandleAddRemote(config Config, unit_id string, folder_path string) {
	remote_units, err := utils.GetRemoteUnits(config.IP, config.Port)
	if err != nil {
		log.Panic(err)
	}

	if utils.CheckExists(remote_units, unit_id) != -1 {
		fmt.Println("Already exists")
		return
	}

	URL := "http://" + config.IP + ":" + config.Port + "/upload"

	zipData, err := utils.Compress(folder_path)
	if err != nil {
		log.Panic(err)
	}

	err = utils.Upload(zipData, URL, unit_id)
	if err != nil {
		log.Panic(err)
	}

	utils.AddUnit(config.Units_metadata, utils.Unit{ID: unit_id, Version: 1, Path: folder_path})
	fmt.Println("Successfully added " + unit_id + " to remote")
}

func HandleList(config Config) {
	units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}
	remote_units, err := utils.GetRemoteUnits(config.IP, config.Port)
	if err != nil {
		log.Panic(err)
	}

	for _, unit := range remote_units {
		exists := utils.CheckExists(units, unit.ID)
		var local utils.Unit
		if exists != -1 {
			local = units[exists]
		}
		fmt.Printf("%s v%d|v%d \t-> %s\n", unit.ID, local.Version, unit.Version, local.Path)
	}
}

func HandleFetch(config Config, unit_id string) {
	URL := "http://" + config.IP + ":" + config.Port + "/download/" + unit_id
	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}
	remote_units, err := utils.GetRemoteUnits(config.IP, config.Port)
	if err != nil {
		log.Panic(err)
	}

	index := utils.CheckExists(local_units, unit_id)
	if index == -1 {
		fmt.Println("You are not subscribed to that " + unit_id)
		return
	}

	local := local_units[index]

	index = utils.CheckExists(remote_units, unit_id)
	if index == -1 {
		fmt.Println("Unit " + unit_id + " does not exist in remote")
		return
	}

	remote := remote_units[index]

	if local.Version < remote.Version {
		path := filepath.Join("/tmp", local.ID+".zip")
		err = utils.Download(URL, path)
		if err != nil {
			log.Panic(err.Error())
		}
		fmt.Println("Successfully downloaded to " + path)

		// Unzip the downloaded file
		os.RemoveAll(local.Path)
		err = utils.Extract(path, local.Path)
		if err != nil {
			fmt.Println("Error extracting file:", err)
			return
		}
		fmt.Println("File extracted successfully")

		utils.UpdateUnit(config.Units_metadata, local, remote.Version)
		fmt.Println("Updated metadata file")

	} else if local.Version > remote.Version {
		HandleUpload(config, unit_id)
	}
}

func HandleUpload(config Config, unit_id string) {
	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}

	exists := utils.CheckExists(local_units, unit_id)
	if exists == -1 {
		fmt.Println("You are not subscribed to that unit!")
		return
	}

	URL := "http://" + config.IP + ":" + config.Port + "/upload"

	zipData, err := utils.Compress(local_units[exists].Path)
	if err != nil {
		log.Panic(err)
	}

	err = utils.Upload(zipData, URL, unit_id)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Successfully uploaded " + unit_id + " to remote")
}
