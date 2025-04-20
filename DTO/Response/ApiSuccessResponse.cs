using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace BerryNameApi.DTO.Response
{
    public class ApiSuccessResponse<T>
    {
        public string Message { get; set; } = "success";
        public T? Data { get; set; }
    }
}