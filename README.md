# JARVIS: Streamer's Butler
An optimized console application comprised of features from a **streamer's wishlist**. 

## Credits
This couldn't of been made without all of the awesome developers out there making their work availble to the public. 
  
Is JARVIS helping you? Show some love for [@reapazor](http://twitch.tv/reapazor) and [@dotbunny](http://twitter.com/dotbunny/) on Twitter!

## Feature Overview
* Simple Notification Log
* Output Data Files
    * Spotify latest song w/ truncation
    * Spotify latest album cover image
    * Twitch latest follower
    * Twitch latest followers
    * Twitch channel views
    * Twitch channel display name
    * Twitch current game
    * Twitch current viewer count
    * Working on text
* Configurable Overlay (For Browser Sources)
    * Accessible local data endpoints
    * Also lets you host any page!
* Twitch Chat Integration
    * Direct message support included
* Interactive Console
* Optimized!
    * < 15 MB memory footprint _- It's 10.4 MB on our test systems!_

## Installation

### Prebuilt Version
While not always the latest and greatest, occasionally we will update the "easy mode".  

[JARVIS-0.2.0.zip](https://github.com/dotBunny/JARVIS/releases/download/0.2.0/JARVIS-0.2.0.zip)

_Make sure to edit the jarvis.toml configuration file!_

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
go get github.com/thoj/go-ircevent
```

Depending on your platform you may need to adjust the build scripts, they work on macOS and Linux varieties. We just haven't made the windows equivalents.

_It's important that GO's 'bin' is in your path (on macOS/Linux) for the build scripts to work. A default GO install often does not include it, so you must manually go back and add it._

## Configuration File
The configuration file needs to be in the same directory as the executable, and named `jarvis.toml`

## [General]
| Option        | Description  | Type  | Example  |
| :------------ |:-------------| :-----| :------- |
| `OutputPath` | The absolute path to where the outputted data files should be placed  | _string_ |`"/Users/reapazor/StreamingData/"`|
| `ServerPort` | The port which the callback/overlay server listens on; this by default should be 8080. If you change it, you must edit your overlay's HTML files to reflect the changed port. | _integer_ | `8080` |

## [Spotify]
| Option        | Description  | Type  | Example  |
| :------------ |:-------------| :-----| :------- |
| `Enabled` | Should JARVIS attempt connections to Spotify; do you want to use the Spotify module? | _boolean_ | `true` |
| `Output` | Should data files for Spotify be output to the `OutputPath` | _boolean_ | `true` |
| `AutoLogin` | Attempt to automate the login process by automatically opening a browser tab during the authentication process. We'd close it too but that is not allowed. | _boolean_ | `true` |
| `PollingFrequency` | How often should Spotify be polled for new information on what's playing/happening. Current recommendation is to keep this at every 5 seconds. | _string_ | `"5s"` |
| `ClientID` | The `ClientID` can be found on your [Spotify Developer](https://developer.spotify.com/my-applications/#!/applications) page for the app; you most likely will need to create an app first to find it. | _string_ | `"7d90d691a1194380a3704dfb818x8cb1"` |
| `ClientSecret` | Same idea as the `ClientID`, it can be found in the same spot, right underneath. | _string_ | `"520dab945cbd4d738df58a124826a91c"` |
| `Callback` | This is the endpoint of the listen server that will take the response from Spotify during the login process. You **must** add the full path (http://localhost:8080/callbackSpotify) on your Spotify developer page to the _Redirect URIs_ section. | _string_ | `"/callbackSpotify"` |
| `TruncateTrackLength` | The character length at which the combined artist and track name will be truncated | _integer_ | `85` |
| `TruncateTrackRunes` | The characters to append when truncating with the `TruncateTrackLength`| _string_ | `"..."` |

## [Twitch]
| Option        | Description  | Type  | Example  |
| :------------ |:-------------| :-----| :------- |
| `Enabled` | Should JARVIS attempt connections to Twitch; do you want to use the Twitch module? | _boolean_ | `true` |
| `Output` | Should data files for Twitch be output to the `OutputPath` | _boolean_ | `true` |
| `PollingFrequency` | How often should Twitch be polled for new information. Current recommendation is to keep this at every 10 seconds. | _string_ | `"10s"` |
| `ClientID` | The `ClientID` can be found on your [Twitch Connections](https://www.twitch.tv/settings/connections) page for the app; you most likely will need to register an app first to find it. | _string_ | `"d9srlt99fyxzrwa9k9ad2zjzjzl3xh"` |
| `ClientSecret` | Same idea as the `ClientID`, it can be found in the same spot, right underneath. You may need to click the _New Secret_ button. | _string_ | `"3owshhw8ukfp2x3i34v3mnh5sjsgo0"` |
| `ChannelID` | This is the numerical identifier of your channel, it isn't so simple to get off hand. Check the section below on one way to get it. | _integer_ | `21139969` |
| `LastFollowersCount` | The number of followers to store in the previously followed list. Twitch typically returns 25 records (even if you request less) and caps it at 100 entirely. | _integer_ | `10` |
| `ChatEnabled` | Should the IRC server be connected too; disabling this will remove your ability to respond to messages inside of JARVIS. | _boolean_ | `false` |
| `ChatEcho` | Should the chat channel content be shown in JARVIS | _boolean_ | `true` |
| `ChatName` | The alias to use when connecting to Twitch's IRC server | _string_ | `"reapazor"` |
| `ChatToken` | The token used for connecting to Twitch's IRC servers, this can be easily found by visiting [here](http://www.twitchapps.com/tmi/). | _string_ | `"oauth:aqwegd3126l25azg0nn70if82nr9d1"` |
| `ChatChannel` | The channel to join when connected to Twitch's IRC server | _string_ | `"#reapazor"` |

## [WorkingOn]
| Option        | Description  | Type  | Example  |
| :------------ |:-------------| :-----| :------- |
| `Enabled` | Should the WorkingOn module be active? | _boolean_ | `true` |
| `Output` | Should data files for what your working on be output to the `OutputPath` | _boolean_ | `true` |

### Get Your Twitch Channel ID
Hop on over to terminal and fill this command out, and it will return some JSON with your "ChannelID" listed in it.
```bash
curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>
```

## Help
Once you have managed to wrangle [GO](https://golang.org/) into compiling the source, you simply need to run the executable. You will be prompted to login to Spotify (to provide a token for the app). You will need to do this every time the app stars, thus, the URL needed is copied to your clipboard for a quick paste into your nearest local browser. If you have left `AutoLogin` on in the Spotify settings of the `jarvis.toml` it will even open a browser automatically for you.

**Type `exit` (and press enter) or press `CTRL-C` to Exit**

### Console Commands

| Command        | Alias | Description  |  Example  |
| :------------- | :---- | :----------- | :-------- |
| `spotify.next` | `next`, `n` | Skips to the next track in the user's Spotify queue. | _next_ |
| `spotify.pause` | `p` | Pause/Play the current track in Spotify. | _p_ |
| `spotify.stats` |  | Display some stats about Spotify. | _spotify.stats_ |
| `spotify.update` | | Force polling Spotify for updates. | _spotify.update_ |
| `twitch.say` | `t` | Send a message to your Twitch channel. | _t Hello World!_ |
| `twitch.stats` |  | Display some stats about the Twitch channel/stream. | _twitch.stats_ |
| `twitch.update` | | Force polling Twitch for updates. | _twitch.update_ |
| `twitch.whisper` | `w` | Send a whisper to someone on the Twitch IRC server. | _w reapazor You are awesome!_ |
| `update` | | Force all active modules to poll their data sources for updates. | _update_ |
| `workingon` | `o` | Set your currently working on text. | _workingon The JARVIS System_ |
| `quit` | `exit`, `x` | Quit the application | _quit_ |

## Overlay
As of 0.1.1, the "Overlay" feature is experimental, but in theory you can create many things with it.  In tools like OBS, you would add a browser source and set it to `http://localhost:8080/overlay` and it will serve the content there.

There is also a neat little feature where you can server other html files in the overlay folder, http://localhost:8080/overlay/page?spotify-image.html for example will serve a browser source version of the spotify image that refreshes every 5 seconds. 

### Data Endpoints
While JARVIS is running, there are numerous endpoints available for extraction data, outside of the file repository:

| Endpoint        | Data |
| :------------- | :---- |
| http://localhost:8080/spotify/track | The current track text from Spotify |
| http://localhost:8080/spotify/image | The raw image data from Spotify |
| http://localhost:8080/twitch/follower/last  | The last person to follow you on Twitch |
| http://localhost:8080/twitch/viewers/current | The number of people watching the Twitch stream |
| http://localhost:8080/workingon | Your last set _Working On_ text |

## Feature Requests
Drop them in the [Issues](https://github.com/dotBunny/JARVIS/issues) section, and mark them as an enhancement (label). Please understand that this is just a side project, resulting from not liking what was currently available.