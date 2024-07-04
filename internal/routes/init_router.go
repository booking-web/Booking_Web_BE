package routes

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	userMux := http.NewServeMux()

	userMux.HandleFunc("/get", GetController)
	userMux.HandleFunc("/post", PostController)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", userMux))

	return mux
}
