<?php

namespace Growth\Action;

class Status
{
    public function __invoke(): void
    {
        $count = 0;
        if (apcu_exists('growth-count')) {
            $count = apcu_fetch('growth-count');
        }

        file_put_contents('php://output', json_encode([
            'msg' => 'complete', // TODO Async
            'count' => $count,
        ]));
    }
}
