package ads

import (
	"strconv"
	"time"

	Core "../../core"
)

// Module Class
type Module struct {
	tickers      []*time.Ticker
	settings     *Config
	warningCount int
	errorCount   int
	j            *Core.JARVIS
}

// Initialize the Ads Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	// Register Status Updator
	m.j.Status.RegisterUpdator("ads", m.StatusUpdate)

	m.loadConfig()

	// Use generic config prefix
	m.j.Log.RegisterChannel("Ads", "green", m.j.Config.GetPrefix())

	// Spin Up Ad Timers
	m.setupTimers()
}

func (m *Module) StatusUpdate() (int, int) {
	return m.warningCount, m.errorCount
}

func (m *Module) setupTimers() {

	for _, ad := range m.settings.Definitions {

		// Create new duration
		newDuration, badDuration := time.ParseDuration(strconv.Itoa(ad.Interval) + "s")

		// Create a new timer and put its reference in the slice
		if badDuration == nil {
			m.tickers = append(m.tickers, time.NewTicker(newDuration))
			m.j.Log.Message("Ads", "Starting ad timer for "+ad.Key)
		}
	}
	go m.loop()
}

func (m *Module) loop() {
	// Loop awaiting ticker
	for {
		for key, ad := range m.settings.Definitions {
			select {
			case <-m.tickers[key].C:
				m.process(ad)
			}
		}
	}
}

func (m *Module) process(ad Ad) {

	var message Core.NotifyMessage

	// Only support Twitch currently
	if Core.StringsContains(ad.Channels, "twitch") {
		message.Twitch = true
	}

	message.Message = Core.RandomFromStrings(ad.Content)

	m.j.Notify.Announce(message)
}
