using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Threading.Tasks;

namespace blueberry_homework_dotnet.DTO.Request
{
    public class ChangeNameRequest
    {
        [Required]
        public Guid? Id { get; set; }

        [Required]
        [StringLength(150, MinimumLength = 1)]
        public required string Name { get; set; }
    }
}