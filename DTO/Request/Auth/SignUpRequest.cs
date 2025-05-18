namespace blueberry_homework_dotnet.DTO.Request.Auth
{
    public class SignUpRequest
    {
        public string Email { get; set; } = "";
        public string Password { get; set; } = "";
        public string Name { get; set; } = "";
        public string Address { get; set; } = "";
        public string Role { get; set; } = "";
    }
}