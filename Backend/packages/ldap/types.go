package ldap

type UserRole int32

const (
	Doctor   UserRole = 1
	Nurse    UserRole = 2
	Employee UserRole = 3
)
