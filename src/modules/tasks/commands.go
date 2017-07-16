package tasks

import (
	Core "../../core"
)

func (m *Module) setupCommands() {
	//	m.j.Discord.RegisterCommand("!work", m.commandWork, "Work Command", Core.CommandAccessModerator, "task")
	m.j.Discord.RegisterCommand("!work", m.commandWorkingOn, "What are you doing?", Core.CommandAccessAdmin, "stats")
}

func (m *Module) commandWorkingOn(message *Core.DiscordMessage) {

	if len(message.Content) > 0 {
		if message.Content == "jira" {
			m.UseJIRAForWork = true

			if !m.jiraInstance.Polling {
				m.jiraInstance.Start()
			}

		} else {
			m.UseJIRAForWork = false
			m.jiraInstance.Stop()
			m.SetWorkingOn(message.Content, true)
		}
	}
}

// SetWorkingOn text
func (m *Module) SetWorkingOn(message string, notify bool) {
	m.data.WorkingOn = message
	Core.SaveFile([]byte(m.data.WorkingOn), m.outputs.WorkingOnPath)
	if notify {
		m.j.Discord.Announcement(m.j.Config.GetPrefix() + "Now working on " + m.data.WorkingOn)
	}
	m.j.Log.Message("Stats", "Working On: "+m.data.WorkingOn)
}

func (m *Module) SetWorkingOnJIRA(message string, notify bool) {
	if !m.UseJIRAForWork {
		return
	}

	m.SetWorkingOn(message, notify)
}
