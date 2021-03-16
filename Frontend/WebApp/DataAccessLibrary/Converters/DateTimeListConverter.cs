using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace DataAccessLibrary.Converters
{
    public class DateTimeListConverter : JsonConverter<List<DateTime>>
    {
        public override List<DateTime> Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        {
            string test = reader.GetString();
            List<DateTime> result = new List<DateTime>();
            return result;
        }

        public override void Write(Utf8JsonWriter writer, List<DateTime> value, JsonSerializerOptions options)
        {
            throw new NotImplementedException();
        }
    }
}
