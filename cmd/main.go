package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Panic("Error reading config")
	}

	if len(os.Args) < 2 {
		help()
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 4 {
			help()
			return
		}
		HandleAdd(config, os.Args[2], os.Args[3])
		break
	case "del":
		if len(os.Args) < 3 {
			help()
			return
		}
		HandleDel(config, os.Args[2])
		break
	case "rm":
		if len(os.Args) < 3 {
			help()
			return
		}
		HandleRemove(config, os.Args[2])
		break
	case "add-remote":
		if len(os.Args) < 4 {
			help()
			return
		}
		HandleAddRemote(config, os.Args[2], os.Args[3])
		break
	case "list":
		HandleList(config)
		break
	case "fetch":
		if len(os.Args) < 3 {
			help()
			return
		}
		HandleFetch(config, os.Args[2])
		break
	case "upload":
		if len(os.Args) < 3 {
			help()
			return
		}
		HandleUpload(config, os.Args[2])
	case "refresh":
		HandleRefresh(config)
	default:
		help()
		return
	}
}

func help() {
	fmt.Println("Usage:")
	fmt.Println("    syncdeck <command> <opts>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("    help                             print this menu")
	fmt.Println()
	fmt.Println("Unit Managment:")
	fmt.Println("    add <unit> <path>                add local unit")
	fmt.Println("    add-remote <unit> <path>         add remote unit")
	fmt.Println("    del <unit>                       remove local unit")
	fmt.Println("    rm  <unit>                       remove local unit and its files (DANGEROUS!)")
	fmt.Println("    list                             list local units")
	fmt.Println("    list-remote                      list remote units")
	fmt.Println("    fetch <unit>                     fetch unit")
	fmt.Println("    upload <unit>                    upload into server")
	fmt.Println("    refresh                          refresh all units versions")
}
