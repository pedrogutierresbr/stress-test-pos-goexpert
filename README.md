# Desafio Stress Test - Pós GoExpert

<br>

## Sobre o projeto

Este é o repositório destinado ao desafio técnico Stress Test do curso Pós GoExpert da faculdade FullCycle. Projeto é um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

<br>

## Funcionalidades

- Apresentar um relatório ao final dos testes contendo:
		- Tempo total gasto na execução
		- Quantidade total de requests realizados.
		- Quantidade de requests com status HTTP 200.
		- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

<br>

## Como executar o projeto

### Pré-requisitos

Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:

- [Git](https://git-scm.com)

- [VSCode](https://code.visualstudio.com/)

- [Docker](https://www.docker.com/)

<br>

#### Acessando o repositório

  

```bash

# Clone este repositório

$  git  clone  https://github.com/pedrogutierresbr/stress-test-pos-goexpert.git

```

 
<br>

  
  

#### Executando a aplicação em ambiente dev

  

```bash

# Contruir imagem docker

$  docker build -t stress_test .

# Executar a imagem

$  docker run stress_test --url=<url a ser testada> --requests=<número requests> --concurrency=<número chamdas simultâneas>

Ex.:

$ docker run stress_test --url=https://braziljournal.com/ --requests=1000 --concurrency=10

```

<br>
  

#### Resultados

```bash

$ docker run stress_test --url=https://braziljournal.com/ --requests=1000 --concurrency=10
Tempo total: 10.584274677s
Total de requests realizados: 1000
Requests com status 200: 0
Distribuição de outros códigos de status:


$ docker run stress_test --url=https://br.investing.com/ --requests=5000 --concurrency=15
Tempo total: 38.427529448s
Total de requests realizados: 5000
Requests com status 200: 0
Distribuição de outros códigos de status:

```

<br>

## Licença

  

Este projeto esta sobe a licença [MIT](./LICENSE). 

Feito por Pedro Gutierres [Entre em contato!](https://www.linkedin.com/in/pedrogabrielgutierres/)