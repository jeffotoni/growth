# API Growth 💙 🐿️ 🐍 🦀

Este repositório foi criado para disponibilizarmos projetos em diversas linguagens de programação com intúito didático e colaborar com a comunidade de desenvolvedores. Uma brincadeira que nasceu nas redes sociais e que se materializou neste repositório ❤️.

As linguagens de programação ❤️ são ferramentas e devem ser utilizadas para resolver problemas específicos do que foram propostas a solucionar. Mas sabemos que é muito além disto 😍, nesta equação temos que adicionar uma pitada de AMOR😍 e quando tem esta combinação as coisas começam a ficarem ainda mais interessantes 😂😂.

---
O escopo do projeto é criar uma API rEST um CRUD e persistir em memória e colocar em uma imagem docker e 
o tamanho desta imagem docker não poderia ultrapassar 6Mb porém sabemos das limitações que cada linguagem possui e neste quesito você poderá enviar uma imagem maior, tente fazer o menor que conseguir bem enxuta ☺️.

O seu POST irá receber um JSON de 1mb ou 3mb e persistir em memória.
Logo abaixo tem o exemplo e a descrição do que irá precisar implementar na API.

Todo repo foi organizado por linguagens de programação, fique a vontade em colaborar enviando um
pull request para nós, logo abaixo vamos deixar na documentação como fazer PR.

O que iremos enviar para o [POST] será um json de 1Mb ou 3Mb com mais de 40k de linhas e o corpo do Json está logo abaixo:
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
## Pull Request

Você poderá organizar seu diretório como os exemplos abaixo:
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

## Endpoints a serem implementados
Os endpoints que devem ser implementados estão listados logo abaixo, vamos seguir o mesmo padrão para todos os projetos:

#### POST
Criando nossa base de dados na memória, esta requisição é assícrona irá ficar rodando em background mas somente implemente este quesito se sua linguagem fornecer suporte. 
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
