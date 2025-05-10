using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace blueberry_homework_dotnet.App
{
    public class AppInitializer
    {
        public static void Init(IServiceCollection services, IConfiguration configuration)
        {
            MongoDBInitializer.Init(services, configuration);
        }
    }
}