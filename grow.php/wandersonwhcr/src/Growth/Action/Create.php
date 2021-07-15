<?php

namespace Growth\Action;

class Create
{
    protected array $server;

    public function __construct(array $server)
    {
        $this->server = $server;
    }

    protected function getRequestContentType(): ?string
    {
        return $this->server['HTTP_CONTENT_TYPE'] ?? null;
    }

    public function __invoke(array $args): void
    {
        if ($this->getRequestContentType() !== 'application/json') {
            header('HTTP/1.1 415 Unsupported Media Type');
            return;
        }

        $data = json_decode(file_get_contents('php://input'));

        if (! $data) {
            header('HTTP/1.1 422 Unprocessable Entity');
            return;
        }

        header('HTTP/1.1 201 Created');

        $key = sprintf('growth-%s-%s-%s', $args['country'], $args['indicator'], $args['year']);

        apcu_store($key, $data);
        apcu_inc('growth-count');
    }
}
