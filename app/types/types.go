package types

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegistration struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       int
	Email    string
	Password string
}
