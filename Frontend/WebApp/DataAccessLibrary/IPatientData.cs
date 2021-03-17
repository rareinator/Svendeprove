using DataAccessLibrary.Models;
using DataAccessLibrary.TransferObjects;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IPatientData
    {
        Task<UserModel> GetPatient(int patientId);
        Task<List<UserModel>> GetPatients();
        Task<List<PatientDiagnoseModel>> GetDiagnosisByPatient(string patient);
        Task<List<SymptomModel>> GetSymptoms();
        Task<List<DiagnoseModel>> GetDiagnosis();
        Task<DiagnosePredictDto> GetDiagnosePrediction(SymptomsDto data);
        Task<PatientDiagnoseModel> InsertPatientDiagnose(PatientDiagnoseModel patientDiagnose);
        Task InsertPatientSymptom(string patient, int patientDiagnoseId, SymptomModel symptom);
        Task DeletePatientDiagnose(PatientDiagnoseModel patientDiagnose);
    }
}