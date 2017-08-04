package core

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"io/ioutil"
	"path"
)

// GeneralConfig Settings
type GeneralConfig struct {
	OutputPath string
	Prefix     string
}

// ConfigCore holds general configuration information
type ConfigCore struct {
	settings   *GeneralConfig
	configPath string

	dataSource  map[string]json.RawMessage
	initialized bool
	j           *JARVIS
	rawSource   []byte
}

// GetConfigData retreives the sub data of a key in the config
func (m *ConfigCore) GetConfigData(key string) json.RawMessage {
	return m.dataSource[key]
}

// GetOutputPath base to use for files
func (m *ConfigCore) GetOutputPath() string {
	return m.settings.OutputPath
}

// GetPrefix for Discord
func (m *ConfigCore) GetPrefix() string {
	return m.settings.Prefix
}

// LoadConfig from config path
func (m *ConfigCore) LoadConfig(fileName string, key string) bool {

	var data []byte
	var errorCheck error

	// Create absolute path to config file
	var absolutePath = path.Join(m.j.configPath, fileName)

	// Check existence
	_, errorCheck = os.Stat(absolutePath)
	if errorCheck != nil {
		log.Println("[Config]\tConfig file is missing: ", absolutePath+"\n"+errorCheck.Error())
		return false
	}
	// Grab Raw Data
	errorCheck = nil
	data, errorCheck = ioutil.ReadFile(absolutePath)
	if errorCheck != nil {
		log.Println("[Config]\tError reading config at: " + absolutePath + "\n" + errorCheck.Error())
		return false
	}

	var dataParsed json.RawMessage
	errorCheck = json.Unmarshal(data, &dataParsed)
	if errorCheck != nil {
		log.Println("[Config]\tSomething went wrong when trying to break down the JSON in the config file: " + m.j.configPath + "\n" + errorCheck.Error())
	}
	m.dataSource[key] = dataParsed

	return true
}

// Initialize the Logging Module
func (m *ConfigCore) Initialize(jarvisInstance *JARVIS) {

	// Create instance of Config Core
	m = new(ConfigCore)

	// Assign JARVIS (circle!)
	jarvisInstance.Config = m
	m.j = jarvisInstance

	// Create Cache
	m.dataSource = make(map[string]json.RawMessage)

	var errorCheck error

	// Load General Settings
	m.LoadConfig("general.json", "General")

	// Create default general settings
	m.settings = new(GeneralConfig)

	// General Config Defaults
	m.settings.OutputPath = path.Join(m.j.GetResourcePath(), "output")
	m.settings.Prefix = ":jarvis: "

	// Check Raw Data
	if m.dataSource["General"] == nil {
		log.Println("[Config] Unable to find \"General\" config file. Using defaults.")
	} else {
		errorCheck = nil
		errorCheck = json.Unmarshal(m.dataSource["General"], &m.settings)
		if errorCheck != nil {
			log.Println("[Config]\tUnable to properly parse General Config, somethings may be wonky.")
			log.Println("[Config]\tGeneral.OutputPath: " + m.settings.OutputPath)
		}
	}

	if m.settings.OutputPath == "<Absolute Path To Where To Store Files>" {
		m.settings.OutputPath = path.Join(m.j.GetResourcePath(), "output")
	}

	// Make sure our output path base is good and ready
	os.MkdirAll(filepath.Dir(m.GetOutputPath()), 0755)

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
