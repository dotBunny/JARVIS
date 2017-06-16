package core

import (
	"io"
	"log"
	"os"
	"path"
	"strings"
)

// LogType Classification
type LogType uint

const (
	// MESSAGE for the sake of logging
	MESSAGE LogType = iota
	// WARNING information to be displayed
	WARNING
	// ERROR needs fixing
	ERROR
	// FATAL application is shutting down
	FATAL
)

var (
	logTypes = []string{"MESSAGE", "WARNING", "ERROR", "FATAL"}
	logFile  *os.File
)

// LogCore Class
type LogCore struct {
	channels map[string]string

	j *JARVIS
}

// Initialize the Logging Module
func (m *LogCore) Initialize(jarvisInstance *JARVIS) {

	// Create isntance of LogCore
	m = new(LogCore)

	// Assign JARVIS (circle!)
	jarvisInstance.Log = m
	m.j = jarvisInstance

	m.channels = make(map[string]string)

	m.RegisterChannel("Core", "white")
	m.RegisterChannel("System", "grey")

	logFile, err := os.OpenFile(path.Join(m.j.Config.Settings.OutputPath, "jarvis.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// Setup echoing to log file and console
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	// TODO: This file is not being closed properly? cause its in a class
	defer logFile.Close()

}

// RegisterChannel for use with loggin
func (m *LogCore) RegisterChannel(tag string, color string) {
	m.channels[strings.ToUpper(tag)] = color
}

// Raw Alet
func (m *LogCore) Raw(logType LogType, channel string, message string) {
	log.Println("[" + channel + "]\t" + message)
}

// Message Level Alart
func (m *LogCore) Message(channel string, message string) {
	log.Println("[" + channel + "]\t" + message)
}

// Warning Level Alart
func (m *LogCore) Warning(channel string, message string) {
	log.Println("[" + channel + "]\t" + message)
}

// Error Level Alart
func (m *LogCore) Error(channel string, message string) {
	log.Println("[" + channel + "]\t" + message)
}

// Fatal Level Alart
func (m *LogCore) Fatal(channel string, message string) {
	log.Println("[" + channel + "]\t" + message)
}

func (pn LogType) name() string {
	return logTypes[pn]
}
func (pn LogType) ordinal() int {
	return int(pn)
}
func (pn LogType) String() string {
	return logTypes[pn]
}
func (pn LogType) values() *[]string {
	return &logTypes
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
