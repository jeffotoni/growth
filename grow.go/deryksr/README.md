# API Growth

Uma simples solução do problema apresentado usando a linguagem de programação **Go** com o framework **Echo**.
A principal motivação dessa implementação é a obtenção de uma camisa personalizada de golang, é sério
eu vim só pela camisa! se eu não ganhar eu vou ficar chateado :cry:

## Instruções básicas

O projeto contém um arquivo *Makefile* para deixar a vida de todo mundo mais fácil.

### Construir a imagem

```shell
make build
```

### Para obter o tamanho da imagem

```shell
make image-size
```

### Executar a aplicação

```shell
make run
```

## Considerações

Até que foi legalzinho pensar em algumas coisas como um banco de dados em memória, a vontade de subir um redis foi grande
porém o tamanho final ficaria um *pouquinho* maior que 6 MB kkkkkkk'.

O endpoint de status estava um pouco confuso na documentação
ele imprimia um valor *test value* que não fazia sentido algum, Por esse motivo eu alterei o retorno da função

### Endpoint /api/v1/growth/post/status

Response

```json
{
    "msg": "still in processing...",
    "count": "0 ~ 42450 // qtd processada até o momento da request"
}
```

OU

```json
{
    "msg": "processing completed",
    "count": "42450 // qtd total no arquivo de teste"
}
```
