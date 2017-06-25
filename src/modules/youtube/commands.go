package youtube

import (
	Core "../../core"
)

func (m *Module) setupCommands() {
	//	m.j.Discord.RegisterCommand("!y", m.commandSay, "Say something on Twitch", Core.CommandAccessModerator)
	m.j.Discord.RegisterCommand("!youtube", m.commandYouTube, "YouTube ([info]/start/stop/auth)", Core.CommandAccessModerator)

}

// func (m *Module) commandSay(message *Core.DiscordMessage) {
// 	if len(message.Content) > 0 && m.irc.Connected() {
// 		m.SendMessage(m.settings.Channel, Core.WrapNicknameForIRC(message.Author)+" "+message.Content)
// 	}
// }

func (m *Module) commandYouTube(message *Core.DiscordMessage) {
	// We have an option
	if len(message.Content) > 0 {
		if message.Content == "start" {
			m.j.Log.Message("YouTube", "Starting YouTube Polling")
			m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"Starting Polling")
			m.Start()
		} else if message.Content == "stop" {
			m.j.Log.Message("YouTube", "Stopping YouTube Polling/IRC")
			m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"Stopping Polling ")
			m.Stop()
		} else if message.Content == "auth" {
			m.j.Log.Message("YouTube", "Authenticating with YouTube")
			m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"Authenticating")
			m.authenticate()
		}
	} else {
		//m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"There are currently "+fmt.Sprintf("%d", m.data.ChannelViewers)+" viewers, and "+fmt.Sprintf("%d", m.data.ChannelFollowers)+".")
	}
}
