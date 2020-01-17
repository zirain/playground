using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using ViewRender.Models;

namespace ViewRender.Controllers
{
    public class HomeController : Controller
    {
        public HomeController(IViewRenderProvider viewRenderProvider)
        {
            _viewRenderProvider = viewRenderProvider;
        }

        private readonly IViewRenderProvider _viewRenderProvider;

        public IActionResult Index()
        {
            var result = _viewRenderProvider.RenderToStringAsync("~/Views/Shared/Error.cshtml", new ErrorViewModel() { RequestId = Guid.NewGuid().ToString() }).Result;
            return View();
        }

        public IActionResult About()
        {
            ViewData["Message"] = "Your application description page.";

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
