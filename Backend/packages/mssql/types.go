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
