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
                    Message = "error",
                    Error = "name must be between 1 and 50 characters"
                });

            _useCase.CreateName(request.Name);

            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Message = "success",
                Data = _useCase.GetAll()
            });

        }

        [HttpGet("getName")]
        public IActionResult Get()
        {
            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Message = "success",
                Data = _useCase.GetAll()
            });
        }

        [HttpDelete("deleteIndex")]
        public IActionResult DeleteByIndex([FromBody] DeleteIndexRequest request)
        {
            if (!request.Index.HasValue)
                return BadRequest(new ApiFailResponse
                {
                    Message = "error",
                    Error = "deleteIndex is required"
                });

            var deleted = _useCase.DeleteByIndex(request.Index.Value);
            if (!deleted)
                return NotFound(new ApiFailResponse
                {
                    Message = "error",
                    Error = $"Invalid index: {request.Index}"
                });

            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Message = "success",
                Data = _useCase.GetAll()
            });
        }

        [HttpDelete("deleteName")]
        public IActionResult DeleteByName([FromBody] DeleteNameRequest request)
        {
            if (!ModelState.IsValid)
                return BadRequest(new ApiFailResponse
                {
                    Message = "error",
                    Error = "name must be between 1 and 50 characters"
                });

            var count = _useCase.DeleteByName(request.Name);
            if (count == 0)
                return NotFound(new ApiFailResponse
                {
                    Message = "error",
                    Error = $"No user with name: {request.Name}"
                });

            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Message = "success",
                Data = _useCase.GetAll()
            });
        }
    }
}