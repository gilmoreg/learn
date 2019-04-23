using System;
using System.Collections.Generic;

namespace DictationProcessor
{
    public class MetaData
    {
        public string Practitioner { get; set; }
        public string Patient { get; set; }
        public DateTime Recorder { get; set; }
        public List<string> Tags { get; set; }
        public AudioFile File { get; set; }
    }
}