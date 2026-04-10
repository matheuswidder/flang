# Flang — Documento de Atualizacao e Roadmap

**Versao atual: v0.5.1** | **Data: Abril 2026**

---

## PARTE 1: O QUE JA FOI FEITO

### Core da Linguagem
- [x] Lexer com 150+ keywords bilingues (PT/EN)
- [x] Parser recursivo descendente gerando AST completo
- [x] 20 idiomas suportados (PT, EN, ES, FR, DE, IT, ZH, JA, KO, AR, HI, BN, RU, ID, TR, VI, PL, NL, TH, SW)
- [x] Sistema de imports com deteccao circular
- [x] Hot reload (re-exec on file change)
- [x] Comando `flang build` para gerar executavel standalone

### Tipos de Dados (15 tipos)
- [x] texto, texto_longo, numero, dinheiro, email, telefone
- [x] data, booleano, imagem, arquivo, upload, link
- [x] status, senha, enum(valores)

### Relacionamentos
- [x] pertence_a (FK com dropdown automatico)
- [x] tem_muitos (1:N)
- [x] muitos_para_muitos (N:N com join table automatica)

### Banco de Dados
- [x] SQLite, MySQL, PostgreSQL
- [x] Auto-criacao de tabelas
- [x] Auto-migration (colunas novas)
- [x] Connection pooling (25 max, 5 idle)
- [x] Soft delete com restore
- [x] Validacao de campos (obrigatorio, unico, email, telefone)
- [x] Validacao customizada (regras do usuario)
- [x] Paginacao, filtros, busca, ordenacao

### Autenticacao e Seguranca
- [x] JWT (HMAC-SHA256) com expiracao 72h
- [x] bcrypt para senhas
- [x] Role-based access (superadmin, admin, gerente, atendente, etc)
- [x] Rate limiting (100 POST/min por IP, 5 tentativas login = bloqueio 5min)
- [x] SSRF protection no proxy
- [x] XSS escaping em todo HTML
- [x] Path traversal prevention nos imports
- [x] File upload whitelist
- [x] Body size limits (1MB POST/PUT, 64KB eval)
- [x] JWT secret via env variable
- [x] CSV injection protection
- [x] CORS com Authorization header
- [x] Login/registro no frontend

### Frontend
- [x] SPA gerado automaticamente com Tailwind CSS
- [x] Dark/light mode
- [x] Sidebar customizavel
- [x] Dashboard com Chart.js (barra, pizza, doughnut)
- [x] Tabelas com paginacao e busca
- [x] Modais de formulario no body level
- [x] FK dropdowns populados automaticamente
- [x] Enum dropdowns com valores
- [x] Status dropdowns (ativo/inativo/pendente/concluido)
- [x] Textarea para texto_longo
- [x] Toggle para booleano
- [x] Senha oculta na tabela
- [x] Tabs de status/enum para filtrar listas
- [x] WebSocket real-time
- [x] Toast notifications
- [x] Export CSV/JSON
- [x] Upload de arquivos com preview

### Tema e Customizacao
- [x] 5 presets: moderno, simples, elegante, corporativo, claro
- [x] Cores por nome: azul, verde, roxo, etc (14 cores)
- [x] 4 estilos visuais: glassmorphism, flat, neumorphism, minimal
- [x] CSS customizado via bloco `css`
- [x] Controle total: fonte, borda, fundo, card, texto_cor

### Scripting (Interpreter)
- [x] Variaveis (definir x = 10)
- [x] Funcoes com parametros e retorno
- [x] Controle de fluxo: se/senao, enquanto, repetir, para cada
- [x] Operadores: +, -, *, /, ==, !=, >, <, >=, <=, e, ou, nao
- [x] Array literals e indexing (arr[0])
- [x] Object field access (obj.campo)
- [x] Try/catch (tentar/erro)
- [x] 30+ funcoes built-in
- [x] DB queries no interpreter (modelo.listar, .criar, .contar, etc)
- [x] HTTP client (chamar)
- [x] JSON parse/stringify

### Async
- [x] paralelo() — funcoes concorrentes
- [x] esperar() — sleep async
- [x] timeout() — deadline para funcao
- [x] chamar_async() — HTTP paralelo
- [x] consultar_paralelo() — DB queries paralelas

