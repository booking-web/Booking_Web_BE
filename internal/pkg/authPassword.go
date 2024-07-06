package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

const salt = bcrypt.DefaultCost

func HashPassword(password string) (string, error) {
	convertPass := []byte(password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(convertPass, salt)
	if err != nil {
		return "", err
	}
	// fmt.Println("Hashed Password:", string(hashedPassword))

	return string(hashedPassword), nil
}

func ValidatePassword(password string, hashedPassword []byte) (bool, error) {
	convertPass := []byte(password)

	// hashedPassword, err := bcrypt.GenerateFromPassword(convertPass, salt)
	// if err != nil {
	// 	return false, err
	// }

	// To compare the password with the hash
	err := bcrypt.CompareHashAndPassword(hashedPassword, convertPass)

	if err != nil {
		return false, err
	}
	return true, nil
}
