using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace BerryNameApi.DTO.Response
{
    public class ApiFailResponse
    {
        public required string Message { get; set; } = "error";
        public string? Error { get; set; }
    }
}