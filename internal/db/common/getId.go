package common

import (
	"database/sql"
	"fmt"
)

func ClinicId(clinicName string, db *sql.DB) int {
	var clinicId int

	query := fmt.Sprintf("SELECT clinic_id FROM clinic WHERE clinic_name = '%s'", clinicName)

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&clinicId)

		if err != nil {
			panic(err)
		}
	}

	return clinicId
}

func WorkId(name string, db *sql.DB) int {
	var resultId int

	query := fmt.Sprintf("SELECT work_location_id FROM work_location WHERE work_location_name = '%s'", name)

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&resultId)

		if err != nil {
			panic(err)
		}
	}

	return resultId
}

func LanguageId(name string, db *sql.DB) int {
	var resultId int

	query := fmt.Sprintf("SELECT language_id FROM language WHERE language_name = '%s'", name)

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&resultId)

		if err != nil {
			panic(err)
		}
	}

	return resultId
}
