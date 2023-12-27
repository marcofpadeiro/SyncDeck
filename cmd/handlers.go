package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marcofpadeiro/SyncDeck/utils"
)

func HandleSubscribe(config Config, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage:     ./syncdeck subscribe <unit> <path>")
		return
	}
	unit_id := args[0]
	path := args[1]

	remote_units, err := utils.GetRemoteUnits(config.IP, config.Port, config.Api_key)
	if err != nil {
		log.Panic(err)
	}
	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}

	index := utils.CheckExists(local_units, unit_id)
	if index != -1 {
		fmt.Println(unit_id + " already exists in your units!")
		return
	}

	index = utils.CheckExists(remote_units, unit_id)
	if index == -1 {
		fmt.Println(unit_id + " does not exist in remote units!")
		return
	}

	remote := remote_units[index]

	remote.Path = path
	remote.Version--

	err = utils.AddUnit(config.Units_metadata, remote)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Unit " + unit_id + " subscribed successfully!")
}

func HandleDel(config Config, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage:     ./syncdeck del <unit>")
		return
	}
	unit_id := args[0]

	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}
	if unit_id == "all" {
		for _, unit := range local_units {
			var user_input string
			fmt.Printf("Are you sure you want to delete all files from %s? [y/N] ", unit.Path)

			fmt.Scanln(&user_input)

			if user_input != "y" && user_input != "Y" {
				return
			}

			err = utils.DeleteUnit(config.Units_metadata, unit.ID)
			if err != nil {
				log.Panic(err)
			}
			os.RemoveAll(unit.Path)
			fmt.Println("Unit " + unit.ID + " deleted successfully!")
			return
		}
	}

	index := utils.CheckExists(local_units, unit_id)
	if index == -1 {
		fmt.Println(unit_id + " does not exist in local units!")
		return
	}

	err = utils.DeleteUnit(config.Units_metadata, unit_id)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Unit " + unit_id + " deleted successfully!")
}

func HandleRemove(config Config, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage:     ./syncdeck rm <unit>")
		return
	}
	unit_id := args[0]

	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}

	if unit_id == "all" {
		for _, unit := range local_units {
			var user_input string
			fmt.Printf("Are you sure you want to delete all files from %s? [y/N] ", unit.Path)

			fmt.Scanln(&user_input)

			if user_input != "y" && user_input != "Y" {
				return
			}

			err = utils.DeleteUnit(config.Units_metadata, unit.ID)
			if err != nil {
				log.Panic(err)
			}
			os.RemoveAll(unit.Path)
			fmt.Println("Unit " + unit.ID + " deleted successfully!")
			return
		}
	}

	index := utils.CheckExists(local_units, unit_id)
	if index == -1 {
		fmt.Println(unit_id + " does not exist in local units!")
		return
	}

	local := local_units[index]

	var user_input string
	fmt.Printf("Are you sure you want to delete all files from %s? [y/N] ", local.Path)

	fmt.Scanln(&user_input)

	if user_input != "y" && user_input != "Y" {
		return
	}

	err = utils.DeleteUnit(config.Units_metadata, unit_id)
	if err != nil {
		log.Panic(err)
	}
	os.RemoveAll(local.Path)
	fmt.Println("Unit " + unit_id + " deleted successfully!")
}

func HandleAddRemote(config Config, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage:     ./syncdeck add-remote <unit> <path>")
		return
	}
	unit_id := args[0]
	folder_path := args[1]

	remote_units, err := utils.GetRemoteUnits(config.IP, config.Port, config.Api_key)
	if err != nil {
		log.Panic(err)
	}

	if utils.CheckExists(remote_units, unit_id) != -1 {
		fmt.Println("Already exists")
		return
	}

	upload(config, unit_id, folder_path)

	utils.AddUnit(config.Units_metadata, utils.Unit{ID: unit_id, Version: 1, Path: folder_path})
	fmt.Println("Successfully added " + unit_id + " to remote")
}

