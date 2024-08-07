package authenticatedb

import (
	"errors"
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/pkg"
)

func ChangePassDB(receiverMail string, input handlers.ForgotPass) (int, error) {
	db, err := db.ConnectPostgres()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	checkedUserExisted, err := pkg.CheckExistedEmail(receiverMail)

	if err != nil {
		return 0, err
	}

	if checkedUserExisted {
		if input.NewPassword == input.ConfirmPassword {
			newPassword, err := pkg.HashPassword(input.NewPassword)

			if err != nil {
				return 0, err
			}

			query := fmt.Sprintf("UPDATE users SET password = '%s' WHERE email = '%v'", newPassword, receiverMail)

			_, err = db.Exec(query)

			if err != nil {
				return 0, err
			}
			return 1, nil
		}
		return -1, errors.New("confirm password is not correct with new password")
	} else {
		return -1, err
	}
}
