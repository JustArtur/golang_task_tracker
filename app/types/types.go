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

type NotePayload struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Note struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
