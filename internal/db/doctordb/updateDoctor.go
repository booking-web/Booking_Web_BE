package doctordb

import (
	"database/sql"
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/db/common"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func UpdateDoctorDB(doctor types.HandlerDoctor) error {
	db, err := db.ConnectPostgres()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = updateHospitalAndLanguage(db, doctor)

	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE doctor SET "+
		"doctor_name = $1, doctor_summary = $2,"+
		"exp_year = $3, edu_location = $4, "+
		"degree = $5, description = $6, "+
		"file_url = $7, clinic_id = $8 WHERE doctor_id = $9",
		doctor.DoctorName, doctor.DoctorSum,
		doctor.ExpYear, doctor.ExpYear,
		doctor.Degree, doctor.Description,
		doctor.FileURL, common.ClinicId(doctor.ClinicName, db), doctor.DoctorId)

	if err != nil {
		return err
	}

	return nil
}

func updateHospitalAndLanguage(db *sql.DB, doctor types.HandlerDoctor) error {

	query := fmt.Sprintf("DELETE FROM doctor_profile WHERE doctor_id = %v", doctor.DoctorId)

	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}

	return common.AddDoctorProfile(db, doctor.DoctorId, doctor.WorkLocation, doctor.Language)

}
