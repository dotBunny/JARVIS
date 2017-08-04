package spotify

import (
	"bufio"
	"bytes"
	"strconv"
	"time"

	Core "../../core"
)

func (m *Module) setupPolling() {
	// Create Ticker
	spotifyPollingFrequency, spotifyPollingError := time.ParseDuration(strconv.Itoa(m.settings.PollingFrequency) + "s")
	if spotifyPollingError != nil {
		spotifyPollingFrequency, _ = time.ParseDuration("5s")
	}
	m.j.Log.Message("Spotify", "Starting polling at "+spotifyPollingFrequency.String())
	m.ticker = time.NewTicker(spotifyPollingFrequency)
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
	m.pollCurrentlyPlaying(notify)
}

func (m *Module) pollCurrentlyPlaying(notify bool) {

	state, err := m.spotifyClient.PlayerCurrentlyPlaying()

	if err != nil {
		m.j.Log.Warning("Spotify", "Unable to retrieve currently playing song.")
	} else {

		// Handle Basic Track Information
		var buffer bytes.Buffer
		var artistCount = len(state.Item.Artists)
		for i := 0; i < artistCount; i++ {
			buffer.WriteString(state.Item.Artists[i].Name)
			if i < (artistCount - 1) {
				buffer.WriteString(", ")
			}
		}

		m.data.ArtistLine = buffer.String()

		buffer.WriteString(" - ")
		buffer.WriteString(state.Item.Name)

		if buffer.String() != m.data.CurrentlyPlayingTrack {
			// Assign Data
			m.data.CurrentlyPlayingTrack = buffer.String()

			m.j.Log.Message("Spotify", "New track detected: "+m.data.CurrentlyPlayingTrack)

			// Check Truncate Length
			if buffer.Len() > m.settings.TruncateTrackLength {
				buffer.Truncate(m.settings.TruncateTrackLength)
				buffer.WriteString(m.settings.TruncateTrackRunes)
			}

			// Safe Track (truncated)
			Core.SaveFile(buffer.Bytes(), m.outputs.SongPath)

			// Get/Save Currently Playing URL
			m.data.CurrentlyPlayingURL = state.Item.ExternalURLs["spotify"]
			m.data.TrackThumbnailURL = state.Item.Album.Images[0].URL
			m.data.TrackName = state.Item.Name
			Core.SaveFile([]byte(m.data.CurrentlyPlayingURL), m.outputs.LinkPath)

			// if notify {

			// 	var message Core.NotifyMessage
			// 	message.Discord = true
			// 	message.DiscordEmbed = &discordgo.MessageEmbed{
			// 		Type:  "rich",
			// 		Title: "Now Playing On Spotify",
			// 		URL:   m.data.CurrentlyPlayingURL,
			// 		//Description: ,
			// 		Color:     1947988,
			// 		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: m.data.TrackThumbnailURL},
			// 		Fields: []*discordgo.MessageEmbedField{
			// 			&discordgo.MessageEmbedField{
			// 				Name:   m.data.TrackName,
			// 				Value:  m.data.ArtistLine,
			// 				Inline: true},
			// 		},
			// 	}
			// 	message.Twitch = true
			// 	message.Message = "Playing " + m.data.ArtistLine + " - " + m.data.TrackName + " (" + m.data.CurrentlyPlayingURL + ")"
			// 	message.DiscordForceChannel = m.j.Discord.GetFeedChannelID()

			// 	m.j.Notify.Announce(message)

			// }

			// New Artwork
			if len(state.Item.Album.Images) > 0 {
				buffer.Reset()
				writer := bufio.NewWriter(&buffer)
				state.Item.Album.Images[0].Download(writer)

				if !bytes.Equal(buffer.Bytes(), m.data.LastImageData) {

					Core.SaveFile(buffer.Bytes(), m.outputs.ImagePath)
					m.data.LastImageData = buffer.Bytes()
				}
				writer.Flush()
			}
		}

		// Store some data from the poll no matter what
		m.data.DurationMS = state.Item.Duration
		m.data.PlayedMS = state.Progress
		m.data.CurrentlyPlaying = state.Playing
	}
	state = nil
}
