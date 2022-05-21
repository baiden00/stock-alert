package handlers

import (
	_ "fmt"
	"net/http"
	 "encoding/json"
)

func HealthCheck(w http.ResponseWriter, rq *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	resp := make(map[string]string)
	resp["running"]="true"
	payload, err := json.Marshal(resp)
	
	if err != nil {
		http.Error(w, "Something Bad Happened", http.StatusExpectationFailed)
		return

	}
	w.Write(payload)

}