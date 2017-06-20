# JARVIS: Streamer's Butler
A digital assistant comprised of features from a **streamer's wishlist** 

## Credits
This couldn't of been made without all of the awesome developers out there making their work availble to the public. 
  
Is JARVIS helping you? Show some love for [@reapazor](http://twitch.tv/reapazor) and [@dotbunny](http://twitter.com/dotbunny/) on Twitter!

## Feature Overview
* Simple Notification Log
* Output Data Files
    * Spotify latest song w/ truncation
    * Spotify latest album cover image
    * Spotify latest URL
    * Twitch latest follower
    * Twitch latest followers
    * Twitch channel views
    * Twitch channel display name
    * Twitch current game
    * Twitch current viewer count
    * Working on text
    * Coffee count
    * Crash count
    * Saves count
    * ... easy to add more counters ...
* Configurable Overlay (For Browser Sources)
    * Accessible local data endpoints
    * Also lets you host any page!
<!--* Twitch Chat Integration
    * Direct message support included-->
<!--* Interactive Console-->
* Optimized!
    * < 15 MB memory footprint _- It's 10.4 MB on our test systems!_

## Installation

### Prebuilt Version
While not always the latest and greatest, occasionally we will update the "easy mode".  

[JARVIS-0.5.0.zip](https://github.com/dotBunny/JARVIS/releases/download/0.5.0/JARVIS-0.5.0.zip)

_Make sure to edit the jarvis.json configuration file!_

### Compile From Source

Make sure you install the necessary libraries for JARVIS to do his thing.
```bash
go get github.com/zmb3/spotify  
go get github.com/chosenken/twitch2go  
go get github.com/andygrunwald/go-jira
go get github.com/fatih/color
go get github.com/atotto/clipboard
go get github.com/akavel/rsrc
go get github.com/skratchdot/open-golang/open
go get github.com/getlantern/filepersist
go get github.com/thoj/go-ircevent
go get github.com/getlantern/systray
go get github.com/matryer/resync
go get github.com/bwmarrin/discordgo
go get github.com/jpillora/backoff
```

Depending on your platform you may need to adjust the build scripts, they work on macOS and Linux varieties. We just haven't made the windows equivalents.

_It's important that GO's 'bin' is in your path (on macOS/Linux) for the build scripts to work. A default GO install often does not include it, so you must manually go back and add it._

## Configuration File
The only arguement used during execution of JARVIS is a path to the configuration JSON file, if that argument is not present, JARVIS assumes it is in the launch directory.

_This is a work in progress, there are a lot more options in the file, you simply need to go looking_

## [General]
| Option        | Description  | Type  | Example  |
| :------------ |:-------------| :-----| :------- |
| `OutputPath` | The absolute path to where the outputted data files should be placed  | _string_ |`"/Users/reapazor/StreamingData/"`|
| `Prefix` | The prefix applied to messages sent to Discord with general purpose, also tied to icons which appear in the log channel. | string | `<:jarvis:326026030458601472> ` |

## [Spotify]
| Option        | Description  | Type  | Example  |
| :------------ |:-------------| :-----| :------- |
| `PollingFrequency` | How often should Spotify be polled for new information on what's playing/happening. Current recommendation is to keep this at every 5 seconds. | _string_ | `"5s"` |
| `ClientID` | The `ClientID` can be found on your [Spotify Developer](https://developer.spotify.com/my-applications/#!/applications) page for the app; you most likely will need to create an app first to find it. | _string_ | `"7d90d691a1194380a3704dfb818x8cb1"` |
| `ClientSecret` | Same idea as the `ClientID`, it can be found in the same spot, right underneath. | _string_ | `"520dab945cbd4d738df58a124826a91c"` |
| `TruncateTrackLength` | The character length at which the combined artist and track name will be truncated | _integer_ | `85` |
| `TruncateTrackRunes` | The characters to append when truncating with the `TruncateTrackLength`| _string_ | `"..."` |

## [Twitch]
| Option        | Description  | Type  | Example  |
| :------------ |:-------------| :-----| :------- |
| `PollingFrequency` | How often should Twitch be polled for new information. Current recommendation is to keep this at every 10 seconds. | _string_ | `"10s"` |
| `ClientID` | The `ClientID` can be found on your [Twitch Connections](https://www.twitch.tv/settings/connections) page for the app; you most likely will need to register an app first to find it. | _string_ | `"d9srlt99fyxzrwa9k9ad2zjzjzl3xh"` |
| `ClientSecret` | Same idea as the `ClientID`, it can be found in the same spot, right underneath. You may need to click the _New Secret_ button. | _string_ | `"3owshhw8ukfp2x3i34v3mnh5sjsgo0"` |
| `ChannelID` | This is the numerical identifier of your channel, it isn't so simple to get off hand. Check the section below on one way to get it. | _integer_ | `21139969` |
| `LastFollowersCount` | The number of followers to store in the previously followed list. Twitch typically returns 25 records (even if you request less) and caps it at 100 entirely. | _integer_ | `10` |

### Get Your Twitch Channel ID
Hop on over to terminal and fill this command out, and it will return some JSON with your "ChannelID" listed in it.
```bash
curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>
```

## Help
Once you have managed to wrangle [GO](https://golang.org/) into compiling the source, you simply need to run the executable. The app will spawn browser tabs to process the OAuth2 login through.You will need to do this every time the app stars, thus, it has been automated as best as can be.

Notice the tray icon, it has a menu.

## Overlay
As of 0.1.1, the "Overlay" feature is experimental, but in theory you can create many things with it.  In tools like OBS, you would add a browser source and set it to `http://localhost:8080/overlay` and it will serve the content there.

There is also a neat little feature where you can server other html files in the www folder, http://localhost:8080/page/?spotify-image.html for example will serve a browser source version of the spotify image that refreshes every 5 seconds. 

### Data Endpoints
While JARVIS is running, there are numerous endpoints available for extraction data (these are only a few of them),  outside of the file repository:

| Endpoint        | Data |
| :------------- | :---- |
| http://localhost:8080/spotify/track | The current track text from Spotify |
| http://localhost:8080/spotify/image | The raw image data from Spotify |
| http://localhost:8080/twitch/follower/last  | The last person to follow you on Twitch |
| http://localhost:8080/twitch/viewers/current | The number of people watching the Twitch stream |
| http://localhost:8080/stats/workingon | Your last set _Working On_ text |
| http://localhost:8080/stats/coffee | Your current coffee count |

## Feature Requests
Drop them in the [Issues](https://github.com/dotBunny/JARVIS/issues) section, and mark them as an enhancement (label). Please understand that this is just a side project, resulting from not liking what was currently available.