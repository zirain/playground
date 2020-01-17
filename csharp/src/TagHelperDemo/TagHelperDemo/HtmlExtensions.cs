using Microsoft.AspNetCore.Html;
using System.IO;
using System.Text.Encodings.Web;

namespace TagHelperDemo
{
    public static class HtmlExtensions
    {

        public static string RenderHtmlContent(this IHtmlContent htmlContent)
        {
            using (var writer = new StringWriter())
            {
                htmlContent.WriteTo(writer, HtmlEncoder.Default);
                var htmlOutput = writer.ToString();
                return htmlOutput;
            }
        }
    }
}
