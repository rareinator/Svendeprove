package models

type UserRole int

const (
	Patient  UserRole = 0
	Doctor   UserRole = 1
	Nurse    UserRole = 2
	Employee UserRole = 3
)
