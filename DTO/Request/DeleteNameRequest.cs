using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Threading.Tasks;

namespace BerryNameApi.DTO.Request
{
    public class DeleteNameRequest
    {
        [Required]
        [StringLength(50, MinimumLength = 1)]
        public required string Name { get; set; }
    }
}