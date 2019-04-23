using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Grades
{
    class Program
    {
        static void Main(string[] args)
        {
            GradeBook book = new GradeBook();
            // book.NameChanged += new NameChangedDelegate(OnNameChanged);
            // Shorthand
            //book.NameChanged += OnNameChanged;
            AssignGradeBookName(book);
            AddGrades(book);
            SaveGrades(book);
            WriteResults(book);
        }

        private static void WriteResults(GradeBook book)
        {
            GradeStatistics stats = book.ComputeStatistics();
            WriteResult("Average", stats.AvgGrade);
            WriteResult("Max", stats.MaxGrade);
            WriteResult("Min", stats.MinGrade);
            WriteResult("Grade", stats.LetterGrade);
            WriteResult("Results", stats.Description);
        }

        private static void SaveGrades(GradeBook book)
        {
            // Automatically closes the outputFile (even if there's an exception)
            using (StreamWriter outputFile = File.CreateText("grades.txt"))
            {
                book.WriteGrades(outputFile);
            }
        }

        private static void AddGrades(GradeBook book)
        {
            book.AddGrade(91);
            book.AddGrade(89.5f);
            book.AddGrade(75);
        }

        private static void AssignGradeBookName(GradeBook book)
        {
            Console.WriteLine("Please enter a gradebook name:");
            try
            {
                book.Name = Console.ReadLine();
            }
            catch (ArgumentNullException ex)
            {
                Console.WriteLine(ex.Message);
            }
            catch (NullReferenceException)
            {
                Console.WriteLine("Something went wrong");
            }
        }

        static void OnNameChanged(object sender, NameChangedEventArgs args)
        {
            Console.WriteLine($"Gradebook changing name from {args.ExistingName} to {args.NewName}.");
        }

        static void WriteResult(string description, float result)
        {
            // Float with two decimal places
            Console.WriteLine($"{description}: {result:F2}");
        }

        static void WriteResult(string description, string result)
        {
            Console.WriteLine($"{description}: {result}");
        }



    }
}
