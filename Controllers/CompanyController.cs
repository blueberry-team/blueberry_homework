using BerryNameApi.DTO.Response;
using BerryNameApi.Utils;
using blueberry_homework_dotnet.DTO.Request;
using blueberry_homework_dotnet.Entities;
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

            var result = _useCase.CreateCompany(request.UserName, request.CompanyName);
            if (!result.Success)
            {
                return BadRequest(new ApiFailResponse
                {
                    Message = Constants.Error,
                    Error = result.ErrorMessage
                });
            }

            return Ok(new ApiSuccessResponse<object> { Message = Constants.Success });
        }

        [HttpGet("getCompany")]
        public IActionResult GetCompany()
        {
            var companies = _useCase.GetAllCompanies();
            return Ok(new ApiSuccessResponse<IEnumerable<CompanyEntity>>
            {
                Message = Constants.Success,
                Data = companies
            });
        }
    }
}