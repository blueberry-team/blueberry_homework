namespace blueberry_homework_dotnet.DTO.Request.Auth
{
    public class ChangeUserRequest
    {
        public Guid Id { get; set; }
        public string? Password { get; set; }
        public string? Address { get; set; }
    }
}