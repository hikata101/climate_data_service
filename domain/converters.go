package domain

import (
	pb "github.com/hikata101/climate_data_service/gen/github.com/hikata101/climate_data_service/v1"
	openmeteo "github.com/innotechdevops/openmeteo"
)

var dic = map[pb.Hourly]string{
	pb.Hourly_HOURLY_TEMPERATURE_2M:             openmeteo.HourlyTemperature2m,
	pb.Hourly_HOURLY_RELATIVE_HUMIDITY_2M:       openmeteo.HourlyRelativeHumidity2m,
	pb.Hourly_HOURLY_DEW_POINT_2M:               openmeteo.HourlyDewPoint2m,
	pb.Hourly_HOURLY_APPARENT_TEMPERATURE:       openmeteo.HourlyApparentTemperature,
	pb.Hourly_HOURLY_PRESSURE_MSL:               openmeteo.HourlyPressureMsl,
	pb.Hourly_HOURLY_SURFACE_PRESSURE:           openmeteo.HourlySurfacePressure,
	pb.Hourly_HOURLY_CLOUDCOVER:                 openmeteo.HourlyCloudCover,
	pb.Hourly_HOURLY_CLOUDCOVER_LOW:             openmeteo.HourlyCloudCoverLow,
	pb.Hourly_HOURLY_CLOUDCOVER_MID:             openmeteo.HourlyCloudCoverMid,
	pb.Hourly_HOURLY_CLOUDCOVER_HIGH:            openmeteo.HourlyCloudCoverHigh,
	pb.Hourly_HOURLY_WINDSPEED_10M:              openmeteo.HourlyWindSpeed10m,
	pb.Hourly_HOURLY_WINDSPEED_80M:              openmeteo.HourlyWindSpeed80m,
	pb.Hourly_HOURLY_WINDSPEED_120M:             openmeteo.HourlyWindSpeed120m,
	pb.Hourly_HOURLY_WINDSPEED_180M:             openmeteo.HourlyWindSpeed180m,
	pb.Hourly_HOURLY_WINDDIRECTION_10M:          openmeteo.HourlyWindDirection10m,
	pb.Hourly_HOURLY_WINDDIRECTION_80M:          openmeteo.HourlyWindDirection80m,
	pb.Hourly_HOURLY_WINDDIRECTION_120M:         openmeteo.HourlyWindDirection120m,
	pb.Hourly_HOURLY_WINDDIRECTION_180M:         openmeteo.HourlyWindDirection180m,
	pb.Hourly_HOURLY_WINDGUSTS_10M:              openmeteo.HourlyWindGusts10m,
	pb.Hourly_HOURLY_SHORTWAVE_RADIATION:        openmeteo.HourlyShortwaveRadiation,
	pb.Hourly_HOURLY_DIRECT_RADIATION:           openmeteo.HourlyDirectRadiation,
	pb.Hourly_HOURLY_DIRECT_NORMAL_IRRADIANCE:   openmeteo.HourlyDirectNormalIrradiance,
	pb.Hourly_HOURLY_DIFFUSE_RADIATION:          openmeteo.HourlyDiffuseRadiation,
	pb.Hourly_HOURLY_VAPOR_PRESSURE_DEFICIT:     openmeteo.HourlyVaporPressureDeficit,
	pb.Hourly_HOURLY_CAPE:                       openmeteo.HourlyCape,
	pb.Hourly_HOURLY_EVAPOTRANSPIRATION:         openmeteo.HourlyEvapotranspiration,
	pb.Hourly_HOURLY_ET0_FAO_EVAPOTRANSPIRATION: openmeteo.HourlyEt0FaoEvapotranspiration,
	pb.Hourly_HOURLY_PRECIPITATION:              openmeteo.HourlyPrecipitation,
	pb.Hourly_HOURLY_SNOWFALL:                   openmeteo.HourlySnowfall,
	pb.Hourly_HOURLY_PRECIPITATION_PROBABILITY:  openmeteo.HourlyPrecipitationProbability,
	pb.Hourly_HOURLY_RAIN:                       openmeteo.HourlyRain,
	pb.Hourly_HOURLY_SHOWERS:                    openmeteo.HourlyShowers,
	pb.Hourly_HOURLY_WEATHERCODE:                openmeteo.HourlyWeatherCode,
	pb.Hourly_HOURLY_SNOW_DEPTH:                 openmeteo.HourlySnowDepth,
	pb.Hourly_HOURLY_FREEZINGLEVEL_HEIGHT:       openmeteo.HourlyFreezingLevelHeight,
	pb.Hourly_HOURLY_VISIBILITY:                 openmeteo.HourlyVisibility,
	pb.Hourly_HOURLY_SOIL_TEMPERATURE_0CM:       openmeteo.HourlySoilTemperature0cm,
	pb.Hourly_HOURLY_SOIL_TEMPERATURE_6CM:       openmeteo.HourlySoilTemperature6cm,
	pb.Hourly_HOURLY_SOIL_TEMPERATURE_18CM:      openmeteo.HourlySoilTemperature18cm,
	pb.Hourly_HOURLY_SOIL_TEMPERATURE_54CM:      openmeteo.HourlySoilTemperature54cm,
	pb.Hourly_HOURLY_SOIL_MOISTURE_01CM:         openmeteo.HourlySoilMoisture01cm,
	pb.Hourly_HOURLY_SOIL_MOISTURE_13CM:         openmeteo.HourlySoilMoisture13cm,
	pb.Hourly_HOURLY_SOIL_MOISTURE_39CM:         openmeteo.HourlySoilMoisture39cm,
	pb.Hourly_HOURLY_SOIL_MOISTURE_927CM:        openmeteo.HourlySoilMoisture927cm,
	pb.Hourly_HOURLY_SOIL_MOISTURE_2781CM:       openmeteo.HourlySoilMoisture2781cm,
	pb.Hourly_HOURLY_IS_DAY:                     "is_day",
}

func convertHourlyParameters(params []pb.Hourly) *[]string {
	var result []string
	for _, p := range params {
		if v, ok := dic[p]; ok {
			result = append(result, v)
		}
	}
	return &result
}

// func convertHourlyResults(data *OpenMeteo_Response) *pb.DownloadReply_OpenMeteo {
// 	if data == nil || data.Hourly == nil {
// 		return &pb.DownloadReply_OpenMeteo{}
// 	}
// 	result := &pb.DownloadReply_OpenMeteo{
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
// 		Hourly: &pb.OpenMeteoHourly{
// 			Time: data.Hourly.Time,
// 		},
// 	}
// 	return result
// }
