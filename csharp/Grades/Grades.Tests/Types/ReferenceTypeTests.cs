using Microsoft.VisualStudio.TestTools.UnitTesting;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Grades.Tests.Types
{
    [TestClass]
    public class TypeTests
    {
        [TestMethod]
        public void GradeBookVariablesHoldReference()
        {
            GradeBook g1 = new GradeBook();
            GradeBook g2 = g1;
            g1.Name = "Grayson's Grade Book";
            Assert.AreEqual(g1.Name, g2.Name);
        }

        [TestMethod]
        public void IntVariablesHoldValue()
        {
            int x1 = 100;
            int x2 = x1;
            x1 = 4;
            Assert.AreNotEqual(x1, x2);
        }

        [TestMethod]
        public void StringComparisons()
        {
            string name1 = "Grayson";
            string name2 = "grayson";
            bool result = String.Equals(name1, name2, StringComparison.InvariantCultureIgnoreCase);
            Assert.IsTrue(result);
        }

        [TestMethod]
        public void ReferenceTypesPassByValue()
        {
            GradeBook book1 = new GradeBook();
            GradeBook book2 = book1;
            GiveBookAName(book2);
            Assert.AreEqual("A Gradebook", book1.Name);

        }

        private void GiveBookAName(GradeBook book)
        {
            book.Name = "A Gradebook";
        }

        [TestMethod]
        public void ValueTypesPassByValue()
        {
            int x = 46;
            IncrementNumber(x);
            Assert.AreEqual(x, 46);
        }

        private void IncrementNumber(int number)
        {
            number++;
        }

        [TestMethod]
        public void RefBehavior()
        {
            int x = 46;
            IncrementNumberByRef(ref x);
            Assert.AreEqual(x, 47);
        }

        private void IncrementNumberByRef(ref int number)
        {
            number++;
        }

        [TestMethod]
        public void AddDaysToDateTime()
        {
            // Value types are immutable
            DateTime date = new DateTime(2015, 1, 1);
            // date.AddDays(1) does not work; original value cannot be modified
            // can only return a new DateTime
            date = date.AddDays(1);
            Assert.AreEqual(date.Day, 2);
        }

        [TestMethod]
        public void UppercaseString()
        {
            // Even though string is a reference type, it behaves like a value type insofar as it is immutable
            string str = "grayson";
            // str.ToUpper() does not work; original value cannot be modified
            // can only return a new string
            str = str.ToUpper();
            Assert.AreEqual(str, "GRAYSON");
        }

        [TestMethod]
        public void UsingArrays()
        {
            float[] grades;
            grades = new float[3];
            // Arrays are reference type so the method can change the object referred to
            AddGrades(grades);
            Assert.AreEqual(89.1f, grades[1], 0.01);
        }

        private void AddGrades(float[] grades)
        {
            grades[1] = 89.1f;
        }
    }
}
