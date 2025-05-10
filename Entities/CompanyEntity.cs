namespace blueberry_homework_dotnet.Entities
{
    public class CompanyEntity
    {
        public Guid Id { get; set; }
        public required string Name { get; set; }
        public required string CompanyName { get; set; }
        public DateTime CreatedAt { get; set; }
    }
}