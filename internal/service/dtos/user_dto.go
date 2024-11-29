package dtos

type RegisterForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
