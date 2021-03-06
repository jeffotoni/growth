<?php

use Growth\Action;
use Growth\HTTP\Router;
use Swoole\HTTP\Request;
use Swoole\HTTP\Response;
use Swoole\HTTP\Server;
use Swoole\Process;
use Swoole\Table;

spl_autoload_register(function ($classname) {
    $filename = sprintf('./src/%s.php', strtr($classname, '\\', '/'));
    if (is_readable($filename)) {
        require $filename;
    }
});

set_error_handler(function ($number, $message, $file, $line) {
    throw new ErrorException($message, 0, $number, $file, $line);
});

$table = new Table(65536);
$table->column('content', Swoole\Table::TYPE_STRING, 128);
$table->create();

$router = new Router();

$router->map('GET', '^/ping', new Action\Ping());
$router->map('POST', '^/api/v1/growth$', new Action\Bulk($table));
$router->map('GET', '^/api/v1/growth/post/status$', new Action\Status($table));
$router->map('GET', '^/api/v1/growth/size$', new Action\Size($table));
$router->map('GET', '^/api/v1/growth/(?<country>[^/]+)/(?<indicator>[^/]+)/(?<year>[0-9]+)$', new Action\Find($table));
$router->map('PUT', '^/api/v1/growth/(?<country>[^/]+)/(?<indicator>[^/]+)/(?<year>[0-9]+)$', new Action\Create($table));
$router->map('DELETE', '^/api/v1/growth/(?<country>[^/]+)/(?<indicator>[^/]+)/(?<year>[0-9]+)$', new Action\Remove($table));

$server = new Server('0.0.0.0', 8080);

$server->set([
    'package_max_length' => 10 * 1024 * 1024, // 10MB
]);

$server->on('Start', function (Server $server) {
    echo sprintf("[%s] Server started at http://0.0.0.0:8080/\n", date('r'));
    Process::signal(SIGINT, fn() => $server->stop());
});

$server->on('Request', function (Request $request, Response $response) use ($router) {
    ($router->match($request))($request, $response);
});

$server->on('Shutdown', function ($server) {
    echo sprintf("[%s] Server stopped\n", date('r'));
});

$server->start();
