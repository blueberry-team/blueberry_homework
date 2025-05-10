using System.ComponentModel.DataAnnotations;

namespace blueberry_homework_dotnet.DTO.Request
{
    public class CreateCompanyRequest
    {
        [Required]
        [StringLength(50, MinimumLength = 1)]
        public string UserName { get; set; }

        [Required]
        [StringLength(100, MinimumLength = 1)]
        public string CompanyName { get; set; }

        public CreateCompanyRequest()
        {
            UserName = string.Empty;
            CompanyName = string.Empty;
        }
    }
}