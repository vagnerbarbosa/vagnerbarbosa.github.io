# Plano de Implementação: Versionamento Automatizado com Tags e Releases

**Branch**: `001-versionamento-automatizado` | **Data**: 2026-04-19 | **Spec**: [spec.md](spec.md)
**Input**: Especificação da funcionalidade de `/specs/001-versionamento-automatizado/spec.md`

**Nota**: Este modelo é preenchido pelo comando `/speckit.plan`. Veja `.specify/templates/plan-template.md` para o fluxo de execução.

## Resumo

Implementar sistema de versionamento automatizado baseado em GitHub Actions que detecta pushes na main, analisa commits seguindo Conventional Commits, gera tags semver automaticamente e cria releases no GitHub com changelog no formato Keep a Changelog em português. Incluir workflow manual para criação de tags retroativas v1.0.0 e v2.0.0.

## Contexto Técnico

**Linguagem/Versão**: YAML (GitHub Actions), Bash scripts auxiliares  
**Dependências Principais**:  
- `semantic-release` ou action GitHub para análise de commits e geração de versões  
- GitHub CLI (`gh`) para criação de releases  
- Action `actions/github-script` ou `softprops/action-gh-release` para releases  
**Armazenamento**: N/A - stateless, usa GitHub API  
**Testes**: Testes manuais em fork ou branch de teste  
**Plataforma Alvo**: GitHub Actions (ubuntu-latest)  
**Tipo de Projeto**: CI/CD Pipeline / DevOps Automation  
**Metas de Performance**: Workflow completo < 2 minutos  
**Restrições**: Depende de token GitHub com permissões `contents:write` e `actions:write`  
**Escala/Alcance**: Um repositório, múltiplos contribuidores, todas as merges na main

## Verificação da Constituição

*PORTÃO: Deve passar antes da pesquisa da Fase 0. Re-verificar após o design da Fase 1.*

| Princípio | Status | Justificativa |
|-----------|--------|---------------|
| I. Minimalismo Intencional | ✅ Passa | Zero dependências de runtime adicionais, usa apenas GitHub Actions nativos |
| II. Bilíngue por Padrão | ✅ Passa | Changelog gerado em português conforme especificado |
| III. Estabilidade Visual | ✅ Passa | Não afeta UI do site, apenas infraestrutura de release |
| IV. Build Reprodutível | ✅ Passa | Tags imutáveis garantem rastreabilidade completa de cada deploy |
| V. Código como Documentação | ✅ Passa | Workflows documentam o processo de release |

## Estrutura do Projeto

### Documentação (esta funcionalidade)

```text
specs/001-versionamento-automatizado/
├── plan.md              # Este arquivo
├── research.md          # Output da Fase 0
├── data-model.md        # Output da Fase 1
├── quickstart.md        # Output da Fase 1
├── contracts/           # Output da Fase 1
└── tasks.md             # Output da Fase 2
```

### Código Fonte (raiz do repositório)

```text
.github/
├── workflows/
│   ├── deploy.yml              # Workflow existente (modificado)
│   ├── release.yml             # NOVO: Criação automática de tags e releases
│   └── create-retro-tag.yml    # NOVO: Workflow manual para tags retroativas
└── scripts/
    └── generate-changelog.sh   # NOVO: Script para gerar changelog em português
```

**Decisão de Estrutura**: Workflows GitHub Actions na pasta `.github/workflows/`. Script auxiliar para tradução do changelog. Sem alterações no código do site em si.

## Rastreamento de Complexidade

> Sem violações da constituição. Toda a complexidade está justificada:

| Complexidade | Por Que Necessário | Alternativa Mais Simples |
|--------------|-------------------|--------------------------|
| Workflow separado para release | Desacoplar criação de tag do deploy permite retry independente | Não há alternativa mais simples que mantenha idempotência |
| Script customizado de changelog | Ferramentas padrão geram em inglês apenas | Usar changelog em inglês violaria Princípio II |

---

## Fase 0: Pesquisa e Decisões

### Decisões Técnicas

**Analisador de Conventional Commits**:
- **Decisão**: Usar `semantic-release` via GitHub Action `cycjimmy/semantic-release-action`
- **Racional**: Mais maduro, amplamente usado, suporta customização de changelog
- **Alternativas consideradas**: 
  - `google-github-actions/release-please-action` - menos flexível para tradução
  - Script bash customizado - mais manutenção, erro-prone

**Geração de Changelog em Português**:
- **Decisão**: Script bash customizado que processa saída do semantic-release
- **Racional**: semantic-release não tem i18n nativa, necessário adaptador
- **Implementação**: Mapear tipos de commit (feat→Adicionado, fix→Corrigido, etc.)

**Criação de Releases**:
- **Decisão**: Action `softprops/action-gh-release`
- **Racional**: Simples, bem mantida, suporta arquivos anexos se necessário futuramente
- **Alternativas**: `actions/create-release` (deprecated), GitHub CLI direto

**Token de Autenticação**:
- **Decisão**: `GITHUB_TOKEN` padrão (auto-provisionado)
- **Racional**: Permissões suficientes para tags e releases no mesmo repo
- **Nota**: Se futuramente precisar criar releases em repo diferente, necessitará PAT

---

## Fase 1: Design e Contratos

### Modelo de Dados

Não aplica-se - sistema é stateless. Estado mantido em:
- Git tags (versões)
- GitHub releases (metadados)
- Git log (histórico de commits)

### Contratos

#### Contract: Trigger de Release

```yaml
# Evento que dispara o workflow release.yml
on:
  push:
    branches: [main]
  workflow_dispatch:
    inputs:
      sha:
        description: 'Commit SHA para tag retroativa'
        required: true
      tag:
        description: 'Nome da tag (ex: v1.0.0)'
        required: true
```

#### Contract: Saída do Semantic Release

```json
{
  "version": "3.1.0",
  "type": "minor",
  "notes": "## Features\n\n* novo recurso..."
}
```

#### Contract: Changelog Traduzido

```markdown
## [3.1.0] - 2026-04-19

### Adicionado
- Nova funcionalidade X

### Corrigido
- Bug no tema escuro

### Alterado
- Melhorias de performance
```

---

## Quickstart

### Para desenvolvedores

1. Commits devem seguir [Conventional Commits](https://www.conventionalcommits.org/):
   - `feat: adiciona suporte a X` → bump minor
   - `fix: corrige bug em Y` → bump patch
   - `feat!: altera API` ou `BREAKING CHANGE:` → bump major

2. Ao fazer merge na main:
   - Workflow `release.yml` executa automaticamente
   - Tag é criada (ex: v3.1.0)
   - Deploy ocorre
   - Release é criada com changelog

### Para criar tags retroativas

1. Acesse Actions → "Create Retroactive Tag"
2. Clique "Run workflow"
3. Informe:
   - SHA: commit hash (ex: `b81034a`)
   - Tag: nome da versão (ex: `v1.0.0`)
4. Execute

---

**Plano criado**: 2026-04-19 | **Pronto para**: `/speckit-tasks`
