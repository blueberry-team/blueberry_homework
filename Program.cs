using BerryNameApi.Repositories;
using BerryNameApi.UseCases;

var builder = WebApplication.CreateBuilder(args);

// 의존성
builder.Services.AddSingleton<NameRepository>();
builder.Services.AddTransient<NameUseCase>();

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

app.UseSwagger();
app.UseSwaggerUI();

app.MapControllers();
app.Run();

