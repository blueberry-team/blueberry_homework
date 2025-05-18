using System.ComponentModel.DataAnnotations;

namespace BerryNameApi.DTO.Request
{
    public class DeleteIndexRequest
    {
        [Required]
        public int Index { get; set; }
    }
}