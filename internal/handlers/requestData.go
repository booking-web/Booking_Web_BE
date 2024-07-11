package handlers

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForgotPass struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
