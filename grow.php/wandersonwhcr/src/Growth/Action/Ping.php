<?php

namespace Growth\Action;

use Swoole\HTTP\Request;
use Swoole\HTTP\Response;

class Ping
{
    public function __invoke(Request $request, Response $response): void
    {
        $response->status(204);
        $response->end();
    }
}
