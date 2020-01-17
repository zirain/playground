using Microsoft.Extensions.Localization;
using Microsoft.Extensions.Logging;

namespace LocalizationDemo.Services
{
    public class LocaleProvider : ILocaleProvider
    {
        public LocaleProvider(
            ILogger<LocaleProvider> logger,
            IStringLocalizer<LocaleProvider> localizer)
        {
            _logger = logger;
            _localizer = localizer;
        }

        private readonly ILogger _logger;
        private readonly IStringLocalizer _localizer;

        public void Test()
        {
            _logger.LogInformation("{guid} {message}", System.Guid.NewGuid(), _localizer["message"]);
        }
    }
}
