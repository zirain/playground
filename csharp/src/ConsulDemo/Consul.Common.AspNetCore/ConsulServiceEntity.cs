namespace Consul.Common.AspNetCore
{
    public class ConsulServiceEntity
    {
        public string IP { get; set; }

        public int Port { get; set; }

        public string ServiceName { get; set; }

        public string ConsulIP { get; set; }

        public int ConsulPort { get; set; }
    }
}
