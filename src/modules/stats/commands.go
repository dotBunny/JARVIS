package stats

import (
	"fmt"
	"strconv"

	Core "../../core"
)

func (m *Module) setupCommands() {
	m.j.Discord.RegisterCommand("!coffee", m.commandCoffee, "How many coffees are you on for the day?", Core.CommandAccessPrivate)
	m.j.Discord.RegisterCommand("!save", m.commandSave, "Did someone save your ass this stream?", Core.CommandAccessPrivate)
	m.j.Discord.RegisterCommand("!work", m.commandWorkingOn, "What are you doing?", Core.CommandAccessPrivate)
	m.j.Discord.RegisterCommand("!crash", m.commandCrash, "A crash happened didn't it?", Core.CommandAccessPrivate)
}

func (m *Module) commandCoffee(message *Core.DiscordMessage) {

	if len(message.Content) > 0 {

		i, err := strconv.Atoi(message.Content)
		if err == nil {
			m.data.CoffeeCount = i
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "Coffees set to " + fmt.Sprintf("%d", m.data.CoffeeCount) + "!")
		}
	} else {
		m.data.CoffeeCount++

		if m.data.CoffeeCount == 1 {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "We are on the first cup of coffee for the day! Watch out!")
		} else {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "Coffee #" + fmt.Sprintf("%d", m.data.CoffeeCount) + "!")
		}

		// Record in Log

	}

	m.j.Log.Message("Stats", "Coffee Recorded ("+fmt.Sprintf("%d", m.data.CoffeeCount)+")")
	Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.CoffeeCount), m.settings.PadCoffeeOutput, "0")), m.outputs.CoffeeCountPath)
}

func (m *Module) commandCrash(message *Core.DiscordMessage) {

	if len(message.Content) > 0 {
		i, err := strconv.Atoi(message.Content)
		if err == nil {
			m.data.CrashCount = i
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "Crashes set to " + fmt.Sprintf("%d", m.data.SavesCount) + "!")
		}
	} else {
		m.data.CrashCount++

		if m.data.CrashCount == 1 {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "Our first crash of the day :(")
		} else {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "CRASHED! (and or burned!) - That's number " + fmt.Sprintf("%d", m.data.CrashCount) + " of the day.")
		}

		// Record in Log

	}
	m.j.Log.Message("Stats", "Crash Recorded ("+fmt.Sprintf("%d", m.data.CrashCount)+")")
	Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.CrashCount), m.settings.PadCrashOutput, "0")), m.outputs.CrashCountPath)
}
func (m *Module) commandSave(message *Core.DiscordMessage) {

	if len(message.Content) > 0 {
		i, err := strconv.Atoi(message.Content)
		if err == nil {
			m.data.SavesCount = i
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "Saves set to " + fmt.Sprintf("%d", m.data.SavesCount) + "!")
		}
	} else {
		m.data.SavesCount++

		if m.data.SavesCount == 1 {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "The first save of the day!")
		} else {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "SAVED!!! We are up to " + fmt.Sprintf("%d", m.data.SavesCount) + "!")
		}

		// Record in Log

	}
	m.j.Log.Message("Stats", "Save Recorded ("+fmt.Sprintf("%d", m.data.SavesCount)+")")
	Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.SavesCount), m.settings.PadSavesOutput, "0")), m.outputs.SavesCountPath)
}

func (m *Module) commandWorkingOn(message *Core.DiscordMessage) {

	m.data.WorkingOn = message.Content

	Core.SaveFile([]byte(m.data.WorkingOn), m.outputs.WorkingOnPath)
	m.j.Discord.Announcement(m.j.Config.GetPrefix() + "Now working on " + m.data.WorkingOn)
	m.j.Log.Message("Stats", "Working On: "+m.data.WorkingOn)
}
