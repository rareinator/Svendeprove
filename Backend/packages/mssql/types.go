package mssql

import "time"

type DBJournal struct {
	JourndId     int
	CreationTime time.Time
	Intro        string
	PatientId    int
	CreatedBy    int
}
