package user

import (
	"net/http"
	"strconv"

	"github.com/billzayy/Booking_Web_BE/internal/db/userdb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	protectroute "github.com/billzayy/Booking_Web_BE/internal/routes/middleware"
)

// @Summary Delete user
// @Description Returns Delete Successful
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userId query int false "User Id"
// @Security  Bearer
// @Success 200 {object} handlers.ResponseDataType
// @Router /api/v1/delete-user [delete]

// @Summary Delete user
// @Description Returns Delete Successful
// @Tags Admins
// @Accept  json
// @Produce  json
// @Param userId query int false "User Id"
// @Security  Bearer
// @Success 200 {object} handlers.ResponseDataType
// @Router /api/v1/admin/delete-user [delete]
func DeleteUserRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid Method")
		return
	}

	checkAuth, err := protectroute.ProtectedRoute(w, r)

	if err != nil {
		handlers.ResponseData(w, http.StatusUnauthorized, err.Error())
		return
	}

	if checkAuth {
		request := r.URL.Query().Get("userId")

		doctorId, err := strconv.Atoi(request)

		if err != nil {
			handlers.ResponseData(w, http.StatusBadRequest, err.Error())
			return
		}

		err = userdb.DeleteUserDB(doctorId)

		if err != nil {
			handlers.ResponseData(w, http.StatusInternalServerError, err.Error())
			return
		}

		handlers.ResponseData(w, http.StatusOK, "Delete User Successful")
		return
	}
}
