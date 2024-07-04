package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type PostData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var postData PostData

func GetSample(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postData)
}

func PostSample(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")

		var result PostData
		err := json.NewDecoder(r.Body).Decode(&result)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		postData = result

		log.Printf("Received POST request with ID: %d\n", postData.ID)
		// Here you would typically save the postData to a database
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(postData)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
