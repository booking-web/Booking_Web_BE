package types

type Users struct {
	UserId   int    `json:"userId"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
	RoleId   int    `json:"roleId"`
}

type UserAttachment struct {
	UserId  int    `json:"userId"`
	FileURL string `json:"fileUrl"`
}

type ResponseUsers struct {
	UserId   int    `json:"userId"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	FileURL  string `json:"fileUrl"`
	Role     string `json:"role"`
}

type ResponseLogin struct {
	UserId int    `json:"userId"`
	Token  string `json:"token"`
}
