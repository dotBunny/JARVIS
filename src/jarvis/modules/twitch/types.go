package twitch

// https://github.com/chosenken/twitch2go/blob/master/types.go

import (
	"encoding/json"
	"time"
)

type Direction string
type SortBy string
type VideoSort string

const (
	ASC           Direction = "asc"
	DESC          Direction = "desc"
	CreatedAt     SortBy    = "created_at"
	LastBroadcast SortBy    = "last_broadcast"
	Login         SortBy    = "login"
	Views         VideoSort = "views"
	Time          VideoSort = "time"
)

// Channel Twitch Channel Data
type Channel struct {
	Mature                       bool        `json:"mature"`
	Status                       string      `json:"status"`
	BroadcasterLanguage          string      `json:"broadcaster_language"`
	DisplayName                  string      `json:"display_name"`
	Game                         string      `json:"game"`
	Language                     string      `json:"language"`
	ID                           json.Number `json:"_id,number"`
	Name                         string      `json:"name"`
	CreatedAt                    time.Time   `json:"created_at"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	Logo                         string      `json:"logo"`
	VideoBanner                  string      `json:"video_banner"`
	ProfileBanner                string      `json:"profile_banner"`
	ProfileBannerBackgroundColor string      `json:"profile_banner_background_color"`
	Partner                      bool        `json:"partner"`
	URL                          string      `json:"url"`
	Views                        uint        `json:"views"`
	Followers                    uint        `json:"followers"`
}

type Post struct {
	ID        json.Number `json:"id,number"`
	CreatedAt time.Time   `json:"created_at"`
	Deleted   bool        `json:"deleted"`
	Emotes    []string    `json:"emotes"`
	Body      string      `json:"body"`
	User      User        `json:"user"`
}

// Follower data for twitch channel
type Follow struct {
	CreatedAt     time.Time         `json:"created_at"`
	Links         map[string]string `json:"_links"`
	Notifications bool              `json:"notifications"`
	User          User              `json:"user"`
	Channel       Channel           `json:"channel"`
}

type Followers struct {
	Total   uint     `json:"_total"`
	Cursor  string   `json:"_cursor"`
	Follows []Follow `json:"follows"`
}

type Subscription struct {
	ID        json.Number `json:"_id,number"`
	CreatedAt time.Time   `json:"created_at"`
	User      User        `json:"user"`
}

type Subscribers struct {
	Total         uint           `json:"_total"`
	Cursor        string         `json:"_cursor"`
	Subscriptions []Subscription `json:"subscriptions"`
}

type UserSearchResult struct {
	Total uint   `json:"_total"`
	Users []User `json:"users"`
}

// User Twitch User Data
type User struct {
	Type             string        `json:"type"`
	Name             string        `json:"name"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
	Logo             string        `json:"logo"`
	ID               json.Number   `json:"_id,number"`
	DisplayName      string        `json:"display_name"`
	Bio              string        `json:"bio"`
	Email            string        `json:"email"`
	EmailVerified    bool          `json:"email_verified"`
	Partnered        bool          `json:"partnered"`
	TwitterConnected bool          `json:"twitter_connected"`
	Notifications    Notifications `json:"notifications"`
}

type Notifications struct {
	Email bool `json:"email"`
	Push  bool `json:"push"`
}

type Editors struct {
	Users []User `json:"Users"`
}

type Videos struct {
	Total  uint    `json:"_total"`
	Videos []Video `json:"videos"`
}

type Video struct {
	ID              string      `json:"_id"`
	BroadcastID     json.Number `json:"broadcast_id,number"`
	BroadcastType   string      `json:"broadcast_type"`
	Channel         Channel     `json:"channel"`
	CreatedAt       time.Time   `json:"created_at"`
	Description     string      `json:"description"`
	DescriptionHTML string      `json:"description_html"`
	Fps             Fps         `json:"fps"`
	Game            string      `json:"game"`
	Language        string      `json:"language"`
	Length          uint        `json:"length"`
	Preview         Preview     `json:"preview"`
	PublishedAt     time.Time   `json:"published_at"`
	Resolutions     Resolutions `json:"resolutions"`
	Status          string      `json:"status"`
	TagList         string      `json:"tag_list"`
	Thumbnails      Thumbnails  `json:"thumbnails"`
	Title           string      `json:"title"`
	URL             string      `json:"url"`
	Viewable        string      `json:"viewable"`
	ViewableAt      interface{} `json:"viewable_at"`
	Views           uint        `json:"views"`
}

type Thumbnail struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Thumbnails struct {
	Large    []Thumbnail `json:"large"`
	Medium   []Thumbnail `json:"medium"`
	Small    []Thumbnail `json:"small"`
	Template []Thumbnail `json:"template"`
}

type Resolutions struct {
	Chunked string `json:"chunked"`
	High    string `json:"high"`
	Low     string `json:"low"`
	Medium  string `json:"medium"`
	Mobile  string `json:"mobile"`
}

type Preview struct {
	Large    string `json:"large"`
	Medium   string `json:"medium"`
	Small    string `json:"small"`
	Template string `json:"template"`
}

type Fps struct {
	Chunked float64 `json:"chunked"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Medium  float64 `json:"medium"`
	Mobile  float64 `json:"mobile"`
}

type Stream struct {
	ID          json.Number `json:"_id,number"`
	Game        string      `json:"game"`
	CommunityID string      `json:"community_id"`
	Viewers     uint        `json:"viewers"`
	VideoHeight uint        `json:"video_height"`
	AverageFps  float64     `json:"average_fps"`
	Delay       uint        `json:"delay"`
	CreatedAt   time.Time   `json:"created_at"`
	IsPlaylist  bool        `json:"is_playlist"`
	Preview     Preview     `json:"preview"`
	Channel     Channel     `json:"channel"`
}

type StreamResponse struct {
	Stream Stream `json:"stream"`
}

type FollowedStream struct {
	Total   uint     `json:"_total"`
	Streams []Stream `json:"streams"`
}

type ResponseError struct {
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Status  json.Number `json:"status,number"`
}

type ChatterResponse struct {
	ChatterCount uint     `json:"chatter_count"`
	Chatters     Chatters `json:"chatters"`
}

type Chatters struct {
	Moderators []string `json:"moderators"`
	Staff      []string `json:"staff"`
	Admins     []string `json:"admins"`
	GlobalMods []string `json:"global_mods"`
	Viewers    []string `json:"viewers"`
}
