<?php

namespace Growth\HTTP;

use Growth\Action\NotFound;
use Swoole\HTTP\Request;
use Swoole\HTTP\Response;

class Router
{
    /**
     * @var array<int, array<string, string>>
     */
    protected array $routes = [];

    public function map(string $method, string $uri, callable $action): self
    {
        $this->routes[] = [
            'method' => $method,
            'uri'    => $uri,
            'action' => $action,
        ];

        return $this;
    }

    public function match(Request $request): callable
    {
        foreach ($this->routes as $route) {
            if ('*' === $route['method'] || $request->getMethod() === $route['method']) {
                if (preg_match(sprintf('#%s#', $route['uri']), $request->server['request_uri'], $matches) === 1) {
                    return function (Request $request, Response $response) use ($route, $matches) {
                        ($route['action'])($request, $response, $matches);
                    };
                }
            }
        }

        return function (Request $request, Response $response) {
            (new NotFound())($request, $response);
        };
    }
}
