using DataAccessLibrary.Models;
using DataAccessLibrary.TransferObjects;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace DataAccessLibrary
{
    public interface IPatientData
    {
        /// <summary>
        /// Sends a GET request to endpoint "/patient/{param}" to get a patient by the specified patientId 
        /// </summary>
        /// <param name="patientId"></param>
        /// <returns>A user object</returns>
        Task<UserModel> GetPatient(int patientId);

        /// <summary>
        /// Sends a GET request to endpoint "/patient" to get all patients 
        /// </summary>
        /// <returns>List of user objects</returns>
        Task<List<UserModel>> GetPatients();

        /// <summary>
        /// Sends a GET request to endpoint "/diagnose/{param}" to get patient diagnosis by specified patient UserId 
        /// </summary>
        /// <param name="patient"></param>
        /// <returns>List of patient diagnose objects</returns>
        Task<List<PatientDiagnoseModel>> GetDiagnosisByPatient(string patient);

        /// <summary>
        /// Sends a GET request to endpoint "/symptom" to get all symptoms 
        /// </summary>
        /// <returns>List of symptom objects</returns>
        Task<List<SymptomModel>> GetSymptoms();

        /// <summary>
        /// Sends a GET request to enpoint "/diagnose" to get all diagnosis 
        /// </summary>
        /// <returns>List of diagnose objects</returns>
        Task<List<DiagnoseModel>> GetDiagnosis();

        /// <summary>
        /// Sends a POST request to endpoint "/journal/symptoms" with {param} in the request body to get diagnose predictions by specified symptoms object
        /// </summary>
        /// <param name="data"></param>
        /// <returns>A diagnose prediction object</returns>
        Task<DiagnosePredictDto> GetDiagnosePrediction(SymptomsDto data);

        /// <summary>
        /// Sends a POST request to endpoint "/patient/{param.Patient}/diagnose" with {param} in the request body to add a new patient diagnose
        /// </summary>
        /// <param name="patientDiagnose"></param>
        /// <returns>New added patient diagnose object</returns>
        Task<PatientDiagnoseModel> InsertPatientDiagnose(PatientDiagnoseModel patientDiagnose);

        /// <summary>
        /// Sends a POST request to endpoint "/patient/{param1}/diagnose/{param2}/symptom" with {param3} in the request body to add a new patient symptom
        /// </summary>
        /// <param name="patient"></param>
        /// <param name="patientDiagnoseId"></param>
        /// <param name="symptom"></param>
        Task InsertPatientSymptom(string patient, int patientDiagnoseId, SymptomModel symptom);

        /// <summary>
        /// Sends a DELETE request to endpoint "/patient/{param.Patient}/diagnose/{param.PatientDiagnoseId}" to delete specified patient diagnose
        /// </summary>
        /// <param name="patientDiagnose"></param>
        Task DeletePatientDiagnose(PatientDiagnoseModel patientDiagnose);
    }
}