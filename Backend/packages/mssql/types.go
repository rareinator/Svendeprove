package mssql

import "time"

type DBAttachment struct {
	AttachmentId    int32             `gorm:"column:AttachmentId;primaryKey"`
	FileName        string            `gorm:"column:FileName"`
	FileStoreId     int32             `gorm:"column:FileStoreId"`
	DocumentId      int32             `gorm:"column:DocumentId"`
	FileTypeId      int32             `gorm:"column:FileTypeId"`
	JournalDocument DBJournalDocument `gorm:"foreignKey:DocumentId;references:DocumentId"`
	FileType        DBFileType        `gorm:"foreignKey:FileTypeId;references:FileTypeId"`
	FileStore       DBFileStore       `gorm:"foreignKey:FileStoreId;references:FileStoreId"`
}

func (DBAttachment) TableName() string {
	return "Attachment"
}

type DBFileStore struct {
	FileStoreId int32  `gorm:"column:FileStoreId;primaryKey"`
	Path        string `gorm:"column:Path"`
}

type DBFileType struct {
	FileTypeId int32  `gorm:"column:FileTypeId;primaryKey"`
	Name       string `gorm:"column:Name"`
}

func (DBFileType) TableName() string {
	return "FileType"
}

func (DBFileStore) TableName() string {
	return "FileStore"
}

type DBJournalDocument struct {
	DocumentId   int32          `gorm:"column:DocumentId;primaryKey"`
	Content      string         `gorm:"column:Content"`
	JournalId    int32          `gorm:"column:Journalid"`
	CreatedBy    string         `gorm:"column:CreatedBy"`
	Title        string         `gorm:"column:Title"`
	Summary      string         `gorm:"column:Summary"`
	CreationTime time.Time      `gorm:"column:CreationTime"`
	Attachments  []DBAttachment `gorm:"foreignKey:DocumentId"`
}

func (DBJournalDocument) TableName() string {
	return "Document"
}

func (DBJournalDocument) GetPatientIDQuery() string {
	return "SELECT j.PatientId FROM Document AS d INNER JOIN Journal AS j on j.JournalId = d.JournalId WHERE d.DocumentId=?"
}

type DBJournal struct {
	JournalId    int32     `gorm:"column:JournalId;primaryKey"`
	CreationTime time.Time `gorm:"column:CreationTime"`
	Intro        string    `gorm:"column:Intro"`
	Patient      string    `gorm:"column:Patient"`
	CreatedBy    string    `gorm:"column:CreatedBy"`
}

func (DBJournal) TableName() string {
	return "Journal"
}

type DBToken struct {
	Token      string    `gorm:"column:Token;primaryKey"`
	IssuedAt   time.Time `gorm:"column:IssuedAt"`
	ValidUntil time.Time `gorm:"column:ValidUntil"`
}

func (DBToken) TableName() string {
	return "Tokens"
}

// type DBPatient struct {
// 	PatientId  int32  `gorm:"column:PatientId;primaryKey"`
// 	Name       string `gorm:"column:Name"`
// 	Address    string `gorm:"column:Address"`
// 	City       string `gorm:"column:City"`
// 	PostalCode string `gorm:"column:PostalCode"`
// 	Country    string `gorm:"column:Country"`
// 	SocialIdNr string `gorm:"column:SocialIdNr"`
// 	Username   string `gorm:"column:Username"`
// 	Password   string `gorm:"column:Password"`
// 	Salt       string `gorm:"column:Salt"`
// }

// func (DBPatient) TableName() string {
// 	return "Patient"
// }

type DBPatientDiagnose struct {
	PatientDiagnoseId int32     `gorm:"column:PatientDiagnoseId;primaryKey"`
	Patient           string    `gorm:"column:Patient"`
	DiagnoseId        int32     `gorm:"column:DiagnoseId"`
	CreationTime      time.Time `gorm:"column:CreationTime"`
}

