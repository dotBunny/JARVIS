package twitch

import (
	Core "../../core"
)

func (m *Module) setupCommands() {
	m.j.Discord.RegisterCommand("!t", m.commandSay, "Say something on Twitch", Core.CommandAccessModerator)
}

func (m *Module) commandSay(message *Core.DiscordMessage) {
	if len(message.Content) > 0 && m.irc.Connected() {
		m.SendMessage(m.settings.Channel, Core.WrapNicknameForIRC(message.Author)+" "+message.Content)
	}
}
