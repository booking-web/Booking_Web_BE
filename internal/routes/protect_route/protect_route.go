package protectroute

import (
	"fmt"
	"net/http"

	jwtauthen "github.com/billzayy/Booking_Web_BE/internal/pkg/jwtAuthen"
)

func ProtectedRoute(w http.ResponseWriter, r *http.Request) (bool, error) {
	w.Header().Set("Content-Type", "application/json")

	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		return false, fmt.Errorf("missing authorization header")
	}

	tokenString = tokenString[len("Bearer "):]

	_, err := jwtauthen.VerifyToken(tokenString)

	if err != nil {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}
