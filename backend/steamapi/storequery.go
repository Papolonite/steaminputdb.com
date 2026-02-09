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

func (c *Client) GetItems(ctx context.Context, req *CStoreBrowse_GetItems_Request) (*CStoreBrowse_GetItems_Response, error) {
	resp := &CStoreBrowse_GetItems_Response{}
	err := GetWithResp(
		ctx,
		Endpoint{Interface: "IStoreBrowseService", Method: "GetItems"},
		req,
		resp,
		&Auth{Key: c.apiKey},
	)
	return resp, err
}
