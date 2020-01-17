using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using NLog.ELK.Demo.Models;
using System.Diagnostics;

namespace NLog.ELK.Demo.Controllers
{
    public class HomeController : Controller
    {
        public HomeController(ILogger<HomeController> logger)
        {
            _logger = logger;
        }

        private readonly Microsoft.Extensions.Logging.ILogger _logger;

        public IActionResult Index()
        {
            _logger.LogInformation("view /Home/Index");
            return View();
        }

        public IActionResult Privacy()
        {
            return View();
        }

        [ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
        public IActionResult Error()
        {
            return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
        }
    }
}
