package authentication

import (
	context "context"
	"fmt"

	"github.com/rareinator/Svendeprove/Backend/packages/ldap"
)

type AuthenticationServer struct {
	UnimplementedAuthenticationServiceServer
	Ldap          *ldap.LDAP
	ListenAddress string
}

func (a *AuthenticationServer) GetHealth(ctx context.Context, e *AEmpty) (*AHealth, error) {
	return &AHealth{Message: fmt.Sprintf("Authentication service is up and running on: %v 🚀", a.ListenAddress)}, nil
}

func (a *AuthenticationServer) LoginEmployee(ctx context.Context, u *User) (*TokenResponse, error) {

	role, err := a.Ldap.AuthenticateUser(u.Username, u.HashedPassword)
	if err != nil {
		fmt.Printf("Failed to authenticate user: %v", err)
		return nil, err
	}

	return &TokenResponse{Token: string(role)}, nil
}
