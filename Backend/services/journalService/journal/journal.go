package journal

import (
	"context"
	"fmt"
	"time"

	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
)

type JournalServer struct {
	UnimplementedJournalServiceServer
	DB            *mssql.MSSQL
	ListenAddress string
}

func (j *JournalServer) GetJournal(ctx context.Context, journal *JournalRequest) (*Journal, error) {
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

func (j *JournalServer) GetJournalsByPatient(ctx context.Context, pr *PatientRequest) (*Journals, error) {
	journals := &Journals{}
	journals.Journals = make([]*Journal, 0)

	dbJournals, err := j.DB.GetJournalsByPatient(pr.PatientId)
	if err != nil {
		return nil, err
	}

	for _, dbJournal := range dbJournals {
		journal := &Journal{
			JournalId:    dbJournal.JournalId,
			CreationTime: dbJournal.CreationTime.Format("02/01/2006 15:04:05"),
			Intro:        dbJournal.Intro,
			PatientId:    dbJournal.PatientId,
			CreatedBy:    dbJournal.CreatedBy,
		}

		journals.Journals = append(journals.Journals, journal)
	}

	return journals, nil

}

func (j *JournalServer) CreateJournal(ctx context.Context, journal *Journal) (*Journal, error) {
	dbJournal := mssql.DBJournal{
		CreationTime: time.Now(),
		Intro:        journal.Intro,
		PatientId:    journal.PatientId,
		CreatedBy:    journal.CreatedBy,
	}

	journal.CreationTime = dbJournal.CreationTime.Format("02/01/2006 15:04:05")

	if err := j.DB.CreateJournal(&dbJournal); err != nil {
		return nil, err
	}

	journal.JournalId = dbJournal.JournalId

	return journal, nil

}

func (j *JournalServer) UpdateJournal(ctx context.Context, journal *Journal) (*Journal, error) {
	parsedtime, err := time.Parse("02/01/2006 15:04:05", journal.CreationTime)
	if err != nil {
		return nil, err
	}

	dbJournal := mssql.DBJournal{
		JournalId:    journal.JournalId,
		CreationTime: parsedtime,
		Intro:        journal.Intro,
		PatientId:    journal.PatientId,
		CreatedBy:    journal.CreatedBy,
	}

	if err := j.DB.UpdateJournal(&dbJournal); err != nil {
		return nil, err
	}

	return journal, nil
}

func (j *JournalServer) DeleteJournal(ctx context.Context, jr *JournalRequest) (*Status, error) {
	dbJournal := mssql.DBJournal{
		JournalId: jr.JournalId,
	}

	if err := j.DB.DeleteJournal(&dbJournal); err != nil {
		return &Status{Success: false}, err
	}

	return &Status{Success: true}, nil
}

func (j *JournalServer) DeleteJournalDocument(ctx context.Context, jdr *JournalDocumentRequest) (*Status, error) {
	dbJournalDocument := mssql.DBJournalDocument{
		DocumentId: jdr.JournalDocumentId,
	}

	if err := j.DB.DeleteJournalDocument(&dbJournalDocument); err != nil {
		return &Status{Success: false}, err
	}

	return &Status{Success: true}, nil
}

func (j *JournalServer) UpdateJournalDocument(ctx context.Context, jdr *JournalDocument) (*JournalDocument, error) {
	dbJournalDocument := mssql.DBJournalDocument{
		DocumentId: jdr.DocumentId,
		Content:    jdr.Content,
		JournalId:  jdr.JournalId,
		CreatedBy:  jdr.CreatedBy,
		Title:      jdr.Title,
		Summary:    jdr.Summary,
	}

	if err := j.DB.UpdateJournalDocument(&dbJournalDocument); err != nil {
		return nil, err
	}

	return jdr, nil

}

func (j *JournalServer) CreateJournalDocument(ctx context.Context, jd *JournalDocument) (*JournalDocument, error) {
	fmt.Println("JournalService got called")
	dbJD := mssql.DBJournalDocument{
		Content:      jd.Content,
		JournalId:    jd.JournalId,
		CreatedBy:    jd.CreatedBy,
		Title:        jd.Title,
		Summary:      jd.Summary,
		CreationTime: time.Now(),
	}

	jd.CreationTime = dbJD.CreationTime.Format("02/01/2006 15:04:05")

	if err := j.DB.CreateJournalDocument(&dbJD); err != nil {
		return nil, err
	}

	fmt.Println("did databaseStuff")

	if len(jd.Attachments) > 0 {
		fmt.Println("Saving journal document Attachments")
		for _, attachment := range jd.Attachments {
			fileType, err := j.DB.GetOrCreateFileTypeByName(*attachment.FileType)
			if err != nil {
				return nil, err
			}
			fmt.Println("did filetype stuff")
			//build up store name
			storeName := fmt.Sprintf("/journal/document/%v", dbJD.DocumentId)
			fmt.Println("buildUpStoreName")
			store, err := j.DB.GetOrCreateFileStoreByPath(storeName)
			if err != nil {
				return nil, err
			}

			dbAttachment := mssql.DBAttachment{
				FileName:    attachment.FileName,
				FileStoreId: store.FileStoreId,
				DocumentId:  dbJD.DocumentId,
				FileTypeId:  fileType.FileTypeId,
			}

			if err := j.DB.CreateAttachment(&dbAttachment); err != nil {
				return nil, err
			}
			path := fmt.Sprintf("http://cloud.m9ssen.me:56060/static%v/%v.%v", store.Path, attachment.FileName, *attachment.FileType)
			fmt.Printf("path: %v\n\r", path)
			attachment.Path = &path
		}
	}

	jd.DocumentId = dbJD.DocumentId

	return jd, nil
}

func (j *JournalServer) CreateAttachment(ctx context.Context, attachment *Attachment) (*Attachment, error) {
	fileType, err := j.DB.GetOrCreateFileTypeByName(*attachment.FileType)
	if err != nil {
		return nil, err
	}
	fileStore, err := j.DB.GetOrCreateFileStoreByPath(*attachment.Path)
	if err != nil {
		return nil, err
	}

	dbAttachment := mssql.DBAttachment{
		FileName:    attachment.FileName,
		FileStoreId: fileStore.FileStoreId,
		DocumentId:  attachment.DocumentId,
		FileTypeId:  fileType.FileTypeId,
	}

	if err := j.DB.CreateAttachment(&dbAttachment); err != nil {
		return nil, err
	}

	return attachment, nil
}

func (j *JournalServer) GetJournalDocumentsByJournal(ctx context.Context, jr *JournalRequest) (*JournalDocuments, error) {
	journalDocuments := JournalDocuments{
		JournalDocuments: make([]*JournalDocument, 0),
	}

	dbJournalDocuments, err := j.DB.GetJournalDocumentsByJournal(jr.JournalId)
	if err != nil {
		return nil, err
	}

	for _, dbJournalDocument := range dbJournalDocuments {
		journalDocument := &JournalDocument{
			DocumentId:   dbJournalDocument.DocumentId,
			Content:      dbJournalDocument.Content,
			JournalId:    dbJournalDocument.JournalId,
			CreatedBy:    dbJournalDocument.CreatedBy,
			Title:        dbJournalDocument.Title,
			Summary:      dbJournalDocument.Summary,
			CreationTime: dbJournalDocument.CreationTime.Format("02/01/2006 15:04:05"),
		}

		journalDocuments.JournalDocuments = append(journalDocuments.JournalDocuments, journalDocument)
	}

	return &journalDocuments, nil

}

func (j *JournalServer) GetJournalDocument(ctx context.Context, jdr *JournalDocumentRequest) (*JournalDocument, error) {
	dbJournalDocument, err := j.DB.GetJournalDocument(jdr.JournalDocumentId)
	if err != nil {
		return nil, err
	}

	var attachments []*Attachment
	fmt.Println("Getting journal")

	if len(dbJournalDocument.Attachments) > 0 {
		fmt.Println("Found attachments")
		for _, attachment := range dbJournalDocument.Attachments {
			resultAttachment := Attachment{
				AttachmentId: attachment.AttachmentId,
				FileName:     attachment.FileName,
				FileType:     new(string),
				Path:         new(string),
			}
			path := fmt.Sprintf("http://cloud.m9ssen.me:56060/static%v/%v.%v", attachment.FileStore.Path, attachment.FileName, attachment.FileType.Name)
			fmt.Printf("path: %v\n\r", path)
			resultAttachment.Path = &path
			resultAttachment.FileType = &attachment.FileType.Name
			attachments = append(attachments, &resultAttachment)
		}
	}

	result := JournalDocument{
		DocumentId:   dbJournalDocument.DocumentId,
		Content:      dbJournalDocument.Content,
		JournalId:    dbJournalDocument.JournalId,
		CreatedBy:    dbJournalDocument.CreatedBy,
		Title:        dbJournalDocument.Title,
		Summary:      dbJournalDocument.Summary,
		CreationTime: dbJournalDocument.CreationTime.Format("02/01/2006 15:04:05"),
		Attachments:  attachments,
	}

	return &result, nil
}
