#!/usr/bin/osascript

# Properties
property Stream : "1"

# Function
property ArgumentCount : 0
property Command : ""
property OutputPath : ""

on CommandNext()
	tell application "iTunes"
		tell app "iTunes" to play next track
	end tell
end CommandNext

on CommandStatus()
	tell application "iTunes"
		if exists name of current track then
			set aTrack to the current track
			set aName to the name of aTrack
			set aArtist to the artist of aTrack
			
			if (count of artwork of aTrack) >= 1 then
				set anArtwork to data of artwork 1 of aTrack
				set imageFile to POSIX path of file (OutputPath & "iTunes_CurrentImage.jpg")
				
				-- Output Current Image
				set the openImageFile to open for access alias imageFile with write permission
				try
					set eof of the openImageFile to 0
					write anArtwork to the openImageFile starting at eof
					close access the openImageFile
				on error
					try
						close access alias imageFile
					end try
				end try
				
			end if
			
			-- Output Current Data
			set outputData to aArtist & "," & aName
			set infoFile to POSIX path of file (OutputPath & "iTunes_CurrentSong.txt")
			set the openInfoFile to open for access alias infoFile with write permission
			try
				set eof of the openInfoFile to 0
				write outputData to the openInfoFile starting at eof
				close access the openInfoFile
			on error
				try
					close access alias infoFile
				end try
			end try
		end if
	end tell
end CommandStatus

on CommandPlayPause()
	tell application "iTunes"
		tell app "iTunes" to playpause	
	end tell
end CommandPlayPause

# Execution
on run argv
	
	# Check Arguments
	try
		set my ArgumentCount to (count of argv)
		if ArgumentCount is greater than 0 then
			set my Arguments to argv
		end if
	on error
		set my ArgumentCount to 0
	end try

	if ArgumentCount is 1 then
		set my Command to item 1 of Arguments
	end if
	if ArgumentCount is 2 then
		set my OutputPath to item 2 of Arguments
	end if
	
	# Action Gate
	if my Command is "next" then
		my CommandNext()
	end if
	if my Command is "status" then
		my CommandStatus()
	end if
	if my Command is "playpause" then
		my CommandPlayPause()
	end if
	
end run
