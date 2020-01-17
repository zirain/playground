using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Http.Internal;
using Microsoft.AspNetCore.StaticFiles;
using System;
using System.IO;
using System.Text;
using System.Threading.Tasks;

namespace MiddlewareDemo
{
    /// <summary>
    /// inspire by https://elanderson.net/2017/02/log-requests-and-responses-in-asp-net-core/
    /// </summary>
    public class RequestResponseLoggingMiddleware
    {
        private readonly RequestDelegate _next;
        //private readonly ITableStorageService _storageService;

        public RequestResponseLoggingMiddleware(RequestDelegate next
            //, ITableStorageService storageService
            )
        {
            _next = next;
            //_storageService = storageService;
        }

        public async Task InvokeAsync(HttpContext context)
        {
            var contentTypeProvider = new FileExtensionContentTypeProvider();
            if (contentTypeProvider.TryGetContentType(context.Request.Path, out string _)) // ignore static file
            {
                await _next(context);
            }
            else
            {
                var requestUrl = $"{context.Request.Method}|{context.Request.Scheme}://{context.Request.Host}{context.Request.Path}";
                //var logEntity = new PlatformApiLogEntity(DateTime.UtcNow, requestUrl)
                //{
                //    RequestContent = await FormatRequest(context.Request)
                //};

                var originalBodyStream = context.Response.Body;
                using (var responseBody = new MemoryStream())
                {
                    context.Response.Body = responseBody;

                    await _next(context);

                    //logEntity.ResponseContent = await FormatResponse(context.Response);
                    //logEntity.ResponseTime = DateTime.UtcNow.ToString(CultureInfo.InvariantCulture);
                    //_storageService.SaveBorkerApiLog(logEntity);

                    await responseBody.CopyToAsync(originalBodyStream);
                }
            }
        }

        private async Task<string> FormatRequest(HttpRequest request)
        {
            request.EnableRewind();

            var buffer = new byte[Convert.ToInt32(request.ContentLength)];
            await request.Body.ReadAsync(buffer, 0, buffer.Length);
            var bodyAsText = Encoding.UTF8.GetString(buffer);
            request.Body.Position = 0;

            return $"QueryString:{request.QueryString},Body:{bodyAsText}";
        }

        private async Task<string> FormatResponse(HttpResponse response)
        {
            response.Body.Seek(0, SeekOrigin.Begin);
            var text = await new StreamReader(response.Body).ReadToEndAsync();
            response.Body.Seek(0, SeekOrigin.Begin);

            return $"{text}";
        }
    }
}
