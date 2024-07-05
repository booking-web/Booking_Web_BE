package types

type User struct {
	UserId   int    `json:"userId"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
	RoleId   int    `json:"roleId"`
}
