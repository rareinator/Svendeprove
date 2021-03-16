using System;
using System.Collections.Generic;
using System.Text;

namespace DataAccessLibrary.Models
{
    public class BedModel
    {
        public int BedId { get; set; }
        public string Name { get; set; }
        public bool IsAvailable { get; set; }
        public int DepartmentId { get; set; }

        public DepartmentModel Department { get; set; }
    }
}
