package authenticate

import (
	"fmt"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db/userdb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/pkg/email"
)

// @Summary Forgot Password
// @Description Returns Sent to Your Email Successful
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userId query int true "userId"
// @Success 200 {object} handlers.ResponseDataType
// @Router /api/v1/forgot-password [get]
func ForgotPassRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

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
		email.SendMail(data[0].Email)
		handlers.ResponseData(w, http.StatusOK, "Sent to your email !")
		return
	}
}
