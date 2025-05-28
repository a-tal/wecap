package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Payload struct {
	DateUTC          time.Time
	TempF            float64
	Humidity         float64
	WindSpeedMPH     float64
	WindGustMPH      float64
	MaxDailyGust     float64
	WindDir          float64
	WindDirAvg10m    float64
	UV               float64
	SolarRadiation   float64
	HourlyRainIn     float64
	EventRainIn      float64
	DailyRainIn      float64
	WeeklyRainIn     float64
	MonthlyRainIn    float64
	YearlyRainIn     float64
	BattOut          float64
	BattRain         float64
	TempInF          float64 // indoors
	HumidityIn       float64 // indoors
	BaromRelIn       float64
	BaromAbsIn       float64
	BattIn           float64
	Leak1            float64
	BatLeak1         float64
	LightningDay     float64
	LightningTime    time.Time // last strike time, in epoch seconds
	LightingDistance float64   // miles, of last strike
	BattLightning    float64
}

func newPayload(query map[string]string) Payload {
	return Payload{
		DateUTC:          dateValue(query, "dateutc"),
		TempF:            floatValue(query, "tempf"),
		Humidity:         floatValue(query, "humidity"),
		WindSpeedMPH:     floatValue(query, "windspeedmph"),
		WindGustMPH:      floatValue(query, "windgustmph"),
		MaxDailyGust:     floatValue(query, "maxdailygust"),
		WindDir:          floatValue(query, "winddir"),
		WindDirAvg10m:    floatValue(query, "winddir_avg10m"),
		UV:               floatValue(query, "uv"),
		SolarRadiation:   floatValue(query, "solarradiation"),
		HourlyRainIn:     floatValue(query, "hourlyrainin"),
		EventRainIn:      floatValue(query, "eventrainin"),
		DailyRainIn:      floatValue(query, "dailyrainin"),
		WeeklyRainIn:     floatValue(query, "weeklyrainin"),
		MonthlyRainIn:    floatValue(query, "monthlyrainin"),
		YearlyRainIn:     floatValue(query, "yearlyrainin"),
		BattOut:          floatValue(query, "battout"),
		BattRain:         floatValue(query, "battrain"),
		TempInF:          floatValue(query, "tempinf"),
		HumidityIn:       floatValue(query, "humidityin"),
		BaromRelIn:       floatValue(query, "baromrelin"),
		BaromAbsIn:       floatValue(query, "baromabsin"),
		BattIn:           floatValue(query, "battin"),
		Leak1:            floatValue(query, "leak1"),
		BatLeak1:         floatValue(query, "batleak1"),
		LightningDay:     floatValue(query, "lightning_day"),
		LightningTime:    tsValue(query, "lightning_time"),
		LightingDistance: floatValue(query, "lightning_distance"),
		BattLightning:    floatValue(query, "batt_lightning"),
	}
}

func floatValue(query map[string]string, key string) float64 {
	v, ok := query[key]
	if ok {
		i, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return i
		}
		ll.Printf("failed to parse float %s: %v", key, err)
	}
	return 0
}

func tsValue(query map[string]string, key string) time.Time {
	v, ok := query[key]
	if ok {
		i, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return time.Unix(i, 0)
		}
		ll.Printf("failed to parse timestamp %s: %v", key, err)
	}
	return time.Unix(0, 0)
}

func dateValue(query map[string]string, key string) time.Time {
	v, ok := query[key]
	if ok {
		ts, err := time.Parse(time.RFC3339, fmt.Sprintf("%sZ", strings.ReplaceAll(v, "+", "T")))
		if err == nil {
			return ts
		}
		ll.Printf("failed to parse datetime %s: %v", key, err)
	}
	return time.Now()
}
