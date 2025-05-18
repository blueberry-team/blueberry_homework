using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace BerryNameApi.Entities
{
    public class UserEntity
    {
        [BsonId]
        [BsonRepresentation(BsonType.String)]
        public Guid Id { get; set; }

        public string Email { get; set; } = string.Empty;
        public string PasswordHashed { get; set; } = string.Empty;
        public string Name { get; set; } = string.Empty;
        public string Address { get; set; } = string.Empty;

        [BsonRepresentation(BsonType.String)]
        public Role Role { get; set; } = Role.Worker;

        public DateTime CreatedAt { get; set; }
        public DateTime UpdatedAt { get; set; }

    }

    public enum Role
    {
        Boss,
        Worker
    }
}