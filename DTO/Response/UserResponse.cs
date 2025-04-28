namespace BerryNameApi.DTO.Response
{
    public class UserResponse
    {
        public required Guid Id { get; set; }
        public required string Name { get; set; }
        public DateTime CreatedAt { get; set; }
        public DateTime UpdatedAt { get; set; }
    }
}