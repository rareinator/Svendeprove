using DataAccessLibrary.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IPatientData
    {
        Task<List<UserModel>> GetPatients();
        Task<UserModel> GetPatient(int patientId);
    }
}