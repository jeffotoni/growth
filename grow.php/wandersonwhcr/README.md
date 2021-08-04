# grow.php/wandersonwhcr

* Tamanho da Imagem: `5.26MB`

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

## Execução

```
docker run --rm \
    --detach \
    --publish 8080:8080 \
    wandersonwhcr/growth:latest
```

### Considerações

* Durante a criação da imagem, efetua-se o _download_ do código-fonte do PHP,
  compilando-o com todos os recursos desabilitados.
* Adiciona-se o módulo externo `swoole` do PHP ao binário de forma estática para
  reduzir o tamanho; este módulo é responsável por trabalhar com _fibers_ e
  _coroutines_.
* O executável gerado é compactado através da ferramenta `upx` e sempre que o
  processo é inicializado, o binário é descompactado e depois executado.
* Utiliza-se uma imagem _scratch_ como base da imagem final.
* O processamento do arquivo é **síncrono**: não se desenvolveu uma rotina no
  Swoole para tal aplicação e isto será desenvolvido nos próximos passos.
* Pela simplicidade do desafio, não se utilizou _frameworks_ e nem testes
  unitários. A API é disponibilizada utilizando o Web Server do Swoole.
* Por fim, sinta-se à vontade para melhorar este projeto! O foco inicial foi a
  redução do tamanho da imagem e criação de uma API funcional.

## Ambiente de Desenvolvimento

```
make dev
```

## Ambiente de Testes

```
make stage
```
