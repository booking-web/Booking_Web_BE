package doctor

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db/doctordb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

// @Summary Update Doctor
// @Description Return A Update Doctor Successful
// @Tags Doctors
// @Accept  json
// @Produce  json
// @Param doctorId body types.HandlerDoctor false "Doctor Information"
// @Success 200 {object} handlers.ResponseDataType
// @Router /api/v1/doctor/update [patch]
func UpdateDoctorRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPatch {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	var doctor types.HandlerDoctor

	body, err := io.ReadAll(r.Body)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &doctor)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, err.Error())
		return
	}

	err = doctordb.UpdateDoctorDB(doctor)

	if err != nil {
		handlers.ResponseData(w, http.StatusInternalServerError, err.Error())
		return
	}

	handlers.ResponseData(w, http.StatusOK, "Updated Doctor Successful!")
}
