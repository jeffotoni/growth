# API Growth <img width="25px" src="https://kotlinlang.org/assets/images/favicon.svg"/>

Esta é a API Growth escrita e Kotlin, utilizando o framework KTOR!

### Aviso

```text
Por conta do build ser dentro do container, ele pode demorar um pouco mais que o normal. 
Mas... dê uma chance ao gradle, ele vai compilar bem =D
```


## Docker
 
Para buildar, basta executar o comando abaixo:

```shell
$ docker build -t growth-kt .
```

Para rodar o projeto, basta rodar o comando a seguir (após fazer o build):

```shell
$ docker run --rm -it -p 8080:8080 growth-kt
```

Caso queira verificar o tamanho da sua imagem, basta dar uma olhada no comando abaixo =]

```shell
$ docker images | grep growth-kt
growth-kt    latest      fdbf94696025   30 seconds ago   75.9MB
```

# Docs de uso da API

O que iremos enviar para o [POST] será um json de 3Mb com mais de 40k de linhas e o corpo do Json está logo abaixo:
```bash
[
   {
      "Country":"BRZ",
      "Indicator":"NGDP_R",
      "Value":183.26,
      "Year":2002
   },
   {
      "Country":"AFG",
      "Indicator":"NGDP_R",
      "Value":198.736,
      "Year":2003
   }
]
```
O arquivo 3mb-growth_json.json que encontra-se no raiz deste repositório.

## Pull Request

Você poderá organizar sua pasta como os exemplos abaixo:
```bash
grow.go/
└── jeffotoni
    ├── grow.fiber
    │   └── README.md
    └── grow.standard.libray
        ├── Dockerfile
        ├── go.mod
        ├── main.go
        ├── main_test.go
        └── README.md

```
Poderá organizar seu projeto escolhendo a linguagem que irá implementar e logo depois seu user do github e dentro de seu
diretório poderá criar e organizar suas contribuições.

Confira mais exemplos:
```bash
grow.python/
└── cassiobotaro
    ├── Dockerfile
    ├── main.py
    ├── README.md
    └── requirements.txt
```

```bash
grow.rust
└── marioidival
    └── actix
        ├── Cargo.toml
        └── src
            └── main.rs
```


```bash
grow.kotlin
└── arieldossantos
    └── src
```

Os endpoints que devem ser implementados estão listados logo abaixo, todos vamos seguir
o mesmo padrão:


#### POST
Criando nossa base de dados na memória, esta requisição é assícrona irá ficar rodando em
background.
```bash
$ curl -i -XPOST -H "Content-Type:application/json" \
localhost:8080/api/v1/growth -d @3mb-growth_json.json
{"msg":"In progress"}
```

#### GET
Com este endpoint conseguimos visualizar o status de como está o processamento que enviamos no [POST]
```bash
$ curl -i -XGET -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/post/status
{"msg":"complete","test value"":183.26, "count":42450}
```
#### GET
Este endpoint faz um busca na memória para retornar o resultado
```bash
$ curl -i -XGET -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/brz/ngdp_r/2002
{"Country":"BRZ","Indicator":"NGDP_R","Value":183.26,"Year":2002}
```
#### PUT
Este endpoint irá fazer uma atualização na base de dados que está em memória,
se não existir o dado ele irá criar um novo.
```bash
$ curl -i -XPUT -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/brz/ngdp_r/2002 \
-d '{"value":333.98}'
```
#### GET
Fazendo um request para checar se o que alteramos ou criamos novo está na base de dados.
```bash
$ curl -i -XGET -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/brz/ngdp_r/2002
{"Country":"BRZ","Indicator":"NGDP_R","Value":333.98,"Year":2002}
```
#### DELETE
Este endpoint irá remove o dado de nossa base de dados memória.
```bash
$ curl -i -XDELETE -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/brz/ngdp_r/2002 
```
#### GET
Este endpoint irá retornar o tamanho que encontra-se a nossa base de dados na memória
```bash
$ curl -i -XGET -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/size
{"size":42450}
```
