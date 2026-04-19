# Pesquisa: Versionamento Automatizado

**Data**: 2026-04-19
**Funcionalidade**: Versionamento Automatizado com Tags e Releases

## Decisões de Pesquisa

### 1. Ferramenta para Análise de Conventional Commits

**Decisão**: Usar `cycjimmy/semantic-release-action`

**Racional**:
- Maturidade: Projeto ativo desde 2020, usado por milhares de repositórios
- Configurabilidade: Suporta plugins para customização
- Integração nativa com GitHub Actions

**Alternativas consideradas**:
| Ferramenta | Prós | Contras |
|------------|------|---------|
| release-please | Mantido pelo Google, simples | Menos flexível para i18n |
| standard-version | CLI simples | Deprecated em favor de release-please |
| Script bash custom | Controle total | Manutenção, edge cases |

### 2. Geração de Changelog em Português

**Decisão**: Script bash customizado como wrapper

**Racional**: Nenhuma ferramenta popular suporta i18n nativamente. Melhor abordagem é traduzir após geração.

**Mapeamento de tipos**:
```
feat → Adicionado
fix → Corrigido
docs → Documentação
style → Estilo
refactor → Alterado
perf → Performance
test → Testes
chore → Manutenção
BREAKING CHANGE → Quebra de compatibilidade
```

### 3. Action para Criar Releases

**Decisão**: `softprops/action-gh-release@v1`

**Racional**:
- Mais atualizada que `actions/create-release` (deprecated)
- Suporta uploads de assets (extensível)
- Simples de configurar

### 4. Detecção de Breaking Changes

**Decisão**: Suportar ambos os padrões Conventional Commits:
- `!` no tipo: `feat!:`, `fix!:`
- Footer `BREAKING CHANGE:`

**Racional**: Cobertura completa do padrão, alinhado com expectativas da comunidade.

### 5. Estratégia de Concorrência

**Decisão**: Usar `concurrency` do GitHub Actions para evitar corridas

```yaml
concurrency:
  group: release
  cancel-in-progress: false
```

**Racional**: Se dois merges ocorrerem simultaneamente, o segundo aguarda o primeiro.

## Notas de Implementação

- Token `GITHUB_TOKEN` tem permissões automáticas para o repo atual
- Para repos forked, pode ser necessário habilitar "Read and write permissions" em Settings → Actions → General
- Semantic-release requer branch protegida com conventional commits enforcement (opcional mas recomendado)
