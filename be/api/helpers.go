package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func respond(dataField string, data interface{}, msg string, statusCode int, w http.ResponseWriter) {
	// Generate response data.
	respData := make(map[string]interface{})

	if dataField != "" {
		respData[dataField] = data
	}

	respData["msg"] = msg
	respData["status"] = statusCode

	// Return JSON type.
	w.Header().Set("Content-Type", "application/json")

	// Respond to foreign domains.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Marshal response.
	resp, respErr := json.Marshal(respData)

	if respErr != nil {
		log.Println("failed to generate the API response, msg:", msg)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(resp)
}
