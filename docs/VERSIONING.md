# Versionamento e Releases

Este documento descreve o processo de versionamento automatizado do projeto.

## Visão Geral

O projeto utiliza um sistema de versionamento automatizado baseado em [Conventional Commits](https://www.conventionalcommits.org/pt-br/v1.0.0/) e [Semantic Versioning](https://semver.org/lang/pt-BR/).

## Como Funciona

### 1. Commits Convencionais

Todo commit deve seguir o formato:

```
<tipo>[(escopo opcional)]: <descrição>

[corpo opcional]

[rodapé(s) opcional(is)]
```

### 2. Tipos de Commit e Versionamento

| Tipo | Quando Usar | Bump de Versão |
|------|-------------|----------------|
| `feat` | Nova funcionalidade | Minor (x.Y.z) |
| `fix` | Correção de bug | Patch (x.y.Z) |
| `docs` | Documentação | Patch (x.y.Z) |
| `style` | Formatação | Patch (x.y.Z) |
| `refactor` | Refatoração | Patch (x.y.Z) |
| `perf` | Performance | Patch (x.y.Z) |
| `test` | Testes | Patch (x.y.Z) |
| `chore` | Manutenção | Patch (x.y.Z) |
| `BREAKING CHANGE` | Quebra de compatibilidade | Major (X.y.z) |

### 3. Processo Automático

1. **Push na main**: Ao fazer merge de uma PR na branch `main`
2. **Análise**: O workflow `release.yml` analisa os commits
3. **Cálculo**: Determina o próximo número de versão semver
4. **Tag**: Cria uma tag Git no formato `vX.Y.Z`
5. **Deploy**: Aguarda o deploy no GitHub Pages
6. **Release**: Cria uma release no GitHub com changelog em português

## Changelog

O changelog é gerado automaticamente no formato [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/) e traduzido para português:

- **Adicionado** - novas funcionalidades
- **Corrigido** - correções de bugs
- **Alterado** - mudanças em funcionalidades existentes
- **Removido** - funcionalidades removidas
- **Obsoleto** - funcionalidades deprecadas
- **Segurança** - correções de segurança

## Tags Retroativas

Para criar tags para versões históricas:

1. Acesse **Actions** → **Create Retroactive Tag**
2. Clique em **Run workflow**
3. Preencha:
   - **SHA**: hash do commit
   - **Tag**: nome da versão (ex: `v1.0.0`)
4. Execute o workflow

## Exemplos

### Commit de Feature

```bash
git commit -m "feat: adiciona seção de blog"
```

Resultado: bump de minor (ex: v3.1.0 → v3.2.0)

### Commit de Fix

```bash
git commit -m "fix: corrige contraste no tema escuro"
```

Resultado: bump de patch (ex: v3.1.0 → v3.1.1)

### Commit Breaking Change

```bash
git commit -m "feat: remove suporte a IE11" -m "BREAKING CHANGE: drop IE11"
```

Ou:

```bash
git commit -m "feat!: altera API de configuração"
```

Resultado: bump de major (ex: v3.1.0 → v4.0.0)

## Workflows

### release.yml

Executado automaticamente em push na main. Cria tags e releases.

### create-retro-tag.yml

Workflow manual para criar tags retroativas.

## Referências

- [Especificação](specs/001-versionamento-automatizado/spec.md)
- [Plano de Implementação](specs/001-versionamento-automatizado/plan.md)
