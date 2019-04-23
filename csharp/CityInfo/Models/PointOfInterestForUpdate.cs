using System.ComponentModel.DataAnnotations;

namespace CityInfo.Models
{
    public class PointOfInterestForUpdate
    {
        [MaxLength(50, ErrorMessage = "Max length for name is 50 characters")]
        public string Name { get; set; }
        [MaxLength(200, ErrorMessage = "Max length for description is 200 characters")]
        public string Description { get; set; }
    }
}