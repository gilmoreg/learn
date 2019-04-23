using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Grades
{
    public class GradeBook
    {
        // default constructor is one that takes no parameters
        public GradeBook()
        {
            _name = "Empty";
            grades = new List<float>();
        }

        public GradeStatistics ComputeStatistics()
        {
            GradeStatistics stats = new GradeStatistics();
            float sum = 0;
            foreach (float grade in grades)
            {
                stats.MaxGrade = Math.Max(grade, stats.MaxGrade);
                stats.MinGrade = Math.Min(grade, stats.MinGrade);
                sum += grade;
            }
            stats.AvgGrade = sum / grades.Count;
            return stats;
        }

        public void WriteGrades(TextWriter destination)
        {
            //for (int i = 0; i < grades.Count; i++)
            //{
            //    destination.WriteLine(grades[i]);
            //}
            foreach (float grade in grades)
            {
                destination.WriteLine(grade);
            }
        }

        public void AddGrade(float grade)
        {
            grades.Add(grade);
        }

        // field initializer syntax
        // List<float> grades = new List<float>(); 
        List<float> grades;

        // Autoimplemented properties
        // public string Name { get; set; }

        // Manually implemented properties
        public string Name
        {
            get { return _name; }
            set
            {
                if (String.IsNullOrEmpty(value)) throw new ArgumentNullException("Name cannot be null or empty");
                if (_name != value && NameChanged != null)
                {
                    NameChangedEventArgs args = new NameChangedEventArgs
                    {
                        ExistingName = _name,
                        NewName = value
                    };
                    NameChanged(this, args);
                }
                _name = value;

            }
        }
        public event NameChangedDelegate NameChanged;
        private string _name;
    }
}
