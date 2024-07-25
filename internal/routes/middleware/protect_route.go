package middleware

import (
	"fmt"
	"net/http"

	middleware "github.com/billzayy/Booking_Web_BE/internal/pkg/middleware"
)

func ProtectedRoute(w http.ResponseWriter, r *http.Request) (bool, error) {
	w.Header().Set("Content-Type", "application/json")

	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		return false, fmt.Errorf("missing authorization header")
	}

	tokenString = tokenString[len("Bearer "):]

	_, err := middleware.VerifyToken(tokenString)

	if err != nil {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}
