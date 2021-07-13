import http, { RequestListener } from 'http';

interface DataGrowth {
    Country: string;
    Indicator: string;
    Value: number;
    Year: number;
}

const dataGrowthKey = (d: DataGrowth) => `${d.Country}${d.Indicator}${d.Year}`.toUpperCase();

const routeKey = (req: http.IncomingMessage): string => {
  // /api/v1/growth/brz/ngdp_r/2002 => BRZNGDP_R2002
  return req.url
    .replace(/^\/api\/v1\/growth\//, '')
    .replace(/\//g, '')
    .toUpperCase();
}

const mapGrow: { [key: string]: DataGrowth } = {};

const GetStatus: RequestListener = (req, res) => {
  if (!('BRZNGDP_R2002' in mapGrow)) {
    res.writeHead(400, JSON.stringify({ msg: 'not finished'}), { 'Content-Type': 'application/json' });
    return res.end();
  }

  const count = Object.keys(mapGrow).length;

  const message = JSON.stringify({
    msg: 'complete',
    'test value': count.toFixed(2),
    count: count.toString(),
  });

  res.writeHead(200, message, { 'Content-Type': 'application/json' });
  res.end();
}

const GetSize: RequestListener = (req, res) => {
  const count = Object.keys(mapGrow).length;
  res.writeHead(200, JSON.stringify({ count: count.toString() }));
  res.end();
}

const RoutePut: RequestListener = (req, res) => {
  const key = routeKey(req);

  if (!(key in mapGrow)) {
    res.writeHead(400, JSON.stringify({ msg: 'not found'}), { 'Content-Type': 'application/json' });
    return res.end();
  }

  let body = '';
  req.on('data', (chunk) => {
    body += chunk;
  });
  req.on('end', () => {
    let input: DataGrowth;

    try {
      input = JSON.parse(body);
    } catch (e) {
      res.writeHead(400, JSON.stringify({ msg: 'error in your json'}), { 'Content-Type': 'application/json' });
      return res.end();
    }

    mapGrow[key] = input;
    res.writeHead(200, JSON.stringify({ msg: 'ok'}), { 'Content-Type': 'application/json' });
    res.end();
  });
}

const RouteDelete: RequestListener = (req, res) => {
  const key = routeKey(req);

  if (!(key in mapGrow)) {
    res.writeHead(400, JSON.stringify({ msg: 'not found'}), { 'Content-Type': 'application/json' });
    return res.end();
  }

  delete mapGrow[key];
  res.writeHead(200, JSON.stringify({ msg: 'deleted'}), { 'Content-Type': 'application/json' });
  res.end();
};

const RouteGet: RequestListener = (req, res) => {
  const key = routeKey(req);

  if (!(key in mapGrow)) {
    res.writeHead(400, JSON.stringify({ msg: 'not found'}), { 'Content-Type': 'application/json' });
    return res.end();
  }

  const d = mapGrow[key];
  res.writeHead(200, JSON.stringify(d));
  res.end();
};

const RoutePost: RequestListener = (req, res) => {
  let body = '';
  req.on('data', (chunk) => {
    body += chunk;
  });
  req.on('end', () => {
    let input: DataGrowth[];

    try {
      input = JSON.parse(body);
    } catch (e) {
      res.writeHead(400, JSON.stringify({ msg: 'error in your json'}), { 'Content-Type': 'application/json' });
      return res.end();
    }

    input.forEach((dataGrowth) => {
      const id = dataGrowthKey(dataGrowth);
      mapGrow[id] = dataGrowth;
    });

    res.writeHead(202, JSON.stringify({ msg: 'In progress' }), { 'Content-Type': 'application/json' });
    res.end();
  });
}

const Route: RequestListener = (req, res) => {
  switch (req.method) {
    case 'GET':
      return RouteGet(req, res);
    case 'DELETE':
      return RouteDelete(req, res);
    case 'PUT':
      return RoutePut(req, res);
    case 'POST':
      return RoutePost(req, res);
    default:
      res.writeHead(404);
      res.end();
  }
}

const requestListener: RequestListener = (req, res) => {
  switch (req.url) {
    case '/ping':
      return res.end('pong😍');
    case '/api/v1/growth':
      return Route(req, res);
    case '/api/v1/growth/post/status':
      return GetStatus(req, res);
    case '/api/v1/growth/size':
      return GetSize(req, res);
    default:
      return Route(req, res);
  }
}

const serverPort = 8080;
const server = http.createServer(requestListener);
console.log('\x1b[36m%s\x1b[0m', `Running on http://0.0.0.0:${serverPort} (Press CTRL+C to quit)`);
server.listen(serverPort);

