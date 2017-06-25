package twitch

import (
	Core "../../core"
	irc "github.com/thoj/go-ircevent"
)

func (m *Module) ircJoinChannels() {
	m.j.Log.Message("Twitch", "Joining IRC Channels ..")
	m.irc.Join(m.settings.Channel)
}

func (m *Module) connectIRC() {

	// Dont do chat
	if !m.settings.Chat {
		return
	}
	// Create IRC Objects - Username must be LOWERCASE

	m.j.Log.Message("Twitch", "Connecting to IRC ...")
	m.irc = irc.IRC(m.settings.ChatUsername, m.settings.ChatUsername)

	m.irc.UseTLS = false
	m.irc.Password = m.settings.ChatOAuth

	// // Set IRC Connection Callback
	m.irc.AddCallback("001", m.handleConnected)
	m.irc.AddCallback("366", func(e *irc.Event) {})
	m.irc.AddCallback("PRIVMSG", m.handleMessage)
	m.irc.AddCallback("NOTICE", m.handleNotice)
	m.irc.AddCallback("PING", m.handlePing)

	// Connect this shit up
	errorIRC := m.irc.Connect("irc.chat.twitch.tv:6667")
	if errorIRC != nil {
		m.j.Log.Error("Twitch", "Unable to connect to Twitch IRC Server. "+errorIRC.Error())
		return
	}
	m.irc.Join(m.settings.Channel)

	go m.irc.Loop()
}

func (m *Module) handleConnected(event *irc.Event) {
	// Join the channel
	m.j.Log.Message("Twitch", "Joining IRC Channel "+m.settings.Channel)
}
func (m *Module) handlePing(event *irc.Event) {
	m.irc.SendRaw("PONG :tmi.twitch.tv")
}

func (m *Module) handleMessage(event *irc.Event) {

	_, _ = m.j.Discord.GetSession().ChannelMessageSend(m.j.Discord.GetChatChannelID(), m.settings.Prefix+Core.WrapNicknameForDiscord(event.Nick)+" "+event.Message())

}
func (m *Module) handleNotice(event *irc.Event) {
	m.j.Log.Message("Twitch", "NOTICE <"+event.Nick+"> "+event.Message())
}

// SendMessage via IRC
func (m *Module) SendMessage(target string, message string) {
	if m.settings.Chat {
		return
	}
	m.irc.Privmsg(target, message)
}
