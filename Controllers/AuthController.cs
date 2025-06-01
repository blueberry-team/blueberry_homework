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
        private readonly AuthUseCase _authUseCase;
        private readonly RefreshUseCase _refreshUseCase;

        public AuthController(AuthUseCase authUseCase, RefreshUseCase refreshUseCase)
        {
            _authUseCase = authUseCase;
            _refreshUseCase = refreshUseCase;
        }

        [HttpPost("sign-up")]
        public IActionResult SignUp([FromBody] SignUpRequest request)
        {
            var result = _authUseCase.SignUp(request);

            if (!result.Success)
            {
                return BadRequest(new ApiFailResponse { Error = result.ErrorMessage });
            }

            return Ok(new ApiSuccessResponse<IEnumerable<AuthResponse>> { Message = Constants.Success });
        }

        [HttpPost("log-in")]
        public IActionResult LogIn([FromBody] LogInRequest request)
        {
            var result = _authUseCase.LogIn(request);

            if (!result.Success || result.Data == null)
            {
                return BadRequest(new ApiFailResponse { Error = result.ErrorMessage });
            }

            var response = result.Data;

            return Ok(new ApiSuccessResponse<AuthResponse> { Message = Constants.Success, Data = response });
        }

        [HttpPatch("change-user")]
        public IActionResult ChangeUser([FromBody] ChangeUserRequest request)
        {
            var result = _authUseCase.ChangeUser(request);

            if (!result.Success)
            {
                return BadRequest(new ApiFailResponse { Error = result.ErrorMessage });
            }

            return Ok(new ApiSuccessResponse<IEnumerable<AuthResponse>> { Message = Constants.Success });
        }

        [HttpGet("get-user")]
        public IActionResult GetUser([FromQuery] Guid id)
        {
            var result = _authUseCase.GetUser(id);
            if (!result.Success || result.Data == null)
            {
                return NotFound(new ApiFailResponse { Error = Constants.UserNotFound });
            }

            var user = result.Data;

            var authResponse = new AuthResponse
            {
                Id = user.Id,
                Email = user.Email,
                Name = user.Name,
                Address = user.Address,
                Role = user.Role.ToString(),
                CreatedAt = user.CreatedAt,
                UpdatedAt = user.UpdatedAt,
                Token = string.Empty
            };

            return Ok(new ApiSuccessResponse<AuthResponse>
            {
                Message = Constants.Success,
                Data = authResponse
            });
        }

        [HttpPost("refresh-token")]
        public IActionResult RefreshToken()
        {
            var token = Request.Headers["Authorization"].ToString().Replace("Bearer ", "");

            var result = _refreshUseCase.Refresh(token);

            if (!result.Success)
            {
                return BadRequest(new ApiFailResponse { Error = result.ErrorMessage });
            }

            return Ok(new ApiSuccessResponse<AuthResponse>
            {
                Message = Constants.Success,
                Data = result.Data
            });
        }
    }
}