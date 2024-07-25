package admin

import (
	"net/http"
	"strconv"

	"github.com/billzayy/Booking_Web_BE/internal/db/admindb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	protectroute "github.com/billzayy/Booking_Web_BE/internal/routes/middleware"
)

// @Summary Delete Doctor
// @Description Return Delete Doctor Successful
// @Tags Admins
// @Accept  json
// @Produce  json
// @Param doctorId query int true "Delete a Doctor"
// @Success 200 {object} handlers.ResponseDataType
// @Router /api/v1/admin/delete-doctor [delete]
func DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	checkAuth, err := protectroute.ProtectedRoute(w, r)

	if err != nil {
		handlers.ResponseData(w, http.StatusUnauthorized, err.Error())
		return
	}

	if checkAuth {
		request := r.URL.Query().Get("doctorId")

		doctorId, err := strconv.Atoi(request)

		if err != nil {
			handlers.ResponseData(w, http.StatusBadRequest, err.Error())
			return
		}

		err = admindb.DeleteDoctorDB(doctorId)

		if err != nil {
			handlers.ResponseData(w, http.StatusInternalServerError, err.Error())
			return
		}

		handlers.ResponseData(w, http.StatusOK, "Delete Doctor Successful !")
	} else {
		handlers.ResponseData(w, http.StatusUnauthorized, err)
		return
	}
}
