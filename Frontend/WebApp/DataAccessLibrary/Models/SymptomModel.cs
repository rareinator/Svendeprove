using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace DataAccessLibrary.Models
{
    public class SymptomModel
    {
        [JsonPropertyName("SymptomId")]
        public int SymptomId { get; set; }
        public string Description { get; set; }
    }
}
