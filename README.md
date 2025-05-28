# wecap

Weather Capture server to intercept Ambient Weather output and expose as prometheus metrics.

## Setup

In your weather console, go to settings, then weather server setup. Choose customized setup, state enabled, protocol type `same as AMBWeather`. IP/HostName set to the IP of your server running wecap. Port set to the port wecap is listening on (default `8000`). Interval set to whatever you want. Path is `/data/report/` (same as default).

## Metrics

The following metrics will be made available on `/metrics`:

- temp_f
- humidity
- wind_speed_mph
- wind_gust_mph
- max_daily_gust
- wind_dir
- wind_dir_avg_10m
- uv
- solar_radiation
- hourly_rain_inches
- event_rain_inches
- daily_rain_inches
- weekly_rain_inches
- monthly_rain_inches
- yearly_rain_inches
- batt_outside
- batt_rain
- temp_indoor_f
- humidity_indoor
- barom_rel_indoor
- barom_abs_indoor
- batt_inside
- leak_1
- batt_leak_1
- lightning_day
- lightning_time
- lighting_distance
- batt_lightning

## Routes

### /data/report/

Captures AMBWeather reports and converts into metrics.

### /metrics

Output prometheus metrics.

### /*

Any other requests will emit a debug log message and respond with a 200.
