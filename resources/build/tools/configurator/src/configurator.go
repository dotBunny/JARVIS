package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var (
	overridePath     string
	configFolderPath string
	dataSource       map[string]string
)

func main() {

	if len(os.Args) < 3 {
		log.Println("Argument 1: Path to override file\nArguement 2: Path to config files.")
		return
	}

	dataSource = make(map[string]string)

	overridePath = os.Args[1]
	configFolderPath = os.Args[2]
	var errorCheck error

	// Read Data File
	// Check existence
	_, errorCheck = os.Stat(overridePath)
	if errorCheck != nil {
		log.Println("Unable to access override file: ", overridePath+"\n"+errorCheck.Error())
		return
	}

	// Grab Raw Data
	file, err := os.Open(overridePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var split = strings.Split(scanner.Text(), ",")
		dataSource[split[0]] = split[1]
	}

	_, errorCheck = os.Stat(configFolderPath)
	if errorCheck != nil {
		log.Println("Unable to access config folder: ", configFolderPath+"\n"+errorCheck.Error())
		return
	}

	files, _ := ioutil.ReadDir(configFolderPath)
	for _, f := range files {

		currentPath := path.Join(configFolderPath, f.Name())
		data, _ := ioutil.ReadFile(currentPath)
		var fileData = string(data)
		for key, value := range dataSource {
			fileData = strings.Replace(fileData, key, value, -1)
		}
		ioutil.WriteFile(currentPath, []byte(fileData), 0644)
	}
}
