using System;
using System.Collections.Generic;
using System.Text;

namespace DataAccessLibrary.Models
{
    public class PredictionModel
    {
        public Data Data { get; set; }
        public string Url { get; set; }
    }

    public class Data
    {
        public int code { get; set; }
        public Prediction prediction { get; set; }
    }

    public class Prediction
    {
        public string negative { get; set; }
        public string positive { get; set; }
    }
}
