using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Net.Http.Json;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public class PatientData : IPatientData
    {
        private readonly HttpClient _client;

        public PatientData(HttpClient client)
        {
            _client = client;
        }

        public Task<List<PatientModel>> GetPatients()
        {
            return _client.GetFromJsonAsync<List<PatientModel>>("/patient");
        }
    }
}
