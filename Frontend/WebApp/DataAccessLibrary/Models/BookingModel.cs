using System;
using System.Collections.Generic;
using System.Text;

namespace DataAccessLibrary.Models
{
    public enum BookingType { Examination, Hospitalization }

    public class BookingModel
    {
        public int BookingId { get; set; }
        public DateTime BookedTime { get; set; }
        public DateTime BookedEnd { get; set; }
        public bool Approved { get; set; }
        public string Patient { get; set; }
        public string Employee { get; set; }
        public string Description { get; set; }
        public BookingType Type { get; set; }

        public int HospitalId { get; set; }
        public HospitalModel Hospital { get; set; }
    }
}
