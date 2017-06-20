package twitch

import (
	"fmt"

	Core "../../core"
)

func (m *Module) setupCommands() {
	m.j.Discord.RegisterCommand("!t", m.commandSay, "Say something on Twitch", Core.CommandAccessPrivate)
}

func (m *Module) commandSay(message *Core.DiscordMessage) {

	fmt.Println("COMMAND " + message.Content)
	if len(message.Content) > 0 && m.irc.Connected() {
		fmt.Println("SEND MESSAGE")
		m.SendMessage(m.settings.Channel, Core.WrapNicknameForIRC(message.Author)+" "+message.Content)
	}
}
