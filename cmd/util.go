package main

import (
	"log"
	"path/filepath"

	"github.com/marcofpadeiro/SyncDeck/utils"
)

func download(config Config, unit_id, path string, version int) {
	URL := "http://" + config.IP + ":" + config.Port + "/download/" + unit_id

	zipPath := filepath.Join("/tmp", unit_id+".zip")
	err := utils.DownloadAPI(URL, zipPath, config.Api_key)
	if err != nil {
		log.Panic(err.Error())
	}

	// Unzip the downloaded file
	err = utils.BackupUnit(path, config.Backup_Path, unit_id, version, config.Backup_Size)

	if err != nil {
		log.Println(err)
	}

	err = utils.Extract(zipPath, path)
	if err != nil {
		log.Println("Error extracting file:", err)
		return
	}
}
func upload(config Config, unit_id, path string) {
	URL := "http://" + config.IP + ":" + config.Port + "/upload"

	zipData, err := utils.Compress(path)
	if err != nil {
		log.Panic(err)
	}

	err = utils.UploadAPI(zipData, URL, unit_id, config.Api_key)
	if err != nil {
		log.Panic(err)
	}
}
