package db

import (
	"fmt"
	"strconv"

	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func GetUserById(userId string) ([]types.User, error) {
	var userList []types.User

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
		var list types.User

		err := rows.Scan(&list.UserId, &list.Email, &list.FullName, &list.Password, &list.RoleId)

		if err != nil {
			return userList, err
		}

		userList = append(userList, list)
	}

	defer rows.Close()

	return userList, nil
}

func CreateUser(input types.User) (int, error) {
	db, err := ConnectPostgres()

	if err != nil {
		return 0, err
	}

	checkQuery := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", input.Email)
	rows, err := db.Query(checkQuery)

	if err != nil {
		return 0, nil
	}

	var checkEmail string
	for rows.Next() {
		var list types.User

		err := rows.Scan(&list.UserId, &list.Email, &list.FullName, &list.Password, &list.RoleId)

		if err != nil {
			return 0, err
		}
		checkEmail = list.Email
	}

	if checkEmail == "" {
		_, err = db.Exec("INSERT INTO users (email, password, full_name, role_id) VALUES ($1, $2, $3, 1)", input.Email, input.Password, input.FullName)

		if err != nil {
			return 0, err
		}

		return 1, nil
	} else {
		return -1, nil
	}
}
