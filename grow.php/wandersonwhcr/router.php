<?php

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
];

$found = false;

foreach ($routes as $route) {
    $found =
        $_SERVER['REQUEST_METHOD'] === $route['method']
        && preg_match(sprintf('#%s#', $route['uri']), $_SERVER['REQUEST_URI'], $matches) === 1;

    if ($found) {
        var_dump($route);
        break;
    }
}

if (! $found) {
    var_dump(null);
}
