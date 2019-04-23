using Microsoft.EntityFrameworkCore;
using OdeToFood.Models;

/*
    To make migrations:
    dotnet ef migrations add NAME -v
    To apply:
    dotnet ef database update -v
 */


namespace OdeToFood.Data
{
    public class OdeToFoodDBContext : DbContext
    {
        public OdeToFoodDBContext(DbContextOptions options) : base(options) {}
        public DbSet<Restaurant> Restaurants { get; set; }
    }
}