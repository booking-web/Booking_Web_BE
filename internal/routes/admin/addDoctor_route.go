package admin

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db/admindb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	protectroute "github.com/billzayy/Booking_Web_BE/internal/routes/middleware"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

// @Summary Add Doctor
// @Description Return Add Doctor Successful
// @Tags Admins
// @Accept  json
// @Produce  json
// @Param doctor body types.HandlerDoctor true "Add a New Doctor"
// @Success 201 {object} handlers.ResponseDataType
// @Router /api/v1/admin/add-doctor [post]
func AddDoctorRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	checkAuth, err := protectroute.ProtectedRoute(w, r)

	if err != nil {
		handlers.ResponseData(w, http.StatusUnauthorized, err.Error())
		return
	}

	if checkAuth {
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

		err = admindb.AddDoctorDB(doctor)

		if err != nil {
			handlers.ResponseData(w, http.StatusBadRequest, "Bad Request")
			return
		}

		handlers.ResponseData(w, http.StatusCreated, "Add Doctor Successful")
	} else {
		handlers.ResponseData(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
}
