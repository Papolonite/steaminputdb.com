package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Alia5/steaminputdb.com/api/ctx"
	"github.com/Alia5/steaminputdb.com/api/memcache"
	"github.com/Alia5/steaminputdb.com/api/steam/auth"
	"github.com/Alia5/steaminputdb.com/steamapi"
	"github.com/danielgtaylor/huma/v2"
)

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

type UserInfoRequest struct {
	UserID string `query:"user_id,omitempty,omitzero"` // this is the user_id query parameter, but it's named something else to avoid confusion with the steamID that can be extracted from the context. if this is 0, we'll attempt to use the steamID from the context instead.
}

func RegisterRoutes(a huma.API, opts ...bool) {
	var useMemCache bool
	if len(opts) > 0 {
		useMemCache = opts[0]
	} else {
		useMemCache = true
	}
	cache := memcache.New(30*time.Minute, 1000)
	huma.Register(
		a,
		huma.Operation{
			Method:  http.MethodGet,
			Path:    "/v1/steam/userinfo",
			Tags:    []string{"steam", "user"},
			Summary: "Get Steam user info",
			Description: `Retrieve user information from Steam for the provided userId,  
or for attempt authenticated user if no userId is provided  
Returns 401 if no id provided and token is invalid and 400 if everything is missing`,
			Errors: []int{
				http.StatusBadGateway, http.StatusUnauthorized, http.StatusNotFound, http.StatusBadRequest,
			},
			Middlewares: huma.Middlewares{
				auth.ExtractSteamIDMiddleware,
			},
		},
		func(c context.Context, req *UserInfoRequest) (*Response, error) {
			var steamID string
			if req.UserID == "" {
				var ok bool
				steamID, ok = c.Value(ctx.KeySteamID).(string)
				if !ok || steamID == "" {
					return nil, huma.Error401Unauthorized("no authentication token provided")
				}
			} else {
				steamID = req.UserID
				if useMemCache {
					cached, ok := memcache.Get[*Response](
						cache,
						steamID,
					)
					if ok {
						return cached, nil
					}
				}
			}
			if steamID == "" {
				return nil, huma.Error400BadRequest("no user ID provided")
			}

			info, err := steamapi.DefaultClient.GetPlayerSummaries(c, steamID)
			if err != nil {
				if errors.Is(err, steamapi.ErrRequest) {
					return nil, huma.Error502BadGateway("failed to get steam user info: %v", err)
				}
				return nil, err
			}

			if len(info.Response.Players) == 0 {
				return nil, huma.Error404NotFound("steam user not found")
			}

			player := info.Response.Players[0]

			res := &Response{
				Body: PlayerInfo{
					CommunityVisibilityState: player.Communityvisibilitystate,
					PersonaName:              player.Personaname,
					ProfileURL:               player.Profileurl,
					Avatar:                   player.Avatar,
					AvatarMedium:             player.Avatarmedium,
					AvatarFull:               player.Avatarfull,
					AvatarHash:               player.Avatarhash,
					LastLogOff:               time.Unix(int64(player.Lastlogoff), 0).UTC(),
					PrimaryClanID:            player.Primaryclanid,
					TimeCreated:              time.Unix(int64(player.Timecreated), 0).UTC(),
					LocCountryCode:           player.Loccountrycode,
				},
			}
			if useMemCache && req.UserID != "" {
				cache.Store(steamID, res)
			}
			return res, nil
		},
	)
}
