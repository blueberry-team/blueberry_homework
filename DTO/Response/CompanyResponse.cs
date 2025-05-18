namespace blueberry_homework_dotnet.DTO.Response
{
    public class CompanyResponse
    {
        public required Guid Id { get; set; }
        public required string CompanyName { get; set; }
        public required string CompanyAddress { get; set; }
        public required int TotalStaff { get; set; }
        public required DateTime CreatedAt { get; set; }
        public required DateTime UpdatedAt { get; set; }
    }
}