using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Net.Http.Json;
using System.Threading.Tasks;
using System.Linq;
using DataAccessLibrary.TransferObjects;
using System.Text.Json;

namespace DataAccessLibrary
{
    public class PatientData : IPatientData
    {
        private readonly HttpClient _client;


        public PatientData(HttpClient client)
        {
            _client = client;
        }

        public async Task<List<UserModel>> GetPatients()
        {
            return await _client.GetFromJsonAsync<List<UserModel>>("/patient"); ;
        }

        public async Task<UserModel> GetPatient(int patientId)
        {
            return await _client.GetFromJsonAsync<UserModel>($"/patient/{patientId}");
        }

        public async Task<List<SymptomModel>> GetSymptoms()
        {
            return await _client.GetFromJsonAsync<List<SymptomModel>>($"/symptom"); ;
        }

        public async Task<List<DiagnoseModel>> GetDiagnosis()
        {
            return await _client.GetFromJsonAsync<List<DiagnoseModel>>($"/diagnose"); ;
        }

        public async Task<DiagnosePredictDto> GetDiagnosePrediction(SymptomsDto data)
        {
            var response = await _client.PostAsJsonAsync($"/journal/symptoms", data);
            string responseMessage = await response.Content.ReadAsStringAsync();

            DiagnosePredictDto prediction = JsonSerializer.Deserialize<DiagnosePredictDto>(responseMessage);

            return prediction;
        }

        public async Task<PatientDiagnoseModel> InsertPatientDiagnose(PatientDiagnoseModel patentDiagnose)
        {
            var response = await _client.PostAsJsonAsync($"/patient/{patentDiagnose.Patient}/diagnose", patentDiagnose);
            string responseMessage = await response.Content.ReadAsStringAsync();

            PatientDiagnoseModel responseModel = JsonSerializer.Deserialize<PatientDiagnoseModel>(responseMessage);

            foreach (var symptom in patentDiagnose.Symptoms)
            {
                await InsertPatientSymptom(responseModel.Patient, responseModel.PatientDiagnoseId, symptom);
            }

            return responseModel;
        }

        public async Task InsertPatientSymptom(string patient, int patientDiagnoseId, SymptomModel symptom)
        {
            await _client.PostAsJsonAsync($"/patient/{patient}/diagnose/{patientDiagnoseId}/symptom", symptom);
        }

        public async Task<List<PatientDiagnoseModel>> GetDiagnosisByPatient(string patient)
        {
            return await _client.GetFromJsonAsync<List<PatientDiagnoseModel>>($"/patient/{patient}/diagnose");
        }

        public async Task DeletePatientDiagnose(PatientDiagnoseModel patientDiagnose)
        {
            await _client.DeleteAsync($"/patient/{patientDiagnose.Patient}/diagnose/{patientDiagnose.PatientDiagnoseId}");
        }
    }
}
