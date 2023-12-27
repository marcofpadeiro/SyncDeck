package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func BackupUnit(src_path, backup_path, unit_id string, version, backup_size int) error {
	if backup_path == "" || backup_size == 0 {
		return nil
	}

	compressedData, err := Compress(src_path)
	if err != nil {
		return err
	}

	unit_backup := filepath.Join(backup_path, unit_id)
	if err := os.MkdirAll(unit_backup, os.ModePerm); err != nil {
		return err
	}

	if err = manageBackupSize(unit_backup, backup_size); err != nil {
		return err
	}

	version_zip := filepath.Join(unit_backup, fmt.Sprintf("v%d.zip", version))
	backupFile, err := os.Create(version_zip)
	if err != nil {
		return err
	}
	defer backupFile.Close()

	_, err = compressedData.WriteTo(backupFile)
	if err != nil {
		return err
	}

	return nil
}

func manageBackupSize(unit_backup string, backup_size int) error {
	for versions := getBackupVersions(unit_backup); len(versions) >= backup_size; versions = getBackupVersions(unit_backup) {
		oldest_version := versions[0]
		oldest_version_path := filepath.Join(unit_backup, oldest_version)
		if err := os.RemoveAll(oldest_version_path); err != nil {
			return err
		}
	}
	return nil
}

func getBackupVersions(unitBackup string) []string {
	files, err := os.ReadDir(unitBackup)
	if err != nil {
		return nil
	}

	var versions []string
	for _, file := range files {
		versions = append(versions, file.Name())
	}

	sort.Slice(versions, func(i, j int) bool {
		numI, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(versions[i], "v"), ".zip"))
		numJ, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(versions[j], "v"), ".zip"))
		return numI < numJ
	})

	return versions
}
