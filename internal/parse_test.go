package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseQuery(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			"empty case",
			"",
			map[string]string{},
		},
		{
			"happy path",
			"/data/report/&PASSKEY=FOOBAR&stationtype=SomeType&dateutc=2025-01-02+03:04:05&tempf=80.0&humidity=60&windspeedmph=2.00&windgustmph=4.00&maxdailygust=8.00&winddir=300&winddir_avg10m=200&uv=1&solarradiation=50.00&hourlyrainin=0.001&eventrainin=0.002&dailyrainin=0.003&weeklyrainin=0.004&monthlyrainin=0.005&yearlyrainin=0.006&battout=1&battrain=2&tempinf=75.0&humidityin=50&baromrelin=29.700&baromabsin=29.200&battin=3&leak1=4&batleak1=5&lightning_day=100&lightning_time=1234567890&lightning_distance=24&batt_lightning=6",
			map[string]string{
				"PASSKEY":            "FOOBAR",
				"stationtype":        "SomeType",
				"dateutc":            "2025-01-02+03:04:05",
				"tempf":              "80.0",
				"humidity":           "60",
				"windspeedmph":       "2.00",
				"windgustmph":        "4.00",
				"maxdailygust":       "8.00",
				"winddir":            "300",
				"winddir_avg10m":     "200",
				"uv":                 "1",
				"solarradiation":     "50.00",
				"hourlyrainin":       "0.001",
				"eventrainin":        "0.002",
				"dailyrainin":        "0.003",
				"weeklyrainin":       "0.004",
				"monthlyrainin":      "0.005",
				"yearlyrainin":       "0.006",
				"battout":            "1",
				"battrain":           "2",
				"tempinf":            "75.0",
				"humidityin":         "50",
				"baromrelin":         "29.700",
				"baromabsin":         "29.200",
				"battin":             "3",
				"leak1":              "4",
				"batleak1":           "5",
				"lightning_day":      "100",
				"lightning_time":     "1234567890",
				"lightning_distance": "24",
				"batt_lightning":     "6",
			},
		},
		{
			"duplicates would be overwritten; last one wins",
			"/data/report/&foo=biz&foo=baz&foo=bar",
			map[string]string{"foo": "bar"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, parseQuery(tt.input))
		})
	}
}
