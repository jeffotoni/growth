<?php

namespace Growth\Action;

use Swoole\HTTP\Request;
use Swoole\HTTP\Response;

class Bulk
{
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

            if (! apcu_exists($key)) {
                apcu_inc('growth-count');
            }

            apcu_store($key, $data);
        }

        $response->status(201);
        $response->end('{"msg":"In progress"}');
    }
}