### Integracoes
- [x] WhatsApp via whatsmeow (QR code, envio de mensagens)
- [x] Email SMTP com deteccao de HTML
- [x] Cron jobs (cada N minutos/horas)
- [x] HTTP client para APIs externas
- [x] WebSocket hub com broadcast
- [x] Proxy endpoint (com SSRF protection)

### Customizacao Avancada
- [x] Rotas customizadas (bloco `rotas`)
- [x] Paginas HTML customizadas (bloco `paginas`)
- [x] Sidebar customizavel (bloco `sidebar`)
- [x] Telas customizadas respeitadas pelo frontend

### CLI
- [x] flang run (executar)
- [x] flang check (validar sintaxe)
- [x] flang new (projeto plano)
- [x] flang init (projeto organizado)
- [x] flang build (gerar executavel)
- [x] flang docker (gerar Dockerfile)
- [x] flang version
- [x] flang help

### Testes
- [x] 59 testes (lexer: 14, parser: 16, AST: 9, interpreter: 20)
- [x] Todos passando

### VS Code Extension
- [x] Syntax highlighting com 130+ keywords
- [x] 22 snippets
- [x] Auto-indentacao
- [x] Comentarios com Ctrl+/

### Exemplos
- [x] Loja (modo plano — 1 arquivo)
- [x] Loja (modo organizado — pastas)
- [x] Evoticket (sistema completo — 34 arquivos, 24 modelos, 18 funcoes)

---

## PARTE 2: O QUE FALTA ADICIONAR

### PRIORIDADE CRITICA — Fundamentos

#### 1. Language Server Protocol (LSP)
**O que e:** Servidor que fornece autocomplete, go-to-definition, hover info, e diagnosticos em tempo real para qualquer editor (VS Code, Neovim, JetBrains).
**Por que:** Toda linguagem serio precisa de LSP. Sem ele, a experiencia de desenvolvimento e ruim.
**Implementacao:** Criar um binario `flang-lsp` que implementa o protocolo LSP usando JSON-RPC. O servidor analisa os .fg files e retorna diagnosticos, completions, e hover info.

#### 2. Type Safety e Validacao em Tempo de Compilacao
**O que e:** Verificar tipos em tempo de compilacao/check, nao so em runtime.
**Por que:** Erros de tipo so aparecem quando o app roda. Deviam aparecer no `flang check`.
**Implementacao:** Verificar que campos pertence_a referenciam modelos existentes, que funcoes chamadas existem, que variaveis usadas foram definidas.

#### 3. Package Manager / Registro de Plugins
**O que e:** `flang install <pacote>` para instalar modulos da comunidade.
**Por que:** Permite reusar modelos, telas, e logica entre projetos.
**Implementacao:** Registry central (GitHub-based ou proprio), resolucao de dependencias, versionamento.

#### 4. Testing Framework Built-in
**O que e:** Permitir que o usuario escreva testes nos .fg files.
**Por que:** Nenhuma linguagem seria sobrevive sem testes.
**Sintaxe proposta:**
```
testes

  teste "criar produto com preco negativo deve falhar"
    definir resultado = produto.criar({"nome": "X", "preco": -10})
    verificar resultado == nulo

  teste "calcular total deve multiplicar"
    definir r = calcular_total(10, 5)
    verificar r == 50
```

### PRIORIDADE ALTA — Features Competitivas

#### 5. AI Integration Nativa
**O que e:** Bloco `ia` que conecta com OpenAI/Claude/Gemini direto no .fg.
**Por que:** Em 2026, toda plataforma low-code tem AI integrada. E o diferencial competitivo.
**Sintaxe proposta:**
```
ia
  provedor openai
  modelo "gpt-4o"
  api_key env("OPENAI_KEY")

logica
  funcao classificar_ticket(mensagem)
    retornar ia.completar("Classifique: " + mensagem)
```

#### 6. Scheduler Avancado (Cron Expressions)
**O que e:** Suporte a expressoes cron reais alem de "cada N minutos".
**Sintaxe:** `cron "0 9 * * 1-5"` (9h de segunda a sexta)

