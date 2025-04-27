using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace blueberry_homework_dotnet.Entities
{
    public class CompanyEntity
    {
        public Guid Id { get; set; }
        public required string Name { get; set; }
        public required string CompanyName { get; set; }
        public DateTime CreatedAt { get; set; }
    }
}