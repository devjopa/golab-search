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

	path := helpers.GetRunningDirectory(args[0])
	//para pruebas
	//path := "D:\\Daniel\\repo-devjodapa\\golab-search\\indexer\\test"
	if helpers.ValidateDirectory(path) {
		helpers.SearchFolderStructure(path)
	} else {
		fmt.Println("The system cannot find the file specified")
	}

}
