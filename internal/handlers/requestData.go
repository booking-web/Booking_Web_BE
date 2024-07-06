package handlers

type RequestCreateUser struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
