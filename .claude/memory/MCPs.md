# MCPs Configurados

## Servidores MCP Ativos

Este projeto utiliza os seguintes MCP servers para interações ricas:

### 1. Playwright MCP
- **Pacote**: `@playwright/mcp@latest`
- **Uso**: Automação de navegador e testes visuais
- **Comando**: `npx -y @playwright/mcp@latest`

### 2. GitHub MCP
- **Pacote**: `@modelcontextprotocol/server-github@latest`
- **Uso**: Interações ricas com o repositório GitHub (apenas quando necessário)
- **Autenticação**: Via `GITHUB_TOKEN` (variável de ambiente do sistema)
- **Comando**: `npx -y @modelcontextprotocol/server-github@latest`
- **Nota**: O pacote oficial do GitHub está em transição para `github/github-mcp-server`

### 3. Context7 MCP
- **Pacote**: `@upstash/context7-mcp@latest`
- **Uso**: Busca de documentação atualizada e boas práticas de código (focado em 2026)
- **Comando**: `npx -y @upstash/context7-mcp@latest`
- **Funcionalidade**: Fornece acesso a documentação versionada de bibliotecas e frameworks atualizados
- **Uso obrigatório**: Toda implementação DEVE consultar o Context7 para boas práticas recentes

## Configuração

Arquivo: `.claude/.mcp.json`

```json
{
  "servers": {
    "playwright": {
      "command": "npx",
      "args": ["-y", "@playwright/mcp@latest"]
    },
    "github": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-github@latest"],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "${GITHUB_TOKEN}"
      }
    },
    "context7": {
      "command": "npx",
      "args": ["-y", "@upstash/context7-mcp@latest"]
    }
  }
}
```

## Pré-requisitos

- `GITHUB_TOKEN` deve estar configurado como variável de ambiente do sistema
- Node.js/npm disponível para execução via npx

## Data de Configuração

2026-04-26
