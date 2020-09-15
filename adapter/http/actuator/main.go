package actuator

import (
	"encoding/json"
	"net/http"
)

//HealthBody estrutura status da app
type HealthBody struct {
	Status string `json:"status"`
}

//Health exibe status da app
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	var status = HealthBody{"alive"}

	_ = json.NewEncoder(w).Encode(status)
}