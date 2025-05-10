namespace blueberry_homework_dotnet.DTO.Response
{
    public class CompanyResponse
    {
        public Guid Id { get; set; }
        public string Name { get; set; }
        public string CompanyName { get; set; }
        public DateTime CreatedAt { get; set; }

        public CompanyResponse()
        {
            Name = string.Empty;
            CompanyName = string.Empty;
        }
    }
}