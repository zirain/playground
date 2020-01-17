using Azure.KeyVault.Web.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using System.Diagnostics;

namespace Azure.KeyVault.Web.Controllers
{
    public class HomeController : Controller
    {
        public HomeController(
            IConfiguration configuration,
            IOptionsSnapshot<CustomOptions> options,
            ILogger<HomeController> logger)
        {
            _configuration = configuration as IConfigurationRoot;
            _options = options;
            _logger = logger;
        }

        private readonly IConfigurationRoot _configuration;
        private readonly IOptionsSnapshot<CustomOptions> _options;
        private readonly ILogger _logger;

        public IActionResult Index()
        {
            _logger.LogInformation("CustomOptions.AppSecret:" + _options.Value.AppSecret);
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

        public IActionResult Reload()
        {
            _configuration.Reload();

            return RedirectToAction("Index");
        }
    }
}
