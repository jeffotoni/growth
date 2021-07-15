<?php

spl_autoload_register(function ($classname) {
    $filename = sprintf('./src/%s.php', strtr($classname, '\\', '/'));
    if (is_readable($filename)) {
        require $filename;
    }
});

set_error_handler(function ($number, $message, $file, $line) {
    throw new ErrorException($message, 0, $number, $file, $line);
});

$routes = [
    [
        'method' => 'POST',
        'uri'    => '^/api/v1/growth$',
        'action' => Growth\Action\Bulk::class,
    ],
    [
        'method' => 'GET',
        'uri'    => '^/api/v1/growth/post/status$',
        'action' => Growth\Action\Status::class,
    ],
    [
        'method' => 'GET',
        'uri'    => '^/api/v1/growth/size$',
        'action' => Growth\Action\Size::class,
    ],
    [
        'method' => 'GET',
        'uri'    => '^/api/v1/growth/(?<country>[^/]+)/(?<indicator>[^/]+)/(?<year>[0-9]+)$',
        'action' => Growth\Action\Find::class,
    ],
    [
        'method' => 'PUT',
        'uri'    => '^/api/v1/growth/(?<country>[^/]+)/(?<indicator>[^/]+)/(?<year>[0-9]+)$',
        'action' => Growth\Action\Create::class,
    ],
    [
        'method' => 'DELETE',
        'uri'    => '^/api/v1/growth/(?<country>[^/]+)/(?<indicator>[^/]+)/(?<year>[0-9]+)$',
        'action' => Growth\Action\Delete::class,
    ],
    [
        'method' => '*',
        'uri'    => '^.*$',
        'action' => Growth\Action\NotFound::class,
    ],
];

try {
    foreach ($routes as $route) {
        if ('*' === $route['method'] || $_SERVER['REQUEST_METHOD'] === $route['method']) {
            if (preg_match(sprintf('#%s#', $route['uri']), $_SERVER['REQUEST_URI'], $matches) === 1) {
                call_user_func(new $route['action']($_SERVER), $matches);
                break;
            }
        }
    }
} catch (Throwable $e) {
    header('HTTP/1.1 500 Internal Server Error');
    error_log($e);
}
