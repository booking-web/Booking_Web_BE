package authenticatedb

import (
	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/pkg"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func SignUp(input types.Users) (int, error) {
	db, err := db.ConnectPostgres()

	if err != nil {
		return 0, err
	}

	checkEmail, err := pkg.CheckExistedEmail(input.Email)

	if err != nil {
		return 0, err
	}

	defer db.Close()

	if !checkEmail {
		hashedPassword, err := pkg.HashPassword(input.Password)

		if err != nil {
			return 0, err
		}

		if input.RoleId == 0 || input.RoleId == 1 {
			_, err = db.Exec("INSERT INTO users (email, password, full_name, role_id) VALUES ($1, $2, $3, 1)", input.Email, hashedPassword, input.FullName)
		} else {
			_, err = db.Exec("INSERT INTO users (email, password, full_name, role_id) VALUES ($1, $2, $3, $4)", input.Email, hashedPassword, input.FullName, input.RoleId)
		}

		if err != nil {
			return 0, err
		}

		return 1, nil
	} else {
		return -1, err
	}
}
