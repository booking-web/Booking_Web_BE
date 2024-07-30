package admindb

import (
	"fmt"

	"github.com/billzayy/Booking_Web_BE/internal/db"
	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func GetLanguageByDoctorId(doctorId int) ([]types.Language, error) {
	db, err := db.ConnectPostgres()

	if err != nil {
		return []types.Language{}, err
	}

	defer db.Close()

	query := fmt.Sprintf("SELECT l.language_id, l.language_name "+
		"FROM Language l "+
		"LEFT JOIN doctor_profile dp ON dp.language = l.language_id "+
		"WHERE dp.doctor_id = %v", doctorId)

	rows, err := db.Query(query)

	if err != nil {
		return []types.Language{}, err
	}

	var sliceLanguage []types.Language

	for rows.Next() {
		var language types.Language

		err := rows.Scan(&language.LanguageId, &language.LanguageName)

		if err != nil {
			return []types.Language{}, err
		}

		sliceLanguage = append(sliceLanguage, language)
	}

	defer rows.Close()

	return sliceLanguage, nil
}

func GetAllLanguage() {}
