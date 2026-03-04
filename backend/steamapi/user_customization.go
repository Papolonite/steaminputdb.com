package steamapi

import "context"

func (c *Client) GetAvatarFrame(ctx context.Context, req *CPlayer_GetAvatarFrame_Request) (*CPlayer_GetAvatarFrame_Response, error) {
	resp := &CPlayer_GetAvatarFrame_Response{}
	err := GetWithResp(
		ctx,
		Endpoint{Interface: "IPlayerService", Method: "GetAvatarFrame"},
		req,
		resp,
		&Auth{Key: c.apiKey},
	)
	return resp, err
}

func (c *Client) GetProfileBackground(ctx context.Context, req *CPlayer_GetProfileBackground_Request) (*CPlayer_GetProfileBackground_Response, error) {
	resp := &CPlayer_GetProfileBackground_Response{}
	err := GetWithResp(
		ctx,
		Endpoint{Interface: "IPlayerService", Method: "GetProfileBackground"},
		req,
		resp,
		&Auth{Key: c.apiKey},
	)
	return resp, err
}

func (c *Client) GetMiniProfileBackground(ctx context.Context, req *CPlayer_GetMiniProfileBackground_Request) (*CPlayer_GetMiniProfileBackground_Response, error) {
	resp := &CPlayer_GetMiniProfileBackground_Response{}
	err := GetWithResp(
		ctx,
		Endpoint{Interface: "IPlayerService", Method: "GetMiniProfileBackground"},
		req,
		resp,
		&Auth{Key: c.apiKey},
	)
	return resp, err
}
