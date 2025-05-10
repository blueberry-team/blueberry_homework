using blueberry_homework_dotnet.Config;
using Microsoft.Extensions.Options;
using MongoDB.Driver;

namespace blueberry_homework_dotnet.App
{
    public class MongoDBInitializer
    {
        public static void Init(IServiceCollection services, IConfiguration configuration)
        {
            services.Configure<AppSettings>(configuration);

            services.AddSingleton<IMongoClient>(sp =>
            {
                var settings = sp.GetRequiredService<IOptions<AppSettings>>().Value;
                return new MongoClient(settings.MongoConnectionString);
            });

            services.AddSingleton<IMongoDatabase>(sp =>
            {
                var settings = sp.GetRequiredService<IOptions<AppSettings>>().Value;
                var client = sp.GetRequiredService<IMongoClient>();
                return client.GetDatabase(settings.MongoDatabaseName);
            });
        }
    }
}