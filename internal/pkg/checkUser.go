package pkg

import (
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
)

func CheckUser(userId int) (bool, error) {
	db, err := db.ConnectPostgres()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT user_id FROM users WHERE user_id = %v", userId)

	rows, err := db.Query(query)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	var checkedUserId bool

	for rows.Next() {
		var checked int

		err := rows.Scan(&checked)

		if err != nil {
			return false, err
		}

		if checked != 0 {
			checkedUserId = true
		} else {
			checkedUserId = false
		}
	}

	return checkedUserId, nil
}
