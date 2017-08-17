package tasks

import (
	Core "../../core"
)

func (m *Module) setupCommands() {
	//	m.j.Discord.RegisterCommand("!work", m.commandWork, "Work Command", Core.CommandAccessModerator, "task")
	m.j.Discord.RegisterCommand("!work", m.commandWorkingOn, "What are you doing?", Core.CommandAccessAdmin, "task")
	m.j.Discord.RegisterCommand("!jira", m.commandJIRA, "User JIRA instead of the work system", Core.CommandAccessAdmin, "task")
}

func (m *Module) commandJIRA(message *Core.DiscordMessage) {

	if !m.UseJIRAForWork {
		// Turn on JIRA
		m.UseJIRAForWork = true
		m.SetWorkingOn("Loading JIRA Task ...", false)

		// Force POLL (with notify)
		m.jiraInstance.Poll(true)

		if !m.jiraInstance.Polling {
			m.jiraInstance.Start()
		}
	}
}

func (m *Module) commandWorkingOn(message *Core.DiscordMessage) {

	if m.UseJIRAForWork || m.jiraInstance.Polling {
		m.UseJIRAForWork = false
		m.jiraInstance.Stop()
	}

	m.SetWorkingOn(message.Content, true)
}
