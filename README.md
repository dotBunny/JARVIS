# JARVIS: Streamer's Butler
An optimized console application comprised of features from a **streamer's wishlist**. 

## Credits
This couldn't of been made without all of the awesome developers out there making their work availble to the public. 
  
Is JARVIS helping you? Show some love for [@reapazor](http://twitch.tv/reapazor) and [@dotbunny](http://twitter.com/dotbunny/) on Twitter!

## Feature Overview
* Simple Notifications

![Console](https://dl.dropboxusercontent.com/u/118962/JARVIS/console.png)

* Text File Data
    * Spotify Latest Song w/ Truncation
    * Twitch Latest Follower
    * Twitch Latest Subscriber (Not Fully Finished)
    * Working On Text
* Image File Data
    * Spotify Latest Song Album Cover
* Configurable Overlay (For Browser Sources)
    * Accessible Local Data Endpoints
    * Page hosting platform
* Twitch Chat Integration
* Console Commands!

![Outputs](https://dl.dropboxusercontent.com/u/118962/JARVIS/outputs.png)

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

For a detailed breakdown of the configuration file, please have a look at the wiki's [Configuration](https://github.com/dotBunny/JARVIS/wiki/Configuration) page.

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
| `twitch.say` | `t` | Send a message to your Twitch channel. | _t Hello World!_ |
| `twitch.whisper` | `w` | Send a whisper to someone on the Twitch IRC server. | _w reapazor You are awesome!_ |
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
| http://localhost:8080/workingon | Your last set _Working On_ text |

## Feature Requests
Drop them in the [Issues](https://github.com/dotBunny/JARVIS/issues) section, and mark them as an enhancement (label). Please understand that this is just a side project, resulting from not liking what was currently available.