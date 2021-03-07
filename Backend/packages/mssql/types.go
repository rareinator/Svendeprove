package mssql

import "time"

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
	Username   string    `gorm:"column:Username"`
	IssuedAt   time.Time `gorm:"column:IssuedAt"`
	ValidUntil time.Time `gorm:"column:ValidUntil"`
}

func (DBToken) TableName() string {
	return "Tokens"
}
