using BerryNameApi.Repositories;
using BerryNameApi.UseCases;
using blueberry_homework_dotnet.App;
using blueberry_homework_dotnet.Repositories;
using blueberry_homework_dotnet.UseCases;
using DotNetEnv;

var builder = WebApplication.CreateBuilder(args);

// Mongo DB
Env.Load();
AppInitializer.Init(builder.Services, builder.Configuration);

// 의존성
builder.Services.AddSingleton<NameRepository>();
builder.Services.AddTransient<NameUseCase>();
builder.Services.AddSingleton<CompanyRepository>();
builder.Services.AddTransient<CompanyUseCase>();

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

app.UseSwagger();
app.UseSwaggerUI();

app.MapControllers();
app.Run();