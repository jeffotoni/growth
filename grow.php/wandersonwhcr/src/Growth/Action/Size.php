<?php

namespace Growth\Action;

use Swoole\HTTP\Request;
use Swoole\HTTP\Response;
use Swoole\Table;

class Size
{
    public function __construct(private Table $table)
    {
    }

    public function __invoke(Request $request, Response $response): void
    {
        $response->header('content-type', 'application/json');
        $response->end(json_encode([
            'size' => $this->table->count(),
        ]));
    }
}
