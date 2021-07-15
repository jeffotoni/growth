<?php

namespace Growth\Action;

class NotFound
{
    public function __invoke()
    {
        header('HTTP/1.1 404 Not Found');
    }
}
