using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using BerryNameApi.DTO.Request;
using BerryNameApi.DTO.Response;
using BerryNameApi.UseCases;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.ApiExplorer;

namespace BerryNameApi.Controllers
{
    [ApiController]
    [Route("names")]
    public class NameController : ControllerBase
    {
        private readonly NameUseCase _useCase;

        public NameController(NameUseCase useCase)
        {
            _useCase = useCase;
        }

        [HttpPost("createName")]
        public IActionResult CreateName([FromBody] CreateNameRequest request)
        {
            if (!ModelState.IsValid)
                return BadRequest(new ApiFailResponse
                {
                    Error = Constnats.NameLengthInvalid
                });

            _useCase.CreateName(request.Name);

            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Data = _useCase.GetAll()
            });

        }

        [HttpGet("getName")]
        public IActionResult Get()
        {
            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Data = _useCase.GetAll()
            });
        }

        [HttpDelete("deleteIndex")]
        public IActionResult DeleteByIndex([FromBody] DeleteIndexRequest request)
        {
            if (!request.Index.HasValue)
                return BadRequest(new ApiFailResponse
                {
                    Error = Constants.DeleteIndexRequired
                });

            var deleted = _useCase.DeleteByIndex(request.Index.Value);
            if (!deleted)
                return NotFound(new ApiFailResponse
                {
                    Error = $"{Constants.InvalidIndex}: {request.Index}"
                });

            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Data = _useCase.GetAll()
            });
        }

        [HttpDelete("deleteName")]
        public IActionResult DeleteByName([FromBody] DeleteNameRequest request)
        {
            if (!ModelState.IsValid)
                return BadRequest(new ApiFailResponse
                {
                    Error = Constants.NameLengthInvalid
                });

            var count = _useCase.DeleteByName(request.Name);
            if (count == 0)
                return NotFound(new ApiFailResponse
                {
                    Error = $"{Constants.NameNotFound}: {request.Name}"
                });

            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Data = _useCase.GetAll()
            });
        }
    }
}