package domain

// User represents a user in the application.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserRepository interface {
	GetUser() *User
}
