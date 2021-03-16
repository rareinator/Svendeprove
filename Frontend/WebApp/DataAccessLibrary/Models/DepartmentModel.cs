using System;
using System.Collections.Generic;
using System.Text;

namespace DataAccessLibrary.Models
{
    public class DepartmentModel
    {
        public int DepartmentId { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public int HospitalId { get; set; }

        public HospitalModel Hospital { get; set; }

        public ICollection<BedModel> Beds { get; set; }
    }
}
