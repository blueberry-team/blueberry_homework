namespace blueberry_homework_dotnet.DTO.Request.Auth
{
    public class SignUpRequest
    {
        public string Email { get; set; }
        public string Password { get; set; }
        public string Name { get; set; }
        public string Address { get; set; }
        public string Role { get; set; }

        public SignUpRequest()
        {
            Email = string.Empty;
            Password = string.Empty;
            Name = string.Empty;
            Address = string.Empty;
            Role = string.Empty;
        }
    }
}