using BerryNameApi.Repositories;
using BerryNameApi.UseCases;
using blueberry_homework_dotnet.App;
using blueberry_homework_dotnet.Repositories;
using blueberry_homework_dotnet.UseCases;

var builder = WebApplication.CreateBuilder(args);

// 의존성
builder.Services.AddSingleton<NameRepository>();
builder.Services.AddTransient<NameUseCase>();
builder.Services.AddSingleton<CompanyRepository>();
builder.Services.AddTransient<CompanyUseCase>();

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// Mongo DB
AppInitializer.Init(builder.Services, builder.Configuration);

var app = builder.Build();

app.UseSwagger();
app.UseSwaggerUI();

app.MapControllers();
app.Run();

