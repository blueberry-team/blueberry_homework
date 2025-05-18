using System.ComponentModel.DataAnnotations;

namespace BerryNameApi.DTO.Request
{
    public class DeleteNameRequest
    {
        [Required]
        [StringLength(50, MinimumLength = 1)]
        public required string Name { get; set; }
    }
}