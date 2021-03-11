using DataAccessLibrary.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IPatientData
    {
        Task<List<PatientModel>> GetPatients();
        Task<PatientModel> GetPatient(int patientId);
    }
}