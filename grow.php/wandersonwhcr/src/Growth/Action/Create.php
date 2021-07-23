<?php

namespace Growth\Action;

use Swoole\HTTP\Request;
use Swoole\HTTP\Response;
use Swoole\Table;
use StdClass;

class Create
{
    public function __construct(private Table $table)
    {
    }

    protected function getRequestContentType(Request $request): ?string
    {
        return $request->header['content-type'] ?? null;
    }

    public function __invoke(Request $request, Response $response, array $args): void
    {
        if ($this->getRequestContentType($request) !== 'application/json') {
            $response->status(415);
            $response->end();
            return;
        }

        $content = json_decode($request->getContent());

        if (! $content) {
            $response->status(422);
            $response->end();
            return;
        }

        $key = sprintf('growth-%s-%s-%s', $args['country'], $args['indicator'], $args['year']);

        $data = new StdClass();
        $data->Country = strtoupper($args['country']);
        $data->Indicator = strtoupper($args['indicator']);
        $data->Value = $content->value;
        $data->Year = strtoupper($args['year']);

        $this->table->set($key, ['content' => json_encode($data)]);

        $response->status(204);
        $response->end();
    }
}
