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
	CoffeePath  string

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
		Core.Touch(m.TextPath)
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
		if err != nil {
			Core.Log("WORKING", "WARNING", err.Error())
		} else {
			m.CoffeeCount = i
		}
	}

	// Setup Endpoints
	Core.AddEndpoint("/workingon", m.endpointWorkingOn)
	Core.AddEndpoint("/coffee", m.endpointCoffee)

	// Setup Console Commands
	console.AddHandler("/coffee", "How many coffees are you on for the day?.", m.consoleCoffee)
	console.AddHandler("/workingon", "Set your currently working on text.", m.consoleWorkingOn)
	console.AddAlias("/w", "/workingon")
}

func (m *WorkingOnModule) consoleCoffee(input string) {

	command, count := Core.GetCommandArguements(input)
	if command == "set" {
		i, err := strconv.Atoi(count)
		if err != nil {
			m.CoffeeCount = i
		}
	} else {
		// Increment
		m.CoffeeCount++
	}

	if m.config.WorkingOn.Output {
		Core.SaveFile([]byte(strconv.Itoa(m.CoffeeCount)), m.CoffeePath)
	}
	Core.Log("WORKING", "WORKING", "Coffee Count: "+strconv.Itoa(m.CoffeeCount))
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
