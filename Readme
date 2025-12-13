# Projeto em go

API REST simples desenvolvida em Go.

## Pré-requisitos

- Go 1.21 ou superior instalado
- Git

## Como rodar o projeto

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/Crud-go.git
cd Crud-go
```

### 2. Instale as dependências

```bash
go mod download
```

### 3. Execute o projeto

```bash
go run main.go
```

O servidor estará rodando em `http://localhost:8080`

## Estrutura do projeto

```
Crud-go/
├── src/
│   ├── configuration/
│   │   └── rest_err.go
│   └── controller/
│       └── routes/
│           ├── routes.go
│           ├── createUser.go
│           ├── deleteUser.go
│           ├── findUser.go
│           └── updateUser.go
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── Readme.md
```

## Endpoints da API

- `POST /users` - Criar usuário
- `GET /users/:id` - Buscar usuário
- `PUT /users/:id` - Atualizar usuário
- `DELETE /users/:id` - Deletar usuário

## Tecnologias

- Go 1.21+
- Gin Framework (ou especifique o framework usado)

## Compilar o projeto

Para gerar um executável:

```bash
go build -o app
```

Para executar:

```bash
./app
```