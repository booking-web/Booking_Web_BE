package types

type Doctor struct {
	DoctorId     int    `json:"doctorId"`
	DoctorName   string `json:"doctorName"`
	DoctorSum    string `json:"doctorSum"`
	ExpYear      int    `json:"expYear"`
	ClinicId     string `json:"clinicId"`
	EduLocation  string `json:"eduLocation"`
	Degree       string `json:"degree"`
	WorkLocation string `json:"workLocation"`
	Language     string `json:"language"`
	FileUrl      string `json:"fileUrl"`
	Description  string `json:"description"`
}

type HandlerDoctor struct {
	DoctorId     int      `json:"doctorId"`
	DoctorName   string   `json:"doctorName"`
	DoctorSum    string   `json:"doctorSum"`
	ExpYear      int      `json:"expYear"`
	ClinicName   string   `json:"clinicName"`
	EduLocation  string   `json:"eduLocation"`
	Degree       string   `json:"degree"`
	WorkLocation []string `json:"workLocation"`
	Language     []string `json:"language"`
	FileURL      string   `json:"fileUrl"`
	Description  string   `json:"description"`
}

type ListDoctor struct {
	DoctorId     int      `json:"doctorId"`
	DoctorName   string   `json:"doctorName"`
	ClinicName   string   `json:"clinicName"`
	WorkLocation []string `json:"workLocation"`
	FileURL      string   `json:"fileUrl"`
	ExpYear      int      `json:"expYear"`
}
