package core

import (
	"io/ioutil"
	"os"
)

// LogCore Class
type MediaCore struct {
	MediaLastPath    string
	MediaLastData    []byte
	MediaLastVersion int

	j *JARVIS
}

// Play Sound
func (m *MediaCore) PlaySound(path string) {
	_, err := os.Stat(path)
	if err == nil {
		m.MediaLastPath = path
		m.MediaLastData, _ = ioutil.ReadFile(path)
		m.MediaLastVersion++
	}

	// 	cmd := exec.Command(m.VLC, path, "--play-and-exit", "--no-loop", "--no-repeat", "--quiet", "--start-time=0")
	// 	errCheck := cmd.Run()
	// 	if errCheck != nil {
	// 		m.j.Log.Warning("Media", errCheck.Error())
	// 	}
	// }
}

// Initialize Media
func (m *MediaCore) Initialize(jarvisInstance *JARVIS) {

	// Create instance of Config Core
	m = new(MediaCore)

	// Assign JARVIS (circle!)
	m.j = jarvisInstance
	m.j.Media = m

	m.MediaLastVersion = 0
	m.MediaLastPath = "Nothing"

}
