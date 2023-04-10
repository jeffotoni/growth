using System.Collections.Concurrent;

namespace Growth;

public static class GrowthDatabase
{
    private static ConcurrentDictionary<string, GrowData> Data { get; } = new();

    public static Task? Process { get; set; }

    public static void AddOrUpdate(GrowData data) =>
        Data.AddOrUpdate(GetKey(data), data, (_, _) => data);

    public static bool TryGetValue(string country, string indicator, int year, out GrowData? data) =>
        Data.TryGetValue(GetKey(country, indicator, year), out data);

    public static void Remove(string country, string indicator, int year) =>
        Data.TryRemove(GetKey(country, indicator, year), out _);

    public static int Count => Data.Count;

    private static string GetKey(GrowData data) =>
        GetKey(data.Country, data.Indicator, data.Year);

    private static string GetKey(string country, string indicator, int year) =>
        $"{country}_{indicator}_{year}".ToLowerInvariant();
}

public record GrowData(string Country, string Indicator, double Value, int Year);
