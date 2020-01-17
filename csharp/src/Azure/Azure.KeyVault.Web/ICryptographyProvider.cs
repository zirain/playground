using System.Threading.Tasks;

namespace Azure.KeyVault.Web
{
    public interface ICryptographyProvider
    {
        Task<string> EncryptAsync(string plaintext);

        Task<string> DecryptAsync(string ciphertext);
    }
}
