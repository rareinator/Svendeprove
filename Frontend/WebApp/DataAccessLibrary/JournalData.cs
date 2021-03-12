using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public class JournalData : IJournalData
    {
        private readonly HTTPService _client;

        public JournalData(HTTPService client)
        {
            _client = client;
        }

        public async Task<List<JournalModel>> GetJournalsByPatient(int patientId)
        {
            return await _client.GetData<List<JournalModel>>($"/journal/byPatient/{patientId}");
        }

        public async Task<List<JournalDocumentModel>> GetJournalDocuments(int journalId)
        {
            return await _client.GetData<List<JournalDocumentModel>>($"/journal/document/byJournal/{journalId}");
        }


        public void InsertJournal(JournalModel journal)
        {
            _client.PostData<JournalModel,JournalModel>($"/journal", journal);
        }

        public void InsertJournalDocument(JournalDocumentModel document)
        {
            _client.PostData<JournalDocumentModel, JournalDocumentModel>($"/journal/document", document);
        }
    }
}
