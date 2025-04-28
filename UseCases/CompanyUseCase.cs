using BerryNameApi.Repositories;
using BerryNameApi.Utils;
using blueberry_homework_dotnet.Entities;
using blueberry_homework_dotnet.Repositories;

namespace blueberry_homework_dotnet.UseCases
{
    public class CompanyUseCase
    {
        private readonly NameRepository _userRepository;
        private readonly CompanyRepository _companyRepository;

        public CompanyUseCase(NameRepository userRepository, CompanyRepository companyRepository)
        {
            _userRepository = userRepository;
            _companyRepository = companyRepository;
        }

        public Result CreateCompany(string userName, string companyName)
        {
            var user = _userRepository.FindByName(userName);
            if (user == null)
            {
                return Result.Fail(Constants.UserNotFound);
            }

            var existingCompany = _companyRepository.FindByUserName(userName);
            if (existingCompany != null)
            {
                return Result.Fail(Constants.UserAlreadyCompany);
            }

            var currentTime = DateTime.UtcNow;

            var company = new CompanyEntity
            {
                Id = Guid.NewGuid(),
                Name = userName,
                CompanyName = companyName,
                CreatedAt = currentTime
            };

            _companyRepository.CreateCompany(company);
            return Result.Ok();
        }

        public IEnumerable<CompanyEntity> GetAllCompanies()
        {
            return _companyRepository.GetAll();
        }
    }
}