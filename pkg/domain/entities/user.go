package entities

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func NewUser(username, password, role string) *User {
	return &User{
		Username: username,
		Password: password,
		Role:     role,
	}
}

func NewFakeUser() User {
	return User{
		Username: "username",
		Password: "password",
		Role:     "admin",
	}
}
