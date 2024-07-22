package doctordb

import (
	"errors"
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/pkg"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func GetDoctorById(doctorId int) (types.HandlerDoctor, error) {
	db, err := db.ConnectPostgres()

	if err != nil {
		return types.HandlerDoctor{}, err
	}

	if doctorId == 0 {
		return types.HandlerDoctor{}, errors.New("doctor id is empty")
	}

	defer db.Close()

	query := fmt.Sprintf("SELECT "+
		"d.doctor_id, "+
		"d.doctor_name, "+
		"d.doctor_summary, "+
		"d.exp_year, "+
		"d.edu_location, "+
		"d.degree, "+
		"w.work_location_name,  "+
		"l.language_name, "+
		"d.description "+
		"FROM doctor d "+
		"LEFT JOIN doctor_profile dp ON dp.doctor_id = d.doctor_id "+
		"LEFT JOIN work_location w ON w.work_location_id = dp.work_location "+
		"LEFT JOIN language l ON l.language_id = dp.language "+
		"WHERE d.doctor_id = %v", doctorId)

	rows, err := db.Query(query)

	if err != nil {
		return types.HandlerDoctor{}, err
	}

	doctorList := []types.Doctor{}

	for rows.Next() {
		var list types.Doctor

		err := rows.Scan(
			&list.DoctorId,
			&list.DoctorName,
			&list.DoctorSum,
			&list.ExpYear,
			&list.EduLocation,
			&list.Degree,
			&list.WorkLocation,
			&list.Language,
			&list.Description,
		)

		if err != nil {
			return types.HandlerDoctor{}, errors.New("not found")
		}

		doctorList = append(doctorList, list)
	}
	defer rows.Close()

	result, err := pkg.ConvertSameDoctor(doctorList, db)

	if err != nil {
		return types.HandlerDoctor{}, err
	}

	if len(result) == 0 {
		return types.HandlerDoctor{}, errors.New("empty")
	} else {
		return result[0], nil
	}
}
