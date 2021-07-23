<?php

namespace Growth\Action;

use Swoole\HTTP\Request;
use Swoole\HTTP\Response;

class NotFound
{
    public function __invoke(Request $request, Response $response): void
    {
        $response->status(404);
        $response->end();
    }
}
