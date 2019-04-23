using CityInfo.Entities;
using Microsoft.AspNetCore.Mvc;

namespace CityInfo.Controllers
{
    public class DummyController : Controller
    {
        private CityInfoContext _ctx;

        public DummyController(CityInfoContext context)
        {
            _ctx = context;
        }

        [HttpGet]
        [Route("api/testdatabase")]
        public IActionResult TestDatabase()
        {
            return Ok();
        }
    }
}