#### 7. Migrations Versionadas
**O que e:** Controle de versao do schema do banco, com rollback.
**Por que:** Em producao, ALTER TABLE sem controle e perigoso.
**Implementacao:** Gerar arquivos de migration, `flang migrate`, `flang migrate rollback`.

#### 8. Multi-tenancy Nativo
**O que e:** Isolamento de dados por empresa/tenant automatico.
**Por que:** Todo SaaS precisa. O Evoticket usa manualmente.
**Sintaxe proposta:**
```
sistema meuapp
  multi_tenant: empresa
```
Todas as queries automaticamente filtram por empresa_id.

#### 9. Permissions Granulares
**O que e:** Controle de acesso por modelo/campo/acao, nao so por tela.
**Sintaxe proposta:**
```
permissoes
  atendente pode ver ticket, contato
  atendente pode criar ticket, mensagem
  atendente nao pode deletar ticket
  admin pode tudo
```

#### 10. Webhooks Outbound
**O que e:** Disparar HTTP POST quando eventos acontecem.
**Sintaxe proposta:**
```
webhooks
  quando criar ticket
    enviar para "https://api.slack.com/webhook" 
      corpo {"evento": "novo_ticket", "protocolo": "{protocolo}"}
```

### PRIORIDADE MEDIA — Experiencia do Desenvolvedor

#### 11. REPL Interativo
**O que e:** `flang repl` para testar codigo interativamente.
**Implementacao:** Loop read-eval-print usando o interpreter existente.

#### 12. Debug Mode
**O que e:** `flang run --debug` com breakpoints, step-through, variable inspection.

#### 13. Formatter / Linter
**O que e:** `flang fmt` para formatar codigo automaticamente, `flang lint` para verificar boas praticas.

#### 14. Documentation Generator
**O que e:** `flang docs` gera documentacao HTML da API a partir dos .fg files.

#### 15. GraphQL Endpoint
**O que e:** Gerar endpoint GraphQL alem de REST.
**Por que:** Muitos frontends modernos preferem GraphQL.

#### 16. Realtime Subscriptions
**O que e:** WebSocket subscriptions para modelos especificos.
**Sintaxe proposta:**
```
logica
  quando ticket.criar
    notificar_websocket "novo_ticket" ticket
```

#### 17. File Storage Abstraction
**O que e:** Suporte a S3, MinIO, Google Cloud Storage para uploads.
**Sintaxe proposta:**
```
config
  storage s3
    bucket "meu-bucket"
    regiao "us-east-1"
```

#### 18. Background Jobs
**O que e:** Executar tarefas pesadas em background com retry.
**Sintaxe proposta:**
```
jobs
  job processar_relatorio
    funcao gerar_relatorio_mensal
    retry 3
    timeout 300
```

### PRIORIDADE BAIXA — Futuro

#### 19. Mobile App Generation
**O que e:** Gerar app mobile (React Native ou Flutter) a partir do mesmo .fg.

#### 20. Deploy One-Click
**O que e:** `flang deploy` para deploy em Fly.io, Railway, Render.

#### 21. Visual Editor (Drag & Drop)
**O que e:** Editor visual web para criar .fg files sem escrever codigo.

#### 22. Plugin System com Marketplace
**O que e:** Plugins em Go que estendem o runtime com novos tipos, componentes, e integracoes.

#### 23. Database Seeding
**O que e:** Bloco `sementes` para popular o banco com dados iniciais.
**Sintaxe proposta:**
```
sementes
  usuario
    nome "Admin", email "admin@app.com", senha "123456", role "admin"
    nome "User", email "user@app.com", senha "123456", role "usuario"
  
  categoria
    nome "Roupas"
    nome "Calcados"
```

#### 24. i18n Nativo
**O que e:** Internacionalizacao do app gerado (nao da linguagem, mas do conteudo).
**Sintaxe proposta:**
```
idioma_app
  pt "Bem-vindo" en "Welcome" es "Bienvenido"
```

#### 25. PDF/Report Generation
**O que e:** Gerar PDFs e relatorios a partir de templates.
**Sintaxe proposta:**
```
relatorio vendas_mensal
  titulo "Relatorio de Vendas"
  dados venda.listar()
  formato pdf
```

