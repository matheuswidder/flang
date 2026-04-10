<p align="center">
  <img src="https://img.shields.io/badge/Flang-v0.4-6366f1?style=for-the-badge&logo=go&logoColor=white" alt="Flang v0.4">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go 1.21+">
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="MIT License">
  <img src="https://img.shields.io/badge/PRs-Welcome-brightgreen?style=for-the-badge" alt="PRs Welcome">
</p>

<h1 align="center">Flang - A Programming Language for Full-Stack Apps</h1>

<p align="center">
  <strong>Descreva sua aplicacao em um arquivo <code>.fg</code> e rode. Backend, frontend, banco de dados, API REST e WhatsApp - tudo automatico.</strong>
</p>

<p align="center">
  <a href="#instalacao">Instalacao</a> |
  <a href="#quick-start">Quick Start</a> |
  <a href="#documentacao">Documentacao</a> |
  <a href="#exemplos">Exemplos</a> |
  <a href="#contribuir">Contribuir</a>
</p>

---

## O que e o Flang?

**Flang** e uma linguagem de programacao **declarativa e bilingue** (Portugues/English) que gera aplicacoes completas a partir de arquivos `.fg`. Voce descreve **o que** a aplicacao faz, e o Flang cuida do **como**.

```
sistema loja

dados
  produto
    nome: texto obrigatorio
    preco: dinheiro
    status: status

telas
  tela produtos
    titulo "Produtos"
    lista produto
      mostrar nome
      mostrar preco
      mostrar status
    botao azul
      texto "Novo Produto"

eventos
  quando clicar "Novo Produto"
    criar produto
```

```bash
flang run inicio.fg
# Aplicacao completa rodando em http://localhost:8080
```

**Resultado:** Dashboard com sidebar, dark mode, CRUD completo, API REST, banco de dados, WebSocket em tempo real e WhatsApp - tudo de um unico arquivo `.fg`.

---

## Por que Flang?

| Problema | Solucao Flang |
|----------|--------------|
| Configurar backend, frontend, DB separadamente | **1 arquivo `.fg` = app completa** |
| Semanas para um CRUD basico | **Segundos** |
| Aprender React + Node + SQL + CSS | **Aprenda Flang em 5 minutos** |
| API REST manual para cada modelo | **Gerada automaticamente** |
| WebSocket complexo de implementar | **Embutido e automatico** |
| Integracao WhatsApp trabalhosa | **3 linhas no `.fg`** |
| Escolher entre Portugues e English | **Bilingue - use os dois** |

---

## Features

- **Linguagem Bilingue** - Escreva em Portugues, English ou misture os dois
- **Backend Embutido** - Servidor HTTP com API REST automatica
- **Frontend Dinamico** - Dashboard moderno com glassmorphism, dark mode, sidebar
- **3 Bancos de Dados** - SQLite (padrao), MySQL, PostgreSQL
- **WebSocket** - Atualizacoes em tempo real entre multiplos usuarios
- **WhatsApp** - Envio automatico de mensagens via whatsmeow
- **Sistema de Imports** - Divida seu projeto em multiplos `.fg`
- **Logica e Validacao** - Regras de negocio declarativas
- **13 Tipos de Dados** - texto, numero, dinheiro, email, telefone, status, senha...
- **Seguranca** - XSS protection, headers, sanitizacao, validacao automatica
- **Responsivo** - Desktop e mobile out-of-the-box
- **Zero Dependencias Externas** - Um unico binario

---

## Instalacao

### Pre-requisitos

