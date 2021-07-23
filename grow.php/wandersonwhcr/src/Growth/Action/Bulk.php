<?php

namespace Growth\Action;

class Bulk
{
    protected array $server;

    protected function getRequestContentType(): ?string
    {
        return $this->server['HTTP_CONTENT_TYPE'] ?? null;
    }

    public function __invoke(): void
    {
        if ($this->getRequestContentType() !== 'application/json') {
            header('HTTP/1.1 415 Unsupported Media Type');
            return;
        }

        $dataset = json_decode(file_get_contents('php://input'));

        if (! $dataset) {
            header('HTTP/1.1 422 Unprocessable Entity');
            return;
        }

        if (! is_array($dataset)) {
            header('HTTP/1.1 422 Unprocessable Entity');
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

        header('HTTP/1.1 201 Accepted');

        file_put_contents('php://output', json_encode(['msg' => 'In progress']));
    }
}
