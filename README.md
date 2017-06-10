## Installation

Make sure you install the necessary libraries for JARVIS to do his thing.
```bash
go get github.com/zmb3/spotify  
go get github.com/chosenken/twitch2go  
go get github.com/andygrunwald/go-jira
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
>OAuth = "TWITCH OAUTH STRING"  
>ChannelID = "YOUR CHANNEL ID"   
  
## Help
Once you have managed to wrangle [GO](https://golang.org/) into compiling the source, you simply need to run the executable.

`CTRL-C` to Exit