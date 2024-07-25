package authenticatedb

import (
	"errors"

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
			var userId int
			rows, err := db.Query("INSERT INTO users (email, password, full_name, role_id) VALUES ($1, $2, $3, 1) RETURNING user_id", input.Email, hashedPassword, input.FullName)

			if err != nil {
				return 0, err
			}

			for rows.Next() {
				err := rows.Scan(&userId)

				if err != nil {
					panic(err)
				}
			}

			_, err = db.Exec("INSERT INTO user_attachment(user_id, file_url) VALUES ($1,$2)", userId, nil)

			if err != nil {
				return 0, err
			}
			return 1, nil
		} else {
			return 0, errors.New("can not sign up on other roles")
		}
	} else {
		return -1, err
	}
}
