package core

import (
	"io"
	"log"
	"os"
	"path"
	"strings"
)

// LogCore Class
type LogCore struct {
	channels map[string]string
	prefix   map[string]string
	logFile  *os.File
	j        *JARVIS
}

// Initialize the Logging Module
func (m *LogCore) Initialize(jarvisInstance *JARVIS) {

	// Create isntance of LogCore
	m = new(LogCore)

	// Assign JARVIS (circle!)
	jarvisInstance.Log = m
	m.j = jarvisInstance

	m.channels = make(map[string]string)
	m.prefix = make(map[string]string)

	m.RegisterChannel("Core", "white", m.j.Config.GetPrefix())
	m.RegisterChannel("Core", "grey", m.j.Config.GetPrefix())

	logFile, err := os.OpenFile(path.Join(m.j.Config.GetOutputPath(), "jarvis.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// Setup echoing to log file and console
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	// TODO: This file is not being closed properly? cause its in a class
	//defer logFile.Close()

}

// RegisterChannel for use with loggin
func (m *LogCore) RegisterChannel(tag string, color string, prefix string) {
	m.channels[strings.ToUpper(tag)] = color
}

// Shutdown LogCore
func (m *LogCore) Shutdown() {

	m.logFile.Close()
}

// Message Level Alart
func (m *LogCore) Message(channel string, message string) {
	if m.j.Discord != nil && m.j.Discord.IsConnected() {
		_, _ = m.j.Discord.GetSession().ChannelMessageSend(m.j.Discord.GetLogChannelID(), "["+channel+"] "+message)
	}
	log.Println("[" + channel + "]\t" + message)
}

// Warning Level Alart
func (m *LogCore) Warning(channel string, message string) {
	if m.j.Discord != nil && m.j.Discord.IsConnected() {
		_, _ = m.j.Discord.GetSession().ChannelMessageSend(m.j.Discord.GetLogChannelID(), "["+channel+"] "+message)
	}
	log.Println("[" + channel + "]\t" + message)
}

// Error Level Alart
func (m *LogCore) Error(channel string, message string) {
	if m.j.Discord != nil && m.j.Discord.IsConnected() {
		_, _ = m.j.Discord.GetSession().ChannelMessageSend(m.j.Discord.GetLogChannelID(), "["+channel+"] "+message)
	}
	log.Println("[" + channel + "]\t" + message)
}

// Fatal Level Alart
func (m *LogCore) Fatal(channel string, message string) {
	if m.j.Discord != nil && m.j.Discord.IsConnected() {
		_, _ = m.j.Discord.GetSession().ChannelMessageSend(m.j.Discord.GetLogChannelID(), "["+channel+"] "+message)
	}
	log.Println("[" + channel + "]\t" + message)
}

/*
	switch class {
	case "ERROR":
		// Full Message Background Color
		color.Set(color.FgHiRed, color.Bold)
		log.Println(channel + "\t" + message)
		break

	case "IMPORTANT":
		// Full Message Colored Text
		if channel == "SPOTIFY" {
			color.Set(color.FgGreen)
		} else if channel == "TWITCH" {
			color.Set(color.FgMagenta)
		} else if channel == "SYSTEM" {
			color.Set(color.FgBlue)
		} else if channel == "OVERLAY" {
			color.Set(color.FgCyan)
		} else if channel == "WORKING" {
			color.Set(color.FgHiBlue)
		}
		log.Println(channel + "\t" + message)
		break
	default:
		if channel == "SPOTIFY" {
			channel = color.GreenString(channel)
		} else if channel == "TWITCH" {
			channel = color.MagentaString(channel)
		} else if channel == "SYSTEM" {
			channel = color.BlueString(channel)
		} else if channel == "OVERLAY" {
			channel = color.CyanString(channel)
		} else if channel == "WORKING" {
			channel = color.HiBlueString(channel)
		}

		// Normal (Just Channel Color)
		log.Println(channel + "\t" + message)
		break
	}
	// Reset Coloring
	color.Unset()*/
