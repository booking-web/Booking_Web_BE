package routes

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Group APIs
	mux.Handle("/api/sample/", http.StripPrefix("/api/sample", sampleMux()))

	return mux
}

func sampleMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/get-sample", GetSample)
	mux.HandleFunc("/post-sample", PostSample)

	return mux
}
