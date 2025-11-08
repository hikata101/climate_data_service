package domain

import (
	"time"

	climate_data "github.com/hikata101/climate_data/v1"
)

type Dataset struct {
	ID        string
	Name      string
	Format    string
	Size      int64
	Path      string
	CreatedAt time.Time
}

type OpenMeteoHourlyUnits struct {
	Time                     string `json:"time,omitempty"`
	Temperature2m            string `json:"temperature_2m,omitempty"`
	RelativeHumidity2m       string `json:"relative_humidity_2m,omitempty"`
	DewPoint2m               string `json:"dew_point_2m,omitempty"`
	ApparentTemperature      string `json:"apparent_temperature,omitempty"`
	PressureMsl              string `json:"pressure_msl,omitempty"`
	SurfacePressure          string `json:"surface_pressure,omitempty"`
	CloudCover               string `json:"cloud_cover,omitempty"`
	CloudCoverLow            string `json:"cloud_cover_low,omitempty"`
	CloudCoverMid            string `json:"cloud_cover_mid,omitempty"`
	CloudCoverHigh           string `json:"cloud_cover_high,omitempty"`
	WindSpeed10m             string `json:"wind_speed_10m,omitempty"`
	WindSpeed80m             string `json:"wind_speed_80m,omitempty"`
	WindSpeed120m            string `json:"wind_speed_120m,omitempty"`
	WindSpeed180m            string `json:"wind_speed_180m,omitempty"`
	WindDirection10m         string `json:"wind_direction_10m,omitempty"`
	WindDirection80m         string `json:"wind_direction_80m,omitempty"`
	WindDirection120m        string `json:"wind_direction_120m,omitempty"`
	WindDirection180m        string `json:"wind_direction_180m,omitempty"`
	WindGusts10m             string `json:"wind_gusts_10m,omitempty"`
	ShortwaveRadiation       string `json:"shortwave_radiation,omitempty"`
	DirectRadiation          string `json:"direct_radiation,omitempty"`
	DirectNormalIrradiance   string `json:"direct_normal_irradiance,omitempty"`
	DiffuseRadiation         string `json:"diffuse_radiation,omitempty"`
	VaporPressureDeficit     string `json:"vapor_pressure_deficit,omitempty"`
	Cape                     string `json:"cape,omitempty"`
	Evapotranspiration       string `json:"evapotranspiration,omitempty"`
	Et0FaoEvapotranspiration string `json:"et0_fao_evapotranspiration,omitempty"`
	Precipitation            string `json:"precipitation,omitempty"`
	Snowfall                 string `json:"snowfall,omitempty"`
	PrecipitationProbability string `json:"precipitation_probability,omitempty"`
	Rain                     string `json:"rain,omitempty"`
	Showers                  string `json:"showers,omitempty"`
	WeatherCode              string `json:"weather_code,omitempty"`
	SnowDepth                string `json:"snow_depth,omitempty"`
	FreezingLevelHeight      string `json:"freezing_level_height,omitempty"`
	Visibility               string `json:"visibility,omitempty"`
	SoilTemperature0cm       string `json:"soil_temperature_0cm,omitempty"`
	SoilTemperature6cm       string `json:"soil_temperature_6cm,omitempty"`
	SoilTemperature18cm      string `json:"soil_temperature_18cm,omitempty"`
	SoilTemperature54cm      string `json:"soil_temperature_54cm,omitempty"`
	SoilMoisture01cm         string `json:"soil_moisture_01cm,omitempty"`
	SoilMoisture13cm         string `json:"soil_moisture_13cm,omitempty"`
	SoilMoisture39cm         string `json:"soil_moisture_39cm,omitempty"`
	SoilMoisture927cm        string `json:"soil_moisture_927cm,omitempty"`
	SoilMoisture2781cm       string `json:"soil_moisture_2781cm,omitempty"`
}

