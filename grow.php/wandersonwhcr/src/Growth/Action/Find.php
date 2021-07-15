<?php

namespace Growth\Action;

class Find
{
    protected array $server;

    public function __invoke(array $args): void
    {
        $key = sprintf('growth-%s-%s-%s', $args['country'], $args['indicator'], $args['year']);

        if (! apcu_exists($key)) {
            header('HTTP/1.1 404 Not Found');
            return;
        }

        $data = apcu_fetch($key);

        header('HTTP/1.1 200 OK');
        header('Content-Type: application/json');

        file_put_contents('php://output', json_encode($data));
    }
}
