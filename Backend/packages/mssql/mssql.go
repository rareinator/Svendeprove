package mssql

import (
	"fmt"

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

func (m *MSSQL) LoginPatient(username string, password string) (*DBPatient, error) {
	var patient DBPatient
	m.db.First(&patient, "Username = ? AND Password = ?", username, password)
	if patient.Username == "" {
		return nil, fmt.Errorf("Could not login")
	}

	return &patient, nil
}

func (m *MSSQL) GetPatientSalt(username string) (string, error) {
	var patient DBPatient
	m.db.First(&patient, "Username = ?", username)
	if patient.Salt == "" {
		return "", fmt.Errorf("Could not find user")
	}

	return patient.Salt, nil
}

func (m *MSSQL) InsertToken(token *DBToken) error {
	if token.Role == 0 {
		m.db.Exec("INSERT INTO Tokens (Token,PatientId,Username,IssuedAt,ValidUntil) VALUES (?,?,?,?,?)",
			token.Token,
			token.PatientID,
			token.Username,
			token.IssuedAt,
			token.ValidUntil)
	} else {
		m.db.Exec("INSERT INTO Tokens (Token,Role,Username,IssuedAt,ValidUntil) VALUES (?,?,?,?,?)",
			token.Token,
			token.Role,
			token.Username,
			token.IssuedAt,
			token.ValidUntil)
	}

	return nil
}

func (m *MSSQL) GetToken(tokenID string) (*DBToken, error) {
	var token DBToken
	m.db.First(&token, "Token = ?", tokenID)
	if token.Username == "" {
		return nil, fmt.Errorf("Could not find a token with ID: %v", tokenID)
	}

	return &token, nil
}
