package handlers

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForgotPass struct {
	Email           string `json:"email"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UpdateUserType struct {
	UserId   int    `json:"userId"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Image    string `json:"images"`
}
