package pkg

import (
	"errors"
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
)

func CheckPassword(userId int, password string) (bool, error) {
	db, err := db.ConnectPostgres()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT password FROM users WHERE user_id = %v", userId)

	rows, err := db.Query(query)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	var checkedPass bool

	for rows.Next() {
		var checked string

		err := rows.Scan(&checked)

		if err != nil {
			return false, err
		}

		hashPass, err := ValidatePassword(password, []byte(checked))

		if err != nil {
			return false, errors.New("password is incorrect")
		}

		if hashPass {
			checkedPass = true
		} else {
			checkedPass = false
		}
	}

	return checkedPass, nil
}
