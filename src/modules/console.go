package modules

import (
	"strings"

	Core "../core"
)

// ConsoleFunction used to pass functions into console
type ConsoleFunction func(input string)

// ConsoleModule Class
type ConsoleModule struct {
	commands     map[string]ConsoleFunction
	descriptions map[string]string
	aliases      map[string]string
}

// Init  Module
func (m *ConsoleModule) Init(config *Core.Config) {
	m.commands = make(map[string]ConsoleFunction)
	m.descriptions = make(map[string]string)
	m.aliases = make(map[string]string)

	m.AddHandler("help", "This list.", m.consoleHelp)
}

// AddHandler for command
func (m *ConsoleModule) AddHandler(command string, description string, function ConsoleFunction) {
	m.commands[command] = function
	m.descriptions[command] = description
}

// AddAlias for a command
func (m *ConsoleModule) AddAlias(command string, alias string) {
	m.aliases[command] = alias
}

// RemoveHandler for command
func (m *ConsoleModule) RemoveHandler(command string) {
	delete(m.commands, command)
	delete(m.descriptions, command)

	// Iterate over aliases
	var keysToDelete []string
	for key, value := range m.aliases {
		if value == command {
			keysToDelete = append(keysToDelete, key)
		}
	}
	for _, value := range keysToDelete {
		delete(m.aliases, value)
	}
}

// Handle a command
func (m *ConsoleModule) Handle(input string) {

	splitLocation := strings.Index(input, " ")
	var command string
	var args string
	if splitLocation > 0 {
		command = input[:splitLocation]
		args = input[(splitLocation + 1):len(input)]
	} else {
		command = input
		args = ""
	}
	command = strings.ToLower(command)

	// Check Alias
	_, alias := m.aliases[command]
	if alias {
		command = m.aliases[command]
	}

	_, ok := m.commands[command]
	if ok {
		execCommand := m.commands[command]
		execCommand(args)
	} else {
		Core.Log("SYSTEM", "LOG", "Invalid command: "+command)
	}
}

func (m *ConsoleModule) consoleHelp(input string) {

	var output = "\n"
	for key, _ := range m.commands {

		if len(key) < 9 {
			output = output + key + "\t\t" + m.descriptions[key] + "\n"
		} else {
			output = output + key + "\t" + m.descriptions[key] + "\n"
		}
	}

	Core.Log("SYSTEM", "LOG", "Registered Commands\n"+output)
}
