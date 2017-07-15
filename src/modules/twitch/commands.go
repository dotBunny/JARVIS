package twitch

import (
	"fmt"

	Core "../../core"
)

func (m *Module) setupCommands() {
	m.j.Discord.RegisterCommand("!t", m.commandSay, "Say something on Twitch", Core.CommandAccessModerator, "twitch")
	m.j.Discord.RegisterCommand("!twitch", m.commandTwitch, "Twitch ([info]/start/stop/auth)", Core.CommandAccessModerator, "twitch")

}

func (m *Module) commandSay(message *Core.DiscordMessage) {
	if len(message.Content) > 0 && m.irc.Connected() {
		m.SendMessage(m.settings.Channel, Core.WrapNicknameForIRC(message.Author)+" "+message.Content)
	}
}

func (m *Module) commandTwitch(message *Core.DiscordMessage) {
	// We have an option
	if len(message.Content) > 0 {
		if message.Content == "start" {
			m.j.Log.Message("Twitch", "Starting Twitch Polling/IRC")
			m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"Starting Polling / IRC")
			m.Start()
		} else if message.Content == "stop" {
			m.j.Log.Message("Twitch", "Stopping Twitch Polling/IRC")
			m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"Stopping Polling / IRC")
			m.Stop()
		} else if message.Content == "auth" {
			m.j.Log.Message("Twitch", "Authenticating with Twitch")
			m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"Authenticating")
			m.authenticate()
		}
	} else {
		m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"There are currently "+fmt.Sprintf("%d", m.data.ChannelViewers)+" viewers, and "+fmt.Sprintf("%d", m.data.ChannelFollowers)+".")
	}
}
