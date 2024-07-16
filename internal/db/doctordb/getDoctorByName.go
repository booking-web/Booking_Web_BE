package doctordb

import (
	"errors"
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/pkg"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func GetDoctorByName(doctorName string) (types.ResponseDoctor, error) {
	db, err := db.ConnectPostgres()

	if err != nil {
		return types.ResponseDoctor{}, err
	}

	if doctorName == "" {
		return types.ResponseDoctor{}, errors.New("doctor name is empty")
	}

	defer db.Close()

	query := fmt.Sprintf("SELECT "+
		"d.doctor_name, d.doctor_summary, d.exp_year, cl.clinic_name, d.edu_location, d.degree, "+
		"w.work_location_name, l.language_name, d.description "+
		"FROM doctor_profile d "+
		"LEFT JOIN work_location w ON w.work_location_id = d.work_location "+
		"LEFT JOIN language l ON l.language_id = d.language "+
		"LEFT JOIN clinic cl ON cl.clinic_id = d.clinic_id "+
		"WHERE d.doctor_name = '%s'", doctorName)

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("Hello")
		return types.ResponseDoctor{}, err
	}

	doctorList := []types.Doctor{}

	for rows.Next() {
		var list types.Doctor

		err := rows.Scan(
			&list.DoctorName,
			&list.DoctorSum,
			&list.ExpYear,
			&list.ClinicName,
			&list.EduLocation,
			&list.Degree,
			&list.WorkLocation,
			&list.Language,
			&list.Description,
		)

		if err != nil {
			return types.ResponseDoctor{}, errors.New("not found")
		}

		doctorList = append(doctorList, list)
	}
	defer rows.Close()

	result, err := pkg.ConvertSameDoctor(doctorList)

	if err != nil {
		return types.ResponseDoctor{}, err
	}

	if len(result) == 0 {
		return types.ResponseDoctor{}, errors.New("empty")
	} else {
		return result[0], nil
	}
}
