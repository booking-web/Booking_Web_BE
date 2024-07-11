package authenticatedb

import (
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/pkg"
)

func ChangePassDB(userId int, input handlers.ForgotPass) (int, error) {
	db, err := db.ConnectPostgres()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	checkedPassExisted, err := pkg.CheckPassword(userId, input.OldPassword)

	if err != nil {
		return 0, err
	}

	if checkedPassExisted {
		newPassword, err := pkg.HashPassword(input.NewPassword)

		if err != nil {
			return 0, err
		}

		query := fmt.Sprintf("UPDATE users SET password = '%s' WHERE user_id = %v", newPassword, userId)

		_, err = db.Exec(query)

		if err != nil {
			return 0, err
		}
		return 1, nil
	} else {
		return -1, err
	}

}
