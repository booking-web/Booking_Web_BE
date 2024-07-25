package userdb

import (
	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
)

func UpdateUserDB(user handlers.UpdateUserType) error {
	db, err := db.ConnectPostgres()

	if err != nil {
		return err
	}

	defer db.Close()

	go func() error {
		_, err = db.Exec("UPDATE user_attachment SET file_url = $1 WHERE user_id = $2", user.Image, user.UserId)

		if err != nil {
			return err
		} else {
			return nil
		}
	}()

	_, err = db.Exec("UPDATE users SET "+
		"email = $1, full_name = $2 WHERE user_id = $3", user.Email, user.FullName, user.UserId)

	if err != nil {
		return err
	}

	return nil
}
