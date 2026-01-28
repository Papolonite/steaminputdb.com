package routes_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alia5/steaminputdb.com/routes"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/assert"
)

func TestUnregisteredMiddleware(t *testing.T) {

	type testCase struct {
		name             string
		setup            func(api *humatest.TestAPI) error
		path             string
		expectedResponse string
		expectedStatus   int
	}

	testCases := []testCase{
		{
			name:             "404 Not Found",
			path:             "/nonexistent",
			expectedResponse: `{"detail":"Resource not found", "status":404, "title":"Not Found"}`,
			expectedStatus:   http.StatusNotFound,
		},
		{
			name: "405 Method Not Allowed",
			setup: func(api *humatest.TestAPI) error {
				huma.Register(
					*api,
					huma.Operation{
						Path:   "/meh",
						Method: http.MethodPost,
					},
					func(ctx context.Context, _ *struct{}) (*struct{}, error) {
						return nil, nil
					},
				)
				return nil
			},
			path:             "/meh",
			expectedResponse: `{"detail":"Method not allowed", "status":405, "title":"Method Not Allowed"}`,
			expectedStatus:   http.StatusMethodNotAllowed,
		},
		{
			name: "GetsToHuma",
			setup: func(api *humatest.TestAPI) error {
				huma.Register(
					*api,
					huma.Operation{
						Path:   "/meh",
						Method: http.MethodGet,
					},
					func(ctx context.Context, _ *struct{}) (*struct {
						Body struct{}
					}, error) {
						return &struct {
							Body struct{}
						}{}, nil
					},
				)
				return nil
			},
			path:             "/meh",
			expectedResponse: `{"$schema":"http://localhost/schemas/Response.json"}`,
			expectedStatus:   http.StatusOK,
		},
	}

	srvMux := http.NewServeMux()
	a := humago.New(srvMux, huma.DefaultConfig("Test", "1"))
	api := humatest.Wrap(t, a)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				err := tc.setup(&api)
				if err != nil {
					t.Fatalf("Failed to setup test case: %v", err)
				}
			}

			req, err := http.NewRequest(
				"GET",
				fmt.Sprintf("http://localhost%s", tc.path),
				nil,
			)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			resp := httptest.NewRecorder()
			routes.UnregisteredMiddleware(
				srvMux,
			).ServeHTTP(resp, req)

			assert.Equal(t, tc.expectedStatus, resp.Code)
			assert.JSONEq(t, tc.expectedResponse, resp.Body.String())
		})
	}

}
