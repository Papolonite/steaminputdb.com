package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Alia5/steaminputdb.com/api/ctx"
	"github.com/Alia5/steaminputdb.com/api/steam/auth"
	"github.com/Alia5/steaminputdb.com/steamapi"
	"github.com/danielgtaylor/huma/v2"
)

const steamSummaryURL = "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2"

type PlayerInfo struct {
	CommunityVisibilityState int       `json:"communityvisibilitystate"` // TODO:enum
	PersonaName              string    `json:"personaname"`
	ProfileURL               string    `json:"profileurl"`
	Avatar                   string    `json:"avatar"`
	AvatarMedium             string    `json:"avatarmedium"`
	AvatarFull               string    `json:"avatarfull"`
	AvatarHash               string    `json:"avatarhash"`
	LastLogOff               time.Time `json:"lastlogoff"`
	PrimaryClanID            string    `json:"primaryclanid"`
	TimeCreated              time.Time `json:"timecreated"`
	LocCountryCode           string    `json:"loccountrycode"`
}

type Response struct {
	Body PlayerInfo
}

func RegisterWithURL(a huma.API, steamURL string) {
	registerRoutes(a, steamURL)
}

func RegisterRoutes(a huma.API) {
	registerRoutes(a, steamSummaryURL)
}
func registerRoutes(a huma.API, loginURL string) {
	huma.Register(
		a,
		huma.Operation{
			Method:      http.MethodPost,
			Path:        "/v1/steam/userinfo",
			Tags:        []string{"steam", "user"},
			Summary:     "Get Steam user info",
			Description: "Retrieve user information from Steam for the currently authenticated user",
			Errors: []int{
				http.StatusBadGateway, http.StatusUnauthorized,
			},
			Middlewares: huma.Middlewares{
				auth.Middleware(a),
			},
		},
		func(c context.Context, _ *struct{}) (*Response, error) {

			steamID, ok := c.Value(ctx.KeySteamID).(string)
			if !ok || steamID == "" {
				return nil, huma.Error401Unauthorized("missing steamid")
			}

			info, err := steamapi.DefaultClient.GetPlayerSummaries(c, steamID)
			if err != nil {
				if errors.Is(err, steamapi.ErrRequest) {
					return nil, huma.Error502BadGateway("failed to get steam user info: %v", err)
				}
				return nil, err
			}

			if len(info.Response.Players) == 0 {
				return nil, huma.Error401Unauthorized("steam user not found")
			}

			player := info.Response.Players[0]

			return &Response{
				Body: PlayerInfo{
					CommunityVisibilityState: player.Communityvisibilitystate,
					PersonaName:              player.Personaname,
					ProfileURL:               player.Profileurl,
					Avatar:                   player.Avatar,
					AvatarMedium:             player.Avatarmedium,
					AvatarFull:               player.Avatarfull,
					AvatarHash:               player.Avatarhash,
					LastLogOff:               time.Unix(int64(player.Lastlogoff), 0),
					PrimaryClanID:            player.Primaryclanid,
					TimeCreated:              time.Unix(int64(player.Timecreated), 0),
					LocCountryCode:           player.Loccountrycode,
				},
			}, nil
		},
	)
}
