package mssql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//MSSQL is the main connector used to talk to the microsoft SQL database
type MSSQL struct {
	db *gorm.DB
}

func NewConnection(dsn string) (MSSQL, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return *&MSSQL{}, err
	}

	mssql := &MSSQL{
		db: db,
	}
	return *mssql, nil
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
	fmt.Printf("Getting token for: %v\n\r", tokenID)
	m.db.First(&token, "Token = ?", tokenID)

	if token.Username == "" {
		fmt.Println("huh")
		return nil, fmt.Errorf("Could not find a token with ID: %v", tokenID)
	}

	fmt.Printf("role: %d patientID: %d username: %v", token.Role, token.PatientID, token.Username)

	return &token, nil
}

func (m *MSSQL) GetJournal(id int32) (*DBJournal, error) {
	var journal DBJournal
	result := m.db.First(&journal, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &journal, nil
}

func (m *MSSQL) GetJournalsByPatient(id int32) ([]*DBJournal, error) {
	var journals []*DBJournal
	result := m.db.Find(&journals).Where("PatientId = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return journals, nil
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

func (m *MSSQL) CreateJournal(journal *DBJournal) error {
	result := m.db.Omit("JournalId").Create(journal)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) UpdateJournal(journal *DBJournal) error {
	result := m.db.Where("JournalId = ?", journal.JournalId).Omit("JournalId").Save(&journal)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) DeleteJournal(journal *DBJournal) error {
	result := m.db.Where("JournalId = ?", journal.JournalId).Delete(journal)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) DeleteJournalDocument(journalDocument *DBJournalDocument) error {
	result := m.db.Where("DocumentId = ?", journalDocument.DocumentId).Delete(journalDocument)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) UpdateJournalDocument(journalDocument *DBJournalDocument) error {
	var result *gorm.DB
	if journalDocument.DocumentStoreId == 0 {
		result = m.db.Where("DocumentId = ?", journalDocument.DocumentId).Omit("DocumentId", "DocumentStoreId").Save(&journalDocument)
	} else {
		result = m.db.Where("DocumentId = ?", journalDocument.DocumentId).Omit("DocumentId").Save(&journalDocument)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) CreateJournalDocument(journalDocument *DBJournalDocument) error {
	var result *gorm.DB
	if journalDocument.DocumentStoreId == 0 {
		result = m.db.Omit("DocumentId", "DocumentStoreId").Create(&journalDocument)
	} else {
		result = m.db.Omit("DocumentId", "DocumentStoreId").Create(&journalDocument)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) GetJournalDocumentsByJournal(journalID int32) ([]*DBJournalDocument, error) {
	var journalDocuments []*DBJournalDocument
	result := m.db.Where("JournalId = ?", journalID).Find(&journalDocuments)
	if result.Error != nil {
		return nil, result.Error
	}

	return journalDocuments, nil
}

func (m *MSSQL) GetJournalDocument(journalDocumentID int32) (*DBJournalDocument, error) {
	var journalDocument DBJournalDocument
	result := m.db.Where("DocumentId = ?", journalDocumentID).First(&journalDocument)
	if result.Error != nil {
		return nil, result.Error
	}

	return &journalDocument, nil
}

func (m *MSSQL) GetPatientID(query string, id int32) (int32, error) {
	var resultData struct {
		PatientId int32
	}

	fmt.Printf("\n\rquery: %v\n\r", query)

	result := m.db.Raw(query, id).Scan(&resultData)
	if result.Error != nil {
		return 0, result.Error
	}

	fmt.Printf("PatientID: %v\n\r", resultData.PatientId)

	if resultData.PatientId == 0 {
		return 0, fmt.Errorf("Could not find a related patient")
	}

	return int32(resultData.PatientId), nil
}

func (m *MSSQL) CreatePatient(patient *DBPatient) error {
	result := m.db.Omit("PatientId").Create(patient)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) GetPatient(id int32) (*DBPatient, error) {
	var patient DBPatient
	result := m.db.First(&patient, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &patient, nil
}
