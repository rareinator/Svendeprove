package mssql

import (
	"errors"
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

func (m *MSSQL) GetToken(tokenID string) (*DBToken, error) {
	var token DBToken
	result := m.db.First(&token, "Token = ?", tokenID)
	if result.Error != nil {
		return nil, result.Error
	}

	if token.Token != tokenID {
		return nil, fmt.Errorf("Could not find a token with ID: %v", tokenID)
	}

	return &token, nil
}

func (m *MSSQL) InsertToken(token *DBToken) error {
	result := m.db.Omit("IssuedAt", "ValidUntil").Create(token)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) GetJournal(id int32) (*DBJournal, error) {
	var journal DBJournal
	result := m.db.First(&journal, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &journal, nil
}

func (m *MSSQL) GetJournalsByPatient(username string) ([]*DBJournal, error) {
	var journals []*DBJournal
	result := m.db.Where("Patient = ?", username).Find(&journals)
	if result.Error != nil {
		return nil, result.Error
	}
	return journals, nil
}

func (m *MSSQL) CreateJournal(journal *DBJournal) error {
	result := m.db.Create(journal)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) UpdateJournal(journal *DBJournal) error {
	result := m.db.Where("JournalId = ?", journal.JournalId).Save(&journal)
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
	result = m.db.Where("DocumentId = ?", journalDocument.DocumentId).Omit("CreationTime").Save(&journalDocument)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) CreateJournalDocument(journalDocument *DBJournalDocument) error {
	var result *gorm.DB
	result = m.db.Create(&journalDocument)
	fmt.Println("huh")

	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("MSSQL DocID: %v\n\r", journalDocument.DocumentId)

	return nil
}

func (m *MSSQL) GetJournalDocumentsByJournal(journalID int32) ([]*DBJournalDocument, error) {
	var journalDocuments []*DBJournalDocument
	result := m.db.Where("JournalId = ?", journalID).Preload("Attachments.FileType").Preload("Attachments.FileStore").Find(&journalDocuments)
	if result.Error != nil {
		return nil, result.Error
	}

	return journalDocuments, nil
}

func (m *MSSQL) GetJournalDocument(journalDocumentID int32) (*DBJournalDocument, error) {
	var journalDocument DBJournalDocument
	result := m.db.Where("DocumentId = ?", journalDocumentID).Preload("Attachments").Preload("Attachments.FileType").Preload("Attachments.FileStore").First(&journalDocument)
	if result.Error != nil {
		return nil, result.Error
	}

	return &journalDocument, nil
}

func (m *MSSQL) GetDiagnose(id int32) (*DBDiagnose, error) {
	var diagnose DBDiagnose
	result := m.db.First(&diagnose, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &diagnose, nil
}

func (m *MSSQL) GetDiagnoses() ([]*DBDiagnose, error) {
	var diagnoses []*DBDiagnose
	result := m.db.Find(&diagnoses)
	if result.Error != nil {
		return nil, result.Error
	}

	return diagnoses, nil
}

func (m *MSSQL) GetHospitals() ([]*DBHospital, error) {
	var hospitals []*DBHospital
	result := m.db.Find(&hospitals)
	if result.Error != nil {
		return nil, result.Error
	}

	return hospitals, nil
}

func (m *MSSQL) GetSymptom(id int32) (*DBSymptom, error) {
	var symptom DBSymptom
	result := m.db.First(&symptom, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &symptom, nil
}

func (m *MSSQL) GetSymptoms() ([]*DBSymptom, error) {
	var symptoms []*DBSymptom
	result := m.db.Find(&symptoms)
	if result.Error != nil {
		return nil, result.Error
	}

	return symptoms, nil
}

func (m *MSSQL) CreatePatientDiagnose(patientDiagnose *DBPatientDiagnose) error {
	var result *gorm.DB
	if patientDiagnose.DiagnoseId == 0 {
		result = m.db.Omit("DiagnoseId").Create(patientDiagnose)
	} else {
		result = m.db.Create(patientDiagnose)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) GetPatientDiagnoses(username string) ([]*DBPatientDiagnose, error) {
	var patientDiagnoses []*DBPatientDiagnose
	result := m.db.Where("Patient = ?", username).Find(&patientDiagnoses)
	if result.Error != nil {
		return nil, result.Error
	}

	return patientDiagnoses, nil
}

func (m *MSSQL) GetPatientDiagnose(id int32) (*DBPatientDiagnose, error) {
	var patientDiagnose DBPatientDiagnose
	result := m.db.First(&patientDiagnose, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &patientDiagnose, nil
}

func (m *MSSQL) UpdatePatientDiagnose(pd *DBPatientDiagnose) error {
	result := m.db.Where("PatientDiagnoseId = ?", pd.PatientDiagnoseId).Omit("CreationTime").Save(&pd)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) DeletePatientDiagnose(patientDiagnose *DBPatientDiagnose) error {
	result := m.db.Where("PatientDiagnoseId = ?", patientDiagnose.PatientDiagnoseId).Delete(patientDiagnose)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) CreatePatientDiagnoseSymptom(patientDiagnoseSymptom *DBPatientDiagnoseSymptom) error {
	result := m.db.Create(patientDiagnoseSymptom)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) GetPatientDiagnoseSymptoms(patientDiagnoseId int32) ([]*DBPatientDiagnoseSymptom, error) {
	var patientDiagnoseSymptoms []*DBPatientDiagnoseSymptom
	result := m.db.Where("PatientDiagnoseId = ?", patientDiagnoseId).Find(&patientDiagnoseSymptoms)
	if result.Error != nil {
		return nil, result.Error
	}

	return patientDiagnoseSymptoms, nil
}

func (m *MSSQL) UpdatePatientDiagnoseSymptom(old *DBPatientDiagnoseSymptom, new *DBPatientDiagnoseSymptom) error {
	result := m.db.Where("PatientdiagnoseId = ? AND SymptomId = ?", old.PatientDiagnoseId, old.SymptomId).Save(&new)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) DeletePatientDiagnoseSymptom(old *DBPatientDiagnoseSymptom) error {
	result := m.db.Where("PatientdiagnoseId = ? AND SymptomId = ?", old.PatientDiagnoseId, old.SymptomId).Delete(old)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) CreateAttachment(attachment *DBAttachment) error {
	result := m.db.Create(attachment)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) GetOrCreateFileTypeByName(name string) (*DBFileType, error) {
	var fileType DBFileType
	result := m.db.Where("Name = ?", name).First(&fileType)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fileType.Name = name
			result = m.db.Create(&fileType)
			if result.Error != nil {
				return nil, result.Error
			}
			return &fileType, nil
		}

		return nil, result.Error
	}

	return &fileType, nil
}

func (m *MSSQL) GetOrCreateFileStoreByPath(path string) (*DBFileStore, error) {
	var fileStore DBFileStore
	fmt.Println("Finding store")
	result := m.db.Where("Path = ?", path).First(&fileStore)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fileStore.Path = path
			fmt.Printf("Creating store%v\n\r", fileStore.Path)
			result = m.db.Create(&fileStore)
			if result.Error != nil {
				return nil, result.Error
			}
			fmt.Println("returning stuff 1")
			return &fileStore, nil
		}
		return nil, result.Error
	}

	fmt.Println("returning stuff 2")
	return &fileStore, nil
}

func (m *MSSQL) CreateBooking(booking *DBBooking) error {
	result := m.db.Create(booking)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) CreateHospitilization(hospitilization *DBHospitilization) error {
	result := m.db.Create(hospitilization)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) CreateExamination(examination *DBExamination) error {
	result := m.db.Create(examination)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) GetBooking(id int32) (*DBBooking, error) {
	var booking DBBooking
	result := m.db.Preload("Hospital").First(&booking, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &booking, nil
}

func (m *MSSQL) GetHospitilizationByBookingId(id int32) (*DBHospitilization, error) {
	var hospitilization DBHospitilization
	result := m.db.Where("BookingId = ?", id).First(&hospitilization)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &hospitilization, nil
}

func (m *MSSQL) GetExaminationByBookingId(id int32) (*DBExamination, error) {
	var examination DBExamination
	result := m.db.Where("BookingId = ?", id).First(&examination)
	if result.Error != nil {
		return nil, result.Error
	}

	return &examination, nil
}

func (m *MSSQL) UpdateBooking(booking *DBBooking) error {
	result := m.db.Where("BookingId = ?", booking.BookingId).Save(&booking)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) DeleteBooking(booking *DBBooking) error {
	result := m.db.Where("BookingId = ?", booking.BookingId).Delete(booking)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MSSQL) GetBookingsByPatient(username string) ([]*DBBooking, error) {
	var bookings []*DBBooking
	result := m.db.Preload("Hospital").Where("Patient = ?", username).Find(&bookings)
	if result.Error != nil {
		return nil, result.Error
	}

	return bookings, nil
}

func (m *MSSQL) GetBookingsByEmployee(username string) ([]*DBBooking, error) {
	var bookings []*DBBooking
	result := m.db.Preload("Hospital").Where("Employee = ?", username).Find(&bookings)
	if result.Error != nil {
		return nil, result.Error
	}

	return bookings, nil
}
