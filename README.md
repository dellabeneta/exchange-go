# Real VS Moedas

![Pipeline](https://github.com/dellabeneta/exchange-go/actions/workflows/main.yaml/badge.svg)

Um serviço web simples para consulta de cotações de moedas estrangeiras em relação ao Real (BRL), desenvolvido em Go. Acesse a interface web para visualizar as taxas de compra, venda e variação das principais moedas.

## Funcionalidades

- Consulta em tempo real das cotações USD, EUR, GBP, AUD e CAD contra BRL
- Interface web responsiva e fácil de usar
- Backend em Go com endpoint `/cotacao`
- Pronto para rodar em Docker e Kubernetes
- CI/CD com GitHub Actions

## Começando

### Pré-requisitos

- Go 1.24.x ou superior
- Docker (opcional)
- Kubernetes/k3s (opcional)

### Instalação Local

```bash
git clone https://github.com/dellabeneta/exchange-go.git
cd exchange-go
go run main.go
```

Acesse em [http://localhost:8080](http://localhost:8080)

### Usando Docker

```bash
docker build -t exchange-go .
docker run -p 8080:8080 exchange-go
```

### Deploy no Kubernetes

```bash
kubectl apply -f k3s/namespace.yaml
kubectl apply -f k3s/deployment.yaml
kubectl apply -f k3s/service.yaml
```

A aplicação estará disponível na porta `30085` do cluster.

## Estrutura do Projeto

```
.
├── Dockerfile
├── k3s
│   ├── deployment.yaml
│   ├── namespace.yaml
│   └── service.yaml
├── main.go
├── nuke.sh
├── README.md
└── static
    ├── favicon.ico
    ├── index.html
    └── flags
        ├── au.svg
        ├── ca.svg
        ├── eu.svg
        ├── gb.svg
        └── us.svg
```

## Como Funciona

1. O usuário acessa a interface web.
2. O frontend faz requisição ao endpoint `/cotacao`.
3. O backend consulta a API [AwesomeAPI](https://docs.awesomeapi.com.br/api-de-moedas) e retorna as cotações.
4. As taxas de compra, venda e variação são exibidas na tabela.

## Scripts Úteis

**nuke.sh**: Script para limpeza completa do Docker (containers, imagens, volumes e redes).

```bash
chmod +x nuke.sh
./nuke.sh
```

## Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE)