package stats

import (
	"fmt"
	"strconv"

	Core "../../core"
)

func (m *Module) setupCommands() {
	m.j.Discord.RegisterCommand("!coffee", m.commandCoffee, "How many coffees are you on for the day?", Core.CommandAccessPrivate)
	m.j.Discord.RegisterCommand("!save", m.commandSave, "Did someone save your ass this stream?", Core.CommandAccessPrivate)
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
		m.j.Log.Message("Stats", "Coffee Recorded ("+fmt.Sprintf("%d", m.data.CoffeeCount)+")")
	}

	Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.CoffeeCount), m.settings.PadCoffeeOutput, "0")), m.outputs.CoffeeCountPath)
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
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "A save has been registered! The first of the day!")
		} else {
			m.j.Discord.Announcement(m.j.Config.GetPrefix() + "A save has been registered! We are up to " + fmt.Sprintf("%d", m.data.SavesCount) + "!")
		}

		// Record in Log
		m.j.Log.Message("Stats", "Save Recorded ("+fmt.Sprintf("%d", m.data.SavesCount)+")")
	}

	Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.SavesCount), m.settings.PadSavesOutput, "0")), m.outputs.SavesCountPath)
}

// TODO : Command WorkingOn

// 	console.AddHandler("/workingon", "Set your currently working on text.", m.consoleWorkingOn)
// 	console.AddAlias("/w", "/workingon")
// }

// func (m *WorkingOnModule) consoleWorkingOn(input string) {
// 	m.Message = input
// 	if m.config.WorkingOn.Output {
// 		Core.SaveFile([]byte(input), m.TextPath)
// 	}
// 	Core.Log("WORKING", "LOG", "Set: "+input)
// }
