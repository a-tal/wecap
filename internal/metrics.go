package internal

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metTempF            = newGauge("temp_f")
	metHumidity         = newGauge("humidity")
	metWindSpeedMPH     = newGauge("wind_speed_mph")
	metWindGustMPH      = newGauge("wind_gust_mph")
	metMaxDailyGust     = newGauge("max_daily_gust")
	metWindDir          = newGauge("wind_dir")
	metWindDirAvg10m    = newGauge("wind_dir_avg_10m")
	metUV               = newGauge("uv")
	metSolarRadiation   = newGauge("solar_radiation")
	metHourlyRainIn     = newGauge("hourly_rain_inches")
	metEventRainIn      = newGauge("event_rain_inches")
	metDailyRainIn      = newGauge("daily_rain_inches")
	metWeeklyRainIn     = newGauge("weekly_rain_inches")
	metMonthlyRainIn    = newGauge("monthly_rain_inches")
	metYearlyRainIn     = newGauge("yearly_rain_inches")
	metBattOut          = newGauge("batt_outside")
	metBattRain         = newGauge("batt_rain")
	metTempInF          = newGauge("temp_indoor_f")
	metHumidityIn       = newGauge("humidity_indoor")
	metBaromRelIn       = newGauge("barom_rel_indoor")
	metBaromAbsIn       = newGauge("barom_abs_indoor")
	metBattIn           = newGauge("batt_inside")
	metLeak1            = newGauge("leak_1")
	metBatLeak1         = newGauge("batt_leak_1")
	metLightningDay     = newGauge("lightning_day")
	metLightningTime    = newGauge("lightning_time")
	metLightingDistance = newGauge("lighting_distance")
	metBattLightning    = newGauge("batt_lightning")
)

func newGauge(name string) prometheus.Gauge {
	return promauto.NewGauge(prometheus.GaugeOpts{
		Name: name,
	})
}

func updateMetrics(payload Payload) {
	metTempF.Set(payload.TempF)
	metHumidity.Set(payload.Humidity)
	metWindSpeedMPH.Set(payload.WindSpeedMPH)
	metWindGustMPH.Set(payload.WindGustMPH)
	metMaxDailyGust.Set(payload.MaxDailyGust)
	metWindDir.Set(payload.WindDir)
	metWindDirAvg10m.Set(payload.WindDirAvg10m)
	metUV.Set(payload.UV)
	metSolarRadiation.Set(payload.SolarRadiation)
	metHourlyRainIn.Set(payload.HourlyRainIn)
	metEventRainIn.Set(payload.EventRainIn)
	metDailyRainIn.Set(payload.DailyRainIn)
	metWeeklyRainIn.Set(payload.WeeklyRainIn)
	metMonthlyRainIn.Set(payload.MonthlyRainIn)
	metYearlyRainIn.Set(payload.YearlyRainIn)
	metBattOut.Set(payload.BattOut)
	metBattRain.Set(payload.BattRain)
	metTempInF.Set(payload.TempInF)
	metHumidityIn.Set(payload.HumidityIn)
	metBaromRelIn.Set(payload.BaromRelIn)
	metBaromAbsIn.Set(payload.BaromAbsIn)
	metBattIn.Set(payload.BattIn)
	metLeak1.Set(payload.Leak1)
	metBatLeak1.Set(payload.BatLeak1)
	metLightningDay.Set(payload.LightningDay)
	metLightningTime.Set(time.Since(payload.LightningTime).Seconds())
	metLightingDistance.Set(payload.LightingDistance)
	metBattLightning.Set(payload.BattLightning)
}
