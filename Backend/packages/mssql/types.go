package mssql

import "time"

type DBJournalDocument struct {
	DocumentId      int32  `gorm:"column:DocumentId"`
	Message         string `gorm:"column:Message"`
	DocumentStoreId int32  `gorm:"column:DocumentStoreId"`
	JournalId       int32  `gorm:"column:Journalid"`
	DocumentType    string `gorm:"column:DocumentType"`
	CreatedBy       int32  `gorm:"column:CreatedBy"`
}

func (DBJournalDocument) TableName() string {
	return "Document"
}

func (DBJournalDocument) GetPatientIDQuery() string {
	return "SELECT j.PatientId FROM Document AS d INNER JOIN Journal AS j on j.JournalId = d.JournalId WHERE d.DocumentId=?"
}

type DBJournal struct {
	JournalId    int32     `gorm:"column:JournalId"`
	CreationTime time.Time `gorm:"column:CreationTime"`
	Intro        string    `gorm:"column:Intro"`
	PatientId    int32     `gorm:"column:PatientId"`
	CreatedBy    int32     `gorm:"column:CreatedBy"`
}

func (DBJournal) TableName() string {
	return "Journal"
}

type DBToken struct {
	Token      string    `gorm:"column:Token"`
	Role       int32     `gorm:"column:Role"`
	PatientID  int32     `gorm:"column:PatientId"`
	Username   string    `gorm:"column:Username"`
	IssuedAt   time.Time `gorm:"column:IssuedAt"`
	ValidUntil time.Time `gorm:"column:ValidUntil"`
}

func (DBToken) TableName() string {
	return "Tokens"
}

type DBPatient struct {
	PatientId  int32  `gorm:"column:PatientId"`
	Name       string `gorm:"column:Name"`
	Address    string `gorm:"column:Address"`
	City       string `gorm:"column:City"`
	PostalCode string `gorm:"column:PostalCode"`
	Country    string `gorm:"column:Country"`
	SocialIdNr string `gorm:"column:SocialIdNr"`
	Username   string `gorm:"column:Username"`
	Password   string `gorm:"column:Password"`
	Salt       string `gorm:"column:Salt"`
}

func (DBPatient) TableName() string {
	return "Patient"
}