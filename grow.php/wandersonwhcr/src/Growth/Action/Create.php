<?php

namespace Growth\Action;

use StdClass;

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

        $content = json_decode(file_get_contents('php://input'));

        if (! $content) {
            header('HTTP/1.1 422 Unprocessable Entity');
            return;
        }

        header('HTTP/1.1 201 Created');

        $key = sprintf('growth-%s-%s-%s', $args['country'], $args['indicator'], $args['year']);

        $data = new StdClass();
        $data->Country = strtoupper($args['country']);
        $data->Indicator = strtoupper($args['indicator']);
        $data->Value = $content->value;
        $data->Year = strtoupper($args['year']);

        apcu_store($key, $data);
        apcu_inc('growth-count');
    }
}
