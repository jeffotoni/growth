<?php

namespace Growth\Action;

use Swoole\HTTP\Request;
use Swoole\HTTP\Response;
use Swoole\Table;

class Status
{
    public function __construct(private Table $table)
    {
    }

    public function __invoke(Request $request, Response $response): void
    {
        $response->end(json_encode([
            'msg' => 'complete', // TODO Async
            'count' => $this->table->count(),
        ]));
    }
}
