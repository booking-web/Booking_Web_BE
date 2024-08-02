package authenticate

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db/authenticatedb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
)

// @Summary Change Password
// @Description Return Change Password Successful
// @Tags Users
// @Accept  json
// @Produce  json
// @Param changePassword body handlers.ForgotPass true "Change Password"
// @Success 200 {object} handlers.ResponseDataType
// @Router /api/v1/change-password [post]
func ChangePassRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		handlers.ResponseData(w, http.StatusBadRequest, "Bad Request")
		return
	}

	var changePass handlers.ForgotPass

	err = json.Unmarshal(body, &changePass)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, "Invalid Body")
		return
	}

	result, err := authenticatedb.ChangePassDB(changePass.Email, changePass)

	if err != nil {
		log.Println(err)
		handlers.ResponseData(w, http.StatusBadRequest, err.Error())
		return
	}

	if result == -1 {
		log.Println(err)
		handlers.ResponseData(w, http.StatusNotFound, fmt.Sprintf("user %v does not existed !", changePass.Email))
		return
	} else if result == 0 {
		log.Println(err)
		handlers.ResponseData(w, http.StatusInternalServerError, err)
		return
	} else {
		handlers.ResponseData(w, http.StatusOK, "Change password Successful")
		return
	}
}
