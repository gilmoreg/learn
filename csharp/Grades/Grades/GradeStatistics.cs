using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Grades
{
    public class GradeStatistics
    {
        public GradeStatistics()
        {
            MaxGrade = 0;
            MinGrade = float.MaxValue;
        }

        public string Description
        {
            get
            {
                switch(LetterGrade)
                {
                    case "A": return "Excellent";
                    case "B": return "Good";
                    case "C": return "Average";
                    case "D": return "Below Average";
                    default: return "Failing";
                }
            }
        }

        public string LetterGrade
        {
            get
            {
                if (AvgGrade >= 90) return "A";
                else if (AvgGrade >= 80) return "B";
                else if (AvgGrade >= 70) return "C";
                else if (AvgGrade >= 60) return "D";
                else return "F";
            }
        }

        public float MaxGrade;
        public float MinGrade;
        public float AvgGrade;
    }
}
