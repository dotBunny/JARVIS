package twitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	Core "../../core"
)

func (m *Module) setupPolling() {
	// Create Ticker
	twitchPollingFrequency, twitchPollingError := time.ParseDuration(strconv.Itoa(m.settings.PollingFrequency) + "s")
	if twitchPollingError != nil {
		twitchPollingFrequency, _ = time.ParseDuration("5s")
	}
	m.j.Log.Message("Twitch", "Starting polling at "+twitchPollingFrequency.String())
	m.ticker = time.NewTicker(twitchPollingFrequency)
	m.Poll(false)
	go m.loop()
}

// Loop awaiting ticker
func (m *Module) loop() {
	for {
		select {
		case <-m.ticker.C:
			m.Poll(true)
		}
	}
}

// Poll For Updates
func (m *Module) Poll(notify bool) {
	m.pollFollowers(notify)
	m.pollStream(notify)
	m.pollSubscribers(notify)
}

func (m *Module) pollFollowers(notify bool) {

	// Sanitize settings (based on Twitch rules)
	var limit = m.settings.LastFollowersCount
	if limit <= 0 {
		limit = 25
	} else if limit > 100 {
		limit = 100
	}
	// Query for data
	var url = twitchRootURL + "channels/" + m.settings.ChannelID + "/follows/?limit=" + strconv.Itoa(limit) + "&direction=desc"
	response, responseError := m.getResponse(url)
	if responseError != nil {
		m.j.Log.Error("Twitch", "Unable to get data valid response from: "+url+", "+responseError.Error())
		return
	}

	followers := &Followers{}
	errorDecoder := json.NewDecoder(response.Body).Decode(&followers)
	if errorDecoder != nil {
		m.j.Log.Error("Twitch", "Error decoding JSON from: "+url+", "+responseError.Error())
		return
	}

	// Only really process if we have ANY followers
	if followers.Total > 0 {

		// Handle Last Follower
		if followers.Follows[0].User.DisplayName != m.data.LastFollower {

			Core.SaveFile([]byte(followers.Follows[0].User.DisplayName), m.outputs.LastFollowerPath)
			m.data.LastFollower = followers.Follows[0].User.DisplayName

			if notify {
				m.j.Discord.GetSession().ChannelMessageSend(
					m.j.Discord.GetPrivateChannelID(),
					m.settings.Prefix+"New Twitch Follower "+followers.Follows[0].User.DisplayName)
			}

			// TODO: Need to make it so it loads so this doesnt ding
			m.j.Log.Message("Twitch", "New Twitch Follower "+followers.Follows[0].User.DisplayName)

			// Since the last follower has changed, it means that we need to update our last list

			// Make sure we have enough for the loop, or dont
			var itemLength = len(followers.Follows)
			if itemLength > m.settings.LastFollowersCount {
				itemLength = m.settings.LastFollowersCount
			}

			m.data.LastFollowers = m.data.LastFollowers[:0]

			var buffer bytes.Buffer
			for i := 0; i < itemLength; i++ {
				m.data.LastFollowers = append(m.data.LastFollowers, followers.Follows[i].User.DisplayName)
				buffer.WriteString(followers.Follows[i].User.DisplayName)
				buffer.WriteString("\n")
			}
			Core.SaveFile(buffer.Bytes(), m.outputs.LastFollowersPath)
		}
	}
	followers = nil
}

