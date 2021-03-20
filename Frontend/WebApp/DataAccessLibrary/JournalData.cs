using DataAccessLibrary.Models;
using DataAccessLibrary.TransferObjects;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public class JournalData : IJournalData
    {
        private readonly HttpClient _client;

        public JournalData(HttpClient client)
        {
            _client = client;
        }

        #region JOURNALS
        /// <summary>
        /// Sends a GET request to enpoint "/journal/byPatient/{param} to get journal by the specified patient 
        /// </summary>
        /// <param name="patient"></param>
        /// <returns>List of patient journals</returns>
        public async Task<List<JournalModel>> GetJournalsByPatient(string patient)
        {
            return await _client.GetFromJsonAsync<List<JournalModel>>($"/journal/byPatient/{patient}");
        }

        public async Task<JournalModel> InsertJournal(JournalModel journal)
        {
            var response = await _client.PostAsJsonAsync($"/journal", journal);
            string responseMessage = await response.Content.ReadAsStringAsync();

            JournalModel responseJournal = JsonSerializer.Deserialize<JournalModel>(responseMessage);

            return responseJournal;
        }
        #endregion

        #region JOURNAL DOCUMENTS
        public async Task<List<JournalDocumentModel>> GetJournalDocuments(int journalId)
        {
            return await _client.GetFromJsonAsync<List<JournalDocumentModel>>($"/journal/document/byJournal/{journalId}");
        }

        public async Task<JournalDocumentModel> GetJournalDocument(int documentId)
        {
            return await _client.GetFromJsonAsync<JournalDocumentModel>($"/journal/document/{documentId}");
        }

        public async Task<JournalDocumentModel> InsertJournalDocument(JournalDocumentModel document)
        {
            var response = await _client.PostAsJsonAsync($"/journal/document", document);
            string responseMessage = await response.Content.ReadAsStringAsync();

            JournalDocumentModel responseDocument = JsonSerializer.Deserialize<JournalDocumentModel>(responseMessage);

            return responseDocument;
        }

        public async Task<JournalDocumentModel> UpdateJournalDocument(JournalDocumentModel document)
        {

            var response = await _client.PostAsJsonAsync($"/journal/document/{document.DocumentId}", document);
            string responseMessage = await response.Content.ReadAsStringAsync();

            JournalDocumentModel responseDocument = JsonSerializer.Deserialize<JournalDocumentModel>(responseMessage);

            return responseDocument;
        }

        public async Task<List<CancerPredictionDto>> GetPredictions(List<string> images)
        {
            var response = await _client.PostAsJsonAsync($"journal/ml", images);
            string responseMessage = await response.Content.ReadAsStringAsync();

            List<CancerPredictionDto> predictions = JsonSerializer.Deserialize<List<CancerPredictionDto>>(responseMessage);

            return predictions;
        }

        public async void DeleteJournalDocument(int documentId)
        {
            await _client.DeleteAsync($"/journal/document/{documentId}");
        }
        #endregion
    }
}
