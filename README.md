# POC Injeção de Dependência

Este projeto tem como objetivo demonstrar como utilizar o [go.uber.org/dig](https://pkg.go.dev/go.uber.org/dig) para gerenciar dependências em uma aplicação Go.

## Objetivo

Demonstrar de forma simples como o `dig` simplifica a injeção de dependências em Go, permitindo configurar e gerenciar componentes de forma modular e testável.

## Como Executar

1. Certifique-se de ter o Go instalado (versão recente).
2. No diretório raiz do projeto, rode:
   ```bash
   go mod tidy
   go run main.go
   ```
3. A aplicação estará disponível em http://localhost:8080.

### Testes

Para executar os testes, rode:
```bash
go test ./...
```

## Estrutura do Projeto

```
/
├── main.go                           # Ponto de entrada da aplicação
├── internal/                         # Código interno da aplicação
│   └── user/                         # Módulo de usuários
│       ├── domain/                   # Camada de domínio - regras de negócio
│       │   ├── user.go               # Entidade de usuário
│       │   └── repository.go         # Interface do repositório de usuários
│       ├── application/              # Camada de aplicação - casos de uso
│       │   └── service.go            # Serviço com lógica de aplicação
│       └── infrastructure/           # Camada de infraestrutura - implementação técnica
│           └── repository.go         # Implementação concreta do repositório
```

