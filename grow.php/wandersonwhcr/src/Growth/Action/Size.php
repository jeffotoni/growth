<?php

namespace Growth\Action;

class Size
{
    public function __invoke(): void
    {
        $count = 0;
        if (apcu_exists('growth-count')) {
            $count = apcu_fetch('growth-count');
        }

        header('HTTP/1.1 200 OK');
        header('Content-Type: application/json');

        file_put_contents('php://output', json_encode(['size' => $count]));
    }
}
