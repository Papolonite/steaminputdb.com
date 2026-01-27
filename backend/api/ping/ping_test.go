package ping_test

import (
	"testing"

	"github.com/Alia5/steaminputdb.com/api/ping"
	"github.com/Alia5/steaminputdb.com/version"
	"github.com/go-fuego/fuego"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {

	type testCase struct {
		name             string
		setupContext     func() (*fuego.MockContext[any, any], error)
		expectedResponse *ping.Ping
		expectedError    error
	}

	testCases := []testCase{
		{
			name: "SUCCESS",
			setupContext: func() (*fuego.MockContext[any, any], error) {
				return fuego.NewMockContextNoBody(), nil
			},
			expectedResponse: &ping.Ping{
				Service: "SteamInputDB.com",
				Version: version.Version,
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, err := tc.setupContext()
			if err != nil {
				t.Fatalf("Failed to setup context: %v", err)
			}
			resp, err := ping.Controller(ctx)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedResponse, resp)
		})
	}

}
