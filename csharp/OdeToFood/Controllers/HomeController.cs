using Microsoft.AspNetCore.Mvc;
using OdeToFood.Models;
using OdeToFood.Services;
using OdeToFood.ViewModels;

namespace OdeToFood.Controllers
{
    // by default it is the HomeController that will receive a request to the root
    public class HomeController : Controller
    {
        public HomeController(IRestaurantData restaurantData, IGreeter greeter)
        {
            _restaurantData = restaurantData;
            _greeter = greeter;
        }

        // MVC default route will look for this method
        public IActionResult Index()
        {
            // return Content("Hello from the HomeController!");
            var results = new HomeIndexViewModel{
                Restaurant = _restaurantData.GetAll(),
                CurrentMessage = _greeter.GetMessageOfTheDay()
            };
            // return new ObjectResult(model);
            
            // Without a name, assume we look for a view with the same name as the action
            // In Views/[controller] or Views/Shared
            return View(results);
        }

        public IActionResult Details(int id)
        {
            var model = _restaurantData.Get(id);
            if (model == null) return RedirectToAction("Index");
            return View(model);
        }

        [HttpGet]
        public IActionResult Create()
        {
            return View();
        }

        [HttpPost]
        [ValidateAntiForgeryToken]
        public IActionResult Create(RestaurantEditModel model)
        {
            if (!ModelState.IsValid) return View();
            var newRestaurant = new Restaurant{
                Name = model.Name,
                Cuisine = model.Cuisine
            };
            newRestaurant = _restaurantData.Add(newRestaurant);
            return RedirectToAction(nameof(Details), new { id=newRestaurant.Id });
        }

        private IRestaurantData _restaurantData;
        private IGreeter _greeter;
    }
}