package common

import (
	"database/sql"
	"fmt"
)

func GetClinicByDoctor(doctorId int, db *sql.DB) string {
	query := fmt.Sprintf("SELECT cl.clinic_name FROM clinic cl "+
		"LEFT JOIN doctor d ON cl.clinic_id = d.clinic_id "+
		"WHERE d.doctor_id = %v", doctorId)

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	var clinicName string
	for rows.Next() {
		rows.Scan(&clinicName)
	}

	return clinicName
}
