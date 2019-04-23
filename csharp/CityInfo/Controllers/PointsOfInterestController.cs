using System;
using System.Collections.Generic;
using System.Linq;
using CityInfo.Models;
using CityInfo.Services;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace CityInfo.Controllers
{
    [Route("api/cities/")]
    public class PointsOfInterestController : Controller
    {
        private ILogger<PointsOfInterestController> _logger;
        private IMailService _mailService;
        public PointsOfInterestController(ILogger<PointsOfInterestController> logger, IMailService mailService)
        {
            _logger = logger;
            _mailService = mailService;
        }

        [HttpGet("{cityId}/pointsofinterest")]
        public IActionResult GetPointsOfInterest(int cityId)
        {
            try {
                var city = CityDataStore.Current.Cities.FirstOrDefault(c => c.Id == cityId);
                if (city == null)
                {
                    _logger.LogInformation($"City with id {cityId} was not found accessing points of interest.");
                    return NotFound();
                }
                return Ok(city.PointsOfInterest);
            } catch (Exception ex) {
                _logger.LogCritical($"Exception while getting points of interest for city id {cityId}");
                return StatusCode(500, "Something went wrong.");
            }
        }

        [HttpGet("{cityId}/pointsofinterest/{id}", Name = "GetPointOfInterest")]
        public IActionResult GetPointOfInterest(int cityId, int id)
        {
            var city = CityDataStore.Current.Cities.FirstOrDefault(c => c.Id == cityId);
            if (city == null) return NotFound();
            var point = city.PointsOfInterest.FirstOrDefault(p => p.Id == id);
            if (point == null) return NotFound();
            return Ok(point);
        }

        [HttpPost("{cityId}/pointsofinterest")]
        public IActionResult CreatePointOfInterest(int cityId, [FromBody] PointOfInterestForCreation pointOfInterest)
        {
            if (pointOfInterest == null) return BadRequest();
            if (pointOfInterest.Name == pointOfInterest.Description) {
                ModelState.AddModelError("Description", "Description and Name should differ");
            }
            if (!ModelState.IsValid) return BadRequest(ModelState);
            var city = CityDataStore.Current.Cities.FirstOrDefault(c => c.Id == cityId);
            if (city == null) return NotFound();
            var max = CityDataStore.Current.Cities.SelectMany(c => c.PointsOfInterest).Max(p => p.Id);
            var finalPointOfInterest = new PointOfInterestDto(){
                Id = ++max,
                Name = pointOfInterest.Name,
                Description = pointOfInterest.Description
            };
            city.PointsOfInterest.Add(finalPointOfInterest);
            return CreatedAtRoute("GetPointOfInterest", new { cityId = cityId, id = finalPointOfInterest.Id }, finalPointOfInterest);
        }

        [HttpPut("{cityId}/pointsofinterest/{id}")]
        public IActionResult UpdatePointOfInterest(int cityId, int id, 
                                                    [FromBody] PointOfInterestForUpdate pointOfInterest)
        {
            if (pointOfInterest == null) return BadRequest();
            if (pointOfInterest.Name == pointOfInterest.Description) {
                ModelState.AddModelError("Description", "Description and Name should differ");
            }
            if (!ModelState.IsValid) return BadRequest(ModelState);
            var city = CityDataStore.Current.Cities.FirstOrDefault(c => c.Id == cityId);
            if (city == null) return NotFound();
            var point = city.PointsOfInterest.FirstOrDefault(p => p.Id == id);
            if (point == null) return NotFound();
            point.Name = pointOfInterest.Name;
            point.Description = pointOfInterest.Description;
            return NoContent();
        }

        [HttpDelete("{cityId}/pointsofinterest/{id}")]
        public IActionResult DeletePointOfInterest(int cityId, int id)
        {
            var city = CityDataStore.Current.Cities.FirstOrDefault(c => c.Id == cityId);
            if (city == null) return NotFound();
            var point = city.PointsOfInterest.FirstOrDefault(p => p.Id == id);
            if (point == null) return NotFound();
            city.PointsOfInterest.Remove(point);
            _mailService.Send("Point of interest deleted",
                $"Point of interest {point.Name} with id {point.Id} was deleted.");
            return NoContent();
        }
    }
}