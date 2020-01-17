using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using System;
using System.Threading;
using System.Threading.Tasks;

namespace AutofacDemo.HostedService
{
    public class AlwaysRunningJob : BackgroundService
    {
        public AlwaysRunningJob(ILogger<AlwaysRunningJob> logger)
        {
            _logger = logger;
        }

        #region Fields

        private readonly ILogger _logger;

        #endregion

        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            try
            {
                _logger.LogDebug($"{DateTime.Now}|{GetType().Name}|Start");

                stoppingToken.Register(() =>
                {
                    _logger.LogInformation($"{DateTime.Now}|{GetType().Name}|Stop");
                });

                while (!stoppingToken.IsCancellationRequested)
                {
                    try
                    {
                        _logger.LogInformation($"{DateTime.Now}|{GetType().Name}|Start");
                        await Task.Delay(3 * 1000);
                    }
                    catch (Exception ex)
                    {
                        _logger.LogError(ex, "{job} get exception", GetType().Name);
                    }
                }

                _logger.LogDebug($"{DateTime.Now}|{GetType().Name}|Stop");
            }
            catch (Exception ex)
            {
                if (!stoppingToken.IsCancellationRequested)
                {
                    _logger.LogCritical(ex, $"{DateTime.Now}|{GetType().Name}|Exception");
                }
                else
                {
                    _logger.LogCritical($"{DateTime.Now}|{GetType().Name}|Stop");
                }
            }
        }
    }
}
