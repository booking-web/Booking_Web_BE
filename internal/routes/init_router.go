package routes

import (
	"net/http"
	"os"

	"github.com/billzayy/Booking_Web_BE/api"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	// programmatically set swagger info
	api.SwaggerInfo.Title = "Swagger GoLang JioHealth"
	api.SwaggerInfo.Description = "This is a Swagger Golang APIs Server."
	api.SwaggerInfo.Version = "3.0"
	api.SwaggerInfo.Host = "localhost:" + os.Getenv("PORT")
	api.SwaggerInfo.Schemes = []string{"http", "https"}

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

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
	mux.HandleFunc("/login", LogInRoute)

	return mux
}
