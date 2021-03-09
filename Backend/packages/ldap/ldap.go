package ldap

import (
	"github.com/go-ldap/ldap/v3"
	"github.com/rareinator/Svendeprove/Backend/packages/models"
)

type LDAP struct {
	AdminUsername string
	AdminPassword string
	Conn          *ldap.Conn
}

func (l *LDAP) AuthenticateUser(username, password string) (models.UserRole, error) {
	if err := l.Conn.Bind(username, password); err != nil {
		return 0, err
	}

	//Should find the group, and then return that instead of hardcoded doctor
	return models.Doctor, nil
}