# Tarefas: Versionamento Automatizado com Tags e Releases

**Input**: Documentos de design de `/specs/001-versionamento-automatizado/`
**Pré-requisitos**: plan.md, spec.md, research.md, data-model.md, contracts/

**Testes**: Não incluídos (testes serão manuais em fork/branch de teste)

**Organização**: Tarefas agrupadas por história de usuário para implementação e teste independentes

## Formato: `[ID] [P?] [Story] Descrição`

- **[P]**: Pode rodar em paralelo
- **[Story]**: História de usuário (HU1, HU2, HU3)

---

## Fase 1: Setup (Infraestrutura Compartilhada)

**Propósito**: Preparar estrutura de workflows e scripts

- [x] T001 [P] Criar diretório `.github/workflows/` se não existir
- [x] T002 [P] Criar diretório `.github/scripts/` para scripts auxiliares
- [x] T003 Verificar permissões do `GITHUB_TOKEN` em Settings → Actions → General (permissões de read/write)

**Checkpoint**: Estrutura de diretórios pronta

---

## Fase 2: Foundacional (Pré-requisitos Bloqueantes)

**Propósito**: Configurações que devem estar prontas antes das histórias de usuário

**⚠️ CRÍTICO**: Estas tarefas devem estar completas antes de implementar qualquer workflow

- [x] T004 [P] Instalar `semantic-release` e plugins necessários no projeto (package.json ou via npx)
- [x] T005 Criar arquivo de configuração `.releaserc.json` com configuração para Conventional Commits
- [x] T006 Criar script `.github/scripts/generate-changelog.sh` para traduzir changelog para português (formato Keep a Changelog)

**Checkpoint**: Ferramentas de versionamento configuradas

---

## Fase 3: User Story 1 - Deploy Automático com Tag (Prioridade: P1) 🎯 MVP

**Objetivo**: Criar workflow que detecta push na main e gera tags semver automaticamente

**Teste Independente**: Fazer merge de PR na main e verificar tag criada no GitHub

### Implementação da HU1

- [x] T007 [P] [HU1] Criar workflow `.github/workflows/release.yml` com trigger em push na main
- [x] T008 [P] [HU1] Adicionar job de análise de commits usando semantic-release no workflow
- [x] T009 [HU1] Configurar job para criar tag semver (major/minor/patch) baseada nos commits
- [x] T010 [HU1] Adicionar configuração de concurrency para evitar corridas entre workflows
- [ ] T011 [HU1] Testar workflow em branch de teste: commit tipo `feat:` deve gerar bump minor
- [ ] T012 [HU1] Testar workflow em branch de teste: commit tipo `fix:` deve gerar bump patch
- [ ] T013 [HU1] Testar: commit com `BREAKING CHANGE:` ou `!` deve gerar bump major

**Checkpoint**: Tags sendo criadas automaticamente em pushes na main

---

## Fase 4: User Story 2 - Release Automático no GitHub (Prioridade: P1)

**Objetivo**: Criar release automaticamente após deploy bem-sucedido com changelog em português

**Teste Independente**: Verificar release criada no GitHub com changelog traduzido

### Implementação da HU2

- [x] T014 [P] [HU2] Adicionar job no `.github/workflows/release.yml` para aguardar deploy bem-sucedido
- [x] T015 [HU2] Integrar script `generate-changelog.sh` para gerar changelog em português
- [x] T016 [HU2] Adicionar step usando `softprops/action-gh-release` para criar release
- [x] T017 [HU2] Configurar release para usar changelog gerado como body
- [x] T018 [HU2] Adicionar verificação: se deploy falha, não criar release (apenas log)
- [ ] T019 [HU2] Testar em fork: release criada com changelog formatado (Adicionado/Corrigido/Alterado)
- [ ] T020 [HU2] Verificar: se deploy falha, tag existe mas release não é criada

**Checkpoint**: Releases sendo criadas automaticamente após deploys bem-sucedidos

---

## Fase 5: User Story 3 - Tags Retroativas (Prioridade: P2)

**Objetivo**: Workflow manual para criar tags retroativas v1.0.0 e v2.0.0

**Teste Independente**: Executar workflow manualmente e verificar tags criadas no GitHub

### Implementação da HU3

- [x] T021 [P] [HU3] Criar workflow `.github/workflows/create-retro-tag.yml` com trigger `workflow_dispatch`
- [x] T022 [HU3] Adicionar inputs ao workflow: `sha` (commit hash) e `tag` (nome da versão)
- [x] T023 [HU3] Implementar validação: SHA deve existir no histórico
- [x] T024 [HU3] Implementar validação: tag deve seguir semver e não existir
- [x] T025 [HU3] Adicionar step para criar tag apontando para o SHA especificado
- [x] T026 [HU3] Adicionar step opcional para criar release da tag retroativa
- [ ] T027 [HU3] Executar manualmente: criar tag v1.0.0 apontando para commit `b81034a` (2015)
- [ ] T028 [HU3] Executar manualmente: criar tag v2.0.0 apontando para commit pré-migração Go
- [ ] T029 [HU3] Verificar no GitHub: tags v1.0.0 e v2.0.0 existem e apontam para commits corretos

**Checkpoint**: Tags retroativas v1.0.0 e v2.0.0 criadas no GitHub

---

## Fase 6: Polish & Cross-Cutting Concerns

**Propósito**: Melhorias e documentação

- [ ] T030 [P] Atualizar `.github/workflows/deploy.yml` se necessário para integração
- [ ] T031 [P] Adicionar badge de versão/latest release no README.md
- [ ] T032 Documentar processo de versionamento em `docs/VERSIONING.md` (opcional)
- [ ] T033 Validar quickstart.md com passos reais
- [ ] T034 [P] Testar edge case: dois merges simultâneos (concurrency funcionando)

**Checkpoint**: Sistema completo documentado e testado

---

## Dependências & Ordem de Execução

### Dependências de Fase

- **Setup (Fase 1)**: Sem dependências - pode começar imediatamente
- **Foundational (Fase 2)**: Depende da conclusão do Setup - BLOQUEIA todas as histórias
- **HU1 (Fase 3)**: Depende da Foundacional completa
- **HU2 (Fase 4)**: Depende da HU1 completa (release precisa da tag criada)
- **HU3 (Fase 5)**: Pode começar após Foundacional, mas recomendado após HU2 para reaproveitar config
- **Polimento (Fase Final)**: Depende de todas as histórias desejadas

### Oportunidades de Paralelismo

- T001-T003 (Setup) podem rodar em paralelo
- T004-T006 (Foundacional) podem rodar em paralelo
- T007-T008 (HU1) podem rodar em paralelo (configuração do workflow)
- T014 (HU2) pode ser preparado enquanto HU1 é testada
- T021-T026 (HU3) podem ser desenvolvidos em paralelo

---

## Implementação Recomendada

### MVP (User Story 1 apenas)

1. Complete Fase 1: Setup
2. Complete Fase 2: Foundacional
3. Complete Fase 3: HU1 (tags automáticas funcionando)
4. **STOP e VALIDE**: Testar em fork/branch

### Incremental Completa

1. Setup + Foundacional → Infraestrutura pronta
2. HU1 → Tags automáticas → Testar
3. HU2 → Releases automáticas → Testar
4. HU3 → Tags retroativas → Executar manualmente
5. Polimento → Documentação final

---

## Notas

- Todas as tarefas usam paths absolutos relativos à raiz do repo
- Testes são manuais: validar em fork do repositório antes de merge na main
- O token `GITHUB_TOKEN` é auto-provisionado, mas verificar permissões
- Commits de teste devem seguir Conventional Commits para trigger correto