#### 26. Audit Log
**O que e:** Log automatico de todas as operacoes CRUD com quem, quando, o que.

#### 27. Rate Limiting por Modelo
**O que e:** Limitar requisicoes por modelo, nao so globalmente.

#### 28. Data Import/Export
**O que e:** `flang import dados.csv --modelo produto` para importar dados em massa.

#### 29. SSR (Server-Side Rendering)
**O que e:** Renderizar HTML no servidor para SEO.

#### 30. PWA Support
**O que e:** Transformar o app em Progressive Web App instalavel.

---

## PARTE 3: COMPARACAO COM CONCORRENTES

| Feature | Flang | Wasp | Retool | Budibase | Appsmith |
|---------|:-----:|:----:|:------:|:--------:|:--------:|
| Open source | sim | sim | nao | sim | sim |
| Declarativo | sim | sim | nao | nao | nao |
| Sem codigo | sim | nao | nao | parcial | nao |
| 20 idiomas | sim | nao | nao | nao | nao |
| Multi-banco | sim | parcial | sim | sim | sim |
| Auth built-in | sim | sim | sim | sim | sim |
| WhatsApp | sim | nao | nao | nao | nao |
| Build executavel | sim | nao | nao | nao | nao |
| Self-hosted | sim | sim | nao | sim | sim |
| AI integrado | planejado | nao | sim | nao | sim |
| Mobile | planejado | nao | sim | nao | nao |
| LSP | planejado | nao | n/a | n/a | n/a |
| Preco | gratis | gratis | $10-50/user | $5-15/user | gratis |

### Diferenciais do Flang
1. **Unica linguagem declarativa que suporta 20 idiomas** — nenhum concorrente faz isso
2. **Gera executavel standalone** — nenhum low-code faz isso
3. **WhatsApp integrado nativamente** — unico no mercado
4. **Zero dependencia em runtime** — o .exe roda sozinho, sem Node, sem Python, sem Docker
5. **Descreve com palavras** — nao precisa de drag-and-drop nem codigo

---

## PARTE 4: METRICAS DO PROJETO

| Metrica | Valor |
|---------|-------|
| Linhas de Go | ~12.000 |
| Arquivos Go | 25+ |
| Keywords bilingues | 150+ |
| Idiomas suportados | 20 |
| Funcoes built-in | 30+ |
| Testes automatizados | 59 |
| Plataformas de build | 6 (Win/Linux/Mac x amd64/arm64) |
| Exemplos inclusos | 3 (loja plano, loja organizado, evoticket) |
| VS Code snippets | 22 |
| Releases publicados | 2 (v0.5.0, v0.5.1) |

---

## PARTE 5: PROXIMAS VERSOES PLANEJADAS

### v0.6.0 — Developer Experience
- [ ] LSP (Language Server Protocol)
- [ ] REPL interativo (`flang repl`)
- [ ] Formatter (`flang fmt`)
- [ ] Validacao de tipos em tempo de compilacao
- [ ] Database seeding (bloco `sementes`)
- [ ] Audit log automatico

### v0.7.0 — AI & Intelligence
- [ ] AI integration nativa (OpenAI, Claude, Gemini)
- [ ] Classificacao automatica via IA
- [ ] Chatbot builder no .fg
- [ ] RAG (knowledge base) built-in

### v0.8.0 — Enterprise
- [ ] Multi-tenancy nativo
- [ ] Permissoes granulares por modelo/campo
- [ ] Migrations versionadas
- [ ] Background jobs com retry
- [ ] Webhooks outbound

### v0.9.0 — Ecosystem
- [ ] Package manager (`flang install`)
- [ ] Plugin system
- [ ] GraphQL endpoint
- [ ] Deploy one-click
- [ ] Testing framework built-in

### v1.0.0 — Production Ready
- [ ] Mobile app generation
- [ ] Visual editor (drag & drop)
- [ ] SSR para SEO
- [ ] PWA support
- [ ] PDF/report generation
- [ ] i18n do app gerado
- [ ] Marketplace de plugins

---

*Documento gerado em Abril 2026. Baseado em pesquisa de mercado e analise de concorrentes.*