- [Go 1.21+](https://go.dev/dl/)

### Build

```bash
git clone https://github.com/flaviokalleu/flang.git
cd flang
go build -o flang .
```

### Verificar instalacao

```bash
./flang version
```

---

## Quick Start

### 1. Crie seu primeiro app

```bash
./flang new meu-app
```

Isso cria `meu-app/inicio.fg` com um template basico.

### 2. Rode

```bash
cd meu-app
../flang run inicio.fg
```

### 3. Abra no navegador

```
http://localhost:8080
```

Pronto. App completa rodando.

---

## Documentacao

### Blocos da Linguagem

Flang usa **blocos** para definir cada parte da aplicacao:

| Bloco | English | O que faz |
|-------|---------|-----------|
| `sistema` | `system` | Nome da aplicacao |
| `dados` | `models` | Modelos de dados (tabelas) |
| `telas` | `screens` | Interface do usuario |
| `eventos` | `events` | Acoes do usuario |
| `logica` | `logic` | Regras de negocio |
| `tema` | `theme` | Customizacao visual |
| `banco` | `database` | Configuracao do banco |
| `integracoes` | `integrations` | WhatsApp, etc |

### Tipos de Dados

| Tipo PT | Tipo EN | SQL | HTML Input | Descricao |
|---------|---------|-----|------------|-----------|
| `texto` | `text` | TEXT | text | Texto simples |
| `numero` | `number` | REAL | number | Numero decimal |
| `dinheiro` | `money` | REAL | number | Valor monetario (R$) |
| `email` | `email` | TEXT | email | Email com validacao |
| `telefone` | `phone` | TEXT | tel | Telefone com validacao |
| `status` | `status` | TEXT | text | Badge colorido automatico |
| `data` | `date` | DATETIME | date | Data |
| `booleano` | `boolean` | INTEGER | checkbox | Verdadeiro/Falso |
| `senha` | `password` | TEXT | password | Mascarado |
| `imagem` | `image` | TEXT | file | Upload de imagem |
| `arquivo` | `file` | TEXT | file | Upload de arquivo |
| `upload` | `upload` | TEXT | file | Upload generico |
| `link` | `link` | TEXT | url | URL |

### Modificadores

| Modificador PT | EN | O que faz |
|----------------|----|-----------|
| `obrigatorio` | `required` | Campo NOT NULL + validacao |
| `unico` | `unique` | Valor unico na tabela |
| `pertence_a` | `belongs_to` | Foreign key |
| `padrao` | `default` | Valor padrao |

### Status Badges Automaticos

Campos do tipo `status` ganham badges coloridos automaticamente:

| Cor | Valores |
|-----|---------|
| Verde | ativo, livre, aberto, ok, disponivel, pronto, entregue, pago |
| Amarelo | pendente, aguardando, em andamento, reservado, preparando |
| Vermelho | inativo, ocupado, fechado, cancelado, bloqueado |
| Azul | outros valores |

### Sistema de Imports

Divida seu projeto em multiplos arquivos `.fg`:

```
importar "dados.fg"              # importa tudo do arquivo
importar dados de "modelos.fg"   # so os dados
importar telas de "paginas.fg"   # so as telas
importar produto de "dados.fg"   # modelo especifico
```

```
# English
import "data.fg"
import models from "models.fg"
import screens from "pages.fg"
```

### Configuracao de Banco

**SQLite (padrao - zero config):**
```
sistema meu_app
dados
  ...
```

**PostgreSQL:**
```
banco
  driver: postgres
  host: "localhost"
  porta: "5432"
  nome: "meu_banco"
  usuario: "postgres"
  senha: "minhasenha"
```

**MySQL:**
```
database
  driver: mysql
  host: "localhost"
  port: "3306"
  name: "my_database"
  user: "root"
  password: "mypassword"
```

### WhatsApp

```
integracoes
  whatsapp
    quando criar pedido
      enviar mensagem para telefone
        texto "Ola {cliente}! Pedido de {prato} recebido! Valor: R${valor}"

    quando atualizar pedido
      enviar mensagem para telefone
        texto "Status atualizado: {status}"
```

Na primeira execucao, escaneie o QR Code no terminal com WhatsApp > Dispositivos Conectados.

**Templates:** Use `{campo}` para inserir valores do registro na mensagem.

### API REST Automatica

Cada modelo gera automaticamente:

| Metodo | Rota | Acao |
|--------|------|------|
| `GET` | `/api/{modelo}` | Listar todos |
| `GET` | `/api/{modelo}/{id}` | Buscar por ID |
| `POST` | `/api/{modelo}` | Criar |
| `PUT` | `/api/{modelo}/{id}` | Atualizar |
| `DELETE` | `/api/{modelo}/{id}` | Deletar |

Exemplo:
```bash
# Listar produtos
curl http://localhost:8080/api/produto

# Criar produto
curl -X POST http://localhost:8080/api/produto \
  -H "Content-Type: application/json" \
  -d '{"nome":"Camiseta","preco":59.90,"status":"ativo"}'
```

### WebSocket

Automatico. Abra a app em 2 abas - quando uma modifica dados, a outra atualiza instantaneamente.

Endpoint: `ws://localhost:8080/ws`

Mensagens recebidas:
```json
{"type":"criar","model":"produto","id":1,"data":{...}}
{"type":"atualizar","model":"produto","id":1,"data":{...}}
{"type":"deletar","model":"produto","id":1}
```

---

## Exemplos

### Loja (Portugues)
```
sistema loja

dados
  produto
    nome: texto obrigatorio
    preco: dinheiro
    categoria: texto
    status: status

  cliente
    nome: texto obrigatorio
    email: email unico
    telefone: telefone

telas
  tela produtos
    titulo "Produtos"
    lista produto
      mostrar nome
      mostrar preco
      mostrar status
    botao azul
      texto "Novo Produto"

  tela clientes
    titulo "Clientes"
    lista cliente
      mostrar nome
      mostrar email
    botao verde
      texto "Novo Cliente"

eventos
  quando clicar "Novo Produto"
    criar produto
  quando clicar "Novo Cliente"
    criar cliente
```

### Store (English)
```
system store

models
  product
    name: text required
    price: money
    status: status

  customer
    name: text required
    email: email unique

screens
  screen products
    title "Products"
    list product
      show name
      show price
      show status
    button blue
      text "New Product"

events
  when click "New Product"
    create product
```

### Projeto Modular (com imports)
```
# inicio.fg
sistema restaurante

importar "tema.fg"
importar "dados.fg"
importar "telas.fg"
importar "eventos.fg"
importar "regras.fg"
```

---

## Comandos CLI

| Comando | Descricao |
|---------|-----------|
| `flang run <arquivo.fg>` | Executa o arquivo .fg |
| `flang <arquivo.fg>` | Atalho para run |
| `flang run <arquivo.fg> <porta>` | Executa na porta especificada |
| `flang check <arquivo.fg>` | Verifica sintaxe sem executar |
| `flang new <nome>` | Cria novo projeto |
| `flang version` | Mostra versao |
| `flang help` | Ajuda |

---

## Arquitetura

```
arquivo.fg
    |
    v
 [Lexer] --> tokens
    |
    v
 [Parser] --> AST
    |
    v
 [Runtime]
    |-- [Banco]      SQLite / MySQL / PostgreSQL
    |-- [Servidor]   HTTP + WebSocket
    |-- [Renderer]   HTML/CSS/JS dinamico
    |-- [WhatsApp]   whatsmeow
    |-- [Validador]  Regras de negocio
```

```
flang/
  main.go                          # Entry point
  cli/                             # CLI (run, check, new, help)
  compiler/
    lexer/                         # Tokenizador bilingue
    parser/                        # Parser -> AST
    ast/                           # Estruturas do AST
  runtime/
    engine.go                      # Motor principal + imports
    banco/                         # SQLite, MySQL, PostgreSQL
    servidor/                      # HTTP + WebSocket + Renderer
    whatsapp/                      # WhatsApp via whatsmeow
  exemplos/
    restaurante-modular/           # Exemplo com imports
    restaurante-whatsapp/          # Exemplo com WhatsApp
    english/                       # Exemplo em ingles
    mixed/                         # Exemplo bilingue
```

---

## Comparacao

| Feature | Flang | Low-Code Tools | Codigo Manual |
|---------|-------|----------------|---------------|
| Tempo para CRUD | Segundos | Minutos | Horas/Dias |
| Linguagem propria | .fg | Visual/Drag | JS/Python/etc |
| Backend + Frontend | Automatico | Parcial | Manual |
| Banco de dados | 3 drivers | Limitado | Manual |
| WebSocket | Automatico | Plugin | Manual |
| WhatsApp | 3 linhas | Nao | Complexo |
| Self-hosted | Sim | Depende | Sim |
| Open source | Sim | Nem sempre | - |
| Bilingue PT/EN | Sim | Nao | Nao |

---

## Roadmap

- [ ] Autenticacao (login/registro)
- [ ] Upload de arquivos
- [ ] Graficos e charts no dashboard
- [ ] Deploy com Docker
- [ ] Flang Cloud (deploy com 1 comando)
- [ ] Mais integracoes (Telegram, Email, SMS)
- [ ] Editor visual para .fg
- [ ] Flang escrito em Flang (self-hosting)

---

## Contribuir

Contribuicoes sao bem-vindas!

1. Fork o repositorio
2. Crie sua branch (`git checkout -b feature/minha-feature`)
3. Commit (`git commit -m 'Add: minha feature'`)
4. Push (`git push origin feature/minha-feature`)
5. Abra um Pull Request

---

## Licenca

MIT License - veja [LICENSE](LICENSE) para detalhes.

---

<p align="center">
  Feito com Go por <a href="https://github.com/flaviokalleu">@flaviokalleu</a>
</p>
