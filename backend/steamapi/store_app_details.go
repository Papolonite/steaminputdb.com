package steamapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

var ErrInfoNotFound = fmt.Errorf("app not found")

func GetStoreInfo(ctx context.Context, appID uint32) (*AppInfo, error) {
	url := "https://store.steampowered.com/api/appdetails?appids=" + strconv.FormatUint(uint64(appID), 10)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpClient := sharedHTTPClient
	if DefaultClient != nil && DefaultClient.httpClient != nil {
		httpClient = DefaultClient.httpClient
	}

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to make request: %w", ErrRequest, err)
	}
	defer (func() {
		err := httpResp.Body.Close()
		if err != nil {
			slog.Error("failed to close response body", "error", err)
		}
	})()

	if httpResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(httpResp.Body)
		return nil, fmt.Errorf("%w: HTTP error %d: %s", ErrRequest, httpResp.StatusCode, string(body))
	}

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to read response: %v", ErrRequest, err)
	}

	slog.Debug("received response from Steam Store API", "app_id", appID, "response", string(body))

	resp := StoreAppDetails{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	for _, appDetails := range resp {
		if appDetails.Success {
			return appDetails.Data, nil
		}
	}
	return nil, ErrInfoNotFound
}

type StoreAppDetails map[string]AppDetails

type Subs struct {
	PackageID                int    `json:"packageid"` //nolint:misspell // Steam Store API uses "packageid"
	PercentSavingsText       string `json:"percent_savings_text"`
	PercentSavings           int    `json:"percent_savings"`
	OptionText               string `json:"option_text"`
	OptionDescription        string `json:"option_description"`
	CanGetFreeLicense        string `json:"can_get_free_license"`
	IsFreeLicense            bool   `json:"is_free_license"`
	PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
}

type PackageGroups struct {
	Name                    string `json:"name"`
	Title                   string `json:"title"`
	Description             string `json:"description"`
	SelectionText           string `json:"selection_text"`
	SaveText                string `json:"save_text"`
	DisplayType             any    `json:"display_type"`
	IsRecurringSubscription string `json:"is_recurring_subscription"`
	Subs                    []Subs `json:"subs"`
}

type Platforms struct {
	Windows bool `json:"windows"`
	Mac     bool `json:"mac"`
	Linux   bool `json:"linux"`
}

type Metacritic struct {
	Score int    `json:"score"`
	URL   string `json:"url"`
}

type Categories struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

type Genres struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Screenshots struct {
	ID            int    `json:"id"`
	PathThumbnail string `json:"path_thumbnail"`
	PathFull      string `json:"path_full"`
}

type Movies struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	DashAv1   string `json:"dash_av1"`
	DashH264  string `json:"dash_h264"`
	HlsH264   string `json:"hls_h264"`
	Highlight bool   `json:"highlight"`
}

type Recommendations struct {
	Total int `json:"total"`
}

type Highlighted struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Achievements struct {
	Total       int           `json:"total"`
	Highlighted []Highlighted `json:"highlighted"`
}

type ReleaseDate struct {
	ComingSoon bool   `json:"coming_soon"`
	Date       string `json:"date"`
}

type SupportInfo struct {
	URL   string `json:"url"`
	Email string `json:"email"`
}

type ContentDescriptors struct {
	Ids   []any  `json:"ids"`
	Notes string `json:"notes"`
}

type Usk struct {
	Rating string `json:"rating"`
}

type Dejus struct {
	RatingGenerated string `json:"rating_generated"`
	Rating          string `json:"rating"`
	RequiredAge     any    `json:"required_age"`
	Banned          string `json:"banned"`
	UseAgeGate      string `json:"use_age_gate"`
	Descriptors     string `json:"descriptors"`
}

type SteamGermany struct {
	RatingGenerated string `json:"rating_generated"`
	Rating          string `json:"rating"`
	RequiredAge     any    `json:"required_age"`
	Banned          string `json:"banned"`
	UseAgeGate      string `json:"use_age_gate"`
	Descriptors     string `json:"descriptors"`
}

type Ratings struct {
	Usk          Usk          `json:"usk"`
	Dejus        Dejus        `json:"dejus"`
	SteamGermany SteamGermany `json:"steam_germany"`
}

type AppInfo struct {
	Type                string             `json:"type"`
	Name                string             `json:"name"`
	SteamAppid          int                `json:"steam_appid"`
	RequiredAge         any                `json:"required_age"`
	IsFree              bool               `json:"is_free"`
	ControllerSupport   string             `json:"controller_support"`
	Dlc                 []int              `json:"dlc"`
	DetailedDescription string             `json:"detailed_description"`
	AboutTheGame        string             `json:"about_the_game"`
	ShortDescription    string             `json:"short_description"`
	SupportedLanguages  string             `json:"supported_languages"`
	HeaderImage         string             `json:"header_image"`
	CapsuleImage        string             `json:"capsule_image"`
	CapsuleImagev5      string             `json:"capsule_imagev5"`
	Website             string             `json:"website"`
	PcRequirements      *any               `json:"pc_requirements"`
	MacRequirements     *any               `json:"mac_requirements"`
	LinuxRequirements   *any               `json:"linux_requirements"`
	Developers          []string           `json:"developers"`
	Publishers          []string           `json:"publishers"`
	Packages            []int              `json:"packages"`
	PackageGroups       []PackageGroups    `json:"package_groups"`
	Platforms           Platforms          `json:"platforms"`
	Metacritic          Metacritic         `json:"metacritic"`
	Categories          []Categories       `json:"categories"`
	Genres              []Genres           `json:"genres"`
	Screenshots         []Screenshots      `json:"screenshots"`
	Movies              []Movies           `json:"movies"`
	Recommendations     Recommendations    `json:"recommendations"`
	Achievements        Achievements       `json:"achievements"`
	ReleaseDate         ReleaseDate        `json:"release_date"`
	SupportInfo         SupportInfo        `json:"support_info"`
	Background          string             `json:"background"`
	BackgroundRaw       string             `json:"background_raw"`
	ContentDescriptors  ContentDescriptors `json:"content_descriptors"`
	Ratings             Ratings            `json:"ratings"`
}

type AppDetails struct {
	Success bool     `json:"success"`
	Data    *AppInfo `json:"data"`
}
