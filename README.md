# MediSystem

Sistema de gestão de cadastro, Edição e Consultas médicas, desenvolvido em Go com PostgreSQL.

## TECNOLOGIAS
-**Backend:** Go (Net/http)
-**Banco de Dados:** PostgreSQL
-**Frontend:** HTML, CSS, JavaScript

## FUNCIONALIDADES
- Agendamento e listagem de consultas médicas
- Cadastro e gestão de médicos
- Cadastro de gestão de pacientes
- Dashboard com resumo geral
- Filtro e busca de consultas

## COMO RODAR O PROJETO

### Pré-requisitos
-Go 1.22+
-PostgreSQL

### Configuração
1. Clone o repositório:
git clone https://github.com/seu-usuário/projeto-saude.git

2. Crie o arquivo '.env' na raiz:
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=saude

3. Execute:
go run main.go

4. Acesse: http://localhost:8080

## ESTRUTURA DO PROJETO

projeto-saude/
├── database/       # Conexão com PostgreSQL
├── handlers/       # Handlers HTTP
├── migrations/     # Scripts SQL
├── models/         # Structs dos modelos
├── static/         # Frontend (HTML, CSS, JS)
├── storage/        # Camada de acesso ao banco
└── main.go         # Entrada da aplicação
