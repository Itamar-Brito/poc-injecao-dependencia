# POC Injeção de Dependência

Este projeto tem como objetivo demonstrar como utilizar o [go.uber.org/dig](https://pkg.go.dev/go.uber.org/dig) para gerenciar dependências em uma aplicação Go.

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

## Objetivo

Demonstrar de forma simples como o `dig` simplifica a injeção de dependências em Go, permitindo configurar e gerenciar componentes de forma modular e testável.
