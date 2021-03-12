using DataAccessLibrary.Converters;
using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace DataAccessLibrary.Models
{
    public class JournalModel
    {
        public int JournalId { get; set; }

        [JsonConverter(typeof(DateTimeFormatter))]
        public DateTime CreationTime { get; set; }
        public string Intro { get; set; }
        public int PatientId { get; set; }
        public int CreatedBy { get; set; }
    }
}
