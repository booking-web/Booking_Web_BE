package db

import (
	"fmt"
	"strconv"

	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/pkg"
	jwtauthen "github.com/billzayy/Booking_Web_BE/internal/pkg/jwtAuthen"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func GetUserById(userId string) ([]types.Users, error) {
	var userList []types.Users

	db, err := ConnectPostgres()

	if err != nil {
		return userList, err
	}

	id, err := strconv.Atoi(userId)

	if err != nil {
		return userList, fmt.Errorf("input id is empty")
	}

	query := fmt.Sprintf("SELECT * FROM users WHERE user_id = %d", id)

	rows, err := db.Query(query)

	if err != nil {
		return userList, err
	}

	defer db.Close()

	for rows.Next() {
		var list types.Users

		err := rows.Scan(&list.UserId, &list.Email, &list.FullName, &list.Password, &list.RoleId)

		if err != nil {
			return userList, err
		}

		userList = append(userList, list)
	}

	defer rows.Close()

	return userList, nil
}

func CreateUser(input types.Users) (int, error) {
	db, err := ConnectPostgres()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	checkEmail, err := pkg.CheckExistedEmail(db, input.Email)

	if err != nil {
		return 0, err
	}

	if !checkEmail {
		hashedPassword, err := pkg.HashPassword(input.Password)

		if err != nil {
			return 0, err
		}

		_, err = db.Exec("INSERT INTO users (email, password, full_name, role_id) VALUES ($1, $2, $3, 1)", input.Email, hashedPassword, input.FullName)

		if err != nil {
			return 0, err
		}

		return 1, nil
	} else {
		return -1, nil
	}
}

func Login(input handlers.Login) (types.ResponseLogin, error) {
	db, err := ConnectPostgres()

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
			return types.ResponseLogin{}, fmt.Errorf("password error")
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
