# Locust

[Locust](https://docs.locust.io) é uma ferramenta de teste de desempenho fácil de usar, programável e escalonável. Você define o comportamento de seus usuários no código Python regular, em vez de usar uma IU desajeitada ou uma linguagem específica de domínio.

### Instalação com pip3
Foi desenvolvido em python e poderá ser instalada usando pip3.
```bash
$ pip3 install locust
```

Para atualizar
```bash
$ pip3 install -e git://github.com/locustio/locust.git@master#egg=locus
```

O Locust nasceu de uma frustração com as soluções existentes. Nenhuma ferramenta de teste de carga existente estava bem equipada para gerar carga realista em um site dinâmico onde a maioria das páginas tinha conteúdo diferente para usuários diferentes. 

As ferramentas existentes usavam interfaces desajeitadas ou arquivos de configuração detalhados para declarar os testes.

O Locust possue uma interface para acompanhar os testes.

Para fazer o teste você precisa criar seu script teste que é todo feito em python.
Ele possui uma declaração e notação para que possa criar o script e tudo poderá encontrado em seu manual que encontra-se no [site oficial](https://docs.locust.io/en/stable/)

### Executando Locust
```bash
$ locust -f locustfile.py
```
Agora é acessar o browser na porta http://localhost:8089 e iniciar os testes.

