# vagnerbarbosa.github.io

[![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)](https://golang.org/)
[![GitHub Pages](https://img.shields.io/badge/GitHub%20Pages-Active-222?logo=github)](https://vagnerbarbosa.github.io)
[![Version](https://img.shields.io/github/v/release/vagnerbarbosa/vagnerbarbosa.github.io?label=vers%C3%A3o)](https://github.com/vagnerbarbosa/vagnerbarbosa.github.io/releases/latest)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

> Personal website and portfolio of Vagner Barbosa - Software Engineer

## Visão Geral

Site pessoal minimalista inspirado no design clean e tipografia do [annamonaco.co](https://annamona.co). Design responsivo, bilíngue (PT/EN) e com foco em acessibilidade e performance.

## Características

- **Design minimalista** - Layout clean inspirado no [annamona.co](https://annamona.co)
- **Bilíngue** - Suporte completo a Português e Inglês com detecção automática do navegador
- **Tema escuro padrão** - Primeira visita sempre inicia no tema escuro (pode ser alterado)
- **Ícones de tema intuitivos** - Lua no modo claro (ativar escuro), sol no modo escuro (ativar claro)
- **Animação fade-in** - Efeito suave de entrada no carregamento da página (estilo annamona.co)
- **Acessibilidade** - Suporte a preferências de movimento reduzido (`prefers-reduced-motion`)
- **Segurança** - Sanitização de conteúdo, permissões mínimas em workflows
- **Performance** - Site 100% estático, sem dependências de runtime, minificação automática
- **Versionamento automatizado** - Tags e releases geradas automaticamente via Conventional Commits
- **LinkedIn Import CLI** - Ferramenta para importar experiências, educação e certificações do LinkedIn

## Tecnologias

### Core
| Tecnologia | Versão | Descrição |
|------------|--------|-----------|
| [Go](https://golang.org/) | 1.21+ | Linguagem do gerador de site estático |
| [html/template](https://pkg.go.dev/html/template) | built-in | Templates seguros para HTML |
| [tdewolff/minify](https://github.com/tdewolff/minify) | v2.24+ | Minificação de JS, CSS, HTML e JSON |

### Frontend
| Tecnologia | Uso |
|------------|-----|
| HTML5 | Estrutura semântica |
| CSS3 | Estilização moderna com variáveis CSS |
| JavaScript (ES6+) | Interatividade modular |

### Módulos JavaScript
| Arquivo | Descrição |
|---------|-----------|
| `utils.js` | Utilitários e storage seguro |
| `theme.js` | Gerenciamento de tema claro/escuro |
| `i18n.js` | Sistema de internacionalização PT/EN |
| `accordion.js` | Componente de acordeão para seções |
| `app.js` | Inicialização e orquestração |

### Deploy
- **GitHub Actions**: Build automático com Go
- **GitHub Pages**: Hospedagem gratuita

## Spec Kit (Speckit)

Este projeto utiliza [Spec Kit](https://speckit-spec.org/) - um toolkit de desenvolvimento dirigido por especificações da GitHub.

### Comandos disponíveis

| Comando | Descrição |
|---------|-----------|
| `/speckit-constitution` | Estabelece princípios do projeto |
| `/speckit-specify` | Cria especificação baseline |
| `/speckit-clarify` | Esclarece áreas ambíguas antes do planejamento |
| `/speckit-plan` | Cria plano de implementação técnica |
| `/speckit-tasks` | Gera tarefas acionáveis |
| `/speckit-analyze` | Análise de consistência entre artefatos |
| `/speckit-checklist` | Gera checklists de qualidade |
| `/speckit-implement` | Executa implementação |

### Git Workflow

| Comando | Descrição |
|---------|-----------|
| `/speckit-git.initialize` | Inicializa repositório Git |
| `/speckit-git.feature` | Cria nova feature branch |
| `/speckit-git.commit` | Commit automático seguindo convenções |
| `/speckit-git.remote` | Configura remote |
| `/speckit-git.validate` | Valida estado do repositório |
- **Domínio customizado**: https://vagnerbarbosa.com

## Estrutura do Projeto

```
vagnerbarbosa.github.io/
├── cmd/
│   ├── generator/            # Gerador de site estático em Go
│   │   ├── main.go           # Entry point do gerador
│   │   ├── main_test.go      # Testes unitários
│   │   └── e2e_test.go       # Testes End-to-End (pipeline completo)
│   └── import-linkedin/      # Ferramenta CLI para importar dados do LinkedIn
│       ├── main.go           # Entry point da ferramenta
│       ├── commands/         # Comandos CLI (import, validate, version)
│       ├── internal/         # Parser, models, comparator, transformer, UI
│       │   └── integration/  # Testes de integração (Golden Files)
│       └── testdata/         # Dados de teste (CSVs malformados, etc.)
├── internal/
│   └── config/
│       ├── config.go         # Parser de configuração YAML
│       └── config_test.go    # Testes unitários de configuração
├── templates/                # Templates Go (html/template)
│   ├── index.html            # Template principal
│   └── partials/             # Componentes reutilizáveis
│       ├── head.html
│       ├── header.html       # Header com tema escuro padrão e ícones sol/lua
│       ├── about.html
│       ├── experience.html
│       ├── skills.html
│       ├── education.html
│       ├── certifications.html # Template de certificações
│       └── footer.html
├── assets/                   # Assets estáticos
│   ├── css/
│   │   └── main.css          # CSS principal com animações fade-in
│   ├── js/                   # JavaScript modular
│   │   ├── utils.js          # Utilitários e storage seguro
│   │   ├── theme.js          # Gerenciamento de tema (escuro padrão)
│   │   ├── i18n.js           # Sistema de internacionalização PT/EN
│   │   ├── accordion.js      # Componente de acordeão
│   │   ├── fadein.js         # Animação fade-in (estilo annamona.co)
│   │   └── app.js            # Inicialização e orquestração
│   ├── fonts/                # Fontes (devicon, fontawesome)
│   └── favicon.png
├── docs/                     # Documentação
│   ├── LINKEDIN-IMPORT.md    # Guia da ferramenta LinkedIn Import CLI
│   ├── VERSIONING.md         # Documentação de versionamento
│   └── TESTING.md            # Guia de estratégia de testes
├── config.yaml               # Configuração do site
├── go.mod                    # Módulo Go
├── go.sum                    # Checksums de dependências
├── SECURITY.md               # Política de segurança do projeto
├── .github/
│   ├── scripts/              # Scripts para GitHub Actions
│   │   ├── package.json      # Dependências Node.js (semantic-release)
│   │   └── package-lock.json
│   └── workflows/
│       ├── deploy.yml        # CI/CD para GitHub Pages
│       ├── release.yml       # Versionamento automatizado com tags
│       └── create-retro-tag.yml  # Criação de tags retroativas
├── scripts/                  # Scripts auxiliares
│   └── check_coverage.sh     # Script de validação de cobertura de testes
├── hooks/                    # Git hooks
│   └── pre-push              # Hook para prevenir push direto na main
├── la-linkedin/              # Dados exportados do LinkedIn (importados via CLI)
│   └── testdata             # Dados de teste para importação
├── specs/                    # Especificações de funcionalidades (Spec Kit)
│   ├── 001-versionamento-automatizado/
│   ├── 002-theme-language-adjustments/
│   ├── 003-linkedin-import/
│   ├── 004-tech-stack-extraction/
│   ├── 005-test-coverage-total/
│   └── 006-integration-e2e-tests/ # Implementação de testes de integração e E2E
├── .claude/
│   ├── CLAUDE.md             # Regras de colaboração
│   └── skills/               # Skills do Spec Kit
├── .specify/                 # Configuração do Spec Kit
│   ├── extensions.yml        # Configuração de extensões
│   ├── init-options.json     # Opções de inicialização
│   ├── feature.json          # Feature atual
│   ├── templates/            # Templates de especificação
│   └── scripts/              # Scripts auxiliares do Spec Kit
├── public/                   # Site gerado (output do generator)
│   ├── index.html            # Página principal gerada
│   ├── assets/               # Assets processados e minificados
│   ├── sitemap.xml           # Sitemap para SEO
│   ├── robots.txt            # Configuração de crawl
│   └── site.webmanifest      # Configuração PWA
├── CNAME                     # Configuração de domínio
├── LICENSE                   # Licença MIT
└── README.md                 # Este arquivo
```

## Desenvolvimento

### Pré-requisitos
- Go 1.25 ou superior
- Git

### Instalação

1. Clone o repositório:
```bash
git clone https://github.com/vagnerbarbosa/vagnerbarbosa.github.io.git
cd vagnerbarbosa.github.io
```

2. Instale as dependências Go:
```bash
go mod download
```

### Comandos

| Comando | Descrição |
|---------|-----------|
| `go run cmd/generator/main.go` | Gera o site em `public/` (com minificação automática) |
| `go build -o generator cmd/generator/main.go` | Compila o gerador |
| `./generator` | Executa o gerador compilado |
| `go run cmd/import-linkedin/main.go import` | Importa dados do LinkedIn (modo interativo) |
| `go run cmd/import-linkedin/main.go import --dry-run` | Visualiza importação sem aplicar |
| `go run cmd/import-linkedin/main.go validate` | Valida os arquivos CSV do LinkedIn |
| `go test ./...` | Executa testes unitários |
| `go test -v -tags=integration ./...` | Executa testes de integração (Golden Files) |
| `go test -v -tags=e2e ./...` | Executa testes de ponta a ponta (E2E) |

O site gerado estará disponível em `public/index.html`.

### Minificação Automática

Durante o build, todos os assets são automaticamente minificados:

| Tipo | Economia típica |
|------|-----------------|
| JavaScript | ~60% |
| CSS | ~25% |
| HTML | ~14% |
| JSON | ~30-40% |

Exemplo de output durante o build:
```
HTML minified: 15563 bytes -> 13378 bytes (saved 2185 bytes, 14.0%)
Minified js/utils.js: 4901 -> 1438 bytes (saved 3463 bytes, 70.7%)
Minified css/main.css: 6308 -> 4712 bytes (saved 1596 bytes, 25.3%)
Site generated successfully in public/
```

## Qualidade e Testes

O projeto utiliza uma estratégia de testes em múltiplas camadas para garantir que os dados importados do de LinkedIn sejam renderizados corretamente:

- **Testes Unitários**: Validam componentes isolados do parser e gerador.
- **Testes de Integração**: Usam **Golden Files** para validar a deterministicidade do pipeline CSV $\rightarrow$ YAML.
- **Testes E2E (End-to-End)**: Validam o fluxo completo CSV $\rightarrow$ YAML $\rightarrow$ HTML, verificando a presença de marcadores de dados no output final.

Para executar todos os testes:
```bash
go test -v ./...
```

Mais detalhes em [docs/TESTING.md](docs/TESTING.md).

## Configuração

### Site (`config.yaml`)
Edite as informações pessoais:

```yaml
site:
  title: "Seu Nome"
  username: "Seu Nome"
  user_description: "Sua descrição"
  user_title: "Seu Título"
  email: "seu@email.com"
```

### Conteúdo (`config.yaml`)
Atualize experiências, educação e tecnologias:

```yaml
content:
  about:
    pt: "Texto em português..."
    en: "Texto em inglês..."
  experiences:
    - title: "Título do cargo"
      company: "Nome da empresa"
      start_date: "2020-01-01"
      end_date: "2022-12-31"
      description:
        - "Descrição da atividade 1"
        - "Descrição da atividade 2"
      tech_stack: "Go, Kubernetes"
      location: "Remoto"
  education:
    - institution: "Universidade X"
      degree: "B.S. Computer Science"
      field: "Ciência da Computação"
      start_date: "2016-01-01"
      end_date: "2020-12-31"
      description:
        - "Foco em Sistemas Distribuídos"
  certifications:
    - name: "AWS Certified Solutions Architect"
      organization: "Amazon Web Services"
      issue_date: "2021-01-01"
      credential_id: "AWS-12345"
      credential_url: "https://aws.amazon.com/cert"
  technologies:
    - "Go"
    - "JavaScript"
```

## Deploy

O deploy é automático via GitHub Actions:

1. Faça push para a branch `main`
2. O GitHub Actions executa o build com Go
3. O site gerado em `public/` é publicado no GitHub Pages
4. O site fica disponível em `https://vagnerbarbosa.com`

### Workflow de CI/CD

O arquivo `.github/workflows/deploy.yml` configura:

1. **Build**: Instala Go, baixa dependências, executa testes e executa o gerador
2. **Minificação**: Durante o build, JS, CSS, HTML e JSON são automaticamente minificados
3. **Deploy**: Publica o conteúdo otimizado de `public/` no GitHub Pages

## História do Projeto

| Era | Tecnologia | Período |
|-----|------------|---------|
| v1.x | Jekyll (inicial) | Fev 2015 – 2019 |
| v2.x | Jekyll (estabilizado) | 2019 – Mar 2026 |
| v3.x | Go gerador estático | Mar 2026 – Abr 2026 |
| **v4.x** | **Go gerador estático + Versionamento + Ajustes de UX** | **Abr 2026 – presente** |

O site nasceu em 2015 como um portfolio Jekyll. Em março de 2026, foi completamente reescrito em Go, motivado pela adoção do design minimalista do [annamona.co](https://annamona.co) e como oportunidade de treinar Go em um projeto real.

Em abril de 2026, foi implementado o sistema de versionamento automatizado com tags e releases, marcando a transição para a v4.x. Este sistema detecta automaticamente merges na branch main, cria tags semver baseadas em Conventional Commits e gera releases no GitHub com changelog em português.

## Migração v2 $\rightarrow$ v3 (Jekyll para Go)

As principais mudanças na migração para v3.0.0:

| Aspecto | Antes (Jekyll) | Agora (Go) |
|---------|----------------|------------|
| Linguagem | Ruby | Go |
| Templates | Liquid | html/template |
| Configuração | `_config.yml` | `config.yaml` |
| Build local | `bundle exec jekyll serve` | `go run cmd/generator/main.go` |
| CI/CD | GitHub Pages nativo | GitHub Actions + Go |
| Performance | ~2-3s build | ~0.5s build |
| Minificação | Requer plugins | Integrado (minify) |
| Dependências | Ruby + gems | Go + yaml + minify |

## Segurança

Os seguintes headers de segurança são configurados via meta tags:

- **CSP (Content-Security-Policy)** - Restrições de conteúdo
- **X-Frame-Options: DENY** - Previne clickjacking
- **X-Content-Type-Options: nosniff** - Evita MIME sniffing
- **X-XSS-Protection** - Proteção XSS do navegador
- **Referrer-Policy** - Política de referrer

## Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

---

Desenvolvido por Vagner Barbosa
