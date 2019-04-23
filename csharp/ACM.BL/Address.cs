using System;
using System.Collections.Generic;
namespace ACM.BL
{
    public class Address
    {
        public Address()
        {
            
        }

        public Address(int id)
        {
            this.Id = id;
        }

        public int Id { get; private set; }
        public int Type { get; set; }
        public string StreetLine1 { get; set; }
        public string StreetLine2 { get; set; }
        public string City { get; set; }    
        public string State { get; set; }
        public string PostalCode { get; set; }
        public string Country { get; set; }
    }
}