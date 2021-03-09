package authentication

import (
	context "context"
	"crypto/sha256"
	"encoding/base64"
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

func (a *AuthenticationServer) LoginPatient(ctx context.Context, u *User) (*TokenResponse, error) {
	salt, err := a.DB.GetPatientSalt(u.Username)
	if err != nil {
		return nil, err
	}
	h := sha256.New()
	h.Write([]byte(u.Password + salt))
	hashedPassword := h.Sum(nil)

	basePassword := base64.StdEncoding.EncodeToString(hashedPassword)

	fmt.Println(basePassword)

	user, err := a.DB.LoginPatient(u.Username, basePassword)
	if err != nil {
		return nil, err
	}

	tokenID := uuid.New()

	dbToken := mssql.DBToken{
		Token:      tokenID.String(),
		PatientID:  user.PatientId,
		Username:   u.Username,
		IssuedAt:   time.Now(),
		ValidUntil: time.Now().Add(time.Minute * 15),
	}

	if err := a.DB.InsertToken(&dbToken); err != nil {
		return nil, err
	}

	return &TokenResponse{Token: dbToken.Token}, nil

}

func (a *AuthenticationServer) LoginEmployee(ctx context.Context, u *User) (*EmployeeTokenResponse, error) {

	role, err := a.Ldap.AuthenticateUser(u.Username, u.Password)
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

	return &EmployeeTokenResponse{Token: dbToken.Token, Role: int32(role)}, nil
}

func (a *AuthenticationServer) ValidateToken(ctx context.Context, tr *TokenRequest) (*ValidatorResponse, error) {

	dbToken, err := a.DB.GetToken(tr.Token)
	if err != nil {
		return &ValidatorResponse{
			Valid:     false,
			Role:      0,
			PatientID: 0,
		}, err
	}

	return &ValidatorResponse{
		Valid:     true,
		Role:      dbToken.Role,
		PatientID: dbToken.PatientID,
	}, nil
}
