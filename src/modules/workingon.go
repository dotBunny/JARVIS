package modules

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"io/ioutil"

	Core "../core"
)

// WorkingOnModule Class
type WorkingOnModule struct {
	TextPath    string
	Message     string
	CoffeeCount int
	SavesCount  int
	CoffeePath  string
	SavesPath   string

	config *Core.Config
}

// Init  Module
func (m *WorkingOnModule) Init(config *Core.Config, console *ConsoleModule) {

	// Assing Config
	m.config = config

	// Only do this if we are going to write files
	if m.config.WorkingOn.Output {
		m.TextPath = filepath.Join(m.config.General.OutputPath, "WorkingOn_Message.txt")
		Core.Touch(m.TextPath)
		m.CoffeePath = filepath.Join(m.config.General.OutputPath, "WorkingOn_Coffee.txt")
		Core.Touch(m.CoffeePath)
		m.SavesPath = filepath.Join(m.config.General.OutputPath, "WorkingOn_Saves.txt")
		Core.Touch(m.SavesPath)
	}

	// Load Saved WorkingOn
	savedMessage, err := ioutil.ReadFile(m.TextPath)
	if err == nil {
		m.Message = string(savedMessage)
	}

	// Load Saved Coffee
	savedCoffee, err := ioutil.ReadFile(m.CoffeePath)
	if err == nil {
		s := string(savedCoffee)
		i, err := strconv.Atoi(s)
		if err == nil {
			m.CoffeeCount = i
		} else {
			m.CoffeeCount = 0
			Core.SaveFile([]byte(fmt.Sprintf("%02d", m.CoffeeCount)), m.CoffeePath)
		}
	}
	// Load Saves
	savedSaves, err := ioutil.ReadFile(m.SavesPath)
	if err == nil {
		s := string(savedSaves)
		i, err := strconv.Atoi(s)
		if err == nil {
			m.SavesCount = i
		} else {
			m.SavesCount = 0
			Core.SaveFile([]byte(strconv.Itoa(m.SavesCount)), m.SavesPath)
		}
	}

	// Setup Endpoints
	Core.AddEndpoint("/workingon", m.endpointWorkingOn)
	Core.AddEndpoint("/coffee", m.endpointCoffee)

	// Setup Console Commands
	console.AddHandler("/coffee", "How many coffees are you on for the day?", m.consoleCoffee)
	console.AddHandler("/save", "Did someone save your ass this stream?", m.consoleSave)
	console.AddHandler("/workingon", "Set your currently working on text.", m.consoleWorkingOn)
	console.AddAlias("/w", "/workingon")
}

func (m *WorkingOnModule) consoleCoffee(input string) {
	i, err := strconv.Atoi(input)

	if err == nil {
		m.CoffeeCount = i
	} else {
		m.CoffeeCount++
	}

	if m.config.WorkingOn.Output {
		Core.SaveFile([]byte(fmt.Sprintf("%02d", m.CoffeeCount)), m.CoffeePath)
	}
	Core.Log("WORKING", "WORKING", "Coffee Count: "+strconv.Itoa(m.CoffeeCount))
}

func (m *WorkingOnModule) consoleSave(input string) {
	i, err := strconv.Atoi(input)

	if err == nil {
		m.SavesCount = i
	} else {
		m.SavesCount++
	}

	if m.config.WorkingOn.Output {
		Core.SaveFile([]byte(fmt.Sprintf("%02d", m.SavesCount)), m.SavesPath)
	}
	Core.Log("WORKING", "WORKING", "Saves Count: "+strconv.Itoa(m.SavesCount))
}

func (m *WorkingOnModule) consoleWorkingOn(input string) {
	m.Message = input
	if m.config.WorkingOn.Output {
		Core.SaveFile([]byte(input), m.TextPath)
	}
	Core.Log("WORKING", "LOG", "Set: "+input)
}

func (m *WorkingOnModule) endpointWorkingOn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.Message))
}

func (m *WorkingOnModule) endpointCoffee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.CoffeeCount))
}
