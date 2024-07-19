package types

type Doctor struct {
	DoctorId     int    `json:"doctorId"`
	DoctorName   string `json:"doctorName"`
	DoctorSum    string `json:"doctorSum"`
	ExpYear      int    `json:"expYear"`
	ClinicName   string `json:"clinicName"`
	EduLocation  string `json:"eduLocation"`
	Degree       string `json:"degree"`
	WorkLocation string `json:"workLocation"`
	Language     string `json:"language"`
	Description  string `json:"description"`
}

type ResponseDoctor struct {
	DoctorId     int      `json:"doctorId"`
	DoctorName   string   `json:"doctorName"`
	DoctorSum    string   `json:"doctorSum"`
	ExpYear      int      `json:"expYear"`
	ClinicName   string   `json:"clinicName"`
	EduLocation  string   `json:"eduLocation"`
	Degree       string   `json:"degree"`
	WorkLocation []string `json:"workLocation"`
	Language     []string `json:"language"`
	Description  string   `json:"description"`
}

type ListDoctor struct {
	DoctorId     int      `json:"doctorId"`
	DoctorName   string   `json:"doctorName"`
	ClinicName   string   `json:"clinicName"`
	WorkLocation []string `json:"workLocation"`
	ExpYear      int      `json:"expYear"`
}
