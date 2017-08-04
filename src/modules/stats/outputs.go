package stats

import (
	"fmt"
	"path/filepath"

	Core "../../core"
)

// GetOutputPath of item
func (m *Module) GetOutputPath(item string, flare string) string {
	return filepath.Join(m.j.Config.GetOutputPath(), "Stats_"+item+"_"+flare+".txt")
}

func (m *Module) OutputNumericalValue(item string, value int) {

	if m.stats[item].NumericalOutput.Enabled {
		Core.SaveFile([]byte(Core.Left(fmt.Sprintf("%d", value), m.stats[item].NumericalOutput.Padding, "0")), m.GetOutputPath(item, "Count"))
	}
}

func (m *Module) OutputTextualValue(item string, value int) {
	if m.stats[item].TextOutput.Enabled {
		Core.SaveFile([]byte(m.stats[item].TextOutput.Prefix+fmt.Sprintf("%d", value)+m.stats[item].TextOutput.Suffix), m.GetOutputPath(item, "Text"))
	}
}
