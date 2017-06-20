package stats

import (
	"fmt"
	"io/ioutil"
	"strconv"

	Core "../../core"
)

// Data Structure
type Data struct {
	WorkingOn   string
	CoffeeCount int
	SavesCount  int
	CrashCount  int
}

func (m *Module) setupData() {
	m.data = new(Data)

	// Default
	m.data.WorkingOn = "JARVIS"
	m.data.CoffeeCount = 0
	m.data.SavesCount = 0
	m.data.CrashCount = 0

	// Load WorkingOn Text
	savedWorkingOn, errorWorkingOn := ioutil.ReadFile(m.outputs.WorkingOnPath)
	if errorWorkingOn == nil {
		m.data.WorkingOn = string(savedWorkingOn)
	} else {
		Core.SaveFile([]byte(m.data.WorkingOn), m.outputs.WorkingOnPath)
	}

	// Load Coffee Count
	savedCoffeeCount, errorCoffeeCount := ioutil.ReadFile(m.outputs.CoffeeCountPath)
	if errorCoffeeCount == nil {
		s := string(savedCoffeeCount)
		i, err := strconv.Atoi(s)
		if err == nil {
			m.data.CoffeeCount = i
		} else {
			Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.CoffeeCount), m.settings.PadCoffeeOutput, "0")), m.outputs.CoffeeCountPath)
		}
	} else {
		Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.CoffeeCount), m.settings.PadCoffeeOutput, "0")), m.outputs.CoffeeCountPath)
	}

	// Load Coffee Count
	savedCrashCount, errorCrashCount := ioutil.ReadFile(m.outputs.CrashCountPath)
	if errorCrashCount == nil {
		s := string(savedCrashCount)
		i, err := strconv.Atoi(s)
		if err == nil {
			m.data.CrashCount = i
		} else {
			Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.CrashCount), m.settings.PadCrashOutput, "0")), m.outputs.CrashCountPath)
		}
	} else {
		Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.CrashCount), m.settings.PadCrashOutput, "0")), m.outputs.CrashCountPath)
	}

	// Load Saves Count
	savedSavesCount, errorSavesCount := ioutil.ReadFile(m.outputs.SavesCountPath)
	if errorSavesCount == nil {
		s := string(savedSavesCount)
		i, err := strconv.Atoi(s)
		if err == nil {
			m.data.SavesCount = i
		} else {
			Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.SavesCount), m.settings.PadSavesOutput, "0")), m.outputs.SavesCountPath)
		}
	} else {
		Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.SavesCount), m.settings.PadSavesOutput, "0")), m.outputs.SavesCountPath)
	}
}
