package steamapi

import "context"

// QueryFiles queries Steam Workshop files via IPublishedFileService/QueryFiles/v1
func (c *Client) QueryFiles(ctx context.Context, req *CPublishedFile_QueryFiles_Request) (*CPublishedFile_QueryFiles_Response, error) {
	resp := &CPublishedFile_QueryFiles_Response{}
	err := GetWithResp(
		ctx,
		Endpoint{Interface: "IPublishedFileService", Method: "QueryFiles"},
		req,
		resp,
		&Auth{Key: c.apiKey},
	)
	return resp, err
}

// GetFileDetails gets details for specific Workshop files via IPublishedFileService/GetDetails/v1
func (c *Client) GetFileDetails(ctx context.Context, req *CPublishedFile_GetDetails_Request) (*CPublishedFile_GetDetails_Response, error) {
	resp := &CPublishedFile_GetDetails_Response{}
	err := GetWithResp(
		ctx,
		Endpoint{Interface: "IPublishedFileService", Method: "GetDetails"},
		req,
		resp,
		&Auth{Key: c.apiKey},
	)
	return resp, err
}
