using DataAccessLibrary.Converters;
using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace DataAccessLibrary.TransferObjects
{
    public class AvailableEmpDto
    {
        [JsonConverter(typeof(DateTimeConverter))]
        public DateTime Day { get; set; }
        public string Doctor { get; set; }
    }
}
