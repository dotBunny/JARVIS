package jira

import (
	"strconv"
	"time"
)

func (m *Module) setupPolling() {
	// Create Ticker
	spotifyPollingFrequency, spotifyPollingError := time.ParseDuration(strconv.Itoa(m.settings.PollingFrequency) + "s")
	if spotifyPollingError != nil {
		spotifyPollingFrequency, _ = time.ParseDuration("5s")
	}
	m.j.Log.Message("JIRA", "Starting polling at "+spotifyPollingFrequency.String())
	m.ticker = time.NewTicker(spotifyPollingFrequency)
	m.Poll(false)
	go m.loop()
}

// Loop awaiting ticker
func (m *Module) loop() {
	for {
		select {
		case <-m.ticker.C:
			m.Poll(true)
		}
	}
}

// Poll For Updates
func (m *Module) Poll(notify bool) {

	m.pollIssues(notify)
}

func (m *Module) pollIssues(notify bool) {

}
