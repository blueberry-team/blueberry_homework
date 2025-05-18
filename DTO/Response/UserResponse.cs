namespace BerryNameApi.DTO.Response
{
    public class UserResponse
    {
        public required Guid Id { get; set; }
        public required string Name { get; set; }
        public required DateTime CreatedAt { get; set; }
        public required DateTime UpdatedAt { get; set; }
    }
}