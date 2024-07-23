package admindb

import (
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
)

func DeleteDoctorDB(doctorId int) error {
	db, err := db.ConnectPostgres()

	if err != nil {
		return err
	}

	defer db.Close()

	query := fmt.Sprintf("DELETE FROM doctor WHERE doctor_id = %v", doctorId)

	go func() {
		query := fmt.Sprintf("DELETE FROM doctor_profile WHERE doctor_id = %v", doctorId)

		_, err = db.Exec(query)

		if err != nil {
			panic(err)
		}
	}()

	_, err = db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
