package pkg

import (
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func CheckExistedEmail(email string) (bool, error) {
	db, err := db.ConnectPostgres()

	if err != nil {
		return false, err
	}
	defer db.Close()

	checkQuery := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email)

	rows, err := db.Query(checkQuery)

	if err != nil {
		return false, err
	}

	var checkEmail bool
	for rows.Next() {
		var list types.Users

		err := rows.Scan(&list.UserId, &list.Email, &list.FullName, &list.Password, &list.RoleId)

		if err != nil {
			return false, err
		}
		if list.Email == "" {
			checkEmail = false
		} else {
			checkEmail = true
		}
	}

	defer rows.Close()
	return checkEmail, nil
}
