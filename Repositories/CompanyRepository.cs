using blueberry_homework_dotnet.Entities;

namespace blueberry_homework_dotnet.Repositories
{
    public class CompanyRepository
    {
        private readonly List<CompanyEntity> _store = new();

        public IEnumerable<CompanyEntity> GetAll() => _store;

        public void CreateCompany(CompanyEntity company) => _store.Add(company);

        public CompanyEntity? FindByUserName(string name)
        {
            return _store.FirstOrDefault(company => company.Name == name);
        }
    }
}