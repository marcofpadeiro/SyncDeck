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
		}
		HandleAdd(config, os.Args[2], os.Args[3])
	case "del":
		if len(os.Args) < 3 {
			help()
		}
		HandleDel(config, os.Args[2])
	case "add-remote":
		if len(os.Args) < 4 {
			help()
		}
		HandleAddRemote(config, os.Args[2], os.Args[3])
	case "list":
		HandleList(config)
	case "fetch":
	case "upload":
	default:
		help()
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
	fmt.Println("    del <unit>                       remove local unit")
	fmt.Println("    add-remote <unit> <path>         add remote unit")
	fmt.Println("    list                             list local units")
	fmt.Println("    list-remote                      list remote units")
	fmt.Println("    fetch                            fetch all units")
	fmt.Println("    fetch <unit>                     fetch unit")
	fmt.Println("    upload <unit>                    upload into server")
}
