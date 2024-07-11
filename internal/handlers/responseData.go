package handlers

import (
	"encoding/json"
	"net/http"
)

type ResponseDataType struct {
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data"`
	Message    string `json:"message"`
}

func ResponseData(w http.ResponseWriter, httpCode int, data any) {
	w.WriteHeader(httpCode)

	if httpCode >= 200 && httpCode <= 299 {
		json.NewEncoder(w).Encode(ResponseDataType{
			StatusCode: httpCode,
			Data:       data,
			Message:    "Successful",
		})
	} else if httpCode >= 400 && httpCode <= 499 {
		json.NewEncoder(w).Encode(ResponseDataType{
			StatusCode: httpCode,
			Data:       data,
			Message:    "Failure",
		})
	} else if httpCode >= 500 && httpCode <= 599 {
		json.NewEncoder(w).Encode(ResponseDataType{
			StatusCode: httpCode,
			Data:       data,
			Message:    "Server Error !",
		})
	}
}
