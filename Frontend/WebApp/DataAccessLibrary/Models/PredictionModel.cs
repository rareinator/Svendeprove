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
        public int Code { get; set; }
        public Prediction Prediction { get; set; }
    }

    public class Prediction
    {
        public string Negative { get; set; }
        public string Positive { get; set; }
    }
}