func (m *Module) pollStream(notify bool) {

	// Query for data
	var url = twitchRootURL + "streams/" + m.settings.ChannelID
	response, responseError := m.getResponse(url)
	if responseError != nil {
		m.j.Log.Error("Twitch", "Unable to get data valid response from: "+url+", "+responseError.Error())
		return
	}

	stream := &StreamResponse{}
	errorDecoder := json.NewDecoder(response.Body).Decode(&stream)
	if errorDecoder != nil {
		m.j.Log.Error("Twitch", "Error decoding JSON from: "+url+", "+responseError.Error())
		return
	}

	// Channel Follower Count
	if stream.Stream.Channel.Followers != m.data.ChannelFollowers {
		m.data.ChannelFollowers = stream.Stream.Channel.Followers
		Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.ChannelFollowers), m.settings.PadChannelFollowersOutput, "0")), m.outputs.ChannelFollowersPath)
	}

	// Channel View Total
	if stream.Stream.Channel.Views != m.data.ChannelViews {
		m.data.ChannelViews = stream.Stream.Channel.Views
		Core.SaveFile([]byte(fmt.Sprint(m.data.ChannelViews)), m.outputs.ChannelViewsPath)
	}

	// How many current viewers
	if stream.Stream.Viewers != m.data.ChannelViewers {
		m.data.ChannelViewers = stream.Stream.Viewers
		Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", m.data.ChannelViewers), m.settings.PadChannelViewersOutput, "0")), m.outputs.ChannelViewersPath)
	}

	// Channel Name - This is the thing you set with all your tags
	if stream.Stream.Channel.DisplayName != m.data.ChannelName {
		m.data.ChannelName = stream.Stream.Channel.DisplayName
		Core.SaveFile([]byte(m.data.ChannelName), m.outputs.ChannelNamePath)
	}

	// Game Currently Playing (in our case we should keep this to "Creative")
	if stream.Stream.Channel.Game != m.data.ChannelGame {
		m.data.ChannelGame = stream.Stream.Channel.Game
		Core.SaveFile([]byte(m.data.ChannelGame), m.outputs.ChannelGamePath)
	}

	stream = nil
}

func (m *Module) pollSubscribers(notify bool) {

	// Sanitize settings (based on Twitch rules)
	var limit = m.settings.LastSubscribersCount
	if limit <= 0 {
		limit = 25
	} else if limit > 100 {
		limit = 100
	}

	// Query for data
	var url = twitchRootURL + "channels/" + m.settings.ChannelID + "/subscriptions/?limit=" + strconv.Itoa(limit) + "&direction=desc"
	response, responseError := m.getResponse(url)
	if responseError != nil {
		m.j.Log.Error("Twitch", "Unable to get data valid response from: "+url+", "+responseError.Error())
		return
	}

	subscribers := &Subscribers{}
	errorDecoder := json.NewDecoder(response.Body).Decode(&subscribers)
	if errorDecoder != nil {
		m.j.Log.Error("Twitch", "Error decoding JSON from: "+url+", "+responseError.Error())
		return
	}

	if subscribers.Total > 0 {
		if subscribers.Subscriptions[0].User.Name != m.data.LastSubscriber {
			m.data.LastSubscriber = subscribers.Subscriptions[0].User.Name
			Core.SaveFile([]byte(m.data.LastSubscriber), m.outputs.LastSubscriberPath)

			if notify {
				m.j.Discord.GetSession().ChannelMessageSend(
					m.j.Discord.GetPrivateChannelID(),
					m.settings.Prefix+"New Twitch Subscriber "+m.data.LastSubscriber)
			}

			// TODO: Need to make it so it loads so this doesnt ding
			m.j.Log.Message("Twitch", "New Twitch Subscriber "+m.data.LastSubscriber)

			var itemLength = len(subscribers.Subscriptions)
			if itemLength > m.settings.LastSubscribersCount {
				itemLength = m.settings.LastSubscribersCount
			}

			m.data.LastSubscribers = m.data.LastSubscribers[:0]

			var buffer bytes.Buffer
			for i := 0; i < itemLength; i++ {
				m.data.LastSubscribers[i] = subscribers.Subscriptions[i].User.DisplayName
				buffer.WriteString(subscribers.Subscriptions[i].User.DisplayName)
				buffer.WriteString("\n")
			}
			Core.SaveFile(buffer.Bytes(), m.outputs.LastSubscriberPath)
		}
	}
}
