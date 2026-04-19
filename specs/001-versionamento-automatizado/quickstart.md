# Quickstart: Versionamento Automatizado

Guia rápido para usar o sistema de versionamento automatizado.

## Para Desenvolvedores

### 1. Escrever Commits Corretamente

Use [Conventional Commits](https://www.conventionalcommits.org/pt-br/v1.0.0/):

```bash
# Nova funcionalidade → bump minor (3.1.0 → 3.2.0)
git commit -m "feat: adiciona seção de blog"

# Correção de bug → bump patch (3.1.0 → 3.1.1)
git commit -m "fix: corrige contraste no tema escuro"

# Documentação → bump patch
git commit -m "docs: atualiza instruções de build"

# Breaking change → bump major (3.1.0 → 4.0.0)
git commit -m "feat!: remove suporte a IE11" -m "BREAKING CHANGE: drop IE11"
# OU
git commit -m "feat!: altera API de configuração"
```

### 2. Merge na Main

Ao fazer merge do seu PR:

```bash
gh pr merge --squash
```

O workflow `release.yml` executará automaticamente:
1. Análise dos commits
2. Cálculo da nova versão
3. Criação da tag
4. Deploy
5. Criação da release com changelog

### 3. Verificar Resultado

```bash
# Ver última tag
git describe --tags --abbrev=0

# Ver releases no GitHub
gh release list --limit 5
```

## Para Criar Tags Retroativas

### Workflow Manual

1. Acesse a aba **Actions** no GitHub
2. Selecione o workflow **"Create Retroactive Tag"**
3. Clique em **"Run workflow"**
4. Preencha os campos:
   - **SHA**: hash do commit (ex: `b81034a`)
   - **Tag**: nome da versão (ex: `v1.0.0`)
5. Clique **"Run workflow"**

### Via CLI

```bash
# Criar tag localmente
git tag -a v1.0.0 b81034a -m "Versão 1.0.0 - Jekyll inicial"

# Push para GitHub
git push origin v1.0.0

# Criar release (opcional)
gh release create v1.0.0 --title "v1.0.0" --notes "Primeira versão com Jekyll"
```

## Troubleshooting

### Tag não foi criada

Verifique se:
- Commits seguem o formato Conventional Commits
- Workflow tem permissões de escrita (Settings → Actions → General)
- Não houve erro no workflow (Actions → release.yml)

### Deploy falhou mas tag existe

Comportamento esperado! A tag permanece como registro histórico. Você pode:
1. Corrigir o problema
2. Fazer novo push para main
3. Nova tag será criada

### Changelog em inglês

O script de tradução deve estar em `.github/scripts/generate-changelog.sh`. Se não estiver presente, verifique se o setup foi completo.

## Referência de Tipos de Commit

| Tipo | Versão | Tradução no Changelog |
|------|--------|----------------------|
| feat | minor | Adicionado |
| fix | patch | Corrigido |
| docs | patch | Documentação |
| style | patch | Estilo |
| refactor | patch | Alterado |
| perf | patch | Performance |
| test | patch | Testes |
| chore | patch | Manutenção |
| BREAKING | major | Quebra de compatibilidade |
