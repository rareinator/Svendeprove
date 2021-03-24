using DataAccessLibrary.Models;
using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace DataAccessLibrary.TransferObjects
{
    public class DiagnosePredictDto
    {
        [JsonPropertyName("code")]
        public int Code { get; set; }
        [JsonPropertyName("diagnosis")]
        public List<DiagnoseDto> Diagnosis { get; set; }
    }

    public class DiagnoseDto
    {
        [JsonPropertyName("name")]
        public string Name { get; set; }
        [JsonPropertyName("symptoms")]
        public List<string> Symptoms { get; set; }
    }
}
