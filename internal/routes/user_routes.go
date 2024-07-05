package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func GetUserByIdRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	userId := r.URL.Query().Get("id")

	data, err := db.GetUserById(userId)

	if err != nil {
		handlers.ResponseData(w, http.StatusNotFound, err.Error())
		return
	}

	handlers.ResponseData(w, http.StatusOK, data[0])
}

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

	var user types.User

	err = json.Unmarshal(body, &user)

	if err != nil {
		handlers.ResponseData(w, http.StatusBadRequest, "Invalid Body")
		return
	}

	num, err := db.CreateUser(user)

	if err != nil {
		handlers.ResponseData(w, http.StatusInternalServerError, err.Error())
		fmt.Println(err)
		return
	}

	if num == 0 {
		handlers.ResponseData(w, http.StatusBadRequest, err)
	} else if num == 1 {
		handlers.ResponseData(w, http.StatusOK, "Sign Up Successful")
	} else {
		handlers.ResponseData(w, http.StatusBadRequest, "Email is existed !")
	}
}
