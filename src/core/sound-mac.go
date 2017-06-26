//+build darwin

package core

import (
	"os"
	"os/exec"
)

// PlaySound File
func PlaySound(filePath string) {
	_, err := os.Stat(filePath)
	if err == nil {
		cmd := exec.Command("afplay", "-v 0.5", filePath)
		cmd.Run()
	}
}
