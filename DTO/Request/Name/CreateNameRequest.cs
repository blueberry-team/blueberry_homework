using System.ComponentModel.DataAnnotations;

namespace BerryNameApi.DTO.Request
{
    public class CreateNameRequest
    {
        [Required]
        [StringLength(50, MinimumLength = 1)]
        public string Name { get; set; }

        public CreateNameRequest()
        {
            Name = string.Empty;
        }
    }
}