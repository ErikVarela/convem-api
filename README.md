# 💰 Banking API em Go (Golang)

Este projeto é uma API REST simples desenvolvida em **Go (Golang)** utilizando o framework **Gin**.  
A aplicação simula operações básicas de um sistema bancário, como consulta de saldo e transferência de dinheiro entre contas.

---

## 🚀 Tecnologias utilizadas

- Go (Golang)
- Gin Gonic (framework web)
- HTTP/REST

---

## 📌 Funcionalidades

- 🔍 Consultar saldo de uma conta
- 💸 Transferir dinheiro entre contas

---

## 📂 Estrutura do Projeto

├── main.go


---

## ⚙️ Como executar o projeto

### 1. Clone o repositório
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio

### 2. Instale as dependências
go mod init banking-api
go get github.com/gin-gonic/gin

### 3. Execute a aplicação
go run main.go

A API estará disponível em:
http://localhost:8080


---

## 📡 Endpoints da API

### 🔍 Consultar saldo

**GET** `/accounts/:id/balance`

#### Exemplo de requisição:
GET http://localhost:8080/accounts/1/balance

#### Resposta:

```json
{
  "id": "1",
  "holder": "Candidato Convem",
  "balance": 100000
}

---

####💸 Transferir dinheiro

POST /transfer

Body (JSON):
{
  "from_id": "1",
  "to_id": "2",
  "amount": 1000
}


Exemplo de requisição:
POST http://localhost:8080/transfer
Resposta:
{
  "message": "Transferência realizada com sucesso!"
}

---

####🧠 Lógica da Aplicação

- As contas são armazenadas em memória (slice de structs)
- A transferência atualiza diretamente os saldos das contas
- Não há persistência em banco de dados (projeto educacional)

####⚠️ Melhorias futuras

- Validação de saldo insuficiente
- Uso de banco de dados (MySQL, PostgreSQL, etc.)
- Autenticação e autorização (JWT)
- Logs e tratamento de erros mais robusto
- Testes automatizados


####📚 Objetivo

- Este projeto foi desenvolvido com fins de estudo para praticar:

- Criação de APIs REST com Go
- Uso do framework Gin
- Manipulação de requisições HTTP
- Estruturação de projetos backend
