package httpResp

import (
	"encoding/json"
	"net/http"
)

func RespondwithJSON(w http.ResponseWriter, code int, Payload interface{}) {
	response, _ := json.Marshal(Payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondwithError(w http.ResponseWriter, code int, message string) {
	RespondwithJSON(w, code, map[string]string{"error":message})

}
