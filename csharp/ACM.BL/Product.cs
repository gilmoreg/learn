using System;
namespace ACM.BL
{
    public class Product
    {
        public Product()
        {
            
        }

        public Product(int productId)
        {
            this.ProductId = productId;
        }

        public Decimal? CurrentPrice { get; set; }
        public string Description { get; set; }
        public string Name { get; set; }

        public int ProductId { get; private set; }

        public Product Retrieve(int productId)
        {
            return new Product();
        }

        public bool Save()
        {
            return true;
        }

        public bool Validate()
        {
            var isValid = true;
            if(string.IsNullOrWhiteSpace(Name)) isValid = false;
            return isValid;
        }
    }
}