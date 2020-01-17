using Autofac.Extras.DynamicProxy;
using GettingStarted.Interceptors;
using System.Threading;

namespace GettingStarted.Services
{
    [Intercept(typeof(CallLoggerInterceptor))]
    public interface ILoggerService
    {

    }

    public interface IDemoService : ILoggerService
    {
        void Test(int sec);
    }

    public class DemoService : IDemoService
    {
        public void Test(int sec)
        {
            Thread.Sleep(sec * 1000);
        }
    }
}
