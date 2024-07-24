package user

import (
	"fmt"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db/userdb"
	protectroute "github.com/billzayy/Booking_Web_BE/internal/routes/protect_route"

	"github.com/billzayy/Booking_Web_BE/internal/handlers"
)

// @Summary Get user details by Id
// @Description Returns user details by Id
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userId query int false "User Id"
// @Security  Bearer
// @Success 200 {object} types.Users
// @Router /api/v1/get-user [get]
func GetUserByIdRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	checkAuth, err := protectroute.ProtectedRoute(w, r)

	if err != nil {
		handlers.ResponseData(w, http.StatusUnauthorized, err.Error())
		return
	}

	if checkAuth {
		userId := r.URL.Query().Get("userId")

		data, err := userdb.GetUserById(userId)

		if err != nil {
			handlers.ResponseData(w, http.StatusNotFound, err.Error())
			return
		}

		if len(data) == 0 {
			handlers.ResponseData(w, http.StatusNotFound, fmt.Sprintf("User %v not found!", userId))
			return
		} else {
			handlers.ResponseData(w, http.StatusOK, data[0])
			return
		}
	} else {
		handlers.ResponseData(w, http.StatusUnauthorized, err)
		return
	}
}
