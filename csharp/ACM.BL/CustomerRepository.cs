using System;
using System.Collections.Generic;
namespace ACM.BL
{
    public class CustomerRepository
    {
        public List<Customer> Retrieve()
        {
            return new List<Customer>();
        }

        public Customer Retrieve(int id)
        {
            var customer = new Customer(id);
            if (id == 1) {
                customer.Email = "fbaggins@hobbiton.me";
                customer.FirstName = "Frodo";
                customer.LastName = "Baggins";

            }
            return customer;
        }

        public bool Save()
        {
            // Code to save to database TODO
            return true;
        }
    }
}