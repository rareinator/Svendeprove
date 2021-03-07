package authentication

import (
	context "context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rareinator/Svendeprove/Backend/packages/ldap"
	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
)

type AuthenticationServer struct {
	UnimplementedAuthenticationServiceServer
	Ldap          *ldap.LDAP
	DB            *mssql.MSSQL
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

	tokenID := uuid.New()

	dbToken := mssql.DBToken{
		Token:      tokenID.String(),
		Role:       int32(role),
		Username:   u.Username,
		IssuedAt:   time.Now(),
		ValidUntil: time.Now().Add(time.Minute * 15),
	}

	a.DB.InsertToken(&dbToken)

	return &TokenResponse{Token: dbToken.Token}, nil
}
