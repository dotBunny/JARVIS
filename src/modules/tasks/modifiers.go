package tasks

// ** JIRA UPDATE STRUCTURE
type JIRAModifier struct {
	tasks *Module
}

func (m *JIRAModifier) GetDataValue() string {

	return m.tasks.GetWorkingOn()
}
func (m *JIRAModifier) SetDataValue(message string, notify bool) {
	m.tasks.SetWorkingOn(message, notify)
}
func (m *JIRAModifier) ShouldUpdate() bool {
	return m.tasks.UseJIRAForWork
}
