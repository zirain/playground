using Autofac;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using System;
using System.Threading;
using System.Threading.Tasks;

namespace AutofacDemo
{
    public class TestAlwaysRunningJob : BackgroundService
    {
        public TestAlwaysRunningJob(IServiceProvider provider,
            ILifetimeScope lifetimeScope,
            ILogger<TestAlwaysRunningJob> logger)
        {
            _provider = provider;
            _logger = logger;
            _lifetimeScope = lifetimeScope;
        }

        private readonly IServiceProvider _provider;
        private readonly ILifetimeScope _lifetimeScope;
        private readonly ILogger _logger;

        protected override Task ExecuteAsync(CancellationToken stoppingToken)
        {
            while (!stoppingToken.IsCancellationRequested)
            {
                //    va _provider.GetService<IScopeEntity>();
                //Console.Write(_scopeEntity);

                using (var scope = _lifetimeScope.BeginLifetimeScope())
                {
                    var scopeEntity = scope.Resolve<IScopeEntity>();
                    _logger.LogInformation(scopeEntity.Id.ToString());

                    Thread.Sleep(3 * 1000);
                }
            }

            return Task.CompletedTask;
        }
    }
}
