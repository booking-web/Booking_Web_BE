package authenticate

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db/authenticatedb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
)

// @Summary Login
// @Description Returns Login Successful
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userLogin body handlers.Login true "Login"
// @Success 200 {object} types.ResponseLogin
// @Router /api/v1/login [post]
func LogInRoute(w http.ResponseWriter, r *http.Request) {
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

	var user handlers.Login

	err = json.Unmarshal(body, &user)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, "Bad Request")
		return
	}

	result, err := authenticatedb.LoginDb(user)

	if err != nil {
		switch err.Error() {
		case "error connect db":
			handlers.ResponseData(w, http.StatusInsufficientStorage, err.Error())
			return

		case "password error":
			handlers.ResponseData(w, http.StatusBadRequest, err.Error())
			return

		case "error gen token":
			handlers.ResponseData(w, http.StatusInternalServerError, err.Error())
			return
		default:
			handlers.ResponseData(w, http.StatusNotFound, err.Error())
			return
		}
	}
	handlers.ResponseData(w, http.StatusOK, result)
}
