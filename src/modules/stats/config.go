package stats

import (
	"encoding/json"
)

// Config elements
type Config struct {
	CoffeeSounds    []string
	CrashSounds     []string
	SaveSounds      []string
	BuildSounds     []string
	PadCoffee       int
	PadCrash        int
	PadSaves        int
	PadCoffeeOutput int
	PadCrashOutput  int
	PadSavesOutput  int
}

// Initialize the Logging Module
func (m *Module) loadConfig() {
	// Create default general settings
	m.settings = new(Config)

	m.settings.PadCoffee = 2
	m.settings.PadSaves = 2
	m.settings.PadCrash = 2
	m.settings.PadCoffeeOutput = 0
	m.settings.PadSavesOutput = 0
	m.settings.PadCrashOutput = 0

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("Stats") {
			m.j.Log.Message("Stats", "Unable to find \"Stats\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("Stats"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Stats Config, somethings may be wonky.")
			}
		}
	}

	// TODO : VALIDATE SOUND EXISTS
}
