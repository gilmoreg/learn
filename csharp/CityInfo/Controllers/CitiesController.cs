using System.Collections.Generic;
using System.Linq;
using Microsoft.AspNetCore.Mvc;

namespace CityInfo.Controllers
{
    [Route("api/[controller]")]
    public class CitiesController : Controller
    {
        [HttpGet]
        public IActionResult GetCities()
        {
            // Ok() supports multiple formats (not just json)
            return Ok(CityDataStore.Current.Cities);
        }

        [HttpGet("{id}")]
        public IActionResult GetCity(int id)
        {
            var city = CityDataStore.Current.Cities.FirstOrDefault(c => c.Id == id);
            if (city == null) return NotFound();
            return Ok(city);
        }
    }
}