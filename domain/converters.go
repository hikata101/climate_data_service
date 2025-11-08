package domain

import (
	climate_data "github.com/hikata101/climate_data"
	openmeteo "github.com/innotechdevops/openmeteo"
)

var dic = map[climate_data.Hourly]string{
	climate_data.Hourly_HOURLY_TEMPERATURE_2M:             openmeteo.HourlyTemperature2m,
	climate_data.Hourly_HOURLY_RELATIVE_HUMIDITY_2M:       openmeteo.HourlyRelativeHumidity2m,
	climate_data.Hourly_HOURLY_DEW_POINT_2M:               openmeteo.HourlyDewPoint2m,
	climate_data.Hourly_HOURLY_APPARENT_TEMPERATURE:       openmeteo.HourlyApparentTemperature,
	climate_data.Hourly_HOURLY_PRESSURE_MSL:               openmeteo.HourlyPressureMsl,
	climate_data.Hourly_HOURLY_SURFACE_PRESSURE:           openmeteo.HourlySurfacePressure,
	climate_data.Hourly_HOURLY_CLOUDCOVER:                 openmeteo.HourlyCloudCover,
	climate_data.Hourly_HOURLY_CLOUDCOVER_LOW:             openmeteo.HourlyCloudCoverLow,
	climate_data.Hourly_HOURLY_CLOUDCOVER_MID:             openmeteo.HourlyCloudCoverMid,
	climate_data.Hourly_HOURLY_CLOUDCOVER_HIGH:            openmeteo.HourlyCloudCoverHigh,
	climate_data.Hourly_HOURLY_WINDSPEED_10M:              openmeteo.HourlyWindSpeed10m,
	climate_data.Hourly_HOURLY_WINDSPEED_80M:              openmeteo.HourlyWindSpeed80m,
	climate_data.Hourly_HOURLY_WINDSPEED_120M:             openmeteo.HourlyWindSpeed120m,
	climate_data.Hourly_HOURLY_WINDSPEED_180M:             openmeteo.HourlyWindSpeed180m,
	climate_data.Hourly_HOURLY_WINDDIRECTION_10M:          openmeteo.HourlyWindDirection10m,
	climate_data.Hourly_HOURLY_WINDDIRECTION_80M:          openmeteo.HourlyWindDirection80m,
	climate_data.Hourly_HOURLY_WINDDIRECTION_120M:         openmeteo.HourlyWindDirection120m,
	climate_data.Hourly_HOURLY_WINDDIRECTION_180M:         openmeteo.HourlyWindDirection180m,
	climate_data.Hourly_HOURLY_WINDGUSTS_10M:              openmeteo.HourlyWindGusts10m,
	climate_data.Hourly_HOURLY_SHORTWAVE_RADIATION:        openmeteo.HourlyShortwaveRadiation,
	climate_data.Hourly_HOURLY_DIRECT_RADIATION:           openmeteo.HourlyDirectRadiation,
	climate_data.Hourly_HOURLY_DIRECT_NORMAL_IRRADIANCE:   openmeteo.HourlyDirectNormalIrradiance,
	climate_data.Hourly_HOURLY_DIFFUSE_RADIATION:          openmeteo.HourlyDiffuseRadiation,
	climate_data.Hourly_HOURLY_VAPOR_PRESSURE_DEFICIT:     openmeteo.HourlyVaporPressureDeficit,
	climate_data.Hourly_HOURLY_CAPE:                       openmeteo.HourlyCape,
	climate_data.Hourly_HOURLY_EVAPOTRANSPIRATION:         openmeteo.HourlyEvapotranspiration,
	climate_data.Hourly_HOURLY_ET0_FAO_EVAPOTRANSPIRATION: openmeteo.HourlyEt0FaoEvapotranspiration,
	climate_data.Hourly_HOURLY_PRECIPITATION:              openmeteo.HourlyPrecipitation,
	climate_data.Hourly_HOURLY_SNOWFALL:                   openmeteo.HourlySnowfall,
	climate_data.Hourly_HOURLY_PRECIPITATION_PROBABILITY:  openmeteo.HourlyPrecipitationProbability,
	climate_data.Hourly_HOURLY_RAIN:                       openmeteo.HourlyRain,
	climate_data.Hourly_HOURLY_SHOWERS:                    openmeteo.HourlyShowers,
	climate_data.Hourly_HOURLY_WEATHERCODE:                openmeteo.HourlyWeatherCode,
	climate_data.Hourly_HOURLY_SNOW_DEPTH:                 openmeteo.HourlySnowDepth,
	climate_data.Hourly_HOURLY_FREEZINGLEVEL_HEIGHT:       openmeteo.HourlyFreezingLevelHeight,
	climate_data.Hourly_HOURLY_VISIBILITY:                 openmeteo.HourlyVisibility,
	climate_data.Hourly_HOURLY_SOIL_TEMPERATURE_0CM:       openmeteo.HourlySoilTemperature0cm,
	climate_data.Hourly_HOURLY_SOIL_TEMPERATURE_6CM:       openmeteo.HourlySoilTemperature6cm,
	climate_data.Hourly_HOURLY_SOIL_TEMPERATURE_18CM:      openmeteo.HourlySoilTemperature18cm,
	climate_data.Hourly_HOURLY_SOIL_TEMPERATURE_54CM:      openmeteo.HourlySoilTemperature54cm,
	climate_data.Hourly_HOURLY_SOIL_MOISTURE_01CM:         openmeteo.HourlySoilMoisture01cm,
	climate_data.Hourly_HOURLY_SOIL_MOISTURE_13CM:         openmeteo.HourlySoilMoisture13cm,
	climate_data.Hourly_HOURLY_SOIL_MOISTURE_39CM:         openmeteo.HourlySoilMoisture39cm,
	climate_data.Hourly_HOURLY_SOIL_MOISTURE_927CM:        openmeteo.HourlySoilMoisture927cm,
	climate_data.Hourly_HOURLY_SOIL_MOISTURE_2781CM:       openmeteo.HourlySoilMoisture2781cm,
	climate_data.Hourly_HOURLY_IS_DAY:                     "is_day",
}

func convertHourlyParameters(params []climate_data.Hourly) *[]string {
	var result []string
	for _, p := range params {
		if v, ok := dic[p]; ok {
			result = append(result, v)
		}
	}
	return &result
}

// func convertHourlyResults(data *OpenMeteo_Response) *climate_data.DownloadReply_OpenMeteo {
// 	if data == nil || data.Hourly == nil {
// 		return &climate_data.DownloadReply_OpenMeteo{}
// 	}
// 	result := &climate_data.DownloadReply_OpenMeteo{
// 		Status:               data.Status,
// 		Latitude:             data.Latitude,
// 		Longitude:            data.Longitude,
// 		GenerationtimeMs:     data.GenerationtimeMs,
// 		UtcOffsetSeconds:     data.UtcOffsetSeconds,
// 		Timezone:             data.Timezone,
// 		TimezoneAbbreviation: data.TimezoneAbbreviation,
// 		Elevation:            data.Elevation,
// 		CurrentWeatherUnits:  data.CurrentWeatherUnits,
// 		CurrentWeather:       data.CurrentWeather,
// 		HourlyUnits:          data.HourlyUnits,
// 		Hourly: &climate_data.OpenMeteoHourly{
// 			Time: data.Hourly.Time,
// 		},
// 	}
// 	return result
// }
