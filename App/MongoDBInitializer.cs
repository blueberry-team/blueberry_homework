using BerryNameApi.Entities;
using blueberry_homework_dotnet.Entities;
using MongoDB.Driver;

namespace blueberry_homework_dotnet.App
{
    public class MongoDBInitializer
    {
        public static void Init(IServiceCollection services, IConfiguration configuration)
        {
            var mongoConnection = Environment.GetEnvironmentVariable("MongoConnectionString");
            var mongoDb = Environment.GetEnvironmentVariable("MongoDatabaseName");

            Console.WriteLine($"✅ Mongo 연결 경로 : {mongoConnection}");
            Console.WriteLine($"✅ Mongo DB 이름: {mongoDb}");

            services.AddSingleton<IMongoClient>(sp => new MongoClient(mongoConnection));
            services.AddSingleton(sp =>
            {
                var client = sp.GetRequiredService<IMongoClient>();
                var db = client.GetDatabase(mongoDb);

                // 현재 DB 로그

                var users = db.GetCollection<UserEntity>("users")
                                .Find(FilterDefinition<UserEntity>.Empty)
                                .ToList();

                var companies = db.GetCollection<CompanyEntity>("companies")
                                    .Find(FilterDefinition<CompanyEntity>.Empty)
                                    .ToList();

                Console.WriteLine("User List:");
                if (users.Count == 0)
                {
                    Console.WriteLine(" - 없음");
                }
                else
                {
                    foreach (var user in users)
                    {
                        Console.WriteLine($" - {user.Name} ({user.CreatedAt:O})");
                    }
                }
                Console.WriteLine("Company List:");
                if (companies.Count == 0)
                {
                    Console.WriteLine(" - 없음");
                }
                else
                {
                    foreach (var company in companies)
                    {
                        Console.WriteLine($" - {company.CompanyName} ({company.CreatedAt:O})");
                    }
                }

                return db;
            });
        }
    }
}