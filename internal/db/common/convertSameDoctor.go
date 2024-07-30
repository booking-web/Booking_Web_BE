package common

import (
	"database/sql"
	"errors"

	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func ConvertSameDoctor(doctorList []types.Doctor, db *sql.DB) ([]types.HandlerDoctor, error) {
	resultSlice := []types.HandlerDoctor{}

	workList := []string{}
	languageList := []string{}

	if len(doctorList) == 0 || doctorList == nil {
		return resultSlice, errors.New("empty data")
	}

	sameList := doctorList[0]

	for _, v := range doctorList {
		workList = append(workList, v.WorkLocation)
		languageList = append(languageList, v.Language)
	}

	test := types.HandlerDoctor{
		DoctorId:     sameList.DoctorId,
		DoctorName:   sameList.DoctorName,
		DoctorSum:    sameList.DoctorSum,
		ExpYear:      sameList.ExpYear,
		ClinicName:   GetClinicByDoctor(sameList.DoctorId, db),
		EduLocation:  sameList.EduLocation,
		Degree:       sameList.Degree,
		WorkLocation: convertSameList(workList),
		Language:     convertSameList(languageList),
		Description:  sameList.Description,
	}

	resultSlice = append(resultSlice, test)

	return resultSlice, nil
}

func convertSameList(slice []string) []string {
	var result []string

	for i := range slice {
		if len(result) == 0 {
			result = append(result, slice[i])
		}

		if result[len(result)-1] != slice[i] {
			result = append(result, slice[i])
		}
	}
	return result
}
