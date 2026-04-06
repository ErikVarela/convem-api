# 💰 Wallet API - Sistema Bancário em Go

Esta é uma API REST desenvolvida em Go (Golang) utilizando o framework Gin. A aplicação simula operações bancárias essenciais com persistência de dados e segurança em transações.

---

## 🚀 Novidades da Versão 2.0

Diferente da versão inicial que utilizava armazenamento em memória (slices), esta atualização traz:

- **Persistência de Dados:** Integração com SQLite para manter os dados mesmo após reiniciar a aplicação.
- **Transações SQL:** Uso de `Begin`, `Commit` e `Rollback` para garantir a integridade financeira (o dinheiro nunca "some" se houver um erro no meio do processo).
- **Driver Pure Go:** Utilização do driver `modernc.org/sqlite`, permitindo a execução em qualquer ambiente sem dependência de compiladores C (CGO).

---

## 🛠️ Tecnologias Utilizadas

- Go (Golang) 1.20+
- Gin Gonic (Framework Web)
- SQLite (Banco de dados embutido)
- Modernc SQLite (Driver nativo para Go)

---

## 📌 Funcionalidades

- **Consultar Saldo:** Busca em tempo real no banco de dados por ID.
- **Transferência Bancária:** Operação atômica entre duas contas.
- **Auto-Bootstrap:** O banco de dados e as tabelas são criados automaticamente na primeira execução.

---

## 📂 Estrutura do Projeto

```plaintext
├── main.go      # Lógica da API e conexão com banco
├── go.mod       # Gerenciamento de dependências
├── banco.db     # Arquivo do banco de dados (gerado automaticamente)
└── .gitignore   # Proteção para não subir o arquivo de banco
´´´

---

## ⚙️ Como executar o projeto

### 1. Clone o repositório

```
git clone https://github.com/ErikVarela/wallet-api.git
cd seu-repositorio
´´´

### 2. Instale as dependências

```
go mod init banking-api
go get github.com/gin-gonic/gin
´´´

### 3. Execute a aplicação
```
go run main.go
´´´

## A API estará disponível em:

```
http://localhost:8080
´´´

---

## 📡 Endpoints da API

### 🔍 Consultar saldo

```
**GET** `/accounts/:id/balance`
´´´

#### Exemplo de requisição:

```
GET http://localhost:8080/accounts/1/balance
´´´

#### Resposta:

```json
{
  "id": "1",
  "holder": "Candidato Convem",
  "balance": 100000
}
```
---

#### 💸 Transferir dinheiro

POST /transfer
```
Body (JSON):
{
  "from_id": "1",
  "to_id": "2",
  "amount": 1000
}
```

Exemplo de requisição:
POST http://localhost:8080/transfer
Resposta:

```
{
  "message": "Transferência realizada com sucesso!"
}
```
---

## 🧠 Lógica da Aplicação

- As contas são armazenadas em memória (slice de structs)
- A transferência atualiza diretamente os saldos das contas
- Não há persistência em banco de dados (projeto educacional)

## ⚠️ Melhorias futuras

- Validação de saldo insuficiente
- Uso de banco de dados (MySQL, PostgreSQL, etc.)
- Autenticação e autorização (JWT)
- Logs e tratamento de erros mais robusto
- Testes automatizados


## 📚 Objetivo

- Este projeto foi desenvolvido com fins de estudo para praticar:

- Criação de APIs REST com Go
- Uso do framework Gin
- Manipulação de requisições HTTP
- Estruturação de projetos backend
