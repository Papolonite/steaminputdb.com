package steamapi

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"reflect"

	"google.golang.org/protobuf/proto"
)

var DefaultClient = &Client{}

// Client is a Steam Web API client
type Client struct {
	apiKey  string
	baseURL string
}

// NewClient creates a new Steam Web API client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: "https://api.steampowered.com",
	}
}

// NewClientWithBaseURL creates a new Steam Web API client with a custom base URL (for testing)
func NewClientWithBaseURL(apiKey string, baseURL string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// APIKey returns the API key used by the client
func (c *Client) APIKey() string {
	return c.apiKey
}

// Auth represents authentication credentials for Steam Web API requests
type Auth struct {
	Key         string
	AccessToken string
}

// AddToParams adds the authentication parameters to the URL query parameters
func (a *Auth) AddToParams(params *url.Values) {
	if a.Key != "" {
		params.Add("key", a.Key)
	}
	if a.AccessToken != "" {
		params.Add("access_token", a.AccessToken)
	}
}

// Endpoint represents a Steam Web API endpoint
type Endpoint struct {
	Interface string
	Method    string
	Version   string // defaults to "1" if empty
}

// URL constructs the full URL for the endpoint
func (e *Endpoint) URL() string {
	version := e.Version
	if version == "" {
		version = "1"
	}
	return fmt.Sprintf("https://api.steampowered.com/%s/%s/v%s/", e.Interface, e.Method, version)
}

// URLWithBase constructs the full URL for the endpoint with a custom base URL
func (e *Endpoint) URLWithBase(baseURL string) string {
	version := e.Version
	if version == "" {
		version = "1"
	}
	return fmt.Sprintf("%s/%s/%s/v%s/", baseURL, e.Interface, e.Method, version)
}

var ErrRequest = errors.New("request error")

// Get makes a type-safe request to the Steam Web API and returns the unmarshaled response.
func Get[Req proto.Message, Resp proto.Message](ctx context.Context, endpoint Endpoint, req Req, auth *Auth) (Resp, error) {
	var zero Resp
	resp := reflect.New(reflect.TypeFor[Resp]().Elem()).Interface().(Resp)
	if err := GetWithResp(ctx, endpoint, req, resp, auth); err != nil {
		return zero, err
	}
	return resp, nil
}

// GetWithResp makes a request to the Steam Web API and fills the provided response message.
func GetWithResp[Req proto.Message, Resp proto.Message](ctx context.Context, endpoint Endpoint, req Req, resp Resp, auth *Auth) error {
	requestProto, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	baseURL := endpoint.URL()
	params := url.Values{
		"input_protobuf_encoded": {base64.StdEncoding.EncodeToString(requestProto)},
	}

	if auth != nil {
		auth.AddToParams(&params)
	}

	fullURL := baseURL + "?" + params.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("%w: failed to make request: %w", ErrRequest, err)
	}
	defer (func() {
		err := httpResp.Body.Close()
		if err != nil {
			slog.Error("failed to close response body", "error", err)
		}
	})()

	if httpResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(httpResp.Body)
		return fmt.Errorf("%w: HTTP error %d: %s", ErrRequest, httpResp.StatusCode, string(body))
	}

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("%w: failed to read response: %v", ErrRequest, err)
	}

	if err := proto.Unmarshal(body, resp); err != nil {
		return fmt.Errorf("failed to unmarshal protobuf response: %w", err)
	}

	return nil
}

func (c *Client) GetJson(ctx context.Context, endpoint Endpoint, req any, params *url.Values, resp any, auth *Auth) error {
	baseURL := endpoint.URLWithBase(c.baseURL)
	if len(*params) == 0 {
		reqJSON, err := json.Marshal(req)
		if err != nil {
			return fmt.Errorf("failed to marshal request to JSON: %w", err)
		}
		params.Add("input_json", string(reqJSON))
	}

	if auth != nil {
		auth.AddToParams(params)
	}

	fullURL := baseURL + "?" + params.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrRequest, err)
	}
	defer (func() {
		err := httpResp.Body.Close()
		if err != nil {
			slog.Error("failed to close response body", "error", err)
		}
	})()

	if httpResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(httpResp.Body)
		return fmt.Errorf("%w: HTTP error %d: %s", ErrRequest, httpResp.StatusCode, string(body))
	}

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("%w: failed to read response: %v", ErrRequest, err)
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	return nil
}

func GetJSON[Req any, Resp any](ctx context.Context, endpoint Endpoint, req Req, params *url.Values, resp Resp, auth *Auth) error {

	baseURL := endpoint.URL()
	if len(*params) == 0 {
		reqJSON, err := json.Marshal(req)
		if err != nil {
			return fmt.Errorf("failed to marshal request to JSON: %w", err)
		}
		params.Add("input_json", string(reqJSON))
	}

	if auth != nil {
		auth.AddToParams(params)
	}

	fullURL := baseURL + "?" + params.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrRequest, err)
	}
	defer (func() {
		err := httpResp.Body.Close()
		if err != nil {
			slog.Error("failed to close response body", "error", err)
		}
	})()

	if httpResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(httpResp.Body)
		return fmt.Errorf("%w: HTTP error %d: %s", ErrRequest, httpResp.StatusCode, string(body))
	}

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("%w: failed to read response: %v", ErrRequest, err)
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	return nil
}
