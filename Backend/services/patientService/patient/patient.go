package patient

import (
	context "context"
	"fmt"
	"time"

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

func (p *PatientServer) UpdatePatient(ctx context.Context, patient *Patient) (*Patient, error) {
	dbPatient := mssql.DBPatient{
		PatientId:  patient.PatientId,
		Name:       patient.Name,
		Address:    patient.Address,
		City:       patient.City,
		PostalCode: patient.PostalCode,
		Country:    patient.Country,
		SocialIdNr: patient.SocialIdNr,
		Username:   patient.Username,
	}

	if err := p.DB.UpdatePatient(&dbPatient); err != nil {
		return nil, err
	}

	return patient, nil
}

func (p *PatientServer) DeletePatient(ctx context.Context, pr *PRequest) (*PStatus, error) {
	dbPatient := mssql.DBPatient{
		PatientId: pr.Id,
	}

	if err := p.DB.DeletePatient(&dbPatient); err != nil {
		return &PStatus{
			Success: false,
		}, err
	}

	return &PStatus{Success: true}, nil
}

func (p *PatientServer) GetDiagnose(ctx context.Context, pr *PRequest) (*Diagnose, error) {
	dbDiagnose, err := p.DB.GetDiagnose(pr.Id)
	if err != nil {
		return nil, err
	}

	result := Diagnose{
		DiagnoseId:  dbDiagnose.DiagnoseId,
		Description: dbDiagnose.Description,
	}

	return &result, nil

}

func (p *PatientServer) CreatePatientDiagnose(ctx context.Context, patientDiagnose *PatientDiagnose) (*PatientDiagnose, error) {
	dbPatientDiagnose := mssql.DBPatientDiagnose{
		PatientDiagnoseId: patientDiagnose.PatientDiagnoseId,
		PatientId:         patientDiagnose.PatientId,
		SymptomId:         patientDiagnose.SymptomId,
		DiagnoseId:        patientDiagnose.DiagnoseId,
		CreationTime:      time.Now(),
	}

	if err := p.DB.CreatePatientDiagnose(&dbPatientDiagnose); err != nil {
		return nil, err
	}

	return patientDiagnose, nil

}
