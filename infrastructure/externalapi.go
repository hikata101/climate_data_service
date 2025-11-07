package infrastructure

import (
	"github.com/hikata101/climate_data_service/logger"
	openmeteo "github.com/innotechdevops/openmeteo"
)

func Execute(request openmeteo.Parameter) (string, error) {
	logger.Logger.Debug("Executing OpenMeteo API request")
	m := openmeteo.New()
	resp, err := m.Execute(request)
	if err != nil {
		logger.Logger.Error(err.Error())
		return "", err
	}
	logger.Logger.Debug("OpenMeteo API request executed successfully")
	return resp, nil
}
