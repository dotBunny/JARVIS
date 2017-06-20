package stats

import (
	"strconv"

	Core "../../core"
)

func (m *Module) setupCommands() {
	m.j.Discord.RegisterCommand("!reset", m.commandReset, "Reset stats for the day.", Core.CommandAccessAdmin)
	m.j.Discord.RegisterCommand("!coffee", m.commandCoffee, "How many coffees are you on for the day?", Core.CommandAccessAdmin)
	m.j.Discord.RegisterCommand("!save", m.commandSave, "Did someone save your ass this stream?", Core.CommandAccessModerator)
	m.j.Discord.RegisterCommand("!work", m.commandWorkingOn, "What are you doing?", Core.CommandAccessAdmin)
	m.j.Discord.RegisterCommand("!crash", m.commandCrash, "A crash happened didn't it?", Core.CommandAccessModerator)
}

func (m *Module) commandCoffee(message *Core.DiscordMessage) {
	if len(message.Content) > 0 {
		i, err := strconv.Atoi(message.Content)
		if err == nil {
			m.ChangeCoffeeCount(i, true)
		}
	} else {
		m.ChangeCoffeeCount(m.data.CoffeeCount+1, true)
	}
}

func (m *Module) commandCrash(message *Core.DiscordMessage) {
	if len(message.Content) > 0 {
		i, err := strconv.Atoi(message.Content)
		if err == nil {
			m.ChangeCrashesCount(i, true)
		}
	} else {
		m.ChangeCrashesCount(m.data.CrashCount+1, true)
	}
}

func (m *Module) commandSave(message *Core.DiscordMessage) {
	if len(message.Content) > 0 {
		i, err := strconv.Atoi(message.Content)
		if err == nil {
			m.ChangeSavesCount(i, true)
		}
	} else {
		m.ChangeSavesCount(m.data.SavesCount+1, true)
	}
}

func (m *Module) commandWorkingOn(message *Core.DiscordMessage) {

	m.data.WorkingOn = message.Content
	Core.SaveFile([]byte(m.data.WorkingOn), m.outputs.WorkingOnPath)
	m.j.Discord.Announcement(m.j.Config.GetPrefix() + "Now working on " + m.data.WorkingOn)
	m.j.Log.Message("Stats", "Working On: "+m.data.WorkingOn)
}

func (m *Module) commandReset(message *Core.DiscordMessage) {

	m.ChangeCoffeeCount(0, false)
	m.ChangeSavesCount(0, false)
	m.ChangeCrashesCount(0, false)

	m.j.Log.Message("Stats", "Stats Reset")
}
