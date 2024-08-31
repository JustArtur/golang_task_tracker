package types

type UserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
