package steamapi

import (
	"context"
	"net/url"
)

// Is somehow not in the Protos....

func (c *Client) GetPlayerSummaries(ctx context.Context, steamID string) (*PlayerSummaries, error) {
	resp := &PlayerSummaries{}
	err := c.GetJSON(
		ctx,
		Endpoint{Interface: "ISteamUser", Method: "GetPlayerSummaries", Version: "2"},
		&struct{}{},
		&url.Values{"steamids": []string{steamID}},
		resp,
		&Auth{Key: c.apiKey},
	)
	return resp, err
}

type PlayerSummaries struct {
	Response Response `json:"response"`
}
type Players struct {
	Steamid                  string `json:"steamid"`
	Communityvisibilitystate int    `json:"communityvisibilitystate"`
	Profilestate             int    `json:"profilestate"`
	Personaname              string `json:"personaname"`
	Profileurl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	Avatarmedium             string `json:"avatarmedium"`
	Avatarfull               string `json:"avatarfull"`
	Avatarhash               string `json:"avatarhash"`
	Lastlogoff               int    `json:"lastlogoff"`
	Personastate             int    `json:"personastate"`
	Realname                 string `json:"realname"`
	Primaryclanid            string `json:"primaryclanid"`
	Timecreated              int    `json:"timecreated"`
	Personastateflags        int    `json:"personastateflags"`
	Loccountrycode           string `json:"loccountrycode"`
}
type Response struct {
	Players []Players `json:"players"`
}
