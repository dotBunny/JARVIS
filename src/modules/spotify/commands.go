package spotify

import (
	Core "../../core"
)

func (m *Module) setupCommands() {
	m.j.Discord.RegisterCommand("!spotify", m.commandSpotify, "Spotify ([info]/start/stop/auth)", Core.CommandAccessModerator)
}

func (m *Module) commandSpotify(message *Core.DiscordMessage) {
	// We have an option
	if len(message.Content) > 0 {
		if message.Content == "start" {
			m.j.Log.Message("Spotify", "Starting Spotify Polling")
			m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"Starting Polling")
			m.Start()
		} else if message.Content == "stop" {
			m.j.Log.Message("Spotify", "Stopping Spotify Polling")
			m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"Stopping Polling")
			m.Stop()
		} else if message.Content == "auth" {
			m.j.Log.Message("Spotify", "Authenticating with Spotify")
			m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"Authenticating")
			m.authenticate()
		}
	} else {
		m.j.Discord.GetSession().ChannelMessageSend(message.ChannelID, m.settings.Prefix+"Currently playing "+m.data.CurrentlyPlayingTrack)
	}
}

/*

!spotify update
!spotify (shows current track info)
!spotify next
!spotify pause
!spotify play

*/

// 	console.AddHandler("/spotify.next", "Skips to the next track in the user's Spotify queue.", m.consoleNextTrack)
// 	console.AddAlias("/next", "/spotify.next")
// 	console.AddAlias("/n", "/spotify.next")
// 	console.AddAlias("/skip", "/spotify.next")
// 	console.AddHandler("/spotify.pause", "Pause/Play the current track in Spotify.", m.consolePausePlay)
// 	console.AddAlias("/p", "/spotify.pause")
// 	console.AddHandler("/spotify.stats", "Display some stats from Spotify.", m.consoleStats)
// 	console.AddHandler("/spotify.update", "Force polling Spotify for updates.", m.consoleUpdate)
// }
