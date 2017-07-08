#!/usr/bin/osascript

# Properties
property Stream : "1"

# Function
property ArgumentCount : 0
property Arguments : ""
property Command : ""

# Arguement Catch
property Shot1Arg : "" # Empty
property Shot2Arg : "" # Overlays
property Shot3Arg : "" # Video
property Shot4Arg : "" # Audio Layer
property Shot5Arg : "" # Mic On/Off

# Setup Variables
on Setup()
	
	# Get Wirecast Document
	tell application "Wirecast"
		set my Stream to last document
	end tell
	
	# Get Command
	if ArgumentCount > 0 then
		set my Command to item 1 of Arguments
		
		# Assign Arguments
		if ArgumentCount > 1 then
			if ArgumentCount is 2 then
				set my Shot1Arg to item 2 of Arguments
			else if ArgumentCount is 3 then
				set my Shot1Arg to item 2 of Arguments
				set my Shot2Arg to item 3 of Arguments
			else if ArgumentCount is 4 then
				set my Shot1Arg to item 2 of Arguments
				set my Shot2Arg to item 3 of Arguments
				set my Shot3Arg to item 4 of Arguments
			else if ArgumentCount is 5 then
				set my Shot1Arg to item 2 of Arguments
				set my Shot2Arg to item 3 of Arguments
				set my Shot3Arg to item 4 of Arguments
				set my Shot4Arg to item 5 of Arguments
			else
				set my Shot1Arg to item 2 of Arguments
				set my Shot2Arg to item 3 of Arguments
				set my Shot3Arg to item 4 of Arguments
				set my Shot4Arg to item 5 of Arguments
				set my Shot5Arg to item 6 of Arguments
			end if
		end if
	end if
end Setup

on SetTransition(number, Speed)
	tell application "Wirecast"
		set my Stream's active transition popup to number
		set my Stream's transition speed to Speed
	end tell
end SetTransition

# Simplified Change Function
on ChangeLayer(LayerName, ShotName)
	tell application "Wirecast"
		set active shot of the layer named LayerName of Stream to the shot named ShotName of Stream
	end tell
end ChangeLayer

on CommandSwitch()
	tell application "Wirecast"
		set my Stream's auto live to true
		
		# Change Layer 1
		if my Shot1Arg is not "" then
			my ChangeLayer("Master Layer 1", Shot1Arg)
		end if
		
		# Change Layer 2
		if my Shot2Arg is not "" then
			my ChangeLayer("Master Layer 2", Shot2Arg)
		end if
		
		# Change Layer 3
		if my Shot3Arg is not "" then
			my ChangeLayer("Master Layer 3", Shot3Arg)
		end if
		
		# Change Layer 4
		if my Shot4Arg is not "" then
			my ChangeLayer("Master Layer 4", Shot4Arg)
		end if
		
		# Change Layer 5
		if my Shot5Arg is not "" then
			my ChangeLayer("Master Layer 5", Shot5Arg)
		end if
		
		set my Stream's auto live to false
	end tell
end CommandSwitch

on CommandToggleMic()
	tell application "Wirecast"
		
		
	end tell
end CommandToggleMic

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
	
	my Setup()
	
	# Action Gate
	if my Command is "switch" then
		my CommandSwitch()
	end if
	
end run
