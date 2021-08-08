#include <drogon/drogon.h>
int main()
{
    drogon::app().loadConfigFile("config.json");
    // Set HTTP listener address and port
    drogon::app().addListener("0.0.0.0", 8080);
    //drogon::app().setLogLevel(trantor::Logger::kWarn);
    drogon::app().setThreadNum(16);
    // Load config file
    // drogon::app().loadConfigFile("../config.json");
    // Run HTTP framework,the method will block in the internal event loop
    drogon::app().run();
    return 0;
}
