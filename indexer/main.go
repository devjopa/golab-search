package main

import (
	"fmt"
	"indexer/helpers"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please give proper command line arguments")
		return
	}

	// Ver el tema asincrono para que sea mas rapido
	// Ver el tema de log y documentacion
	path := helpers.GetRunningDirectory(args[0])
	if helpers.ValidateDirectory(path) {
		helpers.SearchFolderStructure(path)
	} else {
		fmt.Println("The system cannot find the file specified")
	}

}
