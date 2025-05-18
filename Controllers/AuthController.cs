using BerryNameApi.DTO.Response;
using BerryNameApi.Utils;
using blueberry_homework_dotnet.DTO.Request.Auth;
using blueberry_homework_dotnet.DTO.Response;
using blueberry_homework_dotnet.UseCases;
using Microsoft.AspNetCore.Mvc;

namespace blueberry_homework_dotnet.Controllers
{
    [ApiController]
    [Route("auth")]
    public class AuthController : ControllerBase
    {
        private readonly AuthUseCase _useCase;

        public AuthController(AuthUseCase useCase)
        {
            _useCase = useCase;
        }

        [HttpPost("sign-up")]
        public IActionResult SignUp([FromBody] SignUpRequest request)
        {
            var result = _useCase.SignUp(request);

            if (!result.Success)
            {
                return BadRequest(new ApiFailResponse { Error = result.ErrorMessage });
            }

            return Ok(new ApiSuccessResponse<IEnumerable<AuthResponse>> { Message = Constants.Success });
        }

        [HttpPost("log-in")]
        public IActionResult LogIn([FromBody] LogInRequest request)
        {
            var result = _useCase.LogIn(request);

            if (!result.Success)
            {
                return BadRequest(new ApiFailResponse { Error = result.ErrorMessage });
            }

            return Ok(new ApiSuccessResponse<IEnumerable<AuthResponse>> { Message = Constants.Success });
        }

        [HttpPatch("change-user")]
        public IActionResult ChangeUser([FromBody] ChangeUserRequest request)
        {
            var result = _useCase.ChangeUser(request);

            if (!result.Success)
            {
                return BadRequest(new ApiFailResponse { Error = result.ErrorMessage });
            }

            return Ok(new ApiSuccessResponse<IEnumerable<AuthResponse>> { Message = Constants.Success });
        }

        [HttpGet("get-user")]
        public IActionResult GetUser([FromQuery] Guid id)
        {
            var user = _useCase.GetUser(id);

            if (user == null)
            {
                return NotFound(new ApiFailResponse { Error = Constants.UserNotFound });
            }

            var authResponse = new AuthResponse
            {
                Id = user.Id,
                Email = user.Email,
                Name = user.Name,
                Address = user.Address,
                Role = user.Role.ToString(),
                CreatedAt = user.CreatedAt,
                UpdatedAt = user.UpdatedAt
            };

            return Ok(new ApiSuccessResponse<IEnumerable<AuthResponse>>
            {
                Message = Constants.Success,
                Data = [authResponse]
            });
        }
    }
}