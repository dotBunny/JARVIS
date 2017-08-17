package spotify

// Data Structure
type Data struct {
	CurrentlyPlayingTrack string
	LastImageData         []byte
	DurationMS            int
	PlayedMS              int
	CurrentlyPlaying      bool
	CurrentlyPlayingURL   string
	TrackName             string
	ArtistLine            string
	TrackThumbnailURL     string
}

func (m *Module) setupData() {
	m.data = new(Data)
}
