using System;
using System.Collections.Generic;
using System.Globalization;
using System.Text;
using System.Text.Json;
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

    public class DateTimeFormatter : JsonConverter<DateTime>
    {
        public override DateTime Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
            => DateTime.ParseExact(reader.GetString(), "dd/MM/yyyy hh:mm:ss", CultureInfo.InvariantCulture);

        public override void Write(Utf8JsonWriter writer, DateTime value, JsonSerializerOptions options)
            => writer.WriteStringValue(value.ToString("dd/MM/yyyy hh:mm:ss", CultureInfo.InvariantCulture));
    }
}
