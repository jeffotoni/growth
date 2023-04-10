using System.Text.Json.Serialization;
using Growth;
using Growth.Requests;
using Growth.Responses;

var builder = WebApplication.CreateBuilder(args);

var app = builder.Build();

app.MapGet("/ping", () => Results.Json(new StatusResponse {Message = "pong"}, JsonContext.Default.Options));

app.MapGet("/api/v1/growth/size", () => GrowthDatabase.Count);

app.MapGet("/api/v1/growth/post/status", () =>
    Results.Json(GrowthDatabase.Process is {IsCompleted: true}
        ? new StatusResponse {Message = "complete", TestValue = 183.26, Count = GrowthDatabase.Count}
        : new StatusResponse {Message = "In progress"}, JsonContext.Default.Options));

app.MapPost("/api/v1/growth", (IAsyncEnumerable<GrowData> growData) =>
{
    GrowthDatabase.Process = Task.Run(async () =>
    {
        await foreach (var data in growData)
            GrowthDatabase.AddOrUpdate(data);
    });

    return Results.Json(new StatusResponse {Message = "In progress"}, JsonContext.Default.Options);
});

app.MapGet("/api/v1/growth/{country}/{indicator}/{year:int}",
    (string country, string indicator, int year) =>
        GrowthDatabase.TryGetValue(country, indicator, year, out var data)
            ? Results.Ok(data)
            : Results.NotFound());

app.MapDelete("/api/v1/growth/{country}/{indicator}/{year:int}", GrowthDatabase.Remove);

app.MapPut("/api/v1/growth/{country}/{indicator}/{year}",
    (string country, string indicator, int year, UpdateDataRequest request) =>
    {
        var data = new GrowData(country, indicator, request.Value, year);
        GrowthDatabase.AddOrUpdate(data with {Value = request.Value});
    });

app.Run();

[JsonSourceGenerationOptions(PropertyNamingPolicy = JsonKnownNamingPolicy.CamelCase)]
[JsonSerializable(typeof(StatusResponse))]
[JsonSerializable(typeof(GrowData))]
internal partial class JsonContext : JsonSerializerContext
{
}
