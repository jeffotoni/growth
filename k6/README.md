# k6.io

k6 é uma ferramenta de teste de carga de código aberto desenvolvida pela linguagem Go 😍. O k6 vem com recursos, sobre os quais você pode aprender tudo na documentação. Os principais recursos incluem:

- Ferramenta CLI com APIs amigáveis ​​ao desenvolvedor.
- Scripting em JavaScript ES2015 / ES6 - com suporte para módulos locais e remotos
- Verificações e limites - para teste de carga orientado a metas

O k6 criou sua própria lib javascript para comportar como nodejs, então ao construir os scripts irá usar a linguagem javascript porém com libs disponibilizada pela k6.io.

Posso usar npm e suas libs para criação dos meus scripts ?
Sim pode, importar módulos npm ou bibliotecas, você pode [agrupar módulos npm com webpack](https://k6.io/docs/using-k6/modules/#bundling-node-modules) e importá-los em seus testes.

### Github k6.io
[github k6.io](https://github.com/k6io/k6)

### Instalar k6.io

Existe várias formas de instalação, e por ser feito em Go tudo fica mais fácil basta instalar seu binário em sua máquina.

Aqui está o link com todas possibilidades de instalação:
[Install k6.io](https://k6.io/docs/getting-started/installation/)

### Instalar com docker

Vamos mostrar a instalação usando Docker desta forma não irá precisar instalar nadinha na sua máquina.

```bash
$ docker pull loadimpact/k6
```

Agora vamos executar nossa massa de testes.
```bash
$ docker run -v $(pwd):/data \
-e DOMAIN=http://192.168.0.70:8080 \
-i loadimpact/k6 run - <script.js
```
Nosso script já deixamos pré-prontos, fazendo chamada do POST que envia um json e do Get buscando a informação e do nosso famigerado ping 😍.
