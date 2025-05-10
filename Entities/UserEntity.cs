using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace BerryNameApi.Entities
{
    public class UserEntity
    {
        [BsonId]
        [BsonRepresentation(BsonType.String)]
        public Guid Id { get; set; }
        public string Name { get; set; }
        public DateTime CreatedAt { get; set; }
        public DateTime UpdatedAt { get; set; }

        public UserEntity()
        {
            Name = string.Empty;
        }
    }
}