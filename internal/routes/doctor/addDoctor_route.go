package doctor

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db/doctordb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

// @Summary Add Doctor
// @Description Return Add Doctor Successful
// @Tags Doctors
// @Accept  json
// @Produce  json
// @Param doctor body types.HandlerDoctor true "Add a New Doctor"
// @Success 200 {object} handlers.ResponseDataType
// @Router /api/v1/doctor/add [post]
func AddDoctorRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, "Bad Request")
		return
	}
	var doctor types.HandlerDoctor

	err = json.Unmarshal(body, &doctor)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, "Invalid Body")
		return
	}

	err = doctordb.AddDoctor(doctor)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, "Bad Request")
		return
	}

	handlers.ResponseData(w, http.StatusOK, "Add Doctor Successful")
}
