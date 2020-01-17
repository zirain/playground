using Microsoft.AspNetCore.Mvc;
using System.Threading.Tasks;

namespace Azure.KeyVault.Web.Controllers
{
    public class KeyVaultController : Controller
    {
        public KeyVaultController(ICryptographyProvider cryptographyProvider)
        {
            _cryptographyProvider = cryptographyProvider;
        }

        private readonly ICryptographyProvider _cryptographyProvider;
        public async Task<IActionResult> Encrypt(string plaintext)
        {
            var result = await _cryptographyProvider.EncryptAsync(plaintext);
            return Json(result);
        }

        public async Task<IActionResult> Decrypt(string ciphertext)
        {
            var result = await _cryptographyProvider.DecryptAsync(ciphertext);
            return Json(result);
        }
    }
}