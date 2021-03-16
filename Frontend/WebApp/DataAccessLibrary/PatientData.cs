using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Net.Http.Json;
using System.Threading.Tasks;
using System.Linq;

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


    }
}
