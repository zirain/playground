using Castle.DynamicProxy;
using Microsoft.Extensions.Logging;
using System.Linq;

namespace GettingStarted.Interceptors
{
    /// <summary>
    /// 拦截器 需要实现 IInterceptor接口 Intercept方法
    /// </summary>
    public class CallLoggerInterceptor : IInterceptor
    {
        public CallLoggerInterceptor(ILogger<CallLoggerInterceptor> logger)
        {
            _logger = logger;
        }

        private readonly ILogger<CallLoggerInterceptor> _logger;

        /// <summary>
        /// 拦截方法 打印被拦截的方法执行前的名称、参数和方法执行后的 返回结果
        /// </summary>
        /// <param name="invocation">包含被拦截方法的信息</param>
        public void Intercept(IInvocation invocation)
        {
            var args = string.Join(", ", invocation.Arguments.Select(a => (a ?? "").ToString()).ToArray());
            _logger.LogInformation("你正在调用{0}方法 \"{1}\"  参数是 {2}... ",
              invocation.TargetType.FullName,
              invocation.Method.Name,
              args);

            //在被拦截的方法执行完毕后 继续执行
            invocation.Proceed();

            _logger.LogInformation("方法执行完毕，返回结果：{0}", invocation.ReturnValue);
        }
    }
}
