package pkg

import (
	"errors"

	"github.com/billzayy/Booking_Web_BE/internal/types"
)

func ConvertSameDoctor(doctorList []types.Doctor) ([]types.ResponseDoctor, error) {
	resultSlice := []types.ResponseDoctor{}

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

	test := types.ResponseDoctor{
		DoctorName:   sameList.DoctorName,
		DoctorSum:    sameList.DoctorSum,
		ExpYear:      sameList.ExpYear,
		ClinicName:   sameList.ClinicName,
		EduLocation:  sameList.EduLocation,
		Degree:       sameList.Degree,
		WorkLocation: workList,
		Language:     languageList,
		Description:  sameList.Description,
	}

	resultSlice = append(resultSlice, test)

	return resultSlice, nil
}
