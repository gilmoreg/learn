using System.Collections.Generic;

namespace CityInfo.Models
{
    public class CityDto
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public List<PointOfInterestDto> PointsOfInterest { get; set; }
        public int NumberOfPointsOfInterest { get {
            return PointsOfInterest.Count;
        }}
    }
}