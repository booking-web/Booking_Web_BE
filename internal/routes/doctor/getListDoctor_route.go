package doctor

import (
	"log"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db/doctordb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
)

// @Summary Get List Doctor
// @Description Return A List of Doctors
// @Tags Doctors
// @Accept  json
// @Produce  json
// @Success 200 {object} handlers.ResponseDataType
// @Router /api/v1/doctor/list [get]
func GetListDoctorRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	data, err := doctordb.ListDoctor()

	if err != nil {
		log.Println(err)
		handlers.ResponseData(w, http.StatusInternalServerError, err.Error())
		return
	}

	handlers.ResponseData(w, http.StatusOK, data)
}
