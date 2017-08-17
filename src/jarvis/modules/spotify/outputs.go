package spotify

import (
	"path/filepath"

	Core "../../core"
)

// Outputs Pathing
type Outputs struct {
	SongPath  string
	LinkPath  string
	ImagePath string
}

func (m *Module) setupOutputs() {

	m.outputs = new(Outputs)

	m.outputs.SongPath = filepath.Join(m.j.Config.GetOutputPath(), "Spotify_LatestSong.txt")
	m.outputs.LinkPath = filepath.Join(m.j.Config.GetOutputPath(), "Spotify_LatestURL.txt")
	m.outputs.ImagePath = filepath.Join(m.j.Config.GetOutputPath(), "Spotify_LatestImage.jpg")

	// Touch Files
	Core.Touch(m.outputs.SongPath)
	Core.Touch(m.outputs.LinkPath)
	Core.Touch(m.outputs.ImagePath)

	// Write Defaults
	Core.SaveFile(DefaultImage, m.outputs.ImagePath)
	Core.SaveFile([]byte("JARVIS"), m.outputs.SongPath)
	Core.SaveFile([]byte("https://github.com/dotBunny/JARVIS/"), m.outputs.LinkPath)
}
