package userdb

import (
	"fmt"
	"strconv"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func GetUserById(userId string) ([]types.ResponseUsers, error) {
	var userList []types.ResponseUsers

	db, err := db.ConnectPostgres()

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
		return userList, fmt.Errorf("users id %d not found", id)
	}

	defer db.Close()

	for rows.Next() {
		var list types.ResponseUsers

		err := rows.Scan(&list.UserId, &list.Email, &list.FullName, &list.FileURL, &list.Role)

		if err != nil {
			return userList, err
		}

		userList = append(userList, list)
	}

	defer rows.Close()

	return userList, nil
}
