using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Threading.Tasks;

namespace blueberry_homework_dotnet.DTO.Request
{
    public class CreateCompanyRequest
    {
        [Required]
        [StringLength(50, MinimumLength = 1)]
        public required string UserName { get; set; }

        [Required]
        [StringLength(100, MinimumLength = 1)]
        public required string CompanyName { get; set; }
    }
}