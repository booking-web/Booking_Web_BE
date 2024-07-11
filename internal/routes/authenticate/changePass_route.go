package authenticate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/billzayy/Booking_Web_BE/internal/db/authenticatedb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
)

// @Summary Change Password
// @Description Returns Change Password Successful
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userLogin body handlers.ForgotPass true "Login"
// @Param userId query int true "userId"
// @Success 200 {object} handlers.ResponseDataType
// @Router /api/v1/change-password [post]
func ChangePassRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	userId := r.URL.Query().Get("userId")

	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	convertedUserId, err := strconv.Atoi(userId)

	if err != nil {
		panic(err)
	}

	var changePass handlers.ForgotPass

	err = json.Unmarshal(body, &changePass)

	if err != nil {
		panic(err)
	}

	result, err := authenticatedb.ChangePassDB(convertedUserId, changePass)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, err.Error())
		return
	}

	if result == -1 {
		handlers.ResponseData(w, http.StatusNotFound, fmt.Sprintf("user %v does not existed !", convertedUserId))
		return
	} else if result == 0 {
		handlers.ResponseData(w, http.StatusInternalServerError, err)
		return
	} else {
		handlers.ResponseData(w, http.StatusOK, "Change password Successful")
		return
	}
}
