package routes

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Group APIs
	mux.Handle("/api/sample/", http.StripPrefix("/api/sample", sampleMux()))
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", userMux()))

	return mux
}

func sampleMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/get-sample", GetSample)
	mux.HandleFunc("/post-sample", PostSample)

	return mux
}

func userMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/get-user", GetUserByIdRoute)
	mux.HandleFunc("/sign-up", SignUpRoute)

	return mux
}
