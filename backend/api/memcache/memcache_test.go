package memcache_test

import (
	"testing"
	"time"

	"github.com/Alia5/steaminputdb.com/api/memcache"
	"github.com/stretchr/testify/assert"
)

func TestMemcache(t *testing.T) {

	type testCase struct {
		name           string
		validity       time.Duration
		maxItems       int
		storeKeys      []string
		storeValues    []string
		sleepAfter     *time.Duration
		getKey         string
		expectedExists bool
		expectedValue  string
	}

	testCases := []testCase{
		{
			name:           "STORE_AND_GET",
			validity:       1 * time.Hour,
			maxItems:       100,
			storeKeys:      []string{"test_key"},
			storeValues:    []string{"test_value"},
			getKey:         "test_key",
			expectedExists: true,
			expectedValue:  "test_value",
		},
		{
			name:           "GET_NONEXISTENT",
			validity:       1 * time.Hour,
			maxItems:       100,
			storeKeys:      []string{},
			storeValues:    []string{},
			getKey:         "nonexistent",
			expectedExists: false,
			expectedValue:  "",
		},
		{
			name:           "VALUE_WITHIN_TTL",
			validity:       1000 * time.Millisecond,
			maxItems:       100,
			storeKeys:      []string{"ttl_test"},
			storeValues:    []string{"data"},
			sleepAfter:     new(10 * time.Millisecond),
			getKey:         "ttl_test",
			expectedExists: true,
			expectedValue:  "data",
		},
		{
			name:           "VALUE_EXPIRED",
			validity:       20 * time.Millisecond,
			maxItems:       100,
			storeKeys:      []string{"ttl_test"},
			storeValues:    []string{"data"},
			sleepAfter:     new(100 * time.Millisecond),
			getKey:         "ttl_test",
			expectedExists: false,
			expectedValue:  "",
		},
		{
			name:           "CLEANUP_ON_MAX_ITEMS_KEY1",
			validity:       1 * time.Hour,
			maxItems:       3,
			storeKeys:      []string{"key1", "key2", "key3", "key4", "key5"},
			storeValues:    []string{"value1", "value2", "value3", "value4", "value5"},
			sleepAfter:     new(200 * time.Millisecond),
			getKey:         "key1",
			expectedExists: false,
			expectedValue:  "",
		},
		{
			name:           "CLEANUP_ON_MAX_ITEMS_KEY5",
			validity:       1 * time.Hour,
			maxItems:       3,
			storeKeys:      []string{"key1", "key2", "key3", "key4", "key5"},
			storeValues:    []string{"value1", "value2", "value3", "value4", "value5"},
			sleepAfter:     new(200 * time.Millisecond),
			getKey:         "key5",
			expectedExists: true,
			expectedValue:  "value5",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := memcache.New(tc.validity, tc.maxItems)

			for i, key := range tc.storeKeys {
				c.Store(key, tc.storeValues[i])
				if len(tc.storeKeys) > 3 {
					time.Sleep(10 * time.Millisecond)
				}
			}

			if tc.sleepAfter != nil {
				time.Sleep(*tc.sleepAfter)
			}

			got, ok := memcache.Get[string](c, tc.getKey)
			assert.Equal(t, tc.expectedExists, ok)
			if tc.expectedExists {
				assert.Equal(t, tc.expectedValue, got)
			}
		})
	}
}
