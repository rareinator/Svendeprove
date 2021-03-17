using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace DataAccessLibrary.Models
{
    public class DiagnoseModel
    {

        [JsonPropertyName("DiagnoseId")]
        public int DiagnoseId { get; set; }
        public string Description { get; set; }
    }
}
