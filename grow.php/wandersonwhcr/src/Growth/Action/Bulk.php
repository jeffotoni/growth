<?php

namespace Growth\Action;

use Swoole\HTTP\Request;
use Swoole\HTTP\Response;
use Swoole\Table;

class Bulk
{
    public function __construct(private Table $table)
    {
    }

    private function getRequestContentType(Request $request): ?string
    {
        return $request->header['content-type'] ?? null;
    }

    public function __invoke(Request $request, Response $response): void
    {
        if ($this->getRequestContentType($request) !== 'application/json') {
            $response->status(415);
            $response->end();
            return;
        }

        $dataset = json_decode($request->getContent());

        if (! $dataset) {
            $response->status(422);
            $response->end();
            return;
        }

        if (! is_array($dataset)) {
            $response->status(422);
            $response->end();
            return;
        }

        // TODO Async
        foreach ($dataset as $data) {
            $key = strtolower(sprintf('growth-%s-%s-%s', $data->Country, $data->Indicator, $data->Year));

            $this->table->set($key, ['content' => json_encode($data)]);
        }

        $response->status(201);
        $response->header('content-type', 'application/json');
        $response->end(json_encode([
            'msg' => 'In progress',
        ]));
    }
}
