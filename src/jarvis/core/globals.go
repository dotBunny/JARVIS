package core

// LineSpacer over text to align with normal start
const LineSpacer string = "\t\t\t\t"

// LetterBytes is used to generate random hashes
const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//CommandPrefix for commands in channel
const CommandPrefix = "!"

// CommandAccessPublic sets the command to be accessible across all channels the bot is in
const CommandAccessPublic = 1

// CommandAccessModerator sets for moderator list
const CommandAccessModerator = 2

// CommandAccessAdmin sets for admin list
const CommandAccessAdmin = 3

// CommandAccessLog sets the command to be accessible in the log channel only
const CommandAccessLog = 4

// var (
// 	// GameDevMeetupIconURI - PTBO Game Dev Thumbnail
// 	LineSpacer = func() string { return "\t\t\t\t" }

// 	// GameJamDiscordURI - PTBO Game Jam Discord Server Link
// 	GameJamDiscordURI = func() string { return "http://discord.me/ptbogamejam" }

// 	// GameJamIconURI - PTBO Game Jam Thumbnail
// 	GameJamIconURI = func() string { return "http://ptbogamejam.com/files/bot/ptbogamejam-icon.jpg" }

// 	// InternalDiscordURI - dotBunny Discord Server Link
// 	InternalDiscordURI = func() string { return "https://discord.me/dotbunny" }

// 	// InternalIconURI - dotBunny Logo
// 	InternalIconURI = func() string { return "https://dl.dropboxusercontent.com/u/118962/dotBunny/dotBot/icon.png" }

// 	// InternalGuildID - dotBunny Server ID
// 	InternalGuildID = func() string { return "269980984903073794" }

// 	// GameJamGuildID - PTBO Game Jam Server Id
// 	GameJamGuildID = func() string { return "269979472827121666" }
