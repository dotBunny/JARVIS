package twitch

import (
	"path/filepath"

	Core "../../core"
)

// Outputs Pathing
type Outputs struct {
	LastFollowerPath   string
	LastSubscriberPath string
	LastFollowersPath  string

	ChannelFollowersPath string
	ChannelViewsPath     string
	ChannelNamePath      string

	ChannelViewersPath string
	ChannelGamePath    string
}

func (m *Module) setupOutputs() {

	m.outputs = new(Outputs)

	m.outputs.LastFollowerPath = filepath.Join(m.j.Config.GetOutputPath(), "Twitch_LastFollower.txt")
	m.outputs.LastFollowersPath = filepath.Join(m.j.Config.GetOutputPath(), "Twitch_LastFollowers.txt")
	m.outputs.LastSubscriberPath = filepath.Join(m.j.Config.GetOutputPath(), "Twitch_LastSubscriber.txt")
	m.outputs.ChannelFollowersPath = filepath.Join(m.j.Config.GetOutputPath(), "Twitch_ChannelFollowers.txt")
	m.outputs.ChannelViewsPath = filepath.Join(m.j.Config.GetOutputPath(), "Twitch_ChannelViews.txt")
	m.outputs.ChannelNamePath = filepath.Join(m.j.Config.GetOutputPath(), "Twitch_ChannelName.txt")
	m.outputs.ChannelViewersPath = filepath.Join(m.j.Config.GetOutputPath(), "Twitch_ChannelViewers.txt")
	m.outputs.ChannelGamePath = filepath.Join(m.j.Config.GetOutputPath(), "Twitch_ChannelGame.txt")

	Core.Touch(m.outputs.LastFollowerPath)
	Core.Touch(m.outputs.LastFollowersPath)
	Core.Touch(m.outputs.LastSubscriberPath)
	Core.Touch(m.outputs.ChannelFollowersPath)
	Core.Touch(m.outputs.ChannelViewsPath)
	Core.Touch(m.outputs.ChannelNamePath)
	Core.Touch(m.outputs.ChannelViewersPath)
	Core.Touch(m.outputs.ChannelGamePath)
}
