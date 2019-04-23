using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using ACM.BL;

namespace ACM.BL.Tests
{
    [TestClass]
    public class CustomerRepositoryTests
    {
        [TestMethod]
        public void RetrieveExisting()
        {
            var customerRepository = new CustomerRepository();
            var expected = new Customer(1)
            {
                Email = "fbaggins@hobbiton.me",
                FirstName = "Frodo",
                LastName = "Baggins"
            };
            var actual = customerRepository.Retrieve(1);
            Assert.AreEqual(actual.CustomerId, expected.CustomerId);
            Assert.AreEqual(actual.Email, expected.Email);
            Assert.AreEqual(actual.FirstName, expected.FirstName);
            Assert.AreEqual(actual.LastName, expected.LastName);
        }
    }
}