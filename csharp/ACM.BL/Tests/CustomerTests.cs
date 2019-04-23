using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using ACM.BL;

namespace ACM.BL.Tests
{
    [TestClass]
    public class CustomerTests
    {
        [TestMethod]
        public void FullNameTestValid()
        {
            var customer = new Customer();
            customer.FirstName = "First";
            customer.LastName = "Last";
            Assert.AreEqual("Last, First", customer.FullName);
        }

        [TestMethod]
        public void FullNameTestLastNameEmpty()
        {
            var customer = new Customer();
            customer.FirstName = "First";
            Assert.AreEqual("First", customer.FullName);
        }

        [TestMethod]
        public void FullNameTestFirstNameEmpty()
        {
            var customer = new Customer();
            customer.LastName = "Last";
            Assert.AreEqual("Last", customer.FullName);
        }

        [TestMethod]
        public void ValidateValidTest()
        {
            var customer = new Customer();
            customer.LastName = "Last";
            customer.Email = "test@test.com";
            Assert.AreEqual(customer.Validate(), true);
        }

        [TestMethod]
        public void ValidateInvalidTest()
        {
            var customer = new Customer();
            Assert.AreEqual(customer.Validate(), false);
        }
    }
}