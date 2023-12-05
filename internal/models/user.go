package models

type User struct {
	ID       int
	Username string
	Email    string
	Password string // Store a hashed password in a real application
}
