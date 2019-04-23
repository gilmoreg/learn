using System;
using System.IO;
using DictationProcessorLib;

namespace DictationProcessorApp
{
    class Program
    {
        static void Main(string[] args)
        {
            foreach(var subfolder in Directory.GetDirectories("/mnt/uploads"))
            {
                var uploadProcessor = new UploadProcessor(subfolder);
                uploadProcessor.Process();
            }
        }
    }
}
