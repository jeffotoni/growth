using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.OpenApi.Models;
using System.Linq;
using System.Threading.Tasks;
using static System.StringComparison;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
var app = builder.Build();

if (app.Environment.IsDevelopment())
    app.UseDeveloperExceptionPage();

app.UseSwagger();

app.MapGet("/api/v1/growth/size", () => GrowDatabase.Data.Count);

app.MapGet("/api/v1/growth/post/status", () =>
GrowDatabase.Process.IsCompleted 
    ? new { msg = "complete", testValue = 183.26, count = GrowDatabase.Data.Count }
    : (object)new { msg = "in progress" }
);

app.MapPost("/api/v1/growth", ([FromBody] GrowData[] growData) =>
{
    GrowDatabase.Process = Task.Run(() => GrowDatabase.Data.AddRange(growData));
    return new { msg = "in progress" };
});

app.MapGet("/api/v1/growth/{country}/{indicator}/{year}", (string country, string indicator, int year, HttpContext ctx) =>
{
    var data = GrowDatabase.Data.FirstOrDefault(g => g.Country.Equals(country, OrdinalIgnoreCase)
                                                 && g.Indicator.Equals(indicator, OrdinalIgnoreCase)
                                                 && g.Year == year);
    if (data is null)
        ctx.Response.StatusCode = 404;
    else
        ctx.Response.WriteAsJsonAsync(data);
});

app.MapDelete("/api/v1/growth/{country}/{indicator}/{year}", (string country, string indicator, int year) =>
{
    var dataToRemove = GrowDatabase.Data.FirstOrDefault(g => g.Country.Equals(country, OrdinalIgnoreCase)
                                                   && g.Indicator.Equals(indicator, OrdinalIgnoreCase)
                                                   && g.Year == year);
    if (dataToRemove is not null)
        GrowDatabase.Data.Remove(dataToRemove);
});

app.MapPut("/api/v1/growth/{country}/{indicator}/{year}", (string country, string indicator, int year, GrowData growData) =>
{
     var index = GrowDatabase.Data.FindIndex(g => g.Country.Equals(country, OrdinalIgnoreCase)
                                                 && g.Indicator.Equals(indicator, OrdinalIgnoreCase)
                                                 && g.Year == year);

     if (index == -1)
     {
         growData = growData with
         {
             Country = country,
             Indicator = indicator,
             Year = year
         };
         GrowDatabase.Data.Add(growData);
     }
     else
     {
         var current = GrowDatabase.Data[index];
         GrowDatabase.Data[index] = current with { Value = growData.Value };
     }
});

app.UseSwaggerUI();

app.Run();
