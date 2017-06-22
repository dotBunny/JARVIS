package google

import (
	"fmt"

	Core "../../core"
)

// Data Structure
type Data struct {
	LastFollower    string
	LastSubscriber  string
	LastSubscribers []string
	LastFollowers   []string

	ChannelFollowers uint
	ChannelViews     uint

	ChannelViewers uint
}

func (m *Module) setupData() {
	m.data = new(Data)

	m.data.ChannelViewers = 0
	Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.ChannelViewers), m.settings.PadChannelViewersOutput, "0")), m.outputs.ChannelViewersPath)
}
