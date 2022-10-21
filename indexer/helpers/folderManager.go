package helpers

import (
	"encoding/json"
	"fmt"
	"indexer/api"
	"indexer/models"
	"os"
	"path"
	"path/filepath"
)

func SearchFolderStructure(folderPath string) {

	err := filepath.Walk(folderPath, WalkFunction)
	if err != nil {
		fmt.Println(err)
	}
}

func WalkFunction(path string, info os.FileInfo, err error) error {

	isFile := CheckDirectoryFileAndParse(path, info)

	if isFile {

		emailMessage, err := ParseFile(path)
		if err != nil {
			fmt.Println(path)
			fmt.Println(err)
		} else {

			emailMessageJSON, err := json.Marshal(emailMessage)

			if err != nil {
				fmt.Println("Error struct converting to JSON structure", err)
				return err
			}

			errorZincSearch := api.SendDataToZincSearch(string(emailMessageJSON))

			if errorZincSearch != nil {
				fmt.Println("Error sending data to ZincSearch", errorZincSearch)
				return err
			}
		}

	}

	return nil
}

func CheckDirectoryFileAndParse(path string, info os.FileInfo) bool {
	if info.IsDir() {
		return false
	} else {
		fmt.Println(path)
		return true
	}
}

func ParseFile(path string) (*models.EmailMessage, error) {

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("cannot able to read the file", err)
		return nil, err
	}

	msg, err := models.ParseMessage(file)
	if err != nil {
		fmt.Println("cannot parse the file", err)
		return nil, err
	}

	defer file.Close()

	return msg, nil
}

func ValidateDirectory(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func GetRunningDirectory(folderName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}

	dir = path.Join(dir, folderName)
	return dir
}
