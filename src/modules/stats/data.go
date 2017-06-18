package stats

import (
	"io/ioutil"
	"strconv"
)

// Data Structure
type Data struct {
	WorkingOn   string
	CoffeeCount int
	SavesCount  int
}

func (m *Module) setupData() {
	m.data = new(Data)

	// Default
	m.data.WorkingOn = "JARVIS"
	m.data.CoffeeCount = 0
	m.data.SavesCount = 0

	// Load WorkingOn Text
	savedWorkingOn, errorWorkingOn := ioutil.ReadFile(m.outputs.WorkingOnPath)
	if errorWorkingOn == nil {
		m.data.WorkingOn = string(savedWorkingOn)
	}

	// Load Coffee Count
	savedCoffeeCount, errorCoffeeCount := ioutil.ReadFile(m.outputs.CoffeeCountPath)
	if errorCoffeeCount == nil {
		s := string(savedCoffeeCount)
		i, err := strconv.Atoi(s)
		if err == nil {
			m.data.CoffeeCount = i
		} else {
			m.data.CoffeeCount = 0
		}
	}

	// Load Saves Count
	savedSavesCount, errorSavesCount := ioutil.ReadFile(m.outputs.SavesCountPath)
	if errorSavesCount == nil {
		s := string(savedSavesCount)
		i, err := strconv.Atoi(s)
		if err == nil {
			m.data.SavesCount = i
		} else {
			m.data.SavesCount = 0
		}
	}
}
