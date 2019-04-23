using System;
using System.Collections.Generic;
using System.IO;
using System.IO.Compression;
using System.Runtime.Serialization;
using System.Runtime.Serialization.Json;
using System.Security.Cryptography;

namespace DictationProcessorLib
{
    public class UploadProcessor
    {
        private string _subfolder;

        public UploadProcessor(string subfolder)
        {
            _subfolder = subfolder;
        }
        public void Process()
        {
            var metadataCollection = ExtractMetaData(_subfolder);
            foreach (var metadata in metadataCollection)
            {
                var audioFilePath = Path.Combine(_subfolder, metadata.File.FileName);
                var md5Checksum = GetChecksum(audioFilePath);
                if (md5Checksum != metadata.File.Md5Checksum)
                {
                    throw new Exception("Checksum not verified. Audio file corrupt");
                }
                var guid = Guid.NewGuid();
                metadata.File.FileName = guid + ".WAV";
                var newPath = Path.Combine("/mnt/ready_for_transcriptions", guid + ".WAV");
                CreateCompressedFile(audioFilePath, newPath);
                SaveSingleMetaData(metadata, newPath + ".json");
            }
        }

        private void CreateCompressedFile(string inputFilePath, string outputFilePath)
        {
            outputFilePath += ".gz";
            System.Console.WriteLine($"Creating {outputFilePath}");
            var inputFileStream = File.Open(inputFilePath, FileMode.Open);
            var outputFileStream = File.Create(outputFilePath);
            var gzipStream = new GZipStream(outputFileStream, CompressionLevel.Optimal);
            inputFileStream.CopyTo(gzipStream);
            inputFileStream.Dispose();
            outputFileStream.Dispose();
        }

        private string GetChecksum(string audioFilePath)
        {
            var fileStream = File.Open(audioFilePath, FileMode.Open);
            var md5 = MD5.Create();
            var md5bytes = md5.ComputeHash(fileStream);
            fileStream.Dispose();
            return BitConverter.ToString(md5bytes).Replace("-", "").ToLower();
        }

        private List<MetaData> ExtractMetaData(string subfolder)
        {
            var metadataFilePath = Path.Combine(subfolder, "metadata.json");
            Console.WriteLine($"Reading {metadataFilePath}...");
            var metadataFileStream = File.Open(metadataFilePath, FileMode.Open);
            var settings = new DataContractJsonSerializerSettings
            {
                DateTimeFormat = new DateTimeFormat("yyyy-MM-dd'T'HH:mm:ssZ")
            };
            var serializer = new DataContractJsonSerializer(typeof(List<MetaData>), settings);
            return (List<MetaData>)serializer.ReadObject(metadataFileStream);
        }

        private void SaveSingleMetaData(MetaData metadata, string metadataFilePath)
        {
            Console.WriteLine($"Creating {metadataFilePath}...");
            var metadataFileStream = File.Open(metadataFilePath, FileMode.Create);
            var settings = new DataContractJsonSerializerSettings
            {
                DateTimeFormat = new DateTimeFormat("yyyy-MM-dd'T'HH:mm:ssZ")
            };
            var serializer = new DataContractJsonSerializer(typeof(MetaData), settings);
            serializer.WriteObject(metadataFileStream, metadata);
            metadataFileStream.Dispose();
        }
    }
}
