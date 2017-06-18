package spotify

// SpotifyData Pathing
type SpotifyData struct {
	CurrentlyPlayingTrack string
	LastImageData         []byte
	DurationMS            int
	PlayedMS              int
	CurrentlyPlaying      bool
	CurrentlyPlayingURL   string
}

func (m *SpotifyModule) setupData() {
	m.data = new(SpotifyData)
}
