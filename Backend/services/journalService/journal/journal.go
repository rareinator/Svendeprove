package journal

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
)

type JournalServer struct {
	UnimplementedJournalServiceServer
	DB            *mssql.MSSQL
	ListenAddress string
}

func (j *JournalServer) GetJournal(ctx context.Context, journal *JournalRequest) (*Journal, error) {
	log.Printf("Received journal from client: %v", journal.JournalId)

	dbJournal, err := j.DB.GetJournal(journal.JournalId)
	if err != nil {
		return nil, err
	}

	result := &Journal{
		JournalId:    dbJournal.JournalId,
		CreationTime: dbJournal.CreationTime.Format("02/01/2006 15:04:05"),
		Intro:        dbJournal.Intro,
		PatientId:    dbJournal.PatientId,
		CreatedBy:    dbJournal.CreatedBy,
	}

	return result, nil
}

func (j *JournalServer) GetHealth(ctx context.Context, e *JEmpty) (*JHealth, error) {
	return &JHealth{Message: fmt.Sprintf("Journal service is up and running on: %v 🚀", j.ListenAddress)}, nil
}

func (j *JournalServer) GetJournalsByPatient(ctx context.Context, pr *PatientRequest) (*Journals, error) {
	journals := &Journals{}
	journals.Journals = make([]*Journal, 0)

	dbJournals, err := j.DB.GetJournalsByPatient(pr.PatientId)
	if err != nil {
		return nil, err
	}

	for _, dbJournal := range dbJournals {
		journal := &Journal{
			JournalId:    dbJournal.JournalId,
			CreationTime: dbJournal.CreationTime.Format("02/01/2006 15:04:05"),
			Intro:        dbJournal.Intro,
			PatientId:    dbJournal.PatientId,
			CreatedBy:    dbJournal.CreatedBy,
		}

		journals.Journals = append(journals.Journals, journal)
	}

	return journals, nil

}

func (j *JournalServer) CreateJournal(ctx context.Context, journal *Journal) (*Journal, error) {
	parsedTime, err := time.Parse("02/01/2006 15:04:05", journal.CreationTime)
	if err != nil {
		return nil, err
	}

	dbJournal := mssql.DBJournal{
		CreationTime: parsedTime,
		Intro:        journal.Intro,
		PatientId:    journal.PatientId,
		CreatedBy:    journal.CreatedBy,
	}

	if err := j.DB.CreateJournal(&dbJournal); err != nil {
		return nil, err
	}

	return journal, nil

}

func (j *JournalServer) UpdateJournal(ctx context.Context, journal *Journal) (*Journal, error) {
	parsedtime, err := time.Parse("02/01/2006 15:04:05", journal.CreationTime)
	if err != nil {
		return nil, err
	}

	dbJournal := mssql.DBJournal{
		JournalId:    journal.JournalId,
		CreationTime: parsedtime,
		Intro:        journal.Intro,
		PatientId:    journal.PatientId,
		CreatedBy:    journal.CreatedBy,
	}

	if err := j.DB.UpdateJournal(&dbJournal); err != nil {
		return nil, err
	}

	return journal, nil
}

func (j *JournalServer) DeleteJournal(ctx context.Context, jr *JournalRequest) (*Status, error) {
	dbJournal := mssql.DBJournal{
		JournalId: jr.JournalId,
	}

	if err := j.DB.DeleteJournal(&dbJournal); err != nil {
		return &Status{Success: false}, err
	}

	return &Status{Success: true}, nil
}

func (j *JournalServer) DeleteJournalDocument(ctx context.Context, jdr *JournalDocumentRequest) (*Status, error) {
	dbJournalDocument := mssql.DBJournalDocument{
		DocumentId: jdr.JournalDocumentId,
	}

	if err := j.DB.DeleteJournalDocument(&dbJournalDocument); err != nil {
		return &Status{Success: false}, err
	}

	return &Status{Success: true}, nil
}

func (j *JournalServer) UpdateJournalDocument(ctx context.Context, jdr *JournalDocument) (*JournalDocument, error) {
	dbJournalDocument := mssql.DBJournalDocument{
		DocumentId:      jdr.DocumentId,
		Message:         jdr.Message,
		DocumentStoreId: jdr.DocumentStoreId,
		JournalId:       jdr.JournalId,
		DocumentType:    jdr.DocumentType,
		CreatedBy:       jdr.CreatedBy,
	}

	if err := j.DB.UpdateJournalDocument(&dbJournalDocument); err != nil {
		return nil, err
	}

	return jdr, nil

}
