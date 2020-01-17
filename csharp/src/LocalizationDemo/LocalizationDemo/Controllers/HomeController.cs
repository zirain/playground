using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Localization;
using LocalizationDemo.Models;
using LocalizationDemo.Services;

namespace LocalizationDemo.Controllers
{
    public class HomeController : Controller
    {
        private readonly IStringLocalizer<HomeController> _localizer;
        private readonly ILocaleProvider _localeProvider;

        public HomeController(
            ILocaleProvider localeProvider,
            IStringLocalizer<HomeController> localizer)
        {
            _localeProvider = localeProvider;
            _localizer = localizer;
        }

        public IActionResult Index()
        {
            _localeProvider.Test();
            return View(new Models.Home.IndexModel());
        }

        public IActionResult About()
        {
            ViewData["Message"] = _localizer["Your application description page."];

            return View();
        }

        public IActionResult Contact()
        {
            ViewData["Message"] = "Your contact page.";

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
