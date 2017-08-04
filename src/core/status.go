package core

import "strings"

type StatusUpdator func() (int, int)

type StatusCore struct {
	ErrorCount        int
	WarningCount      int
	TotalWarningCount int
	TotalErrorCount   int
	updators          map[string]StatusUpdator
	j                 *JARVIS
}

// Initialize the Logging Module
func (m *StatusCore) Initialize(jarvisInstance *JARVIS) {

	// Create instance of Config Core
	m = new(StatusCore)

	// Assign JARVIS (circle!)
	jarvisInstance.Status = m
	m.j = jarvisInstance

	// Create default general settings
	m.updators = make(map[string]StatusUpdator)

	// Set status to green
	m.TotalWarningCount = 0
	m.TotalErrorCount = 0
	m.WarningCount = 0
	m.ErrorCount = 0

	m.j.Log.Message("Core", "Status System Initialized")
}

func (m *StatusCore) RegisterUpdator(key string, function StatusUpdator) {

	key = strings.ToLower(key)

	// Check for command
	if m.updators[key] != nil {
		m.j.Log.Warning("Core", "Duplicate status updator registration for '"+key+"', ignoring latest.")
		m.WarningCount++
		return
	}
	m.updators[key] = function
}

func (m *StatusCore) ForceUpdate() {
	m.TotalWarningCount = 0
	m.TotalErrorCount = 0
	for _, value := range m.updators {
		warnings, errors := value()
		m.TotalWarningCount += warnings
		m.TotalErrorCount += errors
	}
	m.TotalErrorCount = m.TotalErrorCount + m.ErrorCount
	m.TotalWarningCount = m.TotalWarningCount + m.WarningCount

}

func (m *StatusCore) ForceError() {
	m.TotalErrorCount++
}

func (m *StatusCore) ForceWarning() {
	m.TotalWarningCount++
}
