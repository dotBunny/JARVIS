# JARVIS: Streamer's Butler
A console application comprised of features from a streamer's wishlist.

* Simple Notifications

![Console](https://dl.dropboxusercontent.com/u/118962/JARVIS/console.png)

* Text File Data
    * Spotify Latest Song
    * Twitch Latest Follower
    * Twitch Latest Subscriber (Not Fully Finished)
* Image File Data
    * Spotify Latest Song Album Cover

![Outputs](https://dl.dropboxusercontent.com/u/118962/JARVIS/outputs.png)


## Installation

Make sure you install the necessary libraries for JARVIS to do his thing.
```bash
go get github.com/zmb3/spotify  
go get github.com/chosenken/twitch2go  
go get github.com/andygrunwald/go-jira
go get github.com/fatih/color
```

## Configuration File
The configuration file needs to be in the same directory as the executable, and named `jarvis.toml`

>[General]  
>OutputPath = "/path/where/to/save/files/"  
>ServerPort = "8080"
>  
>[Spotify]  
>ClientID = "SPOTIFY CLIENT ID"  
>ClientSecret = "SPOTIFY SECRET"  
>Callback = "/callbackEndpoint"  
>  
>[Twitch]  
>ClientID = "TWITCH CLIENT ID"  
>ClientSecret = "TWITCH SECRET"  
>OAuth = "TWITCH OAUTH STRING" #Not Used Currently  
>ChannelID = "YOUR CHANNEL ID"   
  
## Help
Once you have managed to wrangle [GO](https://golang.org/) into compiling the source, you simply need to run the executable.

`CTRL-C` to Exit

### Twitch Channel ID
Hop on over to terminal and fill this command out, and it will return some JSON with your "ChannelID" listed in it.
```bash
curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>
```