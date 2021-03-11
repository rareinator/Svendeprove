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

        #region DELETE THIS WHEN CONNECTED TO BACKEND API
        public List<JournalModel> journals = new List<JournalModel>
        {
            new JournalModel
            {
                PatientId = 1,
                CreationTime = DateTime.Now,
                Intro = "1"
            },
            new JournalModel
            {
                PatientId = 1,
                CreationTime = DateTime.Now,
                Intro = "12"
            },
            new JournalModel
            {
                PatientId = 2,
                CreationTime = DateTime.Now,
                Intro = "123"
            },
            new JournalModel
            {
                PatientId = 2,
                CreationTime = DateTime.Now,
                Intro = "1234"
            },
            new JournalModel
            {
                PatientId = 2,
                CreationTime = DateTime.Now,
                Intro = "12345"
            }
        };

        public List<PatientModel> patients = new List<PatientModel>
        {
            new PatientModel
            {
                PatientId = 1,
                Name = "Emile Henriksen",
                SocialIdNr = "568-02-9251"
            },
            new PatientModel
            {
                PatientId = 2,
                Name = "Mimir Bach",
                SocialIdNr = "151152-4330"
            },
            new PatientModel
            {
                PatientId = 3,
                Name = "Heng Chen",
                SocialIdNr = "050688-0114"
            },
        };
        #endregion

        public PatientData(HttpClient client)
        {
            _client = client;
        }

        public List<PatientModel> GetPatients()
        {
            //_client.GetFromJsonAsync<PatientModel>("/Patients");

            return patients;
        }

        public PatientModel GetPatient(int patientId)
        {
            return patients.Where(p => p.PatientId == patientId).FirstOrDefault();
        }

        public List<JournalModel> GetPatientJournals(int patientId)
        {
            _client.DefaultRequestHeaders.Authorization = new System.Net.Http.Headers.AuthenticationHeaderValue("Bearer", "6fe3d770-49a6-4dd1-bb9d-5996b2aea935");

            return journals.Where(j => j.PatientId == patientId).ToList();
        }

        public void InsertPatientJournal(JournalModel journal)
        {
            journals.Add(journal);
        }
    }
}
