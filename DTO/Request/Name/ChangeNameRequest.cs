using System.ComponentModel.DataAnnotations;

namespace blueberry_homework_dotnet.DTO.Request
{
    public class ChangeNameRequest
    {
        [Required]
        public Guid? Id { get; set; }

        [Required]
        [StringLength(150, MinimumLength = 1)]
        public string Name { get; set; }

        public ChangeNameRequest()
        {
            Name = string.Empty;
        }
    }
}