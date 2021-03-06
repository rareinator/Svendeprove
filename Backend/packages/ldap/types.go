package ldap

type UserRole string

const (
	Doctor   UserRole = "Doctor"
	Nurse    UserRole = "Nurse"
	Employee UserRole = "Employee"
)
