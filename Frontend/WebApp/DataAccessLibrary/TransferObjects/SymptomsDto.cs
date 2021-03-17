using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace DataAccessLibrary.TransferObjects
{
    public class SymptomsDto
    {
        [JsonPropertyName("age")]
        public int Age { get; set; }
        [JsonPropertyName("gender")]
        public string Gender { get; set; }
        [JsonPropertyName("symptoms")]
        public List<string> Symptoms { get; set; }
    }
}
