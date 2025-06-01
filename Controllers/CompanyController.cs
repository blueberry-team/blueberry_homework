using BerryNameApi.DTO.Response;
using BerryNameApi.Utils;
using blueberry_homework_dotnet.DTO.Request;
using blueberry_homework_dotnet.DTO.Response;
using blueberry_homework_dotnet.UseCases;
using Microsoft.AspNetCore.Mvc;

namespace blueberry_homework_dotnet.Controllers
{

    [ApiController]
    [Route("companies")]
    public class CompanyController : ControllerBase
    {
        private readonly CompanyUseCase _useCase;

        public CompanyController(CompanyUseCase useCase)
        {
            _useCase = useCase;
        }

        [HttpPost("createCompany")]
        public IActionResult CreateCompany([FromBody] CreateCompanyRequest request)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(new ApiFailResponse
                {
                    Error = Constants.NameLengthInvalid
                });
            }

            var result = _useCase.CreateCompany(request);
            if (!result.Success)
            {
                return BadRequest(new ApiFailResponse
                {
                    Message = Constants.Error,
                    Error = result.ErrorMessage
                });
            }

            return Ok(new ApiSuccessResponse<IEnumerable<CompanyResponse>> { Message = Constants.Success });
        }

        [HttpGet("getCompany")]
        public IActionResult GetCompany([FromQuery] Guid userId)
        {
            // 처리 결과
            var result = _useCase.GetCompany(userId);
            if (result == null || result.Data == null)
            {
                return NotFound(new { message = Constants.Error, error = Constants.CompanyNotFound });
            }

            // 응답 데이터
            var company = result.Data;

            var companyResponse = new CompanyResponse
            {
                Id = company.Id,
                CompanyName = company.CompanyName,
                CompanyAddress = company.CompanyAddress,
                TotalStaff = company.TotalStaff,
                CreatedAt = company.CreatedAt,
                UpdatedAt = company.UpdatedAt
            };

            return Ok(new ApiSuccessResponse<CompanyResponse>
            {
                Message = Constants.Success,
                Data = companyResponse
            });
        }

        [HttpPatch("changeCompany")]
        public IActionResult ChangeCompany([FromQuery] Guid userId, [FromBody] CreateCompanyRequest body)
        {
            var result = _useCase.ChangeCompany(userId, body.CompanyName, body.CompanyAddress, body.TotalStaff);

            if (!result.Success)
            {
                return BadRequest(new ApiFailResponse { Error = result.ErrorMessage });
            }

            return Ok(new ApiSuccessResponse<IEnumerable<CompanyResponse>> { Message = Constants.Success });
        }
    }
}