#!/usr/bin/osascript

--		set the target_file to the target_file as text


--
--on run argv
--    return "hello, " & item 1 of argv & "."
--    end run
--

tell application "iTunes"
	if exists name of current track then
		set aTrack to the current track
		set aName to the name of aTrack
		set aArtist to the artist of aTrack
		
		if (count of artwork of aTrack) ≥ 1 then
			set anArtwork to data of artwork 1 of aTrack
			
			
			set imageFile to (((path to desktop folder) as text) & "iTunes_CurrentImage.jpg")
			
			set currentImageData to (read file imageFile)
			-- Check if its the same
			if currentImageData is not anArtwork then
				-- Output Current Image
				set the openImageFile to open for access file imageFile with write permission
				try
					set eof of the openImageFile to 0
					write anArtwork to the openImageFile starting at eof
					close access the openImageFile
				on error
					try
						close access file imageFile
					end try
				end try
			end if
			
		end if
		
		-- Output Current Data
		set outputData to aArtist & "," & aName
		set infoFile to (((path to desktop folder) as text) & "iTunes_CurrentSong.txt")
		set the openInfoFile to open for access file infoFile with write permission
		try
			set eof of the openInfoFile to 0
			write outputData to the openInfoFile starting at eof
			close access the openInfoFile
		on error
			try
				close access file infoFile
			end try
		end try
	end if
end tell
