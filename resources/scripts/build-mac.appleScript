#!/usr/bin/osascript
tell application "CLion"
	activate
end tell

tell application "System Events"
	tell process "CLion-EAP"
		tell menu bar 1
			tell menu bar item "Run"
				tell menu "Run"
					click menu item "Build"
				end tell
			end tell
		end tell
	end tell
end tell