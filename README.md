# ⚡ Flang

**Flang** é uma linguagem de programação que permite descrever aplicações completas usando apenas arquivos `.fg`. Sem gerar código em outras linguagens — tudo roda direto do `.fg`.

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
```

```bash
flang run inicio.fg
```

**Resultado:** aplicação completa rodando com backend, frontend, banco de dados e API REST.

---

## Instalação

```bash
git clone https://github.com/flaviokalleu/flang.git
cd flang
go build -o flang .
```

## Comandos

| Comando | Descrição |
|---------|-----------|
| `flang run arquivo.fg` | Executa o arquivo .fg |
| `flang arquivo.fg` | Atalho para run |
| `flang check arquivo.fg` | Verifica sintaxe |
| `flang new nome` | Cria novo projeto |
| `flang help` | Ajuda |

## Tipos de Dados

| Tipo | Descrição | Input HTML |
|------|-----------|------------|
| `texto` | Texto simples | text |
| `numero` | Número | number |
| `dinheiro` | Valor monetário (R$) | number |
| `email` | Email com validação | email |
| `telefone` | Telefone com validação | tel |
| `status` | Badge colorido automático | text |
| `data` | Data | date |
| `booleano` | Verdadeiro/Falso | checkbox |
| `senha` | Senha mascarada | password |
| `imagem` | Upload de imagem | file |
| `link` | URL | url |

## Modificadores

```
nome: texto obrigatorio        # campo obrigatório
email: email unico             # valor único
mesa_id: numero pertence_a mesa  # foreign key
```

## Sistema de Imports

Divida seu projeto em múltiplos arquivos `.fg`:

```
importar "dados.fg"
importar "telas.fg"
importar dados de "modelos.fg"
importar tela de "paginas.fg"
```

### Exemplo modular

```
inicio.fg          ← importa tudo
├── dados.fg       ← modelos
├── telas.fg       ← telas
├── eventos.fg     ← eventos
├── regras.fg      ← lógica
└── tema.fg        ← visual
```

## Blocos da Linguagem

### sistema
```
sistema nome_do_app
```

### dados
```
dados
  modelo
    campo: tipo modificadores
```

### telas
```
telas
  tela nome
    titulo "Título"
    lista modelo
      mostrar campo
    botao cor
      texto "Label"
```

### eventos
```
eventos
  quando clicar "Label"
    criar modelo
```

### logica
```
logica
  validar email obrigatorio unico
  se status igual "cancelado"
    mudar cor vermelho
```

### tema
```
tema
  cor primaria "#6366f1"
  cor secundaria "#8b5cf6"
  cor destaque "#f59e0b"
  cor sidebar "#1e1b4b"
```

## Funcionalidades Automáticas

- **API REST** completa (GET, POST, PUT, DELETE) para cada modelo
- **Frontend** com dashboard, sidebar, busca, dark mode
- **Banco de dados** SQLite embutido
- **Validação** automática (email, telefone, obrigatório, único)
- **Badges de status** coloridos (verde=ativo, amarelo=pendente, vermelho=cancelado)
- **Segurança** (XSS protection, headers, sanitização)
- **Responsivo** (desktop e mobile)

## Tech Stack

- **Engine:** Go
- **Banco:** SQLite (embutido)
- **Frontend:** HTML/CSS/JS gerado dinamicamente do AST
- **API:** REST automática

## Estrutura do Projeto

```
flang/
├── main.go                          # Entry point
├── cli/                             # CLI (run, check, new, help)
├── compiler/
│   ├── lexer/                       # Tokenizador
│   ├── parser/                      # Parser → AST
│   └── ast/                         # Estruturas do AST
├── runtime/
│   ├── engine.go                    # Motor principal + imports
│   ├── banco/                       # Banco de dados embutido
│   └── servidor/                    # Servidor HTTP + renderizador
└── exemplos/
    ├── loja/                        # Exemplo simples
    ├── restaurante/                 # Exemplo completo
    └── restaurante-modular/         # Exemplo com imports
```

## Licença

MIT
