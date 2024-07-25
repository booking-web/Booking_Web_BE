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

	defer db.Close()

	go func() error {
		query := fmt.Sprintf("DELETE FROM user_attachment WHERE user_id = %v", userId)

		_, err = db.Exec(query)

		if err != nil {
			return err
		} else {
			return nil
		}
	}()

	query := fmt.Sprintf("DELETE FROM users WHERE user_id = %v", userId)

	_, err = db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
