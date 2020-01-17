using Consul;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Hosting.Server.Features;
using Microsoft.AspNetCore.Http.Features;
using Microsoft.Extensions.Logging;
using System;
using System.Linq;

namespace ConsulTestApi
{
    public static class ConsulAppBuilderExtensions
    {
        public static IApplicationBuilder RegisterConsul(this IApplicationBuilder app,
            ConsulConfig consulConfig,
            IConsulClient consulClient,
            IApplicationLifetime lifetime,
            ILogger logger)
        {
            try
            {
                var registrationID = string.Empty;
                lifetime.ApplicationStarted.Register(() =>
                {
                    var features = app.Properties["server.Features"] as FeatureCollection;
                    var serverAddressesFeature = features.Get<IServerAddressesFeature>();
                    var addresses = serverAddressesFeature.Addresses;
                    foreach (var address in addresses)
                    {
                        logger.LogInformation("listen @ {uri}", address);
                    }
                    var uri = new Uri(addresses.First());

                    // Register service with consul
                    var registration = new AgentServiceRegistration()
                    {
                        ID = $"{consulConfig.ServiceId}-{uri.Port}",
                        Name = consulConfig.ServiceName,
                        Address = consulConfig.ServiceAddress,
                        Port = uri.Port,
                        Tags = consulConfig.ServiceTags
                    };

                    registrationID = registration.ID;
                    consulClient.Agent.ServiceDeregister(registration.ID).Wait();
                    consulClient.Agent.ServiceRegister(registration).Wait();
                });

                lifetime.ApplicationStopping.Register(() =>
                {
                    //服务停止时取消注册
                    consulClient.Agent.ServiceDeregister(registrationID).Wait();
                });
            }
            catch (Exception ex)
            {
                logger.LogCritical(ex, "register to consul fail");
            }

            return app;
        }
    }
}