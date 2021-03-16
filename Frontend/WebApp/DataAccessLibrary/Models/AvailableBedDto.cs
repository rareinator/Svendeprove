using DataAccessLibrary.Converters;
using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace DataAccessLibrary.Models
{
    public class AvailableBedDto
    {
        [JsonConverter(typeof(DateTimeConverter))]
        public DateTime BookedTime { get; set; }
        [JsonConverter(typeof(DateTimeConverter))]
        public DateTime BookedEnd { get; set; }
        public int HospitalId { get; set; }
    }
}
