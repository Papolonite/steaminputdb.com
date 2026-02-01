// Package login provides Steam OpenID authentication and JWT generation.
package login

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/golang-jwt/jwt/v5"
)

const steamLoginURL = "https://steamcommunity.com/openid/login"

type OpenIDRequest struct {
	Body struct {
		NS            string `json:"ns"`
		Mode          string `json:"mode" enum:"id_res"`
		OpEndpoint    string `json:"op_endpoint"`
		ClaimedID     string `json:"claimed_id" minLength:"1"`
		Identity      string `json:"identity"`
		ReturnTo      string `json:"return_to"`
		ResponseNonce string `json:"response_nonce"`
		AssocHandle   string `json:"assoc_handle"`
		Signed        string `json:"signed"`
		Sig           string `json:"sig"`
	}
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Response struct {
	Body LoginResponse
}

const jwtSecret = "TODO:FIXME!"
const jwtValidity = time.Hour * 24

func RegisterWithURL(a huma.API, steamURL string) {
	registerRoutes(a, steamURL)
}

func RegisterRoutes(a huma.API) {
	registerRoutes(a, steamLoginURL)
}
func registerRoutes(a huma.API, loginURL string) {
	huma.Register(
		a,
		huma.Operation{
			Method:      http.MethodPost,
			Path:        "/v1/steam/login",
			Tags:        []string{"auth"},
			Summary:     "Log in with Steam",
			Description: "Authenticate user via Steam OpenID and return JWT token",
			Errors: []int{
				401,
			},
		},
		func(ctx context.Context, req *OpenIDRequest) (*Response, error) {
			params := url.Values{}
			params.Set("openid.mode", "check_authentication")
			params.Set("openid.ns", req.Body.NS)
			params.Set("openid.op_endpoint", req.Body.OpEndpoint)
			params.Set("openid.claimed_id", req.Body.ClaimedID)
			params.Set("openid.identity", req.Body.Identity)
			params.Set("openid.return_to", req.Body.ReturnTo)
			params.Set("openid.response_nonce", req.Body.ResponseNonce)
			params.Set("openid.assoc_handle", req.Body.AssocHandle)
			params.Set("openid.signed", req.Body.Signed)
			params.Set("openid.sig", req.Body.Sig)

			parts := strings.Split(req.Body.ClaimedID, "/")
			if len(parts) == 0 {
				return nil, huma.Error401Unauthorized("invalid claimed_id")
			}
			steamID := parts[len(parts)-1]

			steamReq, err := http.NewRequestWithContext(ctx, http.MethodPost, loginURL, strings.NewReader(params.Encode()))
			if err != nil {
				return nil, huma.Error500InternalServerError("failed to create request")
			}
			steamReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			resp, err := http.DefaultClient.Do(steamReq)
			if err != nil {
				return nil, huma.Error401Unauthorized("failed to verify with steam")
			}
			defer func() {
				err = resp.Body.Close()
				if err != nil {
					slog.Error("failed to close steam response body", "error", err)
				}
			}()

			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, huma.Error500InternalServerError("failed to read response")
			}

			if !strings.Contains(string(respBody), "is_valid:true") {
				return nil, huma.Error401Unauthorized("steam verification failed")
			}

			claims := jwt.MapClaims{
				"sub": steamID,
				"iss": "steaminputdb",
				"iat": time.Now().Unix(),
				"exp": time.Now().Add(jwtValidity).Unix(),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte(jwtSecret))
			if err != nil {
				return nil, huma.Error500InternalServerError("failed to generate token")
			}

			return &Response{
				Body: LoginResponse{
					Token: tokenString,
				},
			}, nil
		},
	)
}
