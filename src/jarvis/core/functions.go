package core

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
)

//CopyToClipboard string
func CopyToClipboard(buffer string) {
	clipboard.WriteAll(buffer)
}

// GetCommandArguements from a string, easily breaking it up
func GetCommandArguements(input string) (string, string) {
	splitLocation := strings.Index(strings.Trim(input, " "), " ")
	var command string
	var subcommand string
	if splitLocation > 0 {
		command = input[:splitLocation]
		subcommand = strings.Trim(input[(splitLocation+1):len(input)], " ")
	} else {
		command = input
	}
	return command, subcommand
}

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

// RandomString Generator
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = LetterBytes[rand.Intn(len(LetterBytes))]
	}
	return string(b)
}

// Return random string from collection
func RandomFromStrings(strings []string) string {
	return strings[rand.Intn(len(strings))]
}

func StringsContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ReadLines grabs the contents of a text file, and allows conditional includes
func ReadLines(filePath string, parse func(string) (string, bool)) ([]string, error) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		//	Log("SYSTEM", "ERROR", "Error opening file: "+filePath)
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

// Round number
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// SaveFile writes a file no matter what
func SaveFile(data []byte, path string) {
	ioutil.WriteFile(path, data, 0755)
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

// SyncFile with data, only writing if the data is different)
func SyncFile(data []byte, path string) bool {
	buffer, error := ioutil.ReadFile(path)

	if error != nil {
		//	Log("SYSTEM", "ERROR", error.Error())
	} else {
		if !bytes.Equal(buffer, data) {
			ioutil.WriteFile(path, data, 0755)
			return true
		}
	}

	return false
}

// Touch a file
func Touch(filePath string) {
	os.MkdirAll(filepath.Dir(filePath), 0755)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ioutil.WriteFile(filePath, nil, 0755)
	}
}

// WriteLines outputs the lines to a file, creating the directory structure as needed
func WriteLines(lines []string, path string) error {

	// Check directory
	os.MkdirAll(filepath.Dir(path), 0755)

	// Make file
	file, err := os.Create(path)
	if err != nil {
		//	Log("SYSTEM", "ERROR", "Error occured when making file "+err.Error())
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

// Left left-pads the string with pad up to len runes
// len may be exceeded if
func Left(str string, length int, pad string) string {
	return times(pad, length-len(str)) + str
}

// Right right-pads the string with pad up to len runes
func Right(str string, length int, pad string) string {
	return str + times(pad, length-len(str))
}

// WrapNicknameForDiscord use
func WrapNicknameForDiscord(nick string) string {
	return "<**" + nick + "**>"
}

// WrapNicknameForIRC use
func WrapNicknameForIRC(nick string) string {
	return "<" + nick + ">"
}
