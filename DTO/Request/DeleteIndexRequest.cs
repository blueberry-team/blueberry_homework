using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Threading.Tasks;

namespace BerryNameApi.DTO.Request
{
    public class DeleteIndexRequest
    {
        [Required]
        public int? Index { get; set; }
    }
}