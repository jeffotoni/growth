<?php

namespace Growth\Action;

class Remove
{
    protected array $server;

    public function __invoke(array $args): void
    {
        $key = sprintf('growth-%s-%s-%s', $args['country'], $args['indicator'], $args['year']);

        if (! apcu_exists($key)) {
            header('HTTP/1.1 404 Not Found');
            return;
        }

        apcu_delete($key);
        apcu_dec('growth-count');

        header('HTTP/1.1 204 No Content');
    }
}
