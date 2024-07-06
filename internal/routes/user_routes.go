package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/types"
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

	checkAuth, err := ProtectedRoute(w, r)

	if err != nil {
		handlers.ResponseData(w, http.StatusUnauthorized, err)
	}

	if checkAuth {
		userId := r.URL.Query().Get("userId")

		data, err := db.GetUserById(userId)

		if err != nil {
			handlers.ResponseData(w, http.StatusNotFound, err.Error())
			return
		}

		handlers.ResponseData(w, http.StatusOK, data[0])
		return
	} else {
		handlers.ResponseData(w, http.StatusUnauthorized, err)
		return
	}
}

// @Summary Sign Up Account
// @Description Returns  Sign Up Successful
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userId body types.Users true "New User"
// @Success 200 {object} types.Users
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
		handlers.ResponseData(w, http.StatusBadRequest, "Invalid Body")
		return
	}

	num, err := db.CreateUser(user)

	if err != nil {
		handlers.ResponseData(w, http.StatusInternalServerError, err.Error())
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

// @Summary Login
// @Description Returns Login Successful
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userLogin body handlers.Login true "Login"
// @Success 200 {object} types.Users
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

	result, err := db.Login(user)

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
