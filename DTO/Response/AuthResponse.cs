namespace blueberry_homework_dotnet.DTO.Response
{
    public class AuthResponse
    {
        public required Guid Id { get; set; }
        public required string Email { get; set; }
        public required string Name { get; set; }
        public required string Address { get; set; }
        public required string Role { get; set; }
        public required DateTime CreatedAt { get; set; }
        public required DateTime UpdatedAt { get; set; }
    }


}