using BerryNameApi.DTO.Request;
using BerryNameApi.DTO.Response;
using BerryNameApi.UseCases;
using Microsoft.AspNetCore.Mvc;
using BerryNameApi.Utils;
using blueberry_homework_dotnet.DTO.Request;

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
                    Error = Constants.NameLengthInvalid
                });

            var result = _useCase.CreateName(request.Name);

            if (!result.Success)
            {
                return BadRequest(new ApiFailResponse
                {
                    Error = result.ErrorMessage
                });
            }

            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Message = Constants.Success
            });

        }

        [HttpPut("changeName")]
        public IActionResult ChangeName([FromBody] ChangeNameRequest request)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(new ApiFailResponse
                {
                    Error = Constants.NameLengthInvalid
                });
            }

            var result = _useCase.ChangeName(request.Id!.Value, request.Name);
            if (!result.Success)
            {
                return BadRequest(new ApiFailResponse
                {
                    Error = result.ErrorMessage
                });
            }

            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Message = Constants.Success
            });
        }

        [HttpGet("getName")]
        public IActionResult GetNames()
        {
            // 처리 결과
            var result = _useCase.GetAll();
            if (result == null || result.Data == null)
            {
                return NotFound(new { message = Constants.Error, error = Constants.NameNotFound });
            }

            // 응답 데이터
            var names = result.Data;

            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Data = names
            });
        }

        [HttpDelete("deleteIndex")]
        public IActionResult DeleteByIndex([FromBody] DeleteIndexRequest request)
        {
            // index 필수 || 400 에러 노출
            // if (!request.Index.HasValue)
            //     return BadRequest(new ApiFailResponse
            //     {
            //         Error = Constants.DeleteIndexRequired
            //     });

            var result = _useCase.DeleteByIndex(request.Index);
            if (!result.Success)
                return BadRequest(new ApiFailResponse
                {
                    Error = result.ErrorMessage
                });

            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Message = Constants.Success
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

            var result = _useCase.DeleteByName(request.Name);
            if (!result.Success)
                return BadRequest(new ApiFailResponse
                {
                    Error = result.ErrorMessage
                });

            return Ok(new ApiSuccessResponse<IEnumerable<UserResponse>>
            {
                Message = Constants.Success
            });
        }
    }
}