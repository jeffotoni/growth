# API Growth

Este repositório foi criado para colocarmos projetos em diversas linguagens com intúito totalmente didático 
para colaborar com a comunidade e desenvolvedores como resolver o problema proposto com objetivo de
visualizarmos as vantagens e desvantagens de cada uma para resolver o problema.

O escopo do projeto é criar uma API rEST um CRUD e persistir em memória e colocar em uma imagem docker e 
o tamanho desta imagem não poderá ultrapassar 6Mb.

Todo repo foi organizado por linguagens de programação, fique a vontade em colaborar enviando um
pull request para nós, logo abaixo vamos deixar na documentação como fazer PR.

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

