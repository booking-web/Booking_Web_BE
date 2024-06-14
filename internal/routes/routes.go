package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PostData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	userMux := http.NewServeMux()

	userMux.HandleFunc("/get", GetController)
	userMux.HandleFunc("/post", PostController)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", userMux))

	return mux
}

func GetController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")

	var postData PostData
	fmt.Println("Hello")
	fmt.Fprint(w, postData)
}

func PostController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")

		var postData PostData
		err := json.NewDecoder(r.Body).Decode(&postData)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Received POST request with ID: %d\n", postData.ID)
		// Here you would typically save the postData to a database
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(postData)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
