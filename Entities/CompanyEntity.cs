using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace blueberry_homework_dotnet.Entities
{
    public class CompanyEntity
    {
        [BsonId]
        [BsonRepresentation(BsonType.String)]
        public Guid Id { get; set; }
        public required string Name { get; set; }
        public required string CompanyName { get; set; }
        public DateTime CreatedAt { get; set; }
    }
}