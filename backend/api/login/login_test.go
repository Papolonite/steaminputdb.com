package login_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alia5/steaminputdb.com/api/login"
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/assert"
)

func TestSteamLogin(t *testing.T) {

	type testCase struct {
		name             string
		setupMock        func() *httptest.Server
		requestBody      map[string]any
		expectedStatus   int
		expectedResponse string
	}

	testCases := []testCase{
		{
			name: "SUCCESS",
			setupMock: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("ns:http://specs.openid.net/auth/2.0\nis_valid:true\n"))
				}))
			},
			requestBody: map[string]any{
				"ns":             "http://specs.openid.net/auth/2.0",
				"mode":           "id_res",
				"op_endpoint":    "https://steamcommunity.com/openid/login",
				"claimed_id":     "https://steamcommunity.com/openid/id/13371337696942069",
				"identity":       "https://steamcommunity.com/openid/id/13371337696942069",
				"return_to":      "http://localhost:5173/login/callback",
				"response_nonce": "2026-01-31T23:46:40ZfbncPsCvbGlcAnm7O7AOGbvecyg=",
				"assoc_handle":   "1234567890",
				"signed":         "signed,op_endpoint,claimed_id,identity,return_to,response_nonce,assoc_handle",
				"sig":            "5z/7Gwh1MhMbJ8JeU2pPNSit9l8=",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "INVALID_MODE",
			requestBody: map[string]any{
				"ns":             "http://specs.openid.net/auth/2.0",
				"mode":           "invalid_mode",
				"op_endpoint":    "https://steamcommunity.com/openid/login",
				"claimed_id":     "https://steamcommunity.com/openid/id/13371337696942069",
				"identity":       "https://steamcommunity.com/openid/id/13371337696942069",
				"return_to":      "http://localhost:5173/login/callback",
				"response_nonce": "2026-01-31T23:46:40ZfbncPsCvbGlcAnm7O7AOGbvecyg=",
				"assoc_handle":   "1234567890",
				"signed":         "signed,op_endpoint,claimed_id,identity,return_to,response_nonce,assoc_handle",
				"sig":            "5z/7Gwh1MhMbJ8JeU2pPNSit9l8=",
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedResponse: `{
				"title": "Unprocessable Entity",
				"status": 422,
				"detail": "validation failed",
				"errors": [
					{
						"message": "expected value to be one of \"id_res\"",
						"location": "body.mode",
						"value": "invalid_mode"
					}
				]
			}`,
		},
		{
			name: "INVALID_SIGNATURE",
			setupMock: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("ns:http://specs.openid.net/auth/2.0\nis_valid:false\n"))
				}))
			},
			requestBody: map[string]any{
				"ns":             "http://specs.openid.net/auth/2.0",
				"mode":           "id_res",
				"op_endpoint":    "https://steamcommunity.com/openid/login",
				"claimed_id":     "https://steamcommunity.com/openid/id/13371337696942069",
				"identity":       "https://steamcommunity.com/openid/id/13371337696942069",
				"return_to":      "http://localhost:5173/login/callback",
				"response_nonce": "2026-01-31T23:46:40ZfbncPsCvbGlcAnm7O7AOGbvecyg=",
				"assoc_handle":   "1234567890",
				"signed":         "signed,op_endpoint,claimed_id,identity,return_to,response_nonce,assoc_handle",
				"sig":            "invalid_signature",
			},
			expectedStatus:   http.StatusUnauthorized,
			expectedResponse: `{"detail":"steam verification failed", "status":401, "title":"Unauthorized"}`,
		},
		{
			name: "EMPTY_CLAIMED_ID",
			setupMock: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("ns:http://specs.openid.net/auth/2.0\nis_valid:true\n"))
				}))
			},
			requestBody: map[string]any{
				"ns":             "http://specs.openid.net/auth/2.0",
				"mode":           "id_res",
				"op_endpoint":    "https://steamcommunity.com/openid/login",
				"claimed_id":     "",
				"identity":       "https://steamcommunity.com/openid/id/13371337696942069",
				"return_to":      "http://localhost:5173/login/callback",
				"response_nonce": "2026-01-31T23:46:40ZfbncPsCvbGlcAnm7O7AOGbvecyg=",
				"assoc_handle":   "1234567890",
				"signed":         "signed,op_endpoint,claimed_id,identity,return_to,response_nonce,assoc_handle",
				"sig":            "5z/7Gwh1MhMbJ8JeU2pPNSit9l8=",
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedResponse: `{
				"title": "Unprocessable Entity",
				"status": 422,
				"detail": "validation failed",
				"errors": [
					{
						"message": "expected length >= 1",
						"location": "body.claimed_id",
						"value": ""
					}
				]
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, api := humatest.New(t)

			var mockUrl string
			if tc.setupMock != nil {
				mockServer := tc.setupMock()
				defer mockServer.Close()
				mockUrl = mockServer.URL
			}
			login.RegisterWithURL(api, mockUrl)

			resp := api.Post("/v1/steam/login", tc.requestBody)
			assert.Equal(t, tc.expectedStatus, resp.Code)

			if tc.expectedResponse != "" {
				assert.JSONEq(t, tc.expectedResponse, resp.Body.String())
			}

		})
	}

}
