//+build darwin

package core

import (
	"os/exec"
)

// PlaySound File
func PlaySound(filePath string) {
	cmd := exec.Command("afplay", filePath)
	cmd.Run()
}
