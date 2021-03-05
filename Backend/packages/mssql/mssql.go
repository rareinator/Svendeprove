package mssql

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

//MSSQL is the main connector used to talk to the microsoft SQL database
type MSSQL struct {
	db *gorm.DB
}

func NewConnection(dsn string) (MSSQL, error) {
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return *&MSSQL{}, err
	}

	mssql := &MSSQL{
		db: db,
	}
	return *mssql, nil
}

func (m *MSSQL) GetJournal(id int32) (DBJournal, error) {
	var journal DBJournal
	m.db.First(&journal, 1)

	return journal, nil
}
