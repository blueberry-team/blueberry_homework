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

        public void CreateCompany(CompanyEntity company)
        {
            _collection.InsertOne(company);
        }

        public CompanyEntity? FindByUserId(Guid userId)
        {
            return _collection.Find(company => company.UserId == userId).FirstOrDefault();
        }

        public void UpdateCompany(CompanyEntity company)
        {
            _collection.ReplaceOne(company => company.Id == company.Id, company);
        }

        public void DeleteCompany(Guid id)
        {
            _collection.DeleteOne(company => company.Id == id);
        }

    }
}