using System.ComponentModel.DataAnnotations;

namespace blueberry_homework_dotnet.DTO.Request
{
    public class CreateCompanyRequest
    {
        public Guid UserId { get; set; }

        [Required]
        [StringLength(100, MinimumLength = 1)]
        public string CompanyName { get; set; }

        public string CompanyAddress { get; set; }

        public int TotalStaff { get; set; }

        public CreateCompanyRequest()
        {
            CompanyName = string.Empty;
            CompanyAddress = string.Empty;
        }
    }
}