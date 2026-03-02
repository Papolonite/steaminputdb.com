package steamapi_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Alia5/steaminputdb.com/steamapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Test_GET(t *testing.T) {

	type testCase struct {
		name         string
		endpoint     steamapi.Endpoint
		req          proto.Message
		expectedResp proto.Message
	}

	tests := []testCase{
		{
			name: "SUCCESS_SearchApps",
			endpoint: steamapi.Endpoint{
				Interface: "IStoreQueryService",
				Method:    "SearchSuggestions",
			},
			req: &steamapi.CStoreQuery_SearchSuggestions_Request{
				Context: &steamapi.StoreBrowseContext{
					Language:    new("english"),
					CountryCode: new("US"),
				},
				SearchTerm: new("isaac"),
				MaxResults: new(uint32(5)),
				Filters: &steamapi.CStoreQueryFilters{
					TypeFilters: &steamapi.CStoreQueryFilters_TypeFilters{
						IncludeGames: new(true),
					},
				},
			},
			expectedResp: &steamapi.CStoreQuery_SearchSuggestions_Response{
				Metadata: &steamapi.CStoreQueryResultMetadata{
					TotalMatchingRecords: new(int32(4)),
					Start:                new(int32(0)),
					Count:                new(int32(4)),
				},
				Ids: []*steamapi.StoreItemID{
					{
						Appid: new(uint32(250900)),
					},
					{
						Appid: new(uint32(113200)),
					},
					{
						Appid: new(uint32(1273600)),
					},
					{
						Appid: new(uint32(341260)),
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Helper()

			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/"+tc.endpoint.Interface+"/"+tc.endpoint.Method+"/v1/" {
					http.Error(w, "wrong path", http.StatusNotFound)
					return
				}
				body, err := proto.Marshal(tc.expectedResp)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/octet-stream")
				_, _ = w.Write(body)
			}))
			defer srv.Close()

			orig := steamapi.DefaultClient
			steamapi.DefaultClient = steamapi.NewClientWithBaseURL("", srv.URL)
			t.Cleanup(func() {
				steamapi.DefaultClient = orig
			})

			resp := reflect.New(reflect.TypeOf(tc.expectedResp).Elem()).Interface().(proto.Message)
			err := steamapi.GetWithResp(
				t.Context(),
				tc.endpoint,
				tc.req,
				resp,
				nil,
			)
			if err != nil {
				t.Fatalf("Get() error = %v", err)
			}

			if !proto.Equal(resp, tc.expectedResp) {
				marshaler := protojson.MarshalOptions{
					Multiline: true,
					Indent:    "  ",
				}
				gotJSON, _ := marshaler.Marshal(resp)
				wantJSON, _ := marshaler.Marshal(tc.expectedResp)
				t.Errorf("Get() mismatch:\nGot:\n%s\n\nWant:\n%s", gotJSON, wantJSON)
			}
		})
	}

}
