package core

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetFiles returns a list of all files under the basePath (recursively) that have the passed extensions
func GetFiles(basePath string, extensions []string) []string {

	// Create empty list
	fileList := []string{}

	// Walk our path (recursively) to find our files
	filepath.Walk(basePath, func(path string, f os.FileInfo, err error) error {

		// Make sure its not a directory, noone likes them
		if !f.IsDir() {

			// Get current files extension
			currentExtension := filepath.Ext(path)

			if StringInArray(currentExtension, extensions) {
				fileList = append(fileList, path)
			}
		}
		return nil
	})

	return fileList
}

// Log message as Jarvis
func Log(channel string, message string) {
	fmt.Println("[JARVIS]\t", "(", channel, ")\t", message)
}

// ReadLines grabs the contents of a text file, and allows conditional includes
func ReadLines(filePath string, parse func(string) (string, bool)) ([]string, error) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		Log("System", "Error opening file: "+filePath)
		return nil, err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var results []string
	for scanner.Scan() {
		if output, add := parse(scanner.Text()); add {
			results = append(results, output)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

// StringInArray checks if the target is in the list
func StringInArray(target string, list []string) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

// WriteLines outputs the lines to a file, creating the directory structure as needed
func WriteLines(lines []string, path string) error {

	// Check directory
	os.MkdirAll(filepath.Dir(path), 0755)

	// Make file
	file, err := os.Create(path)
	if err != nil {
		Log("System", "Error occured when making file "+err.Error())
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// SyncFile with data, only writing if the data is different)
func SyncFile(data []byte, path string) bool {
	buffer, error := ioutil.ReadFile(path)

	if error != nil {
		Log("System", error.Error())
	} else {
		if !bytes.Equal(buffer, data) {
			ioutil.WriteFile(path, data, 0755)
			return true
		}
	}
	return false
}

// SaveFile writes a file no matter what
func SaveFile(data []byte, path string) {
	ioutil.WriteFile(path, data, 0755)
}
