using System.Collections.Generic;
using CityInfo.Models;

namespace CityInfo
{
    public class CityDataStore
    {
        public static CityDataStore Current { get; } = new CityDataStore();
        public List<CityDto> Cities { get; set; }

        public CityDataStore()
        {
            Cities = new List<CityDto>(){
                new CityDto(){
                    Id = 1,
                    Name = "New York City",
                    Description = "The one with the big park.",
                    PointsOfInterest = new List<PointOfInterestDto>(){
                        new PointOfInterestDto() {
                            Id = 1,
                            Name = "Central Park",
                            Description = "Most visited park in urban US"
                        },
                        new PointOfInterestDto() {
                            Id = 2,
                            Name = "Empire State Building",
                            Description = "102-story skyscraper"
                        }
                    }
                },
                new CityDto(){
                    Id = 2,
                    Name = "Antwerp",
                    Description = "The one with the unfinished cathedral.",
                    PointsOfInterest = new List<PointOfInterestDto>(){
                        new PointOfInterestDto() {
                            Id = 1,
                            Name = "Cathedral of the Lady",
                            Description = "Gothic Cathedral"
                        },
                        new PointOfInterestDto() {
                            Id = 2,
                            Name = "Antwerp Central Station",
                            Description = "Finest example of railway architecture in Belgium"
                        }
                    }
                },
                new CityDto(){
                    Id = 3,
                    Name = "Paris",
                    Description = "The one with the big tower.",
                    PointsOfInterest = new List<PointOfInterestDto>(){
                        new PointOfInterestDto() {
                            Id = 1,
                            Name = "Eiffel Tower",
                            Description = "Wrought-iron lattice tower"
                        },
                        new PointOfInterestDto() {
                            Id = 2,
                            Name = "The Louvre",
                            Description = "World's largest museum"
                        }
                    }
                },
            };
        }
    }
}