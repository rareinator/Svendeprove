package authentication

import (
	context "context"
	"fmt"

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
	// salt, err := a.DB.GetPatientSalt(u.Username)
	// if err != nil {
	// 	return nil, err
	// }
	// h := sha256.New()
	// h.Write([]byte(u.Password + salt))
	// hashedPassword := h.Sum(nil)

	// basePassword := base64.StdEncoding.EncodeToString(hashedPassword)

	// fmt.Println(basePassword)

	// user, err := a.DB.LoginPatient(u.Username, basePassword)
	// if err != nil {
	// 	return nil, err
	// }

	// response := TokenResponse{
	// 	FullName: user.Name,
	// 	Role:     0,
	// 	Username: user.Username,
	// 	UserID:   user.PatientId,
	// }

	// return &response, nil
	return nil, nil
	//TODO: oauth fix

}

func (a *AuthenticationServer) LoginEmployee(ctx context.Context, u *User) (*TokenResponse, error) {
	role, err := a.Ldap.AuthenticateUser(u.Username, u.Password)
	if err != nil {
		fmt.Printf("Failed to authenticate user: %v", err)
		return nil, err
	}

	//TODO: get actual info from LDAP
	response := TokenResponse{
		FullName: "Morten Nissen",
		Role:     int32(role),
		Username: u.Username,
		UserID:   1,
	}

	return &response, nil
}

func (a *AuthenticationServer) ValidateToken(ctx context.Context, tr *TokenRequest) (*ValidatorResponse, error) {

	_, err := a.DB.GetToken(tr.Token)
	if err != nil {
		return &ValidatorResponse{
			Valid: false,
		}, err
	}

	return &ValidatorResponse{
		Valid: true,
	}, nil
}

func (a *AuthenticationServer) GetRelatedPatient(ctx context.Context, rpr *RelatedPatientRequest) (*RelatedPatient, error) {
	// var result RelatedPatient

	// switch rpr.Type {
	// case "DBJournalDocument":
	// 	document := mssql.DBJournalDocument{}
	// 	patientID, err := a.DB.GetPatientID(document.GetPatientIDQuery(), rpr.Id)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	result.PatientId = patientID
	// default:
	// 	return nil, fmt.Errorf("Could not find a valid type")
	// }

	// return &result, nil
	return nil, nil
}

func (a *AuthenticationServer) InsertToken(ctx context.Context, tr *TokenRequest) (*ValidatorResponse, error) {
	dbToken := mssql.DBToken{
		Token: tr.Token,
	}

	if err := a.DB.InsertToken(&dbToken); err != nil {
		return &ValidatorResponse{Valid: false}, err
	}

	return &ValidatorResponse{
		Valid: true,
	}, nil
}
