using System;
using System.IO;
using DictationProcessorLib;

namespace DictationProcessorSvc
{
    class Program
    {
        static void Main(string[] args)
        {
            var fileSystemWatcher = new FileSystemWatcher("/mnt/uploads", "metadata.json");
            fileSystemWatcher.IncludeSubdirectories = true;
            while(true)
            {
                var result = fileSystemWatcher.WaitForChanged(WatcherChangeTypes.Created);
                Console.WriteLine($"{result} created.");
                var fullMetadataFilePath = Path.Combine("/mnt/uploads", result.Name);
                var subfolder = Path.GetDirectoryName(fullMetadataFilePath);
                var processor = new UploadProcessor(subfolder);
                processor.Process();
            }
        }
    }
}
