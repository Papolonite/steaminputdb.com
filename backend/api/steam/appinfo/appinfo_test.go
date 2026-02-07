package appinfo_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/Alia5/steaminputdb.com/api/steam/appinfo"
	"github.com/Alia5/steaminputdb.com/steamapi"
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/assert"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func redirectSteamStoreTo(t *testing.T, targetBaseURL string) {
	t.Helper()
	orig := http.DefaultClient

	u, err := url.Parse(targetBaseURL)
	if err != nil {
		t.Fatalf("failed to parse target base URL: %v", err)
	}

	http.DefaultClient = &http.Client{
		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			if req.URL.Host == "store.steampowered.com" {
				req.URL.Scheme = u.Scheme
				req.URL.Host = u.Host
			}
			return http.DefaultTransport.RoundTrip(req)
		}),
	}

	t.Cleanup(func() {
		http.DefaultClient = orig
	})
}

func TestSteamAppInfo(t *testing.T) {
	type testCase struct {
		name             string
		setupMock        func() *httptest.Server
		path             string
		expectedStatus   int
		expectedResponse string
		expectedContains string
		assertBody       func(t *testing.T, body []byte)
	}

	testCases := []testCase{
		{
			name: "SUCCESS",
			path: "/v1/steam/appinfo?app_id=440",
			setupMock: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path != "/api/appdetails" {
						http.Error(w, "wrong path", http.StatusNotFound)
						return
					}
					if r.URL.Query().Get("appids") != "440" {
						http.Error(w, "wrong appids", http.StatusBadRequest)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"440":{"success":true,"data":{"type":"game","name":"Team Fortress 2","steam_appid":440}}}`))
				}))
			},
			expectedStatus: http.StatusOK,
			assertBody: func(t *testing.T, body []byte) {
				t.Helper()
				var got steamapi.AppInfo
				assert.NoError(t, json.Unmarshal(body, &got))
				assert.Equal(t, "Team Fortress 2", got.Name)
				assert.Equal(t, 440, got.SteamAppid)
			},
		},
		{
			name:             "MISSING_APP_ID",
			path:             "/v1/steam/appinfo",
			expectedStatus:   http.StatusUnprocessableEntity,
			expectedContains: "app_id",
		},
		{
			name: "APP_NOT_FOUND",
			path: "/v1/steam/appinfo?app_id=440",
			setupMock: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"440":{"success":false,"data":null}}`))
				}))
			},
			expectedStatus:   http.StatusNotFound,
			expectedResponse: `{"title":"Not Found","status":404,"detail":"app not found"}`,
		},
		{
			name: "STEAM_STORE_DOWN",
			path: "/v1/steam/appinfo?app_id=440",
			setupMock: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "nope", http.StatusServiceUnavailable)
				}))
			},
			expectedStatus:   http.StatusBadGateway,
			expectedContains: "failed to get steam app info",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, api := humatest.New(t)
			appinfo.RegisterRoute(api)

			if tc.setupMock != nil {
				store := tc.setupMock()
				defer store.Close()
				redirectSteamStoreTo(t, store.URL)
			}

			resp := api.Get(tc.path)

			assert.Equal(t, tc.expectedStatus, resp.Code)
			if tc.expectedResponse != "" {
				assert.JSONEq(t, tc.expectedResponse, resp.Body.String())
			}
			if tc.expectedContains != "" {
				assert.True(t, strings.Contains(resp.Body.String(), tc.expectedContains), "response body should contain %q, got: %s", tc.expectedContains, resp.Body.String())
			}
			if tc.assertBody != nil {
				tc.assertBody(t, resp.Body.Bytes())
			}

		})
	}

}
