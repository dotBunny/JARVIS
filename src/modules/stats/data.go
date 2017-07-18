package stats

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"

	Core "../../core"
)

func (m *Module) setupData() {

	// Initialize map
	m.stats = make(map[string]Stat)
	m.dashboardDefinitions = nil

	// Take definitions and put them into our map
	for _, definition := range m.settings.Definitions {

		var ourStat Stat
		ourStat.Decrease.Sounds = definition.Decrease.Sounds

		// Assign it so we can immediately reference it
		m.stats[definition.Key] = definition

		// Get our working path
		var workingPath = m.GetOutputPath(definition.Key, "Count")

		// We check usage here, but normally it happens on the save
		if definition.NumericalOutput.Enabled {
			Core.Touch(workingPath)
		}
		if definition.TextOutput.Enabled {
			Core.Touch(m.GetOutputPath(definition.Key, "Text"))
		}

		// Load Existing Data
		loadedValue, errorLoading := ioutil.ReadFile(workingPath)
		if errorLoading == nil {
			s := string(loadedValue)
			i, err := strconv.Atoi(s)
			if err == nil {
				m.j.Log.Message("STATS", "Loaded value "+s+" for \""+definition.Key+"\"")

				// Go Issue (fix is coming in Go2)
				var tmp = m.stats[definition.Key]
				tmp.Value = i
				m.stats[definition.Key] = tmp
			} else {
				m.OutputNumericalValue(definition.Key, 0)
			}
		} else {
			m.OutputNumericalValue(definition.Key, 0)
		}

		// Always top off textual value
		m.OutputTextualValue(definition.Key, m.stats[definition.Key].Value)

		// Make Dashboard
		dashboardItem := new(DashboardCounterDefinition)
		dashboardItem.ID = definition.Key
		dashboardItem.BackgroundColor = definition.Dashboard.BackgroundColor
		dashboardItem.ForegroundColor = definition.Dashboard.ForegroundColor
		dashboardItem.IconClass = definition.Dashboard.IconClass
		dashboardItem.Description = definition.Dashboard.Description
		m.dashboardDefinitions = append(m.dashboardDefinitions, *dashboardItem)
	}
}

func (m *Module) ChangeData(item string, value int, notify bool) {

	// Flags
	var increase = false
	var decrease = false

	if value < m.stats[item].Value {
		decrease = true
	} else if value > m.stats[item].Value {
		increase = true
	}

	// Go Issue (fix is coming in Go2)
	var tmp = m.stats[item]
	tmp.Value = value
	m.stats[item] = tmp

	m.OutputNumericalValue(item, value)
	m.OutputTextualValue(item, value)

	if increase && notify && m.stats[item].Increase.Notify {

		// Check for callback
		if len(m.stats[item].Increase.NotifyCallback) > 0 {
			m.j.WebServer.TouchEndpoint(m.stats[item].Increase.NotifyCallback)
		}

		// Legacy Sound Events
		if len(m.stats[item].Increase.Sounds) > 0 {
			m.j.Media.PlaySound(m.stats[item].Increase.Sounds[rand.Intn(len(m.stats[item].Increase.Sounds))])
		}

		var message Core.NotifyMessage
		message.Discord = true
		message.DiscordPrefix = m.j.Config.GetPrefix()
		message.Twitch = true
		message.Message = strings.Replace(m.stats[item].Increase.NotifyMessage, "###", fmt.Sprintf("%d", value), -1)
		m.j.Notify.Announce(message)

	} else if decrease && notify && m.stats[item].Decrease.Notify {

		// Check for callback
		if len(m.stats[item].Decrease.NotifyCallback) > 0 {
			m.j.WebServer.TouchEndpoint(m.stats[item].Decrease.NotifyCallback)
		}

		// Legacy Sound Events
		if len(m.stats[item].Decrease.Sounds) > 0 {
			m.j.Media.PlaySound(m.stats[item].Increase.Sounds[rand.Intn(len(m.stats[item].Increase.Sounds))])
		}

		var message Core.NotifyMessage
		message.Discord = true
		message.DiscordPrefix = m.j.Config.GetPrefix()
		message.Twitch = true
		message.Message = strings.Replace(m.stats[item].Decrease.NotifyMessage, "###", fmt.Sprintf("%d", value), -1)
		m.j.Notify.Announce(message)
	}

	// Log
	m.j.Log.Message("Stats", "\""+item+"\""+" set to "+fmt.Sprintf("%d", value))
}
