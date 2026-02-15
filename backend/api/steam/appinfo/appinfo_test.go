package appinfo_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Alia5/steaminputdb.com/api/steam/appinfo"
	"github.com/Alia5/steaminputdb.com/steamapi"
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func setupMockSteamAPI(t *testing.T, baseURL string) {
	t.Helper()
	orig := http.DefaultClient
	u, err := url.Parse(baseURL)
	require.NoError(t, err)

	http.DefaultClient = &http.Client{
		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			if req.URL.Host == "api.steampowered.com" {
				req.URL.Scheme = u.Scheme
				req.URL.Host = u.Host
			}
			return http.DefaultTransport.RoundTrip(req)
		}),
	}

	steamapi.DefaultClient = steamapi.NewClientWithBaseURL("test-key", baseURL)
	t.Cleanup(func() {
		http.DefaultClient = orig
	})
}

func TestSteamAppInfo(t *testing.T) {
	type testCase struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string
		contains       string
		setupMock      func(t *testing.T) *httptest.Server
	}

	testCases := []testCase{
		{
			name: "SUCCESS",
			setupMock: func(t *testing.T) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					resp := &steamapi.CStoreBrowse_GetItems_Response{
						StoreItems: []*steamapi.StoreItem{{
							Appid:        new(uint32((250900))),
							Name:         new("The Binding of Isaac: Rebirth"),
							StoreUrlPath: new("/app/250900/"),
						}},
					}
					w.Header().Set("Content-Type", "application/octet-stream")
					b, err := proto.Marshal(resp)
					require.NoError(t, err)
					w.Write(b)
				}))
			},
			path:           "/v1/steam/appinfo?app_id=250900",
			expectedStatus: http.StatusOK,
			contains:       "The Binding of Isaac: Rebirth",
		},
		{
			name:           "MISSING_APP_ID",
			setupMock:      nil,
			path:           "/v1/steam/appinfo",
			expectedStatus: http.StatusUnprocessableEntity, contains: "app_id",
		},
		{
			name: "APP_NOT_FOUND",
			setupMock: func(t *testing.T) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path != "/IStoreBrowseService/GetItems/v1/" {
						http.Error(w, "", http.StatusNotFound)
					}
				}))
			},
			path:           "/v1/steam/appinfo?app_id=999999",
			expectedStatus: http.StatusNotFound,
			expectedBody: `{
				"title": "Not Found",
				"status": 404,
				"detail": "item not found"
			}`,
		},
		{
			name: "STEAM_API_ERROR",
			setupMock: func(t *testing.T) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					http.Error(w, "", http.StatusServiceUnavailable)
				}))
			},
			path:           "/v1/steam/appinfo?app_id=250900",
			expectedStatus: http.StatusBadGateway,
			contains:       "failed to get steam app info",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMock != nil {
				srv := tc.setupMock(t)
				setupMockSteamAPI(t, srv.URL)
				t.Cleanup(srv.Close)
			}

			_, api := humatest.New(t)
			appinfo.RegisterRoute(api)
			resp := api.Get(tc.path)

			assert.Equal(t, tc.expectedStatus, resp.Code, "body: %s", resp.Body.String())
			if tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, resp.Body.String())
			}
			if tc.contains != "" {
				assert.Contains(t, resp.Body.String(), tc.contains)
			}
		})
	}
}
