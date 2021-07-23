<?php

use Swoole\HTTP\Request;
use Swoole\HTTP\Response;
use Swoole\HTTP\Server;
use Swoole\Process;

$server = new Server('0.0.0.0', 8080);

$server->on('Start', function (Server $server) {
    echo sprintf("[%s] Server started at http://0.0.0.0:8080/\n", date('r'));
    Process::signal(SIGINT, fn() => $server->stop());
});

$server->on('Request', function (Request $request, Response $response) {
    $response->header('Content-Type', 'application/json');
    $response->end('{"message":"ping"}');
});

$server->on('Shutdown', function ($server) {
    echo sprintf("[%s] Server stopped\n", date('r'));
});

$server->start();
