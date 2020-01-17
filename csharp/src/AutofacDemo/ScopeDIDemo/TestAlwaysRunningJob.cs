using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using System;
using System.Threading;
using System.Threading.Tasks;

namespace ScopeDIDemo
{
    public class TestAlwaysRunningJob : BackgroundService
    {
        public TestAlwaysRunningJob(IServiceProvider provider,
            ILogger<TestAlwaysRunningJob> logger)
        {
            _provider = provider;
            _logger = logger;
        }

        private readonly IServiceProvider _provider;
        private readonly ILogger _logger;

        protected override Task ExecuteAsync(CancellationToken stoppingToken)
        {
            while (!stoppingToken.IsCancellationRequested)
            {
                //    va _provider.GetService<IScopeEntity>();
                //Console.Write(_scopeEntity);

                //using (var scope = _provider.CreateScope())
                {
                    var scopeEntity = _provider.GetService<IScopeEntity>();
                    _logger.LogInformation(scopeEntity.Id.ToString());

                    Thread.Sleep(3 * 1000);
                }
            }

            return Task.CompletedTask;
        }
    }
}
