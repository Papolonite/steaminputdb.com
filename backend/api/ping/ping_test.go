package ping_test

import (
	"net/http"
	"testing"

	"github.com/Alia5/steaminputdb.com/api/ping"
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {

	type testCase struct {
		name             string
		expectedResponse string
		expectedStatus   int
	}

	testCases := []testCase{
		{
			name:             "SUCCESS",
			expectedResponse: `{"service":"SteamInputDB.com", "version":"v1.0.0"}`,
			expectedStatus:   http.StatusOK,
		},
	}
	_, api := humatest.New(t)
	ping.RegisterRoutes(api)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := api.Get("/v1/ping")
			assert.Equal(t, tc.expectedStatus, resp.Code)
			assert.JSONEq(t, tc.expectedResponse, resp.Body.String())
		})
	}

}
