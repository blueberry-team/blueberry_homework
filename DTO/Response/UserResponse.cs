using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace BerryNameApi.DTO.Response
{
    public class UserResponse
    {
        public required string Name { get; set; }
        public DateTime CreatedAt { get; set; }
    }
}