package stats

import (
	"path/filepath"

	Core "../../core"
)

// SpotifyOutputs Pathing
type StatsOutputs struct {
	WorkingOnPath  string
	CoffeeCountPath  string
	SavesCountPath string
}

func (m *StatsModule) setupOutputs() {

	m.outputs = new(StatsOutputs)

	m.outputs.WorkingOnPath = filepath.Join(m.j.Config.GetOutputPath(), "Stats_WorkingOn.txt")
	m.outputs.CoffeeCountPath = filepath.Join(m.j.Config.GetOutputPath(), "Stats_CoffeeCount.txt")
	m.outputs.SavesCountPath = filepath.Join(m.j.Config.GetOutputPath(), "Spotify_SavesCount.txt")

	// Touch Files
	Core.Touch(m.outputs.WorkingOnPath)
	Core.Touch(m.outputs.CoffeeCountPath)
	Core.Touch(m.outputs.SavesCountPath)
}
