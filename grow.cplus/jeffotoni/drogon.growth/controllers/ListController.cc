#include <drogon/drogon.h>

using namespace drogon;

HttpResponsePtr makeFailedResponse()
{
    Json::Value json;
    json["ok"] = false;
    auto resp = HttpResponse::newHttpJsonResponse(json);
    resp->setStatusCode(k500InternalServerError);
    return resp;
}

HttpResponsePtr makeSuccessResponse()
{
    Json::Value json;
    json["ok"] = true;
    auto resp = HttpResponse::newHttpJsonResponse(json);
    return resp;
}

std::string getRandomString(size_t n)
{
    std::vector<unsigned char> random(n);
    utils::secureRandomBytes(random.data(), random.size());

    // This is cryptographically safe as 256 mod 16 == 0
    static const std::string alphabets = "0123456789abcdef";
    assert(256 % alphabets.size() == 0);
    std::string randomString(n, '\0');
    for (size_t i = 0; i < n; i++)
        randomString[i] = alphabets[random[i] % alphabets.size()];
    return randomString;
}

struct DataItem
{
    Json::Value item;
    std::mutex mtx;
};

class JsonStore : public HttpController<JsonStore>
{
  public:
    METHOD_LIST_BEGIN
    ADD_METHOD_TO(JsonStore::createItem, "/api/v1/growth", Post);
    METHOD_LIST_END

    void createItem(const HttpRequestPtr& req,
                  std::function<void(const HttpResponsePtr&)>&& callback)
    {   
        if(req->jsonObject())
        {
            auto &root=*(req->jsonObject());
            {
                std::string key = "";
                for(auto &item: root){
                   // std::cout << item["Country"] << "\n";
                    key = item["Country"].asString() + item["Indicator"].asString() + item["Year"].asString();
                    auto itemVal = std::make_shared<DataItem>();
                    itemVal->item = std::move(item["Value"]);
                    dataStore_.insert({key, std::move(itemVal)});
                }
            }
           // co_return;
        }

        callback(makeSuccessResponse());
    }

    std::unordered_map<std::string, std::shared_ptr<DataItem>> dataStore_;
    std::mutex storageMtx_;
};
