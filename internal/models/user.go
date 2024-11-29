package models

// Role: Student, Parent, Teacher
type User struct {
	ID           int
	Role         string
	Name         string
	PasswordHash string
}
