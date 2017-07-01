package jira

type JIRAFields struct {
	summary     string
	description string
}
type JIRAIssue struct {
	expand string
	id     int
	self   string
	key    string
	fields []JIRAFields
}
type JIRAResponse struct {
	expand     string
	startAt    int
	maxResults int
	total      int
	issues     []JIRAIssue
}

// Data Structure
type Data struct {
	LastResponse JIRAResponse
}

func (m *Module) setupData() {
	m.data = new(Data)
}
