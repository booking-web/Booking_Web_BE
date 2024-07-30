package routes

import (
	"net/http"
	"os"

	"github.com/billzayy/Booking_Web_BE/api"
	"github.com/billzayy/Booking_Web_BE/internal/routes/admin"
	"github.com/billzayy/Booking_Web_BE/internal/routes/authenticate"
	"github.com/billzayy/Booking_Web_BE/internal/routes/doctor"
	"github.com/billzayy/Booking_Web_BE/internal/routes/user"

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

	mux.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./assets"))))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Group APIs
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", userMux()))
	mux.Handle("/api/v1/doctor/", http.StripPrefix("/api/v1/doctor", doctorMux()))
	mux.Handle("/api/v1/admin/", http.StripPrefix("/api/v1/admin", adminMux()))

	return mux
}

func userMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/change-password", authenticate.ChangePassRoute)
	mux.HandleFunc("/forgot-password", authenticate.ForgotPassRoute)
	mux.HandleFunc("/get-user", user.GetUserByIdRoute)
	mux.HandleFunc("/sign-up", authenticate.SignUpRoute)
	mux.HandleFunc("/login", authenticate.LogInRoute)
	mux.HandleFunc("/delete-user", user.DeleteUserRoute)
	mux.HandleFunc("/update-user", user.UpdateUser)

	return mux
}

func doctorMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/detail", doctor.GetDoctorRoute)
	mux.HandleFunc("/list", doctor.GetListDoctorRoute)
	mux.HandleFunc("/update", doctor.UpdateDoctorRoute)

	return mux
}

func adminMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/add-doctor", admin.AddDoctorRoute)
	mux.HandleFunc("/delete-doctor", admin.DeleteDoctor)
	mux.HandleFunc("/delete-user", user.DeleteUserRoute)

	return mux
}
