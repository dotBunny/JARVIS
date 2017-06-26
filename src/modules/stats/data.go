package stats

import (
	"fmt"
	"io/ioutil"
	"math/rand"
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

// ChangeCoffeeCount to specific value
func (m *Module) ChangeCoffeeCount(value int, notify bool) {

	// Check action
	if value > m.data.CoffeeCount && len(m.settings.CoffeeSounds) > 0 {
		// yup increase play a sound
		m.j.Media.PlaySound(m.settings.CoffeeSounds[rand.Intn(len(m.settings.CoffeeSounds))])
	}

	// Set Value
	m.data.CoffeeCount = value

	// Save File
	Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.CoffeeCount), m.settings.PadCoffeeOutput, "0")), m.outputs.CoffeeCountPath)

	if notify {
		if m.data.CoffeeCount == 1 {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "We are on the first cup of coffee for the day! Watch out!")
		} else {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "Coffee #" + fmt.Sprintf("%d", m.data.CoffeeCount) + "!")
		}
	}

	// Log Change
	m.j.Log.Message("Stats", "Coffee Count set to "+fmt.Sprintf("%d", m.data.CoffeeCount))
}

// ChangeCrashCount to specific value
func (m *Module) ChangeCrashesCount(value int, notify bool) {

	// Check action
	if value > m.data.CrashCount && len(m.settings.CrashSounds) > 0 {
		// yup increase play a sound
		m.j.Media.PlaySound(m.settings.CrashSounds[rand.Intn(len(m.settings.CrashSounds))])
	}

	// Set Value
	m.data.CrashCount = value

	// Save File
	Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.CrashCount), m.settings.PadCrashOutput, "0")), m.outputs.CrashCountPath)

	if notify {
		if m.data.CrashCount == 1 {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "Our first crash of the day :(")
		} else {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "CRASHED! (and or burned!) - That's number " + fmt.Sprintf("%d", m.data.CrashCount) + " of the day.")
		}
	}

	// Log Change
	m.j.Log.Message("Stats", "Crash Count set to "+fmt.Sprintf("%d", m.data.CrashCount))
}

// ChangeSavesCount to specific value
func (m *Module) ChangeSavesCount(value int, notify bool) {

	// Check action
	if value > m.data.SavesCount && len(m.settings.SaveSounds) > 0 {
		// yup increase play a sound
		m.j.Media.PlaySound(m.settings.SaveSounds[rand.Intn(len(m.settings.SaveSounds))])
	}

	// Set Value
	m.data.SavesCount = value

	// Save File
	Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.SavesCount), m.settings.PadSavesOutput, "0")), m.outputs.SavesCountPath)

	if notify {
		if m.data.SavesCount == 1 {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "The first save of the day!")
		} else {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "SAVED!!! We are up to " + fmt.Sprintf("%d", m.data.SavesCount) + "!")
		}
	}

	// Log Change
	m.j.Log.Message("Stats", "Save Count set to "+fmt.Sprintf("%d", m.data.SavesCount))
}
