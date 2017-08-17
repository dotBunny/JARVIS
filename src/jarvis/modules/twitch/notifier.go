package twitch

// ** JIRA UPDATE STRUCTURE
type TwitchNotifier struct {
	Twitch *Module
}

func (m *TwitchNotifier) Notify(message string) {
	m.Twitch.Say(message)
}
