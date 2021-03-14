using DataAccessLibrary.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IJournalData
    {
        // JOURNAL
        Task<List<JournalModel>> GetJournalsByPatient(int patientId);
        Task<JournalModel> InsertJournal(JournalModel journal);

        // JOURNAL DOCUMENTS
        Task<List<JournalDocumentModel>> GetJournalDocuments(int journalId);
        Task<JournalDocumentModel> GetJournalDocument(int documentId);
        Task<List<PredictionModel>> GetPredictions(List<string> images);
        Task<JournalDocumentModel> InsertJournalDocument(JournalDocumentModel document);
        Task<JournalDocumentModel> UpdateJournalDocument(JournalDocumentModel document);
        void DeleteJournalDocument(int documentId);
    }
}