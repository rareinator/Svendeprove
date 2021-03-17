using DataAccessLibrary.TransferObjects;
using System;
using System.Collections.Generic;
using System.Text;

namespace DataAccessLibrary.Models
{
    public class AttachmentModel
    {
        public int AttachmentId { get; set; }
        public string FileName { get; set; }
        public string Path { get; set; }
        public string FileType { get; set; }
        public string Content { get; set; }

        public Prediction Prediction { get; set; }
    }
}
