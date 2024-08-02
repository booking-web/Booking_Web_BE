package authenticate

import (
	"fmt"
	"log"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/pkg"
	"github.com/billzayy/Booking_Web_BE/internal/pkg/email"
)

// @Summary Forgot Password
// @Description Return Sent to Your Email Successful
// @Tags Users
// @Accept  json
// @Produce  json
// @Param email query string true "User's Email"
// @Success 200 {object} handlers.ResponseDataType
// @Router /api/v1/forgot-password [post]
func ForgotPassRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	userMail := r.URL.Query().Get("email")

	data, err := pkg.CheckExistedEmail(userMail)

	if err != nil {
		log.Println(err)
		handlers.ResponseData(w, http.StatusNotFound, err.Error())
		return
	}

	if !data {
		log.Println(err)
		handlers.ResponseData(w, http.StatusNotFound, fmt.Sprintf("User %v not found!", userMail))
		return
	} else {
		email.SendMail(userMail)
		handlers.ResponseData(w, http.StatusOK, "Sent to your email !")
		return
	}
}
