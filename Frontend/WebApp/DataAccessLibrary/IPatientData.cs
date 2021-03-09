using DataAccessLibrary.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IPatientData
    {
        List<PatientModel> GetPatients();
        List<JournalModel> GetPatientJournals(int patientId);
    }
}