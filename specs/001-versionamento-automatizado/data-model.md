# Modelo de Dados: Versionamento Automatizado

**Data**: 2026-04-19

## Nota

Esta funcionalidade é **stateless**. Não há persistência de dados própria - todo o estado é mantido via:

- **Git tags**: identificadores imutáveis de versões
- **GitHub releases**: metadados associados às tags
- **Git commit log**: histórico usado para gerar changelogs

## Entidades Virtuais

### Release

Representação de uma versão do software.

| Campo | Tipo | Origem |
|-------|------|--------|
| tag_name | string | Input do usuário ou auto-gerado (semver) |
| target_commitish | string | SHA do commit na main |
| name | string | "Release {tag_name}" ou customizado |
| body | markdown | Changelog gerado dos commits |
| draft | boolean | false (sempre publicado diretamente) |
| prerelease | boolean | false (não usamos pré-releases) |
| created_at | datetime | Gerado pelo GitHub |
| published_at | datetime | Gerado pelo GitHub |

### Tag

Ponteiro imutável para um commit específico.

| Campo | Tipo | Origem |
|-------|------|--------|
| ref | string | refs/tags/{nome} |
| object.sha | string | SHA do commit apontado |
| object.type | string | "commit" |

### Changelog Entry

Item individual do changelog.

| Campo | Tipo | Exemplos |
|-------|------|----------|
| type | string | feat, fix, docs, refactor, etc. |
| scope | string | opcional, ex: "theme", "i18n" |
| description | string | descrição do commit |
| breaking | boolean | true se BREAKING CHANGE ou ! |

## Fluxo de Dados

```
Push na main
    ↓
Análise de commits (conventional-commits)
    ↓
Cálculo de próxima versão (semver)
    ↓
Criação da tag (git tag)
    ↓
Deploy para GitHub Pages
    ↓
Criação da release (GitHub API)
    ↓
Notificação (implícita via GitHub UI)
```

## Permissões Necessárias

| Recurso | Permissão | Uso |
|---------|-----------|-----|
| contents | write | Criar tags e releases |
| actions | read | Verificar status do deploy |
| metadata | read | Acessar informações do repo |
