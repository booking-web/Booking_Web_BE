package handlers

type RequestCreateUser struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}
