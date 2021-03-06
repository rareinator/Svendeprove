package journal

import (
	"context"
	"fmt"
	"log"

	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
)

type JournalServer struct {
	UnimplementedJournalServiceServer
	DB            mssql.MSSQL
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
