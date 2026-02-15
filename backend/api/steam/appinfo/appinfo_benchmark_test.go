package appinfo_test

import (
	"testing"

	"github.com/Alia5/steaminputdb.com/api/steam/appinfo"
	"github.com/danielgtaylor/huma/v2/humatest"
)

func BenchmarkAppInfo(b *testing.B) {

	type testCase struct {
		name        string
		useMemCache bool
	}

	testCases := []testCase{
		{name: "With MemCache", useMemCache: true},
		{name: "Without MemCache", useMemCache: false},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {

			bWithoutLogging := BSupressedLogs{b}
			_, api := humatest.New(bWithoutLogging)
			appinfo.RegisterRoute(api, tc.useMemCache)
			for b.Loop() {
				resp := api.Get("/v1/steam/appinfo?app_id=250900")
				b.StopTimer()
				if resp.Body == nil {
					b.Fatal("response body is nil")
				}
				if resp.Code != 200 {
					b.Fatalf("expected status 200, got %d", resp.Code)
				}
				b.StartTimer()
			}

		})
	}

}

type BSupressedLogs struct {
	*testing.B
}

func (b BSupressedLogs) Log(args ...any) {}
