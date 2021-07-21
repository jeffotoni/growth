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

### Considerações

* Durante a criação da imagem, efetua-se o _download_ do código-fonte do PHP,
  compilando-o com todos os recursos desabilitados.
* Adiciona-se o módulo externo `apcu` do PHP ao binário de forma estática para
  reduzir o tamanho; este módulo é responsável por trabalhar com tabelas de
  _hash_ em memória.
* O executável gerado é compactado através da ferramenta `upx` e sempre que o
  processo é inicializado, o binário é descompactado e depois executado.
* Utiliza-se uma imagem _scratch_ como base da imagem final; inclui-se um volume
  `/tmp` para _upload_ de arquivos temporários.

## Ambiente de Desenvolvimento

```
make dev
```

## Ambiente de Testes

```
make stage
```
