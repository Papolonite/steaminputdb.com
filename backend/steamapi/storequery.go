package steamapi

import "context"

// SearchSuggestions searches for Steam apps via IStoreQueryService/SearchSuggestions/v1
func (c *Client) SearchSuggestions(ctx context.Context, req *CStoreQuery_SearchSuggestions_Request) (*CStoreQuery_SearchSuggestions_Response, error) {
	resp := &CStoreQuery_SearchSuggestions_Response{}
	err := GetWithResp(
		ctx,
		Endpoint{Interface: "IStoreQueryService", Method: "SearchSuggestions"},
		req,
		resp,
		&Auth{Key: c.apiKey},
	)
	return resp, err
}
