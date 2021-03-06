package stats

import (
	"strconv"

	Core "../../core"
)

func (m *Module) setupCommands() {

	for _, definition := range m.stats {
		if len(definition.Increase.Command) > 0 {
			m.j.Discord.RegisterCommand(definition.Increase.Command, m.commandHandler, definition.Increase.CommandDescription, definition.Increase.CommandLevel, definition.Key)
		}

		if len(definition.Decrease.Command) > 0 {
			m.j.Discord.RegisterCommand(definition.Decrease.Command, m.commandHandler, definition.Decrease.CommandDescription, definition.Decrease.CommandLevel, definition.Key)
		}

		if len(definition.Set.Command) > 0 {
			m.j.Discord.RegisterCommand(definition.Set.Command, m.commandHandler, definition.Set.CommandDescription, definition.Set.CommandLevel, definition.Key)
		}
	}

	m.j.Discord.RegisterCommand("!stats.reset", m.commandReset, "Reset stats for the day.", Core.CommandAccessAdmin, "stats")

}

func (m *Module) commandHandler(message *Core.DiscordMessage) {
	if len(message.Content) > 0 {
		i, err := strconv.Atoi(message.Content)
		if err == nil {
			m.ChangeData(message.CommandKey, i, true)
		}
	} else if m.stats[message.CommandKey].Increase.Command == message.Command {
		m.ChangeData(message.CommandKey, m.stats[message.CommandKey].Value+1, true)
	} else if m.stats[message.CommandKey].Decrease.Command == message.Command {
		m.ChangeData(message.CommandKey, m.stats[message.CommandKey].Value-1, true)
	}
}

func (m *Module) commandReset(message *Core.DiscordMessage) {

	for _, definition := range m.stats {
		m.ChangeData(definition.Key, 0, false)
	}
	m.j.Log.Message("Stats", "Stats Reset")
}
