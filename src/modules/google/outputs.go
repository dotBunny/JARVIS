package google

import (
	"path/filepath"

	Core "../../core"
)

// Outputs Pathing
type Outputs struct {
	LastFollowerPath    string
	LastSubscriberPath  string
	LastFollowersPath   string
	LastSubscribersPath string

	ChannelFollowersPath string
	ChannelViewsPath     string
	ChannelViewersPath   string
}

func (m *Module) setupOutputs() {

	m.outputs = new(Outputs)

	m.outputs.LastFollowerPath = filepath.Join(m.j.Config.GetOutputPath(), "YouTube_LastFollower.txt")
	m.outputs.LastFollowersPath = filepath.Join(m.j.Config.GetOutputPath(), "YouTube_LastFollowers.txt")
	m.outputs.LastSubscriberPath = filepath.Join(m.j.Config.GetOutputPath(), "YouTube_LastSubscriber.txt")
	m.outputs.LastSubscribersPath = filepath.Join(m.j.Config.GetOutputPath(), "YouTube_LastSubscribers.txt")
	m.outputs.ChannelFollowersPath = filepath.Join(m.j.Config.GetOutputPath(), "YouTube_ChannelFollowers.txt")
	m.outputs.ChannelViewsPath = filepath.Join(m.j.Config.GetOutputPath(), "YouTube_ChannelViews.txt")
	m.outputs.ChannelViewersPath = filepath.Join(m.j.Config.GetOutputPath(), "YouTube_ChannelViewers.txt")

	Core.Touch(m.outputs.LastFollowerPath)
	Core.Touch(m.outputs.LastFollowersPath)
	Core.Touch(m.outputs.LastSubscriberPath)
	Core.Touch(m.outputs.LastSubscribersPath)
	Core.Touch(m.outputs.ChannelFollowersPath)
	Core.Touch(m.outputs.ChannelViewsPath)
	Core.Touch(m.outputs.ChannelViewersPath)
}
