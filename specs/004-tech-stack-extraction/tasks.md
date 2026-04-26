# Tarefas: Extração Inteligente de Tech Stack

**Input**: Documentos de design de `/specs/004-tech-stack-extraction/`
**Pré-requisitos**: plan.md, spec.md, data-model.md, research.md, quickstart.md

**Testes**: Incluídos - Testes unitários são essenciais para validar a lógica de extração

**Organização**: Tarefas agrupadas por história de usuário para permitir implementação e teste independentes

## Formato: `[ID] [P?] [Story] Descrição`

- **[P]**: Pode rodar em paralelo (arquivos diferentes, sem dependências)
- **[Story]**: A qual história de usuário esta tarefa pertence (US1, US2, US3)
- Inclua caminhos de arquivo exatos nas descrições

---

## Fase 1: Setup (Infraestrutura Compartilhada)

**Propósito**: Criar estrutura base para a funcionalidade

- [x] T001 [P] Criar diretório e arquivo para transformer de tech stack em `cmd/import-linkedin/internal/transformer/techstack.go`

---

## Fase 2: Fundacional (Pré-requisitos Bloqueantes)

**Propósito**: Modelo de dados que DEVE estar completo antes das histórias de usuário

**⚠️ CRÍTICO**: Nenhum trabalho de história de usuário pode começar até esta fase estar completa

- [x] T002 Adicionar campo `TechStack` ao struct `Experience` em `cmd/import-linkedin/internal/models/experience.go`
- [x] T003 Atualizar mapeamento YAML para incluir `tech_stack` em `cmd/import-linkedin/internal/config/yaml.go`

**Checkpoint**: Fundação pronta - implementação de histórias de usuário pode agora começar

---

## Fase 3: História de Usuário 1 - Extração Automática de Tech Stack (Prioridade: P1) 🎯 MVP

**Objetivo**: Implementar a extração automática de tech stack de bullets de descrição

**Teste Independente**: Importar CSV com descrições contendo padrões de tech stack e verificar que o campo `TechStack` é preenchido corretamente e o bullet é removido da descrição

### Testes para História de Usuário 1 ⚠️

> **NOTA: Escreva estes testes PRIMEIRO, garanta que FALHEM antes da implementação**

- [ ] T004 [P] [US1] Criar testes unitários para extração de tech stack em `cmd/import-linkedin/internal/transformer/techstack_test.go` - testar padrão "Tecnologias:"
- [ ] T005 [P] [US1] Criar teste para extração com vírgula como separador
- [ ] T006 [P] [US1] Criar teste para remoção do bullet de tech stack da descrição
- [ ] T007 [P] [US1] Criar teste para múltiplos separadores (pipe, hífen, bullet)
- [ ] T008 [P] [US1] Criar teste para tech stack no meio da descrição (preservar outros bullets)

### Implementação da História de Usuário 1

- [ ] T009 [US1] Implementar struct `TechStackResult` em `cmd/import-linkedin/internal/transformer/techstack.go`
- [ ] T010 [US1] Implementar função `ExtractTechStack(bullets []string) TechStackResult` em `cmd/import-linkedin/internal/transformer/techstack.go` - detectar padrões PT
- [ ] T011 [US1] Implementar regex para padrões de tech stack (case-insensitive) em `cmd/import-linkedin/internal/transformer/techstack.go`
- [ ] T012 [US1] Implementar parsing de tecnologias com separadores variados em `cmd/import-linkedin/internal/transformer/techstack.go`
- [ ] T013 [US1] Implementar formatação de saída com separador " • " em `cmd/import-linkedin/internal/transformer/techstack.go`
- [ ] T014 [US1] Integrar `ExtractTechStack` no parser de experiências em `cmd/import-linkedin/internal/parser/experience.go` - chamar após `SplitDescription`

**Checkpoint**: Neste ponto, a História de Usuário 1 deve estar totalmente funcional e testável independentemente

---

## Fase 4: História de Usuário 2 - Suporte a Múltiplos Padrões e Formatos (Prioridade: P2)

**Objetivo**: Adicionar suporte a padrões em inglês e mais formatos de separadores

**Teste Independente**: Importar CSV com variações de padrões ("Tech Stack:", "Technologies:", etc.) e diferentes separadores

### Testes para História de Usuário 2 ⚠️

- [ ] T015 [P] [US2] Criar teste para padrão "Tech Stack:" em `cmd/import-linkedin/internal/transformer/techstack_test.go`
- [ ] T016 [P] [US2] Criar teste para padrão "Technologies:" em `cmd/import-linkedin/internal/transformer/techstack_test.go`
- [ ] T017 [P] [US2] Criar teste para padrão "Tools:" em `cmd/import-linkedin/internal/transformer/techstack_test.go`
- [ ] T018 [P] [US2] Criar teste para separador pipe `|` em `cmd/import-linkedin/internal/transformer/techstack_test.go`
- [ ] T019 [P] [US2] Criar teste para padrão "As principais tecnologias e ferramentas utilizadas:" em `cmd/import-linkedin/internal/transformer/techstack_test.go`

### Implementação da História de Usuário 2

