using System;

namespace ScopeDIDemo
{
    public interface IScopeEntity
    {
        Guid Id { get; set; }
    }

    public class ScopeEntity : IScopeEntity
    {
        public ScopeEntity()
        {
            Id = Guid.NewGuid();
        }

        public Guid Id { get; set; }
    }
}
