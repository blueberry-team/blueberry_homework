using BerryNameApi.Entities;
using MongoDB.Driver;

namespace BerryNameApi.Repositories
{
    public class NameRepository
    {
        private readonly IMongoCollection<UserEntity> _collection;

        public NameRepository(IMongoDatabase db)
        {
            _collection = db.GetCollection<UserEntity>("users");
        }
        public IEnumerable<UserEntity> GetAll() => _collection.Find(FilterDefinition<UserEntity>.Empty).ToList();

        // 생성
        public void CreateName(UserEntity user) => _collection.InsertOne(user);

        // 중복 검색 ID
        public UserEntity? FindById(Guid id)
        {
            return _collection.Find(user => user.Id == id).FirstOrDefault();
        }

        // 중복 검색 Name
        public UserEntity? FindByName(string name) => _collection.Find(user => user.Name == name).FirstOrDefault();

        // 이름 변경
        public bool ChangeName(UserEntity user)
        {
            var result = _collection.ReplaceOne(u => u.Id == user.Id, user);
            return result.IsAcknowledged && result.ModifiedCount > 0;
        }

        // 인덱스로 삭제
        public bool DeleteByIndex(int index)
        {
            var all = _collection.Find(FilterDefinition<UserEntity>.Empty).ToList();

            if (index < 0 || index >= all.Count)
                return false;

            var target = all[index];
            var result = _collection.DeleteOne(u => u.Id == target.Id);
            return result.DeletedCount == 1;
        }

        // 이름으로 삭제 (삭제 수 반환)
        public long DeleteByName(string name)
        {
            var result = _collection.DeleteMany(u => u.Name == name);
            return result.DeletedCount;
        }
    }
}