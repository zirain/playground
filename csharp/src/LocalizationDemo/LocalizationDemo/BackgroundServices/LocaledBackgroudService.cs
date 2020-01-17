using LocalizationDemo.Services;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System.Threading;
using System.Threading.Tasks;

namespace LocalizationDemo.BackgroundServices
{
    public class LocaledBackgroudService : BackgroundService
    {
        public LocaledBackgroudService(IServiceScopeFactory factory)
        {
            _factory = factory;
        }

        private readonly IServiceScopeFactory _factory;

        protected override Task ExecuteAsync(CancellationToken stoppingToken)
        {
            while (!stoppingToken.IsCancellationRequested)
            {
                using (var scope = _factory.CreateScope())
                {
                    var culture = new System.Globalization.CultureInfo("en-US");
                    //System.Globalization.CultureInfo.CurrentCulture = culture;
                    System.Globalization.CultureInfo.CurrentUICulture = culture;
                    //System.Globalization.CultureInfo.DefaultThreadCurrentCulture = culture;
                    //System.Globalization.CultureInfo.DefaultThreadCurrentUICulture = culture;

                    var provider = scope.ServiceProvider.GetService<ILocaleProvider>();
                    provider.Test();
                }

                Thread.Sleep(3 * 1000);
            }

            return Task.CompletedTask;
        }
    }
}
