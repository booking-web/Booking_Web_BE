package common

import (
	"database/sql"
	"errors"
)

func AddDoctorProfile(db *sql.DB, doctorId int, workLocation []string, language []string) error {

	var resultErr error

	sliceHospital := len(workLocation)
	sliceLanguage := len(language)

	addQuery := "INSERT INTO doctor_profile(doctor_id, work_location, language) VALUES ($1,$2,$3)"

	if sliceHospital == 1 {
		for _, v := range language {
			_, err := db.Exec(addQuery, doctorId, WorkId(workLocation[0], db), LanguageId(v, db))

			if err != nil {
				resultErr = err
				break
			}
		}
	} else if sliceLanguage == 1 {
		for _, v := range workLocation {
			_, err := db.Exec(addQuery, doctorId, WorkId(v, db), LanguageId(language[0], db))

			if err != nil {
				resultErr = err
				break
			}
		}
	} else if sliceHospital == sliceLanguage {
		for i := range workLocation {
			_, err := db.Exec(addQuery, doctorId, WorkId(workLocation[i], db), LanguageId(language[i], db))

			if err != nil {
				resultErr = err
				break
			}
		}
	} else if sliceLanguage == 2 {
		for i := range sliceHospital {
			if i < sliceHospital-1 {
				_, err := db.Exec(addQuery, doctorId, WorkId(workLocation[i], db), LanguageId(language[0], db))

				if err != nil {
					break
				}
			} else if i == sliceHospital-1 {
				_, err := db.Exec(addQuery, doctorId, WorkId(workLocation[i], db), LanguageId(language[len(language)-1], db))

				if err != nil {
					break
				}
			}
		}
		return nil
	} else {
		resultErr = errors.New("doctor can not sign more than 2 languages")
	}

	if resultErr != nil {
		return resultErr
	} else {
		return nil
	}
}
