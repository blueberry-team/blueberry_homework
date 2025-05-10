namespace BerryNameApi.DTO.Response
{
    public class ApiFailResponse
    {
        public string Message { get; set; } = "error";
        public string? Error { get; set; }
    }
}