package twitch

// Data Structure
type Data struct {
	LastFollower    string
	LastSubscriber  string
	LastSubscribers []string
	LastFollowers   []string

	ChannelFollowers uint
	ChannelViews     uint
	ChannelName      string

	ChannelViewers uint
	ChannelGame    string
}

func (m *Module) setupData() {
	m.data = new(Data)
}
