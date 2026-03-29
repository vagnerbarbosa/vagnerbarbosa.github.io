# vagnerbarbosa.github.io

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://golang.org/)
[![GitHub Pages](https://img.shields.io/badge/GitHub%20Pages-Active-222?logo=github)](https://vagnerbarbosa.github.io)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

> Personal website and portfolio of Vagner Barbosa - Software Engineer

## Visão Geral

Site pessoal minimalista inspirado no design clean e tipografia do [annamonaco.co](https://annamona.co). Design responsivo, bilíngue (PT/EN) e com foco em acessibilidade e performance.

## Características

- **Design minimalista** - Layout clean inspirado no annamonaco.co
- **Bilíngue** - Suporte completo a Português e Inglês com toggle de idioma
- **Tema claro/escuro** - Alternância automática de tema
- **Acessibilidade** - Suporte a preferências de movimento reduzido
- **Segurança** - Headers de segurança (CSP, X-Frame-Options)
- **Performance** - Site 100% estático, sem dependências de runtime

## Tecnologias

### Core
| Tecnologia | Versão | Descrição |
|------------|--------|-----------|
| [Go](https://golang.org/) | 1.21+ | Linguagem do gerador de site estático |
| [html/template](https://pkg.go.dev/html/template) | built-in | Templates seguros para HTML |

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
- **Domínio customizado**: https://vagnerbarbosa.com

## Estrutura do Projeto

```
vagnerbarbosa.github.io/
├── cmd/
│   └── generator/            # Gerador de site estático em Go
│       ├── main.go           # Entry point do gerador
│       └── main_test.go      # Testes unitários do gerador
├── internal/
│   └── config/
│       ├── config.go         # Parser de configuração YAML
│       └── config_test.go    # Testes unitários de configuração
├── templates/                # Templates Go (html/template)
│   ├── index.html            # Template principal
│   └── partials/             # Componentes reutilizáveis
│       ├── head.html
│       ├── header.html
│       ├── about.html
│       ├── experience.html
│       ├── skills.html
│       ├── education.html
│       ├── contact.html
│       └── footer.html
├── assets/                   # Assets estáticos
│   ├── css/
│   │   └── main.css          # CSS principal
│   ├── js/                   # JavaScript modular
│   │   ├── utils.js
│   │   ├── theme.js
│   │   ├── i18n.js
│   │   ├── accordion.js
│   │   └── app.js
│   ├── fonts/                # Fontes (devicon, fontawesome)
│   └── favicon.png
├── config.yaml               # Configuração do site (novo)
├── go.mod                    # Módulo Go
├── go.sum                    # Checksums de dependências
├── .github/
│   └── workflows/
│       └── deploy.yml        # CI/CD para GitHub Pages
├── CNAME                     # Configuração de domínio
├── LICENSE                   # Licença MIT
└── README.md                 # Este arquivo
```

## Desenvolvimento

### Pré-requisitos
- Go 1.21 ou superior
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
| `go run cmd/generator/main.go` | Gera o site em `public/` |
| `go build -o generator cmd/generator/main.go` | Compila o gerador |
| `./generator` | Executa o gerador compilado |

O site gerado estará disponível em `public/index.html`

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
      period: "Jan 2020 – Presente"
      details:
        - "Descrição da atividade 1"
        - "Descrição da atividade 2"
  education:
    - title: "Curso"
      school: "Instituição"
      period: "2018 – 2022"
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

1. **Build**: Instala Go, baixa dependências e executa o gerador
2. **Deploy**: Publica o conteúdo de `public/` no GitHub Pages

## Migração de Jekyll para Go

Este projeto foi migrado de Jekyll (Ruby) para um gerador estático em Go. As principais mudanças:

| Aspecto | Antes (Jekyll) | Agora (Go) |
|---------|----------------|------------|
| Linguagem | Ruby | Go |
| Templates | Liquid | html/template |
| Configuração | `_config.yml` | `config.yaml` |
| Build local | `bundle exec jekyll serve` | `go run cmd/generator/main.go` |
| CI/CD | GitHub Pages nativo | GitHub Actions + Go |
| Performance | ~2-3s build | ~0.5s build |
| Dependências | Ruby + gems | Apenas Go standard library + yaml |

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
