package core

import "github.com/bwmarrin/discordgo"

type NotifyMessage struct {
	Author string

	Discord             bool
	DiscordEmbed        *discordgo.MessageEmbed
	DiscordPrefix       string
	DiscordForceChannel string

	Twitch bool

	Message string
}

// DiscordFunc for IRC
//type NotifyNotifier func(string)

type NotifyCore struct {
	twitch TextNotifier
	j      *JARVIS
}

// Initialize Media
func (m *NotifyCore) Initialize(jarvisInstance *JARVIS) {
	m = new(NotifyCore)
	m.j = jarvisInstance
	m.j.Notify = m
}
func (m *NotifyCore) ConnectTwitch(twitch TextNotifier) {
	m.twitch = twitch
}

func (m *NotifyCore) Announce(message NotifyMessage) {

	// Handle Discord Notification
	if message.Discord && m.j.Discord.IsConnected() {

		if len(message.DiscordForceChannel) > 0 {
			if message.DiscordEmbed != nil {
				m.j.Discord.GetSession().ChannelMessageSendEmbed(message.DiscordForceChannel, message.DiscordEmbed)
			} else {
				m.j.Discord.GetSession().ChannelMessageSend(message.DiscordForceChannel, message.DiscordPrefix+" "+message.Message)
			}
		} else {
			if message.DiscordEmbed != nil {
				m.j.Discord.AnnoucementEmbed(message.DiscordEmbed)
			} else {
				m.j.Discord.Announcement(message.DiscordPrefix + " " + message.Message)
			}
		}
	}

	// Handle Twitch Notification
	if message.Twitch && m.twitch != nil {
		m.twitch.Notify(message.Message)
	}
}