func (DBPatientDiagnose) TableName() string {
	return "PatientDiagnose"
}

type DBDiagnose struct {
	DiagnoseId  int32  `gorm:"column:DiagnoseId;primaryKey"`
	Description string `gorm:"column:Description"`
}

func (DBDiagnose) TableName() string {
	return "Diagnose"
}

type DBSymptom struct {
	SymptomId   int32  `gorm:"column:SymptomId;primaryKey"`
	Description string `gorm:"column:Description"`
}

func (DBSymptom) TableName() string {
	return "Symptom"
}

type DBPatientDiagnoseSymptom struct {
	PatientDiagnoseId int32 `gorm:"column:PatientDiagnoseId"`
	SymptomId         int32 `gorm:"column:SymptomId"`
}

func (DBPatientDiagnoseSymptom) TableName() string {
	return "PatientDiagnoseSymptom"
}

type DBHospital struct {
	HospitalId int32  `gorm:"column:HospitalId"`
	Name       string `gorm:"column:Name"`
	Address    string `gorm:"column:Address"`
	City       string `gorm:"column:City"`
	PostalCode string `gorm:"column:PostalCode"`
	Country    string `gorm:"column:Country"`
}

func (DBHospital) TableName() string {
	return "Hospital"
}

type DBBooking struct {
	BookingId          int32      `gorm:"column:BookingId;primaryKey"`
	Bookedtime         time.Time  `gorm:"column:Bookedtime"`
	BookedEnd          time.Time  `gorm:"column:BookedEnd"`
	Patient            string     `gorm:"column:Patient"`
	Employee           string     `gorm:"column:Employee"`
	ApprovedByEmployee bool       `gorm:"column:ApprovedByEmployee"`
	HospitalId         int32      `gorm:"column:HospitalId"`
	Hospital           DBHospital `gorm:"foreignKey:HospitalId;references:HospitalId"`
}

func (DBBooking) TableName() string {
	return "Booking"
}

// type DBEmployee struct {
// 	EmployeeId   int32  `gorm:"column:EmployeeId;primaryKey"`
// 	Name         string `gorm:"column:Name"`
// 	WorktitleId  int32  `gorm:"column:WorktitleId"`
// 	DepartmentId int32  `gorm:"column:DepartmentId"`
// 	Username     string `gorm:"column:Username"`
// }

// func (DBEmployee) TableName() string {
// 	return "Employee"
// }

type DBBed struct {
	BedId        int32  `gorm:"column:BedId"`
	Name         string `gorm:"column:Name"`
	DepartmentId int32  `gorm:"column:DepartmentId"`
	IsAvailable  bool   `gorm:"column:IsAvailable"`
}

func (DBBed) TableName() string {
	return "Bed"
}

type DBHospitilization struct {
	HospitilizationId int32     `gorm:"column:HospitilizationId,primaryKey"`
	BookingId         int32     `gorm:"column:BookingId"`
	Description       string    `gorm:"column:Description"`
	StartedTime       time.Time `gorm:"column:StartedTime"`
	EndedTime         time.Time `gorm:"column:EndedTime"`
	BedId             int32     `gorm:"column:BedId"`
	Bed               DBBed     `gorm:"foreignKey:BedId;references;BedId"`
	Booking           DBBooking `gorm:"foreignKey:BookingId;references;BookingId"`
}

func (DBHospitilization) TableName() string {
	return "Hospitilization"
}

type DBExamination struct {
	ExaminationId int32     `gorm:"column:ExaminationId"`
	Description   string    `gorm:"column:Description"`
	StartedTime   time.Time `gorm:"column:StartedTime"`
	EndedTime     time.Time `gorm:"column:EndedTime"`
	BookingId     int32     `gorm:"column:BookingId"`
	Booking       DBBooking `gorm:"foreignKey:BookingId;references:BookingId"`
}

func (DBExamination) TableName() string {
	return "Examination"
}