type OpenMeteoHourly struct {
	Time                     []string  `json:"time,omitempty"`
	Temperature2m            []float32 `json:"temperature_2m,omitempty"`
	RelativeHumidity2m       []float32 `json:"relative_humidity_2m,omitempty"`
	DewPoint2m               []float32 `json:"dew_point_2m,omitempty"`
	ApparentTemperature      []float32 `json:"apparent_temperature,omitempty"`
	PressureMsl              []float32 `json:"pressure_msl,omitempty"`
	SurfacePressure          []float32 `json:"surface_pressure,omitempty"`
	CloudCover               []float32 `json:"cloud_cover,omitempty"`
	CloudCoverLow            []float32 `json:"cloud_cover_low,omitempty"`
	CloudCoverMid            []float32 `json:"cloud_cover_mid,omitempty"`
	CloudCoverHigh           []float32 `json:"cloud_cover_high,omitempty"`
	WindSpeed10m             []float32 `json:"wind_speed_10m,omitempty"`
	WindSpeed80m             []float32 `json:"wind_speed_80m,omitempty"`
	WindSpeed120m            []float32 `json:"wind_speed_120m,omitempty"`
	WindSpeed180m            []float32 `json:"wind_speed_180m,omitempty"`
	WindDirection10m         []float32 `json:"wind_direction_10m,omitempty"`
	WindDirection80m         []float32 `json:"wind_direction_80m,omitempty"`
	WindDirection120m        []float32 `json:"wind_direction_120m,omitempty"`
	WindDirection180m        []float32 `json:"wind_direction_180m,omitempty"`
	WindGusts10m             []float32 `json:"wind_gusts_10m,omitempty"`
	ShortwaveRadiation       []float32 `json:"shortwave_radiation,omitempty"`
	DirectRadiation          []float32 `json:"direct_radiation,omitempty"`
	DirectNormalIrradiance   []float32 `json:"direct_normal_irradiance,omitempty"`
	DiffuseRadiation         []float32 `json:"diffuse_radiation,omitempty"`
	VaporPressureDeficit     []float32 `json:"vapor_pressure_deficit,omitempty"`
	Cape                     []float32 `json:"cape,omitempty"`
	Evapotranspiration       []float32 `json:"evapotranspiration,omitempty"`
	Et0FaoEvapotranspiration []float32 `json:"et0_fao_evapotranspiration,omitempty"`
	Precipitation            []float32 `json:"precipitation,omitempty"`
	Snowfall                 []float32 `json:"snowfall,omitempty"`
	PrecipitationProbability []float32 `json:"precipitation_probability,omitempty"`
	Rain                     []float32 `json:"rain,omitempty"`
	Showers                  []float32 `json:"showers,omitempty"`
	WeatherCode              []float32 `json:"weather_code,omitempty"`
	SnowDepth                []float32 `json:"snow_depth,omitempty"`
	FreezingLevelHeight      []float32 `json:"freezing_level_height,omitempty"`
	Visibility               []float32 `json:"visibility,omitempty"`
	SoilTemperature0cm       []float32 `json:"soil_temperature_0cm,omitempty"`
	SoilTemperature6cm       []float32 `json:"soil_temperature_6cm,omitempty"`
	SoilTemperature18cm      []float32 `json:"soil_temperature_18cm,omitempty"`
	SoilTemperature54cm      []float32 `json:"soil_temperature_54cm,omitempty"`
	SoilMoisture01cm         []float32 `json:"soil_moisture_01cm,omitempty"`
	SoilMoisture13cm         []float32 `json:"soil_moisture_13cm,omitempty"`
	SoilMoisture39cm         []float32 `json:"soil_moisture_39cm,omitempty"`
	SoilMoisture927cm        []float32 `json:"soil_moisture_927cm,omitempty"`
	SoilMoisture2781cm       []float32 `json:"soil_moisture_2781cm,omitempty"`
}

type OpenMeteo_Response struct {
	Status               int32                                      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Latitude             float32                                    `protobuf:"fixed32,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32                                    `protobuf:"fixed32,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	GenerationtimeMs     float32                                    `protobuf:"fixed32,4,opt,name=generationtime_ms,json=generationtimeMs,proto3" json:"generationtime_ms,omitempty"`
	UtcOffsetSeconds     int32                                      `protobuf:"varint,5,opt,name=utc_offset_seconds,json=utcOffsetSeconds,proto3" json:"utc_offset_seconds,omitempty"`
	Timezone             string                                     `protobuf:"bytes,6,opt,name=timezone,proto3" json:"timezone,omitempty"`
	TimezoneAbbreviation string                                     `protobuf:"bytes,7,opt,name=timezone_abbreviation,json=timezoneAbbreviation,proto3" json:"timezone_abbreviation,omitempty"`
	Elevation            float32                                    `protobuf:"fixed32,8,opt,name=elevation,proto3" json:"elevation,omitempty"`
	CurrentWeatherUnits  *climate_data.OpenMeteoCurrentWeatherUnits `protobuf:"bytes,9,opt,name=current_weather_units,json=currentWeatherUnits,proto3" json:"current_weather_units,omitempty"`
	CurrentWeather       *climate_data.OpenMeteoCurrentWeather      `protobuf:"bytes,10,opt,name=current_weather,json=currentWeather,proto3" json:"current_weather,omitempty"`
	HourlyUnits          *OpenMeteoHourlyUnits                      `protobuf:"bytes,11,opt,name=hourly_units,json=hourlyUnits,proto3" json:"hourly_units,omitempty"`
	Hourly               *OpenMeteoHourly                           `protobuf:"bytes,12,opt,name=hourly,proto3" json:"hourly,omitempty"`
}
