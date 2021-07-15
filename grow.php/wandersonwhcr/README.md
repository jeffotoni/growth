# grow.php/wandersonwhcr

* Tamanho da Imagem: `2.75MB`

## Imagem

A imagem Docker `wandersonwhcr/growth:latest` pode ser gerada através do comando
abaixo, responsável também por construir a imagem base para desenvolvimento.

```
make
```

O tamanho final da imagem pode ser verificado através do seguinte comando:

```
docker inspect \
    --type image \
    --format '{{ .Size }}' \
    wandersonwhcr/growth:latest \
    | numfmt --to iec --format '%.2f'
```

## Ambiente de Desenvolvimento

```
make dev
```

## Ambiente de Testes

```
make stage
```
