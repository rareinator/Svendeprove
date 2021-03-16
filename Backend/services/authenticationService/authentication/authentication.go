package authentication

import (
	context "context"
	"fmt"

	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
)

type AuthenticationServer struct {
	UnimplementedAuthenticationServiceServer
	DB            *mssql.MSSQL
	ListenAddress string
}

func (a *AuthenticationServer) GetHealth(ctx context.Context, e *AEmpty) (*AHealth, error) {
	return &AHealth{Message: fmt.Sprintf("Authentication service is up and running on: %v 🚀", a.ListenAddress)}, nil
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