func HandleEdit(config Config, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage:     ./syncdeck edit <unit> <path>")
		return
	}
	unit_id := args[0]
	path := args[1]

	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}

	index := utils.CheckExists(local_units, unit_id)
	if index == -1 {
		fmt.Println(unit_id + " does not exist in your units!")
		return
	}

	old_path := local_units[index].Path

	err = utils.EditUnit(config.Units_metadata, local_units[index], path)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Unit %s edited successfully! %s -> %s\n", unit_id, old_path, path)
}

func HandleList(config Config, args []string) {
	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}
	remote_units, err := utils.GetRemoteUnits(config.IP, config.Port, config.Api_key)
	if err != nil {
		log.Panic(err)
	}

	for _, unit := range remote_units {
		index := utils.CheckExists(local_units, unit.ID)

		var local utils.Unit
		if index != -1 {
			local = local_units[index]
		}
		fmt.Printf("%s v%d|v%d \t-> %s\n", unit.ID, local.Version, unit.Version, local.Path)
	}
}

func HandleFetch(config Config, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage:     ./syncdeck fetch <unit>")
		return
	}
	unit_id := args[0]

	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}
	remote_units, err := utils.GetRemoteUnits(config.IP, config.Port, config.Api_key)
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
		fmt.Printf("%s is outdated! (v%d->v%d)\n", local.ID, local.Version, remote.Version)
		download(config, local.ID, local.Path, local.Version)
		utils.UpdateUnit(config.Units_metadata, local, remote.Version)
	} else if local.Version > remote.Version {
		fmt.Printf("%s is ahead of remote! (v%d->v%d)\n", local.ID, local.Version, remote.Version)
		HandleUpload(config, args)
	}
	fmt.Printf("%s is up-to-date! (v%d)\n", local.ID, remote.Version)
}

func HandleUpload(config Config, args []string) {
	if len(args) > 1 {
		fmt.Println("Usage:     ./syncdeck upload <unit>")
		return
	}
	unit_id := args[0]

	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}

	if unit_id == "all" {
		for _, unit := range local_units {
			upload(config, unit.ID, unit.Path)
			version, err := utils.GetUnitVersion(config.IP, config.Port, config.Api_key, unit.ID)
			if err != nil {
				log.Panic(err.Error())
			}
			utils.UpdateUnit(config.Units_metadata, unit, version)
			fmt.Printf("%s has been uploaded successfully (v%d)\n", unit.ID, version)
		}
		return
	}

	index := utils.CheckExists(local_units, unit_id)
	if index == -1 {
		fmt.Println("You are not subscribed to that unit!")
		return
	}

	local := local_units[index]

	upload(config, unit_id, local.Path)
	version, _ := utils.GetUnitVersion(config.IP, config.Port, config.Api_key, local.ID)
	utils.UpdateUnit(config.Units_metadata, local, version)

	fmt.Printf("Successfully uploaded %s v%d to remote\n", local.ID, local.Version+1)
}

func HandleRefresh(config Config, args []string) {
	local_units, err := utils.GetUnits(config.Units_metadata)
	if err != nil {
		log.Panic(err)
	}
	remote_units, err := utils.GetRemoteUnits(config.IP, config.Port, config.Api_key)
	if err != nil {
		log.Panic(err)
	}

	for _, local := range local_units {
		index := utils.CheckExists(remote_units, local.ID)
		if index == -1 {
			log.Panic("Should exist but does not wtf")
		}
		remote := remote_units[index]
		if local.Version < remote.Version {
			fmt.Printf("%s is outdated! (v%d->v%d)   ::Updating...\n", local.ID, local.Version, remote.Version)
			download(config, local.ID, local.Path, local.Version)
			utils.UpdateUnit(config.Units_metadata, local, remote.Version)
		} else if local.Version > remote.Version {
			HandleUpload(config, []string{local.ID})
			fmt.Printf("%s is ahead of server! (v%d->v%d)   ::Uploading...\n", local.ID, local.Version, remote.Version)
		}
		fmt.Printf("%s is up to date! (v%d)\n", local.ID, local.Version)
	}
}
