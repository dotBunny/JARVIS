package core

import (
	"os"
	"os/exec"
)

// LogCore Class
type MediaCore struct {
	VLC string
	j   *JARVIS
}

// Play Sound
func (m *MediaCore) PlaySound(path string) {
	_, err := os.Stat(path)
	if err == nil {

		cmd := exec.Command(m.VLC, "file://"+path, "--play-and-exit", "--no-loop", "--no-repeat", "--quiet", "--start-time=0")
		errCheck := cmd.Run()
		if errCheck != nil {
			m.j.Log.Warning("Media", errCheck.Error())
		}
	}
}

// Initialize Media
func (m *MediaCore) Initialize(jarvisInstance *JARVIS) {

	// Create instance of Config Core
	m = new(MediaCore)

	// Assign JARVIS (circle!)
	jarvisInstance.Media = m
	m.VLC = jarvisInstance.Config.GetVLCPath()
	m.j = jarvisInstance
}
