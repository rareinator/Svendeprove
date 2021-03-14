using DataAccessLibrary.Converters;
using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace DataAccessLibrary.Models
{
    public class JournalDocumentModel
    {
        public int DocumentId { get; set; }

        [JsonConverter(typeof(DateTimeFormatter))]
        public DateTime CreationTime { get; set; }
        public string Title { get; set; }
        public string Summary { get; set; }
        public string Content { get; set; }
        public int CreatedBy { get; set; }
        public int JournalId { get; set; }

        public List<AttachmentModel> Attachments { get; set; }
    }
}
