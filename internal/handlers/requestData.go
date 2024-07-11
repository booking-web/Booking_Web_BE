package handlers

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForgotPass struct {
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
