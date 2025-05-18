using BerryNameApi.Entities;
using MongoDB.Driver;

namespace blueberry_homework_dotnet.Repositories
{
    public class AuthRepository
    {
        private readonly IMongoCollection<UserEntity> _collection;

        public AuthRepository(IMongoDatabase db)
        {
            _collection = db.GetCollection<UserEntity>("users");
        }

        public void CreateUser(UserEntity user)
        {
            _collection.InsertOne(user);
        }

        public UserEntity? FindByEmail(string email)
        {
            return _collection.Find(user => user.Email == email).FirstOrDefault();
        }

        public UserEntity? FindById(Guid id)
        {
            return _collection.Find(user => user.Id == id).FirstOrDefault();
        }

        public void UpdateUser(UserEntity user)
        {
            _collection.ReplaceOne(user => user.Id == user.Id, user);
        }
    }
}