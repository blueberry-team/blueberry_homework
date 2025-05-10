using blueberry_homework_dotnet.Entities;
using MongoDB.Driver;

namespace blueberry_homework_dotnet.Repositories
{
    public class CompanyRepository
    {
        private readonly IMongoCollection<CompanyEntity> _collection;

        public CompanyRepository(IMongoDatabase database)
        {
            _collection = database.GetCollection<CompanyEntity>("companies");
        }

        public IEnumerable<CompanyEntity> GetAll()
        {
            return _collection.Find(FilterDefinition<CompanyEntity>.Empty).ToList();
        }

        public void CreateCompany(CompanyEntity company)
        {
            _collection.InsertOne(company);
        }

        public CompanyEntity? FindByUserName(string name)
        {
            return _collection.Find(c => c.Name == name).FirstOrDefault();
        }
    }
}