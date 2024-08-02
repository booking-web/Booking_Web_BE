package authenticate

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db/authenticatedb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

// @Summary Sign Up Account
// @Description Return Sign Up Successful
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userId body types.Users true "userId to Add"
// @Success 201 {object} handlers.ResponseDataType
// @Router /api/v1/sign-up [post]
func SignUpRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, "Bad Request")
		return
	}

	var user types.Users

	err = json.Unmarshal(body, &user)

	if err != nil {
		log.Println(err)
		handlers.ResponseData(w, http.StatusBadRequest, "Invalid Body")
		return
	}

	num, err := authenticatedb.SignUp(user)

	if err != nil {
		log.Println(err)
		handlers.ResponseData(w, http.StatusInternalServerError, err.Error())
		return
	}

	if num == 0 {
		log.Println(err)
		handlers.ResponseData(w, http.StatusBadRequest, err)
		return
	} else if num == 1 {
		handlers.ResponseData(w, http.StatusCreated, "Sign Up Successful")
		return
	} else {
		log.Println(err)
		handlers.ResponseData(w, http.StatusBadRequest, "Email is existed !")
		return
	}
}
