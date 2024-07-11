package authenticatedb

import (
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/pkg"
	"github.com/billzayy/Booking_Web_BE/internal/types"

	jwtauthen "github.com/billzayy/Booking_Web_BE/internal/pkg/jwtAuthen"
)

func LoginDb(input handlers.Login) (types.ResponseLogin, error) {
	db, err := db.ConnectPostgres()

	if err != nil {
		return types.ResponseLogin{}, err
	}

	defer db.Close()

	query := fmt.Sprintf("SELECT user_id, password FROM users WHERE email = '%s'", input.Email)

	rows, err := db.Query(query)

	if err != nil {
		return types.ResponseLogin{}, fmt.Errorf("error connect db")
	}
	defer rows.Close()

	var login types.ResponseLogin

	for rows.Next() {
		var userId int
		var password string

		err := rows.Scan(&userId, &password)

		if err != nil {
			return types.ResponseLogin{}, err
		}

		checkPass, err := pkg.ValidatePassword(input.Password, []byte(password))

		if err != nil {
			return types.ResponseLogin{}, fmt.Errorf("password is incorrect")
		}

		if checkPass {
			login.UserId = userId
		}
	}

	if login.UserId == 0 {
		return login, fmt.Errorf("email or password is not found")
	}

	result, err := jwtauthen.GenerateToken(string(login.UserId))

	if err != nil {
		return types.ResponseLogin{}, fmt.Errorf("error gen token")
	}

	login.Token = result

	return login, nil
}
