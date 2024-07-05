package handlers

import (
	"encoding/json"
	"net/http"
)

func ResponseData(w http.ResponseWriter, httpCode int, data any) {
	w.WriteHeader(httpCode)

	if httpCode >= 200 && httpCode <= 299 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": httpCode,
			"data":       data,
			"message":    "Successful",
		})
	} else if httpCode >= 400 && httpCode <= 499 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": httpCode,
			"data":       data,
			"message":    "Failure",
		})
	} else if httpCode >= 500 && httpCode <= 599 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": httpCode,
			"data":       data,
			"message":    "Server Error !",
		})
	}
}
