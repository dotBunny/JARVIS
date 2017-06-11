# JARVIS: Streamer's Butler
A optimized console application comprised of features from a **streamer's wishlist**. 

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
* Image File Data
    * Spotify Latest Song Album Cover
* Configurable Overlay (For Browser Sources)
    * Accessible Local Data Endpoints

![Outputs](https://dl.dropboxusercontent.com/u/118962/JARVIS/outputs.png)

## Installation

### Prebuilt Version
While not always the latest and greatest, occasionally we will update the "easy mode".  

[JARVIS-0.1.zip](https://github.com/dotBunny/JARVIS/releases/download/0.1/JARVIS-0.1.zip)

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
```

Depending on your platform you may need to adjust the build scripts, they work on macOS and Linux varieties. We just haven't made the windows equivalents.

_It's important that GO's 'bin' is in your path (on macOS/Linux) for the build scripts to work. A default GO install often does not include it, so you must manually go back and add it._

## Configuration File
The configuration file needs to be in the same directory as the executable, and named `jarvis.toml`

### [General]
| Option        | Description           | Type | Example  |
| ------------- |:-------------| :-----| :-----|
| `OutputPat` | The absolute path to where the outputted data files should be placed  | _string_ |"/Users/reapazor/StreamingData/"|
| `ServerPort` | The port which the callback/overlay server listens on; this by default should be 8080. If you change it, you must edit your overlay's HTML files to reflect the changed port. | _integer_ | 8080 |

### [Spotify]
| Option        | Description           | Type | Example  |
| ------------- |:-------------| :-----| :-----|
| `Enabled` | Should JARVIS attempt connections to Spotify; do you want to use the Spotify module? | _boolean_ | true |
| `Output` | Should data files for Spotify be output to the `OutputPath` | _boolean_ | true |
| `PollingFrequency` | How often should Spotify be polled for new information on what's playing/happening. Current recommendation is to keep this at every 5 seconds. | _string_ | "5s" |
| `ClientID` | The `ClientID` can be found on your [Spotify Developer](https://developer.spotify.com/my-applications/#!/applications) page for the app; you most likely will need to create an app first to find it. | _string_ | "7d90d691a1194380a3704dfb818x8cb1" |
| `ClientSecret` | Same idea as the `ClientID`, it can be found in the same spot, right underneath. | _string_ | "520dab945cbd4d738df58a124826a91c" |
| `Callback` | This is the endpoint of the listen server that will take the response from Spotify during the login process. You **must** add the full path (http://localhost:8080/callbackSpotify) on your Spotify developer page to the _Redirect URIs_ section. | _string_ | "/callbackSpotify" |
| `TruncateTrackLength` | The character length at which the combined artist and track name will be truncated | _integer_ | 85 |
| `TruncateTrackRunes` | The characters to append when truncating with the `TruncateTrackLength`| _string_ | "..." |

### [Twitch]
| Option        | Description           | Type | Example  |
| ------------- |:-------------| :-----| :-----|
| `Enabled` | Should JARVIS attempt connections to Twitch; do you want to use the Twitch module? | _boolean_ | true |
| `Output` | Should data files for Twitch be output to the `OutputPath` | _boolean_ | true |
| `PollingFrequency` | How often should Twitch be polled for new information. Current recommendation is to keep this at every 10 seconds. | _string_ | "10s" |
| `ClientID` | The `ClientID` can be found on your [Twitch Connections](https://www.twitch.tv/settings/connections) page for the app; you most likely will need to register an app first to find it. | _string_ | "d9srlt99fyxzrwa9k9ad2zjzjzl3xh" |
| `ClientSecret` | Same idea as the `ClientID`, it can be found in the same spot, right underneath. You may need to click the _New Secret_ button. | _string_ | "3owshhw8ukfp2x3i34v3mnh5sjsgo0" |
| `ChannelID` | This is the numerical identifier of your channel, it isn't so simple to get off hand. Check the section below on one way to get it. | _integer_ | 21139969 |

  
### Get Your Twitch Channel ID
Hop on over to terminal and fill this command out, and it will return some JSON with your "ChannelID" listed in it.
```bash
curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>
```

## Help
Once you have managed to wrangle [GO](https://golang.org/) into compiling the source, you simply need to run the executable. You will be prompted to login to Spotify (to provide a token for the app). You will need to do this every time the app stars, thus, the URL needed is copied to your clipboard for a quick paste into your nearest local browser.

`CTRL-C` to Exit

## Overlay
As of 0.1.1, the "Overlay" feature is experimental, but in theory you can create many things with it.  In tools like OBS, you would add a browser source and set it to `http://localhost:[ServerPort]/overlay` and it will serve the content there

### Data Endpoints
While JARVIS is running, there are numerous endpoints available for extraction data, outside of the file repository:

>http://localhost:[ServerPort]/spotify/track  
>http://localhost:[ServerPort]/spotify/image  
>http://localhost:[ServerPort]/twitch/follower/last  

## Feature Requests
Drop them in the [Issues](https://github.com/dotBunny/JARVIS/issues) section, and mark them as an enhancement (label). Please understand that this is just a side project, resulting from not liking what was currently available.