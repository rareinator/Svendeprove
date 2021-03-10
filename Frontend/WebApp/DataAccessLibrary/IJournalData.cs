using DataAccessLibrary.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IJournalData
    {
        Task<List<JournalModel>> GetJournalsByPatient(int patientId);
        void InsertJournal(JournalModel journal);
        Task<List<JournalDocumentModel>> GetJournalDocuments(int journalId);
        void InsertJournalDocument(JournalDocumentModel journalDocument);
    }
}