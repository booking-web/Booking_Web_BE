package doctor

import (
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db/doctordb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
)

func GetListDoctorRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	data, err := doctordb.ListDoctor()

	if err != nil {
		handlers.ResponseData(w, http.StatusInternalServerError, err.Error())
		return
	}

	handlers.ResponseData(w, http.StatusOK, data)
}
