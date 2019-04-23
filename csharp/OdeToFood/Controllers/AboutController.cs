using Microsoft.AspNetCore.Mvc;

namespace OdeToFood.Controllers
{
    [Route("[controller]/[action]")]
    public class AboutController
    {
        // Empty string is a default
        // [Route("")]
        public string Phone()
        {
            return "Fake phone number";
        }

        public string Address()
        {
            return "fake address";
        }
    }
}