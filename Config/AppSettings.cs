namespace blueberry_homework_dotnet.Config
{
    public class AppSettings
    {
        public string MongoConnectionString { get; set; }
        public string MongoDatabaseName { get; set; }

        public AppSettings()
        {
            MongoConnectionString = string.Empty;
            MongoDatabaseName = string.Empty;
        }
    }
}