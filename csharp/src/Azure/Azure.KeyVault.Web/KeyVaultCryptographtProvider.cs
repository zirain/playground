using Microsoft.Azure.KeyVault;
using Microsoft.Azure.KeyVault.WebKey;
using Microsoft.Extensions.Options;
using Microsoft.IdentityModel.Clients.ActiveDirectory;
using System;
using System.Text;
using System.Threading.Tasks;

namespace Azure.KeyVault.Web
{
    public class KeyVaultCryptographtProvider : ICryptographyProvider
    {
        public KeyVaultCryptographtProvider(IOptionsSnapshot<KeyVaultOptions> keyVaultOptions)
        {
            _options = keyVaultOptions;
            _client = new KeyVaultClient(async (string authority, string resource, string scope) =>
            {
                var authContext = new AuthenticationContext(authority);
                var clientCred = new ClientCredential(_options.Value?.ClientId, _options.Value?.ClientSecret);
                var result = await authContext.AcquireTokenAsync(resource, clientCred);
                if (result == null)
                {
                    throw new InvalidOperationException("Failed to retrieve access token for Key Vault");
                }

                return result.AccessToken;
            });
        }

        private readonly IKeyVaultClient _client;
        private readonly IOptionsSnapshot<KeyVaultOptions> _options;

        private readonly string _algorithm = JsonWebKeyEncryptionAlgorithm.RSAOAEP;

        public async Task<string> DecryptAsync(string ciphertext)
        {
            var encryptedBytes = Convert.FromBase64String(ciphertext);
            var decryptResult = await _client.DecryptAsync(_options.Value?.CryptographyIdentity, _algorithm, encryptedBytes);
            if (decryptResult != null)
            {
                return Encoding.UTF8.GetString(decryptResult.Result);
            }

            return null;
        }

        public async Task<string> EncryptAsync(string plaintext)
        {
            var plaintextBytes = Encoding.UTF8.GetBytes(plaintext);
            var encryptResult = await _client.EncryptAsync(_options.Value?.CryptographyIdentity, _algorithm, plaintextBytes);
            return encryptResult == null ? null : Convert.ToBase64String(encryptResult.Result);
        }
    }
}
