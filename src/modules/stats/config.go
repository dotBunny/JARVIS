package stats

import (
	"encoding/json"
)

// Config elements
type Config struct {
	Definitions []Stat
}
type Stat struct {
	Key        string `json:"Key"`
	Value      int    `json:"Value"`
	TextOutput struct {
		Enabled bool   `json:"Enabled"`
		Prefix  string `json:"Prefix"`
		Suffix  string `json:"Suffix"`
	} `json:"TextOutput"`
	NumericalOutput struct {
		Enabled bool `json:"Enabled"`
		Padding int  `json:"Padding"`
	} `json:"NumericalOutput"`
	Increase struct {
		Command            string   `json:"Command"`
		CommandDescription string   `json:"CommandDescription"`
		CommandLevel       int      `json:"CommandLevel"`
		Notify             bool     `json:"Notify"`
		NotifyCallback     []string `json:"NotifyCallback"`
		NotifyMessage      string   `json:"NotifyMessage"`
		Sounds             []string `json:"Sounds"`
	} `json:"Increase"`
	Decrease struct {
		Command            string   `json:"Command"`
		CommandDescription string   `json:"CommandDescription"`
		CommandLevel       int      `json:"CommandLevel"`
		Notify             bool     `json:"Notify"`
		NotifyCallback     []string `json:"NotifyCallback"`
		NotifyMessage      string   `json:"NotifyMessage"`
		Sounds             []string `json:"Sounds"`
	} `json:"Decrease"`
	Set struct {
		Command            string   `json:"Command"`
		CommandDescription string   `json:"CommandDescription"`
		CommandLevel       int      `json:"CommandLevel"`
		Notify             bool     `json:"Notify"`
		NotifyCallback     []string `json:"NotifyCallback"`
		NotifyMessage      string   `json:"NotifyMessage"`
		Sounds             []string `json:"Sounds"`
	} `json:"Set"`
	Dashboard struct {
		ForegroundColor string `json:"ForegroundColor"`
		BackgroundColor string `json:"BackgroundColor"`
		IconClass       string `json:"IconClass"`
		Description     string `json:"Description"`
	} `json:"Dashboard"`
}

// Initialize the Logging Module
func (m *Module) loadConfig() {

	// Create default general settings
	m.settings = new(Config)

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("Stats") {
			m.j.Log.Message("Config", "Unable to find \"Stats\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("Stats"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Stats Config, somethings may be wonky.")
			}
		}
	}
}

type DashboardCounterDefinition struct {
	ID              string `json:"ID"`
	ForegroundColor string `json:"ForegroundColor"`
	BackgroundColor string `json:"BackgroundColor"`
	IconClass       string `json:"IconClass"`
	Description     string `json:"Description"`
}
