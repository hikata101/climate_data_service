package infrastructure

import (
	"github.com/hikata101/climate_data_service/logger"
	openmeteo "github.com/innotechdevops/openmeteo"
)

var maxRequests int = 2 // can be configured
var counterCh chan int = make(chan int, maxRequests)

func Execute(request openmeteo.Parameter) (string, error) {
	counterCh <- 1
	defer func() { <-counterCh }()
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
