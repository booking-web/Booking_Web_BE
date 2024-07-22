package doctordb

import (
	"database/sql"
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func ListDoctor() ([]types.ListDoctor, error) {
	var doctorList []types.ListDoctor

	db, err := db.ConnectPostgres()

	if err != nil {
		return doctorList, err
	}

	defer db.Close()

	query := "SELECT d.doctor_id, d.doctor_name, " +
		"d.exp_year, " +
		"cl.clinic_name " +
		"FROM doctor d " +
		"LEFT JOIN clinic cl ON cl.clinic_id = d.clinic_id ORDER BY d.doctor_id ASC"

	rows, err := db.Query(query)

	if err != nil {
		return doctorList, err
	}

	for rows.Next() {
		var result types.ListDoctor

		err := rows.Scan(&result.DoctorId, &result.DoctorName, &result.ExpYear, &result.ClinicName)

		if err != nil {
			return doctorList, err
		}

		doctorList = append(doctorList, result)
	}

	defer rows.Close()

	for i := range doctorList {
		locationSlice, err := getSliceLocation(doctorList[i].DoctorId, db)

		if err != nil {
			return doctorList, err
		}

		doctorList[i].WorkLocation = locationSlice
	}
	return doctorList, nil
}

// This function is using for get a string slice of work_location_name
func getSliceLocation(doctorId int, db *sql.DB) ([]string, error) {
	locationList := []string{}

	query := fmt.Sprintf("SELECT w.work_location_name "+
		"FROM doctor d "+
		"LEFT JOIN doctor_profile dp ON dp.doctor_id = d.doctor_id "+
		"LEFT JOIN work_location w ON w.work_location_id = dp.work_location "+
		"WHERE d.doctor_id = %v", doctorId)

	rows, err := db.Query(query)

	if err != nil {
		return locationList, err
	}

	for rows.Next() {
		var locationName string

		err := rows.Scan(&locationName)

		if err != nil {
			return locationList, err
		}

		if len(locationList) == 0 { // If a slice is empty, add the value
			locationList = append(locationList, locationName)
		}
		if locationList[len(locationList)-1] != locationName { // If current value of slice doesn't equal with work location name
			locationList = append(locationList, locationName)
		}
	}

	return locationList, nil
}
