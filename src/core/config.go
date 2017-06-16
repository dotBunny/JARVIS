package core

import (
	"encoding/json"
	"log"
	"os"

	"io/ioutil"
	"path"
)

// GeneralConfig Settings
type GeneralConfig struct {
	Mode       string
	OutputPath string
}

// ConfigCore holds general configuration information
type ConfigCore struct {
	Settings *GeneralConfig

	dataSource  map[string]*json.RawMessage
	initialized bool
	j           *JARVIS
	rawSource   []byte
}

// GetConfigData retreives the sub data of a key in the config
func (m *ConfigCore) GetConfigData(key string) *json.RawMessage {
	return m.dataSource[key]
}

// Initialize the Logging Module
func (m *ConfigCore) Initialize(jarvisInstance *JARVIS) {

	// Create instance of Config Core
	m = new(ConfigCore)

	// Assign JARVIS (circle!)
	jarvisInstance.Config = m
	m.j = jarvisInstance

	var errorCheck error

	// Check for config's existence
	errorCheck = nil
	_, errorCheck = os.Stat(m.j.configPath)
	if errorCheck != nil {
		log.Println("[Config]\tConfig file is missing: ", m.j.configPath)
	}

	// Grab Raw Data
	errorCheck = nil
	m.rawSource, errorCheck = ioutil.ReadFile(m.j.configPath)
	if errorCheck != nil {
		log.Println("[Config]\tError reading config at: " + m.j.configPath + "\n" + errorCheck.Error())
	}

	// Load into pseudo map
	errorCheck = nil
	errorCheck = json.Unmarshal(m.rawSource, &m.dataSource)
	if errorCheck != nil {
		log.Println("[Config]\tSomething went wrong when trying to break down the JSON in the config file: " + m.j.configPath + "\n" + errorCheck.Error())
	}

	// Create default general settings
	m.Settings = new(GeneralConfig)

	// General Config
	m.Settings.Mode = "bot"
	m.Settings.OutputPath = path.Join(m.j.GetApplicationPath(), "output")

	// Check Raw Data
	if m.dataSource["General"] == nil {
		log.Println("[Config] Unable to find \"General\" config section. Using defaults.")
	} else {
		errorCheck = nil
		errorCheck = json.Unmarshal(*m.dataSource["General"], &m.Settings)
		if errorCheck != nil {
			log.Println("[Config]\tUnable to properly parse General Config, somethings may be wonky.")
			log.Println("[Config]\tGeneral.Mode: " + m.Settings.Mode)
			log.Println("[Config]\tGeneral.OutputPath: " + m.Settings.OutputPath)
		}
	}

	// Flag class as loaded
	m.initialized = true
	log.Println("[Config]\tInitialized")
}

// IsInitialized yet?
func (m *ConfigCore) IsInitialized() bool {
	return m.initialized
}

// IsValidKey in root of config
func (m *ConfigCore) IsValidKey(key string) bool {
	if m.dataSource[key] == nil {
		return false
	}
	return true
}
