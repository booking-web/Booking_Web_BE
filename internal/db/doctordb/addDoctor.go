package doctordb

import (
	"errors"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/pkg"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func AddDoctor(input types.HandlerDoctor) error {
	db, err := db.ConnectPostgres()

	if err != nil {
		return err
	}

	defer db.Close()

	row, err := db.Query("INSERT INTO doctor(doctor_name, doctor_summary, exp_year, edu_location, degree, description, file_url, clinic_id) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING doctor_id",
		input.DoctorName, input.DoctorSum,
		input.ExpYear, input.EduLocation,
		input.Degree, input.Description, input.FileURL, pkg.ClinicId(input.ClinicName, db))

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

	workSlice := input.WorkLocation
	languageSlice := input.Language
	var resultErr error

	if len(workSlice) == 1 {
		for i := range input.Language {
			_, err := db.Exec("INSERT INTO doctor_profile(doctor_id, work_location, language) VALUES($1, $2,$3)", doctorId, pkg.WorkId(workSlice[0], db), pkg.LanguageId(languageSlice[i], db))

			if err != nil {
				resultErr = err
				break
			}
		}
	} else if len(languageSlice) == 1 {
		for i := range input.WorkLocation {
			_, err := db.Exec("INSERT INTO doctor_profile(doctor_id, work_location, language) VALUES($1, $2,$3)", doctorId, pkg.WorkId(workSlice[i], db), pkg.LanguageId(languageSlice[0], db))

			if err != nil {
				resultErr = err
				break
			}
		}
	} else if len(workSlice) == len(languageSlice) {
		for i := range input.WorkLocation {
			_, err := db.Exec("INSERT INTO doctor_profile(doctor_id, work_location, language) VALUES($1, $2,$3)", doctorId, pkg.WorkId(workSlice[i], db), pkg.LanguageId(languageSlice[i], db))

			if err != nil {
				resultErr = err
				break
			}
		}
	} else if len(workSlice) == 0 || len(languageSlice) == 0 {
		return errors.New("slice must not empty")
	} else {
		for i := range len(workSlice) {
			if i < len(workSlice)-1 {
				_, err := db.Exec("INSERT INTO doctor_profile(doctor_id, work_location, language) VALUES($1, $2,$3)", doctorId, pkg.WorkId(workSlice[i], db), pkg.LanguageId(languageSlice[0], db))

				if err != nil {
					resultErr = err
					break
				}
			} else if i == len(workSlice)-1 {
				_, err := db.Exec("INSERT INTO doctor_profile(doctor_id, work_location, language) VALUES($1, $2,$3)", doctorId, pkg.WorkId(workSlice[i], db), pkg.LanguageId(languageSlice[len(languageSlice)-1], db))

				if err != nil {
					resultErr = err
					break
				}
			}
		}
	}
	if resultErr != nil {
		return resultErr
	}
	return nil
}
