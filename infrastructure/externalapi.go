package infrastructure

import (
	"log/slog"

	openmeteo "github.com/innotechdevops/openmeteo"
)

var maxRequests int = 2 // can be configured
var counterCh chan int = make(chan int, maxRequests)

func Execute(request openmeteo.Parameter) (string, error) {
	counterCh <- 1
	defer func() { <-counterCh }()
	slog.Debug("Executing OpenMeteo API request")
	m := openmeteo.New()
	resp, err := m.Execute(request)
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}
	slog.Debug("OpenMeteo API request executed successfully")
	return resp, nil
}
