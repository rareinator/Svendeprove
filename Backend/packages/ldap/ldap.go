package ldap

import (
	"github.com/go-ldap/ldap/v3"
	"github.com/rareinator/Svendeprove/Backend/packages/models"
)

type LDAP struct {
	AdminUsername string
	AdminPassword string
	Uri           string
}

func (l *LDAP) NewConnection() (*ldap.Conn, error) {
	lConn, err := ldap.DialURL(l.Uri)
	if err != nil {
		return nil, err
	}

	return lConn, nil
}

func (l *LDAP) AuthenticateUser(username, password string) (models.UserRole, error) {
	lConn, err := l.NewConnection()
	if err != nil {
		return 0, err
	}
	defer lConn.Close()

	if err := lConn.Bind(username, password); err != nil {
		return 0, err
	}

	//Should find the group, and then return that instead of hardcoded doctor
	return models.Doctor, nil
}
