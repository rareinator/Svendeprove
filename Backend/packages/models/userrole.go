package models

type UserRole string

const (
	Patient UserRole = "Patient"
	Doctor  UserRole = "Doctor"
	Nurse   UserRole = "Nurse"
	Office  UserRole = "Office"
)
