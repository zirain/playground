using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.AspNetCore.Mvc.ViewFeatures;
using Microsoft.AspNetCore.Razor.TagHelpers;
using System.Threading.Tasks;

namespace TagHelperDemo.TagHelpers
{
    [HtmlTargetElement("bs-datetime", Attributes = ForAttributeName, TagStructure = TagStructure.WithoutEndTag)]
    public class BootstrapDatetimePickerTagHelper : TagHelper
    {
        private const string TemplateName = "DatetimePicker";
        private const string ForAttributeName = "asp-for";
        private const string ReadonlyAttributeName = "asp-readonly";
        private const string StartViewAttributeName = "asp-start-view";
        private const string FormatViewAttributeName = "asp-format";

        private readonly IHtmlHelper _htmlHelper;

        public BootstrapDatetimePickerTagHelper(IHtmlHelper htmlHelper)
        {
            _htmlHelper = htmlHelper;
        }

        #region Properties

        [HtmlAttributeName(ForAttributeName)]
        public ModelExpression For { get; set; }

        [HtmlAttributeName(ReadonlyAttributeName)]
        public string Readonly { get; set; }

        /// <summary>
        /// Default: 'days'
        /// Accepts: 'decades','years','months','days'
        /// See detail: https://eonasdan.github.io/bootstrap-datetimepicker/Options/#viewmode
        /// </summary>
        [HtmlAttributeName(StartViewAttributeName)]
        public string StartView { get; set; } = "days";

        #endregion

        [HtmlAttributeNotBound]
        [ViewContext]
        public ViewContext ViewContext { get; set; }

        public override async Task ProcessAsync(TagHelperContext context, TagHelperOutput output)
        {
            await base.ProcessAsync(context, output);

            //Contextualize the html helper
            (_htmlHelper as IViewContextAware).Contextualize(ViewContext);

            var data = new ViewDataDictionary(ViewContext.ViewData)
            {
                ["Readonly"] = Readonly,
                ["startView"] = StartView,
                ["ID"] = For.Name
            };


            var htmlContent = await _htmlHelper.PartialAsync("TagHelpers/BootstrapDatetimePicker", data);

            output.TagName = null; // Remove the <default> element
            output.Content.SetHtmlContent(htmlContent);
        }
    }
}
