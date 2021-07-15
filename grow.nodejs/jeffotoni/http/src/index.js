var http = require("http");
const port = 8080;
var map = new Map();

const requestListener = (RequestListener = (req, res) => {
  switch (req.url) {
    case "/ping":
      return res.end("pong😍");
    case "/api/v1/growth":
      return Route(req, res);
    case "/api/v1/growth/post/status":
      return GetStatus(req, res);
    case "/api/v1/growth/size":
      return GetSize(req, res);
    default:
      return Route(req, res);
  }
});

const Route = (RequestListener = (req, res) => {
  switch (req.method) {
    case "GET":
      return RouteGet(req, res);
    case "DELETE":
      return RouteDelete(req, res);
    case "PUT":
      return RoutePut(req, res);
    case "POST":
      return RoutePost(req, res);
    default:
      res.writeHead(404);
      res.end();
  }
});

const routeKey = (req = http.IncomingMessage) => {
  return req.url
    .replace(/^\/api\/v1\/growth\//, "")
    .replace(/\//g, "")
    .toUpperCase();
};


const GetSize = RequestListener = (req, res) => {
  const count = Object.keys(mapGrow).length;
  res.writeHead(200, JSON.stringify({ count: count.toString() }));
  res.end();
};

const RoutePost = (RequestListener = (req, res) => {
  let body = "";
  req.on("data", (chunk) => {
    body += chunk;
  });
  req.on("end", () => {
    var array = [];
    try {
      array = JSON.parse(body);
    } catch (e) {
      res.writeHead(400, JSON.stringify({ msg: "error in your json" }), {
        "Content-Type": "application/json",
      });
      return res.end();
    }
    var i = 0;
    array.forEach(function (object) {
      var key =
        object.Country.toUpperCase() +
        object.Indicator.toUpperCase() +
        object.Year;
      var value = object.Value;
      map.set(key, value);
     // console.log(key);
      i++;
    });

    res.writeHead(202, JSON.stringify({ msg: "In progress" }), {
      "Content-Type": "application/json",
    });
    res.end();
  });
});

const RouteGet = RequestListener = (req, res) => {
  const key = routeKey(req);
  //console.log(key)
  const val = map.get(key);
  if (!val) {
    res.writeHead(400, JSON.stringify({ msg: "not found" }), {
      "Content-Type": "application/json",
    });
    return res.end();
  }

  res.writeHead(200, JSON.stringify({Value: val}));
  res.end();
};

const server = http.createServer(requestListener);
console.log(
  "\x1b[36m%s\x1b[0m",
  `Running on http://0.0.0.0:${port} (Press CTRL+C to quit)`
);
server.listen(port);
