<?php

use Swoole\HTTP\Request;
use Swoole\HTTP\Response;
use Swoole\HTTP\Server;

$server = new Server('0.0.0.0', 8080);

$server->on('Request', function (Request $request, Response $response) {
    $response->header('Content-Type', 'application/json');
    $response->end('{"message":"ping"}');
});

$server->start();
