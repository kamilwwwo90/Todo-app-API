package structs

type User struct {
	UserID   int    `json:"user_id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
