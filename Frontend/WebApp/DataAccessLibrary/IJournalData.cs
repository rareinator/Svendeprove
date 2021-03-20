using DataAccessLibrary.Models;
using DataAccessLibrary.TransferObjects;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IJournalData
    {
        // JOURNAL
        /// <summary>
        /// Sends a GET request to enpoint "/journal/byPatient/{param} to get journal by the specified patient 
        /// </summary>
        /// <param name="patient"></param>
        /// <returns>List of patient journal objects</returns>
        Task<List<JournalModel>> GetJournalsByPatient(string patient);

        /// <summary>
        /// Sends a POST request to enpoint "/journal" with {param} in the request body to add a new journal to a patient  
        /// </summary>
        /// <param name="journal"></param>
        /// <returns>New added journal object</returns>
        Task<JournalModel> InsertJournal(JournalModel journal);

        // JOURNAL DOCUMENTS

        /// <summary>
        /// Sends a GET request to enpoint "/journal/document/byJournal/{param} to get journal documents by specified id of a existing journal
        /// </summary>
        /// <param name="journalId"></param>
        /// <returns>List of journal document objects</returns>
        Task<List<JournalDocumentModel>> GetJournalDocuments(int journalId);

        /// <summary>
        /// Sends a GET request to enpoint "/journal/document/{param} to get a journal document by specified id of a existing journal document
        /// </summary>
        /// <param name="documentId"></param>
        /// <returns>A journal document object</returns>
        Task<JournalDocumentModel> GetJournalDocument(int documentId);

        /// <summary>
        /// Sends a POST request to endpoint "journal/ml" with {param} in the request body to get prediction objects back
        /// </summary>
        /// <param name="images"></param>
        /// <returns>List of cancer prediction objects</returns>
        Task<List<CancerPredictionDto>> GetPredictions(List<string> images);

        /// <summary>
        /// Sends a POST request to endpoint "journal/document" with {param} in the request body to add a new jorunal document to a journal
        /// </summary>
        /// <param name="document"></param>
        /// <returns>New added journal document object</returns>
        Task<JournalDocumentModel> InsertJournalDocument(JournalDocumentModel document);

        /// <summary>
        /// Sends a POST request to endpoint "journal/document/{param.documentId} with {param} int the request body to edit a existing journal document
        /// </summary>
        /// <param name="document"></param>
        /// <returns>Edited journal document object</returns>
        Task<JournalDocumentModel> UpdateJournalDocument(JournalDocumentModel document);

        /// <summary>
        /// Sends a DELETE request to endpoint "journal/document/{param}" to delete the journal document with the specified documentId
        /// </summary>
        /// <param name="documentId"></param>
        void DeleteJournalDocument(int documentId);
    }
}