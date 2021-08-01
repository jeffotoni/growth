using System.Collections.Generic;
using System.Threading.Tasks;

public static class GrowDatabase
{
    public static List<GrowData> Data { get; } = new();
    public static Task Process { get; set; }
}

public record GrowData(string Country, string Indicator, decimal Value, int Year);