package infrastructure

import (
	"fmt"

	openmeteo "github.com/innotechdevops/openmeteo"
)

func Execute(request openmeteo.Parameter) (string, error) {
	m := openmeteo.New()
	resp, err := m.Execute(request)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return resp, nil
}
