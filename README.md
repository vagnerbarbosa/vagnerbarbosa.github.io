# vagnerbarbosa.github.io

[![Jekyll](https://img.shields.io/badge/Jekyll-4.0-CC0000?logo=jekyll)](https://jekyllrb.com/)
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
- **Segurança** - Headers de segurança (CSP, HSTS, X-Frame-Options) via Cloudflare
- **Performance** - JavaScript modular com carregamento defer e SRI hashes

## Tecnologias

### Core
| Tecnologia | Versão | Descrição |
|------------|--------|-----------|
| [Jekyll](https://jekyllrb.com/) | 4.x | Gerador de sites estáticos |
| [Ruby](https://www.ruby-lang.org/) | 3.2+ | Linguagem base do Jekyll |

### Frontend
| Tecnologia | Uso |
|------------|-----|
| HTML5 | Estrutura semântica |
| CSS3 | Estilização moderna com variáveis |
| JavaScript (ES6+) | Interatividade modular |

### Módulos JavaScript
| Arquivo | Descrição |
|---------|-----------|
| `utils.js` | Utilitários e polyfills |
| `theme.js` | Gerenciamento de tema claro/escuro |
| `i18n.js` | Sistema de internacionalização PT/EN |
| `accordion.js` | Componente de acordeão para experiências |
| `app.js` | Inicialização e orquestração |

### Deploy
- **GitHub Pages**: Hospedagem gratuita integrada ao GitHub
- **Cloudflare**: CDN e headers de segurança
- **Domínio customizado**: https://vagnerbarbosa.com

## Estrutura do Projeto

```
vagnerbarbosa.github.io/
├── _config.yml              # Configuração do Jekyll
├── _includes/               # Componentes reutilizáveis
│   ├── about.html          # Seção principal com acordeão
│   ├── footer.html         # Rodapé minimalista
│   ├── head.html           # Meta tags e headers de segurança
│   └── header.html         # Cabeçalho com toggles
├── _layouts/                # Templates de página
│   └── default.html        # Layout principal único
├── assets/                  # Assets compilados
│   ├── css/main.css        # CSS principal
│   ├── js/                 # JavaScript modular
│   │   ├── utils.js
│   │   ├── theme.js
│   │   ├── i18n.js
│   │   ├── accordion.js
│   │   └── app.js
│   └── favicon.png
├── scripts/                 # Scripts auxiliares
│   ├── generate-sri.sh     # Gerador de SRI hashes
│   └── sri-hashes.txt      # Hashes atuais
├── CNAME                    # Configuração de domínio
├── index.html              # Página inicial
├── LICENSE                 # Licença MIT
└── README.md               # Este arquivo
```

## Desenvolvimento

### Pré-requisitos
- Ruby 3.2 ou superior
- Bundler (`gem install bundler`)

### Instalação

1. Clone o repositório:
```bash
git clone https://github.com/vagnerbarbosa/vagnerbarbosa.github.io.git
cd vagnerbarbosa.github.io
```

2. Instale as dependências Ruby:
```bash
bundle install
```

### Comandos

| Comando | Descrição |
|---------|-----------|
| `bundle exec jekyll serve` | Inicia servidor de desenvolvimento |
| `bundle exec jekyll serve --livereload` | Com hot reload |

O site estará disponível em `http://localhost:4000`

## Configuração

### Site (`_config.yml`)
Edite as informações pessoais:

```yaml
username: Seu Nome
user_description: Sua descrição
user_title: Seu Título
email: seu@email.com
```

### Internacionalização (`assets/js/i18n.js`)
Para atualizar textos em PT/EN, edite o objeto `translations` no arquivo `i18n.js`.

### Experiências (`_includes/about.html`)
Para atualizar experiências profissionais, edite as seções dentro do arquivo `about.html`.

### SRI Hashes
Ao modificar arquivos JS, gere novos hashes:

```bash
# Linux/Mac
bash scripts/generate-sri.sh

# Windows (PowerShell)
scripts/generate-sri.ps1
```

Depois atualize os hashes no `_layouts/default.html`.

## Segurança

Os seguintes headers de segurança são configurados via Cloudflare:

- **CSP (Content-Security-Policy)** - Restrições de conteúdo
- **HSTS** - Força HTTPS
- **X-Frame-Options: DENY** - Previne clickjacking
- **X-Content-Type-Options: nosniff** - Evita MIME sniffing
- **X-XSS-Protection** - Proteção XSS do navegador
- **Referrer-Policy** - Política de referrer

## Deploy

O deploy é automático via GitHub Pages:
1. Faça push para a branch `master`
2. O GitHub Pages irá buildar e publicar automaticamente
3. O Cloudflare aplica os headers de segurança
4. O site fica disponível em `https://vagnerbarbosa.com`

## Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

---

Desenvolvido por Vagner Barbosa
