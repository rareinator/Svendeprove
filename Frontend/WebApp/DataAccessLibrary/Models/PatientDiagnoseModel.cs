using DataAccessLibrary.Converters;
using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace DataAccessLibrary.Models
{
    public class PatientDiagnoseModel
    {
        public int PatientDiagnoseId { get; set; }
        public string Patient { get; set; }
        public int DiagnoseId { get; set; }

        [JsonConverter(typeof(DateTimeConverter))]
        public DateTime CreationTime { get; set; }

        public DiagnoseModel Diagnose { get; set; }
        public IList<SymptomModel> Symptoms { get; set; }
    }
}
