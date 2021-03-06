package tasks

import (
	"io/ioutil"

	Core "../../core"
)

// Data Structure
type Data struct {
	WorkingOn string
}

func (m *Module) GetWorkingOn() string {
	return m.data.WorkingOn
}

// SetWorkingOn text
func (m *Module) SetWorkingOn(message string, notify bool) {

	if len(message) <= 0 {
		return
	}

	m.data.WorkingOn = message
	Core.SaveFile([]byte(m.data.WorkingOn), m.outputs.WorkingOnPath)

	if notify {
		if m.UseJIRAForWork {
			var message Core.NotifyMessage
			message.Discord = true
			message.DiscordPrefix = m.jiraInstance.GetPrefix()
			message.Twitch = true
			message.Message = "Now working on " + m.data.WorkingOn
		} else {
			var message Core.NotifyMessage
			message.Discord = true
			message.DiscordPrefix = m.j.Config.GetPrefix()
			message.Twitch = true
			message.Message = "Now working on " + m.data.WorkingOn
			m.j.Notify.Announce(message)
		}
	}
	m.j.Log.Message("Stats", "Working on: "+m.data.WorkingOn)
}

func (m *Module) setupData() {

	m.data = new(Data)

	// Default
	m.data.WorkingOn = "JARVIS"

	// Load WorkingOn Text
	savedWorkingOn, errorWorkingOn := ioutil.ReadFile(m.outputs.WorkingOnPath)
	if errorWorkingOn == nil {
		m.data.WorkingOn = string(savedWorkingOn)
	} else {
		Core.SaveFile([]byte(m.data.WorkingOn), m.outputs.WorkingOnPath)
	}

	// Load Issues
}
