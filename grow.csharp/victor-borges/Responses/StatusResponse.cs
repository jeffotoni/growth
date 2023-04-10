namespace Growth.Responses;

internal readonly record struct StatusResponse(string Message, int Count, double TestValue);
