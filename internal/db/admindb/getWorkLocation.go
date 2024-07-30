package admindb

import (
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func GetWorkLocationByDoctorId(doctorId int) ([]types.WorkLocation, error) {
	db, err := db.ConnectPostgres()

	if err != nil {
		return []types.WorkLocation{}, err
	}

	defer db.Close()

	query := fmt.Sprintf("SELECT w.work_location_id, w.work_location_name "+
		"FROM work_location w "+
		"LEFT JOIN doctor_profile dp ON dp.language = w.work_location_id "+
		"WHERE dp.doctor_id = %v", doctorId)

	rows, err := db.Query(query)

	if err != nil {
		return []types.WorkLocation{}, err
	}

	var sliceHospital []types.WorkLocation

	for rows.Next() {
		var workLocation types.WorkLocation

		err := rows.Scan(&workLocation.LocationId, &workLocation.LocationName)

		if err != nil {
			return []types.WorkLocation{}, err
		}

		sliceHospital = append(sliceHospital, workLocation)
	}

	defer rows.Close()

	return sliceHospital, nil
}

// func GetAllHospital(){}
