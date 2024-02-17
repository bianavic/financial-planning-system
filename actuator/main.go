package actuator

import (
	"encoding/json"
	"net/http"
)

type HealthBody struct {
	Status string `json:"status"`
}

func Health(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	var status = HealthBody{"alive"}

	_ = json.NewEncoder(writer).Encode(status)
}
