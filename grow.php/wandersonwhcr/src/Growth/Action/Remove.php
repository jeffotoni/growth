<?php

namespace Growth\Action;

use Swoole\HTTP\Request;
use Swoole\HTTP\Response;
use Swoole\Table;

class Remove
{
    public function __construct(private Table $table)
    {
    }

    public function __invoke(Request $request, Response $response, array $args): void
    {
        $key = sprintf('growth-%s-%s-%s', $args['country'], $args['indicator'], $args['year']);

        if (! $this->table->exists($key)) {
            $response->status(404);
            $response->end();
            return;
        }

        $this->table->del($key);

        $response->status(204);
        $response->end();
    }
}
