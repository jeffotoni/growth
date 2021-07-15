<?php

spl_autoload_register(function ($classname) {
    require sprintf('./src/%s.php', strtr($classname, '\\', '/'));
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

foreach ($routes as $route) {
    if ('*' === $route['method'] || $_SERVER['REQUEST_METHOD'] === $route['method']) {
        if (preg_match(sprintf('#%s#', $route['uri']), $_SERVER['REQUEST_URI'], $matches) === 1) {
            call_user_func(new $route['action']($_SERVER), $matches);
            break;
        }
    }
}
