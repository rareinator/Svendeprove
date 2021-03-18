package patient

import (
	context "context"
	"fmt"
	"time"

	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
	. "github.com/rareinator/Svendeprove/Backend/packages/protocol"
)

type PatientServer struct {
	UnimplementedPatientServiceServer
	DB            *mssql.MSSQL
	ListenAddress string
}

func (p *PatientServer) GetHealth(ctx context.Context, e *Empty) (*Health, error) {
	return &Health{Message: fmt.Sprintf("Patient service is up and running on: %v 🚀", p.ListenAddress)}, nil
}

func (p *PatientServer) GetDiagnose(ctx context.Context, pr *Request) (*Diagnose, error) {
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

func (p *PatientServer) GetDiagnoses(ctx context.Context, e *Empty) (*Diagnoses, error) {
	diagnoses := Diagnoses{
		Diagnoses: make([]*Diagnose, 0),
	}

	dbDiagnoses, err := p.DB.GetDiagnoses()
	if err != nil {
		return nil, err
	}

	for _, dbDiagnose := range dbDiagnoses {
		diagnose := Diagnose{
			DiagnoseId:  dbDiagnose.DiagnoseId,
			Description: dbDiagnose.Description,
		}

		diagnoses.Diagnoses = append(diagnoses.Diagnoses, &diagnose)
	}

	return &diagnoses, nil
}

func (p *PatientServer) GetSymptom(ctx context.Context, pr *Request) (*Symptom, error) {
	dbSymptom, err := p.DB.GetSymptom(pr.Id)
	if err != nil {
		return nil, err
	}

	symptom := Symptom{
		SymptomId:   dbSymptom.SymptomId,
		Description: dbSymptom.Description,
	}

	return &symptom, nil
}

func (p *PatientServer) GetSymptoms(ctx context.Context, e *Empty) (*Symptoms, error) {
	symptoms := Symptoms{
		Symptoms: make([]*Symptom, 0),
	}

	dbSymptoms, err := p.DB.GetSymptoms()
	if err != nil {
		return nil, err
	}

	for _, dbSymptom := range dbSymptoms {
		symptom := Symptom{
			SymptomId:   dbSymptom.SymptomId,
			Description: dbSymptom.Description,
		}

		symptoms.Symptoms = append(symptoms.Symptoms, &symptom)
	}

	return &symptoms, nil
}

func (p *PatientServer) CreatePatientDiagnose(ctx context.Context, patientDiagnose *PatientDiagnose) (*PatientDiagnose, error) {
	dbPatientDiagnose := mssql.DBPatientDiagnose{
		Patient:      patientDiagnose.Patient,
		DiagnoseId:   patientDiagnose.Diagnose.DiagnoseId,
		CreationTime: time.Now(),
	}

	patientDiagnose.CreationTime = dbPatientDiagnose.CreationTime.Format("02/01/2006 15:04:05")

	if err := p.DB.CreatePatientDiagnose(&dbPatientDiagnose); err != nil {
		return nil, err
	}

	patientDiagnose.PatientDiagnoseId = dbPatientDiagnose.PatientDiagnoseId

	return patientDiagnose, nil

}

func (p *PatientServer) GetPatientDiagnoses(ctx context.Context, pr *Request) (*PatientDiagnoses, error) {
	patientDiagnoses := PatientDiagnoses{
		PatientDiagnoses: make([]*PatientDiagnose, 0),
	}

	dbPatientDiagnoses, err := p.DB.GetPatientDiagnoses(pr.UserId)
	if err != nil {
		return nil, err
	}

	for _, dbPatientDiagnose := range dbPatientDiagnoses {
		patientDiagnose := PatientDiagnose{
			PatientDiagnoseId: dbPatientDiagnose.PatientDiagnoseId,
			Patient:           dbPatientDiagnose.Patient,
			DiagnoseId:        dbPatientDiagnose.DiagnoseId,
			CreationTime:      dbPatientDiagnose.CreationTime.Format("02/01/2006 15:04:05"),
			Diagnose: &Diagnose{
				DiagnoseId:  dbPatientDiagnose.Diagnose.DiagnoseId,
				Description: dbPatientDiagnose.Diagnose.Description,
			},
			Symptoms: []*Symptom{},
		}

		dbSymptoms, err := p.DB.GetPatientDiagnoseSymptoms(patientDiagnose.PatientDiagnoseId)
		if err != nil {
			return nil, err
		}

		for _, dbSymptom := range dbSymptoms {
			patientDiagnose.Symptoms = append(patientDiagnose.Symptoms, &Symptom{
				SymptomId:   dbSymptom.Symptom.SymptomId,
				Description: dbSymptom.Symptom.Description,
			})
		}

		patientDiagnoses.PatientDiagnoses = append(patientDiagnoses.PatientDiagnoses, &patientDiagnose)
	}

	return &patientDiagnoses, nil

}

func (p *PatientServer) GetPatientDiagnose(ctx context.Context, pr *Request) (*PatientDiagnose, error) {
	dbPatientDiagnose, err := p.DB.GetPatientDiagnose(pr.Id)
	if err != nil {
		return nil, err
	}

	patientDiagnose := PatientDiagnose{
		PatientDiagnoseId: dbPatientDiagnose.PatientDiagnoseId,
		Patient:           dbPatientDiagnose.Patient,
		DiagnoseId:        dbPatientDiagnose.DiagnoseId,
		CreationTime:      dbPatientDiagnose.CreationTime.Format("02/01/2006 15:04:05"),
	}

	return &patientDiagnose, nil
}

func (p *PatientServer) UpdatePatientDiagnose(ctx context.Context, pd *PatientDiagnose) (*PatientDiagnose, error) {
	dbPatientDiagnose := mssql.DBPatientDiagnose{
		PatientDiagnoseId: pd.PatientDiagnoseId,
		Patient:           pd.Patient,
		DiagnoseId:        pd.DiagnoseId,
	}

	if err := p.DB.UpdatePatientDiagnose(&dbPatientDiagnose); err != nil {
		return nil, err
	}

	return pd, nil
}

func (p *PatientServer) DeletePatientDiagnose(ctx context.Context, pr *Request) (*Status, error) {
	dbPatientDiagnose := mssql.DBPatientDiagnose{
		PatientDiagnoseId: pr.Id,
	}

	if err := p.DB.DeletePatientDiagnose(&dbPatientDiagnose); err != nil {
		return &Status{Success: false}, err
	}

	return &Status{Success: true}, nil
}

func (p *PatientServer) CreateDiagnoseSymptom(ctx context.Context, diagnoseSymptom *DiagnoseSymptom) (*DiagnoseSymptom, error) {
	dbDiagnoseSymptom := mssql.DBPatientDiagnoseSymptom{
		PatientDiagnoseId: diagnoseSymptom.PatientDiagnoseId,
		SymptomId:         diagnoseSymptom.SymptomId,
	}

	if err := p.DB.CreatePatientDiagnoseSymptom(&dbDiagnoseSymptom); err != nil {
		return nil, err
	}

	return diagnoseSymptom, nil
}

func (p *PatientServer) GetDiagnoseSymptoms(ctx context.Context, pr *Request) (*DiagnoseSymptoms, error) {
	diagnoseSymptoms := DiagnoseSymptoms{
		DiagnoseSymptoms: make([]*DiagnoseSymptom, 0),
	}

	dbDiagnoseSymptoms, err := p.DB.GetPatientDiagnoseSymptoms(pr.Id)
	if err != nil {
		return nil, err
	}

	for _, dbDiagnoseSymptom := range dbDiagnoseSymptoms {
		diagDescription, err := p.DB.GetSymptom(dbDiagnoseSymptom.SymptomId)
		if err != nil {
			return nil, err
		}

		diagnoseSymptom := DiagnoseSymptom{
			SymptomId:         dbDiagnoseSymptom.SymptomId,
			PatientDiagnoseId: dbDiagnoseSymptom.PatientDiagnoseId,
			Description:       diagDescription.Description,
		}

		diagnoseSymptoms.DiagnoseSymptoms = append(diagnoseSymptoms.DiagnoseSymptoms, &diagnoseSymptom)
	}

	return &diagnoseSymptoms, nil

}

func (p *PatientServer) UpdateDiagnoseSymptom(ctx context.Context, dsur *DiagnoseSymptomUpdateRequest) (*DiagnoseSymptom, error) {
	dbNewPatientDiagnoseSymptom := mssql.DBPatientDiagnoseSymptom{
		PatientDiagnoseId: dsur.New.PatientDiagnoseId,
		SymptomId:         dsur.New.SymptomId,
	}

	dbOldPatientDiagnoseSymptom := mssql.DBPatientDiagnoseSymptom{
		PatientDiagnoseId: dsur.Old.PatientDiagnoseId,
		SymptomId:         dsur.Old.SymptomId,
	}

	if err := p.DB.UpdatePatientDiagnoseSymptom(&dbOldPatientDiagnoseSymptom, &dbNewPatientDiagnoseSymptom); err != nil {
		return nil, err
	}

	return dsur.New, nil
}

func (p *PatientServer) DeleteDiagnoseSymptom(ctx context.Context, diagnoseSymptom *DiagnoseSymptom) (*Status, error) {
	dbDiagnoseSymptom := mssql.DBPatientDiagnoseSymptom{
		PatientDiagnoseId: diagnoseSymptom.PatientDiagnoseId,
		SymptomId:         diagnoseSymptom.SymptomId,
	}

	if err := p.DB.DeletePatientDiagnoseSymptom(&dbDiagnoseSymptom); err != nil {
		return &Status{Success: false}, err
	}

	return &Status{Success: true}, nil
}