- [ ] T020 [US2] Adicionar suporte a padrão "Tech Stack:" em `cmd/import-linkedin/internal/transformer/techstack.go`
- [ ] T021 [US2] Adicionar suporte a padrão "Technologies:" em `cmd/import-linkedin/internal/transformer/techstack.go`
- [ ] T022 [US2] Adicionar suporte a padrão "Tools:" em `cmd/import-linkedin/internal/transformer/techstack.go`
- [ ] T023 [US2] Adicionar suporte a padrão "As principais tecnologias e ferramentas utilizadas:" em `cmd/import-linkedin/internal/transformer/techstack.go`
- [ ] T024 [US2] Adicionar suporte a separador pipe `|` em `cmd/import-linkedin/internal/transformer/techstack.go`
- [ ] T025 [US2] Adicionar suporte a padrão "Ferramentas:" em `cmd/import-linkedin/internal/transformer/techstack.go`

**Checkpoint**: Neste ponto, as Histórias de Usuário 1 E 2 devem ambas funcionar independentemente

---

## Fase 5: História de Usuário 3 - Preservação de Descrições sem Tech Stack (Prioridade: P3)

**Objetivo**: Garantir que descrições sem tech stack sejam processadas normalmente

**Teste Independente**: Importar experiências sem padrões de tech stack e verificar que a descrição permanece intacta

### Testes para História de Usuário 3 ⚠️

- [ ] T026 [P] [US3] Criar teste para descrição sem tech stack - preservar todos os bullets em `cmd/import-linkedin/internal/transformer/techstack_test.go`
- [ ] T027 [P] [US3] Criar teste para tech stack vazio após prefixo em `cmd/import-linkedin/internal/transformer/techstack_test.go`

### Implementação da História de Usuário 3

- [ ] T028 [US3] Garantir que bullets sem padrões de tech stack são preservados inalterados em `cmd/import-linkedin/internal/transformer/techstack.go`
- [ ] T029 [US3] Tratar caso de tech stack vazio após prefixo em `cmd/import-linkedin/internal/transformer/techstack.go`

**Checkpoint**: Todas as histórias de usuário devem agora ser independentemente funcionais

---

## Fase 6: Polimento & Preocupações Transversais

**Propósito**: Melhorias que afetam múltiplas histórias de usuário

- [ ] T030 [P] Atualizar documentação do LinkedIn Import CLI em `docs/LINKEDIN-IMPORT.md`
- [ ] T031 Atualizar README.md seção de comandos do import-linkedin
- [ ] T032 Executar todos os testes: `go test ./cmd/import-linkedin/...`
- [ ] T033 Verificar cobertura de testes: `go test -cover ./cmd/import-linkedin/...`

---

## Dependências & Ordem de Execução

### Dependências de Fase

- **Setup (Fase 1)**: Sem dependências - pode começar imediatamente
- **Fundacional (Fase 2)**: Depende da conclusão do Setup - BLOQUEIA todas as histórias de usuário
- **Histórias de Usuário (Fase 3+)**: Todas dependem da conclusão da fase Fundacional
  - US1 deve ser concluída antes de US2 (dependência de código)
  - US2 pode ser feita em paralelo com US3
- **Polimento (Fase Final)**: Depende que todas as histórias de usuário estejam completas

### Dependências entre Tarefas

```
T001 → T002 → T003 → T004-T008 (testes US1) → T009-T014 (impl US1)
                                      ↓
                                T015-T019 (testes US2) → T020-T025 (impl US2)
                                      ↓
                                T026-T027 (testes US3) → T028-T029 (impl US3)
                                      ↓
                                T030-T033 (polimento)
```

### Oportunidades de Paralelismo

- Todas as tarefas de teste marcadas [P] dentro de uma história podem rodar em paralelo
- T001 pode rodar a qualquer momento
- Após T003, US1, US2 e US3 podem começar (mas US2 depende de US1 em termos de código)

---

## Exemplo de Execução Sequencial (MVP)

```bash
# Fase 1-2: Setup e Fundacional
T001, T002, T003

# Fase 3: US1 (MVP - apenas isto entrega valor)
T004, T005, T006, T007, T008  # Testes
T009, T010, T011, T012, T013, T014  # Implementação

# Neste ponto, pode-se PARAR e usar a funcionalidade
# Se quiser mais padrões, continue com US2...

# Fase 4: US2 (opcional - mais padrões)
T015-T019, T020-T025

# Fase 5: US3 (opcional - casos de borda)
T026-T027, T028-T029

# Fase 6: Polimento
T030-T033
```

---

## Estratégia de Implementação

### MVP Primeiro (Apenas História de Usuário 1)

1. Completar Fase 1: Setup (T001)
2. Completar Fase 2: Fundacional (T002-T003)
3. Completar Fase 3: História de Usuário 1 (T004-T014)
4. **PARE e VALIDE**: Teste a História de Usuário 1 independentemente
5. Deploy/demo se estiver pronto

**Valor entregue no MVP**: Extração automática de tech stack com os padrões mais comuns ("Tecnologias:", vírgula como separador)

### Entrega Incremental

1. Completar Setup + Fundacional → Fundação pronta
2. Adicionar US1 → Testar → Funcionalidade core pronta
3. Adicionar US2 → Testar → Mais padrões suportados
4. Adicionar US3 → Testar → Casos de borda cobertos
5. Polimento → Documentação atualizada

---

## Notas

- Tarefas [P] = arquivos diferentes, sem dependências
- Label [Story] mapeia tarefa para história de usuário específica
- Cada história de usuário deve ser completável e testável independentemente
- Verifique que testes falham antes de implementar (TDD)
- Commit após cada tarefa ou grupo lógico
- Pare em qualquer checkpoint para validar história independentemente
