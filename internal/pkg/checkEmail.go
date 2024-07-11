package pkg

import (
	"database/sql"
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func CheckExistedEmail(db *sql.DB, email string) (bool, error) {
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
