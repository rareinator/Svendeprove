using DataAccessLibrary.Converters;
using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace DataAccessLibrary.Models
{
    public enum BookingType { Examination, Hospitilization }

    public class BookingModel
    {
        public int BookingId { get; set; }

        [JsonConverter(typeof(DateTimeConverter))]
        public DateTime BookedTime { get; set; }

        [JsonConverter(typeof(DateTimeConverter))]
        public DateTime BookedEnd { get; set; }

        public bool Approved { get; set; }
        public string Patient { get; set; }
        public string Employee { get; set; }
        public string Description { get; set; }

        [JsonConverter(typeof(JsonStringEnumConverter))]
        public BookingType Type { get; set; }

        public HospitalModel Hospital { get; set; }
        public BedModel Bed { get; set; }
    }
}
