package patient

import (
	context "context"
	"fmt"

	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
)

type PatientServer struct {
	UnimplementedPatientServiceServer
	DB            *mssql.MSSQL
	ListenAddress string
}

func (p *PatientServer) GetHealth(ctx context.Context, e *PEmpty) (*PHealth, error) {
	return &PHealth{Message: fmt.Sprintf("Patient service is up and running on: %v 🚀", p.ListenAddress)}, nil
}

func (p *PatientServer) CreatePatient(ctx context.Context, patient *Patient) (*Patient, error) {
	dbPatient := mssql.DBPatient{
		Name:       patient.Name,
		Address:    patient.Address,
		City:       patient.City,
		PostalCode: patient.PostalCode,
		Country:    patient.Country,
		SocialIdNr: patient.SocialIdNr,
		Username:   patient.Username,
		Password:   "",
		Salt:       "",
	}

	if err := p.DB.CreatePatient(&dbPatient); err != nil {
		return nil, err
	}

	return patient, nil
}

func (p *PatientServer) ReadPatient(ctx context.Context, pr *PRequest) (*Patient, error) {
	dbPatient, err := p.DB.GetPatient(pr.Id)
	if err != nil {
		return nil, err
	}

	result := Patient{
		Name:       dbPatient.Name,
		Address:    dbPatient.Address,
		City:       dbPatient.City,
		PostalCode: dbPatient.PostalCode,
		Country:    dbPatient.Country,
		SocialIdNr: dbPatient.SocialIdNr,
		Username:   dbPatient.Username,
	}

	return &result, nil
}