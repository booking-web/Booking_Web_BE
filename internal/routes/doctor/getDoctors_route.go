package doctor

import (
	"net/http"
	"strconv"

	"github.com/billzayy/Booking_Web_BE/internal/db/doctordb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
)

// @Summary Get Doctors List
// @Description Returns List of Doctors
// @Tags Doctors
// @Accept  json
// @Produce  json
// @Param name query string false "Doctor Name"
// @Success 200 {object} types.ResponseDoctor
// @Router /api/v1/doctor/list [get]
func GetListRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	doctorId := r.URL.Query().Get("doctorId")

	input, err := strconv.Atoi(doctorId)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := doctordb.GetDoctorById(input)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, err.Error())
		return
	}

	handlers.ResponseData(w, http.StatusOK, data)
}
