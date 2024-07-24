package userdb

import (
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
)

func DeleteUserDB(userId int) error {
	db, err := db.ConnectPostgres()

	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM users WHERE user_id = %v", userId)

	_, err = db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
