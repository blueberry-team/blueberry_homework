using Microsoft.AspNetCore.Identity.Data;

namespace blueberry_homework_dotnet.DTO.Request.Auth
{
    public class LogInRequest
    {
        public string Email { get; set; }
        public string Password { get; set; }

        public LogInRequest()
        {
            Email = string.Empty;
            Password = string.Empty;
        }
    }
}