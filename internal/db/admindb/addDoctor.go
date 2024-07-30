package admindb

import (
	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/db/common"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func AddDoctorDB(input types.HandlerDoctor) error {
	db, err := db.ConnectPostgres()

	if err != nil {
		return err
	}

	defer db.Close()

	row, err := db.Query("INSERT INTO doctor(doctor_name, doctor_summary, exp_year, edu_location, degree, description, file_url, clinic_id) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING doctor_id",
		input.DoctorName, input.DoctorSum,
		input.ExpYear, input.EduLocation,
		input.Degree, input.Description, input.FileURL, common.ClinicId(input.ClinicName, db))

	if err != nil {
		return err
	}

	var doctorId int

	for row.Next() {
		err := row.Scan(&doctorId)

		if err != nil {
			return err
		}
	}

	return common.AddDoctorProfile(db, doctorId, input.WorkLocation, input.Language)
}
