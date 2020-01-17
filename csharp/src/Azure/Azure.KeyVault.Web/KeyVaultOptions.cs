namespace Azure.KeyVault.Web
{
    public class KeyVaultOptions
    {
        public string ClientId { get; set; }

        public string ClientSecret { get; set; }

        public string DnsName { get; set; }

        public string CryptographyIdentity { get; set; }
    }
}
