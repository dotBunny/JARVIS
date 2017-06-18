package twitch

import (
	irc "./irc"
)

func (m *TwitchModule) ircJoinChannels() {
	m.j.Log.Message("Twitch", "Joining IRC Channels ..")
	m.irc.Join(m.settings.Channel)
}
func (m *TwitchModule) ircWelcomeMessage() {
	m.irc.Privmsg(m.settings.Channel, "Hello World!")
}

func (m *TwitchModule) connectIRC() {
	// Create IRC Objects - Username must be LOWERCASE

	m.j.Log.Message("Twitch", "Connecting to IRC ...")
	m.irc = irc.NewClient(m.settings.Username, m.settings.Username)

	m.irc.TLS = false
	m.irc.Password = "oauth:" + m.twitchToken

	// Cue up commands
	// go m.ircSetPass(100, &m.irc.)
	// go m.ircJoinChannels(200, &m.irc.ready)

	// Connect this shit up
	m.irc.Connect("irc.chat.twitch.tv:6667")

	// // Set IRC Connection Callback
	// m.irc.AddCallback("001", m.handleConnected)
	// m.irc.AddCallback("366", func(e *irc.Event) {})
	// m.irc.AddCallback("PRIVMSG", m.handleMessage)
	// m.irc.AddCallback("NOTICE", m.handleNotice)

	// err := m.irc.Connect("irc.chat.twitch.tv:6667")
	// if err != nil {
	// 	m.j.Log.Warning("Twitch", "Unable to connect to IRC")
	// 	return
	// }

	// go m.irc.Loop()
}

// // Main loop to control the connection.
// func (irc *Connection) Loop() {
// 	errChan := irc.ErrorChan()
// 	for !irc.isQuitting() {
// 		err := <-errChan
// 		close(irc.end)
// 		irc.Wait()
// 		for !irc.isQuitting() {
// 			irc.Log.Printf("Error, disconnected: %s\n", err)
// 			if err = irc.Reconnect(); err != nil {
// 				irc.Log.Printf("Error while reconnecting: %s\n", err)
// 				time.Sleep(60 * time.Second)
// 			} else {
// 				errChan = irc.ErrorChan()
// 				break
// 			}
// 		}
// 	}
// }

// func (m *TwitchModule) handleConnected(event *irc.Event) {
// 	// Join the channel
// 	m.j.Log.Message("Twitch", "Joining channel "+m.settings.Channel)
// 	m.irc.Join(m.settings.Channel)
// 	m.irc.Privmsg(m.settings.Channel, "Hello World!")
// }

//:nickname!nickname@nickname.tmi.twitch.tv PRIVMSG #channel :message
// func (m *TwitchModule) handleMessage(event *irc.Event) {

// 	m.j.Log.Message("Twitch", "Message Received "+event.Message())
// 	if m.settings.ChatSync {
// 		_, _ = m.discord.GetSession().ChannelMessageSend(m.settings.ChatSyncChannelID, Core.WrapNickname(event.Nick)+" "+event.Message())
// 	}
// }
// func (m *TwitchModule) handleNotice(event *irc.Event) {

// 	// if m.settings.ChatSync {
// 	// 	_, _ = m.discord.GetSession().ChannelMessageSend(m.settings.ChatSyncChannelID, "[NOTICE] "+Core.WrapNickname(event.Nick)+" "+event.Message())
// 	// }
// }
