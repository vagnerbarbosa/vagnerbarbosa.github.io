---
description: "Tarefas de implementação para LinkedIn Import CLI"
---

# Tarefas: LinkedIn Import CLI

**Feature**: LinkedIn Import CLI  
**Branch**: `003-linkedin-import`  
**Input**: Ferramenta CLI para importar experiências, educação e certificações do LinkedIn a partir do export CSV manual, converter datas para português, dividir descrições em bullets, comparar com config.yaml atual e aplicar mudanças com confirmação interativa.

---

## Formato: `[ID] [P?] [Story] Descrição`

- **[P]**: Pode rodar em paralelo (arquivos diferentes, sem dependências)
- **[Story]**: A qual história de usuário esta tarefa pertence (ex: US1, US2, US3, US4)
- Inclui caminhos de arquivo exatos nas descrições

---

## Fase 1: Setup (Infraestrutura Compartilhada)

**Propósito**: Inicialização do projeto e estrutura básica do CLI

- [X] T001 Criar diretório cmd/import-linkedin/ na raiz do repositório
- [X] T002 Inicializar módulo Go em cmd/import-linkedin/ com `go mod init`
- [X] T003 Criar estrutura de diretórios: cmd/import-linkedin/internal/{parser,transformer,comparator,config,models,ui}/
- [X] T004 Criar diretório cmd/import-linkedin/commands/ para implementação dos comandos
- [X] T005 Criar diretório cmd/import-linkedin/testdata/ com arquivos CSV de exemplo
- [X] T006 [P] Adicionar dependências ao go.mod: go.yaml.in/yaml/v4, charmbracelet/huh v2, charmbracelet/lipgloss v2, sergi/go-diff

---

## Fase 2: Fundacional (Pré-requisitos Bloqueantes)

**Propósito**: Infraestrutura central que DEVE estar completa antes de QUALQUER história de usuário ser implementada

**⚠️ CRÍTICO**: Nenhum trabalho de história de usuário pode começar até esta fase estar completa

- [X] T007 [P] Criar struct Experience em cmd/import-linkedin/internal/models/experience.go com campos: company, role, start_date, end_date, description ([]string), location
- [X] T008 [P] Criar struct Education em cmd/import-linkedin/internal/models/education.go com campos: institution, degree, field, start_date, end_date, description ([]string)
- [X] T009 [P] Criar struct Certification em cmd/import-linkedin/internal/models/certification.go com campos: name, organization, issue_date, expiration_date, credential_id, credential_url
- [X] T010 [P] Criar struct Change em cmd/import-linkedin/internal/models/change.go com campos: entity_type, entity_id, old_value, new_value, change_type, fields_changed
- [X] T011 Criar struct ConfigPortfolio em cmd/import-linkedin/internal/models/config.go mapeando estrutura do config.yaml
- [X] T012 Implementar parser CSV em cmd/import-linkedin/internal/parser/csv.go usando encoding/csv com streaming (Read em loop)
- [X] T013 Implementar conversão de datas em cmd/import-linkedin/internal/parser/date.go (mapa en→pt: Jan→Jan, Feb→Fev, Mar→Mar, Apr→Abr, May→Mai, Jun→Jun, Jul→Jul, Aug→Ago, Sep→Set, Oct→Out, Nov→Nov, Dec→Dez; Present→Presente)
- [X] T014 Implementar split de descrições em bullets em cmd/import-linkedin/internal/transformer/description.go (heurística: \n\n ou ponto final + maiúscula)
- [X] T015 Implementar leitura/escrita YAML em cmd/import-linkedin/internal/config/yaml.go usando go.yaml.in/yaml/v4 com yaml.Node para preservar comentários
- [X] T016 Implementar backup em cmd/import-linkedin/internal/config/backup.go (cria backup com timestamp antes de modificar)

**Checkpoint**: Fundação pronta - implementação de histórias de usuário pode agora começar em paralelo

---

## Fase 3: História de Usuário 1 - Importar Experiências Profissionais (Prioridade: P1) 🎯 MVP

**Objetivo**: Permitir importação de experiências profissionais do CSV do LinkedIn para estrutura interna, com conversão de datas e divisão de descrições em bullets

**Teste Independente**: Executar a ferramenta com arquivo Experiences.csv de exemplo e verificar se todas as entradas são parseadas corretamente com company, role, datas em pt-BR, e descrição como lista de bullets

### Implementação da História de Usuário 1

- [X] T017 [US1] Implementar parse de Experiences.csv em cmd/import-linkedin/internal/parser/experience.go mapeando colunas: Company Name→company, Title→role, Started On→start_date, Finished On→end_date, Description→description, Location→location
- [X] T018 [US1] Implementar validação de experiências (company e role obrigatórios, datas em formato válido)
- [X] T019 [US1] Integrar conversão de datas no parser de experiências
- [X] T020 [US1] Integrar split de descrições em bullets no parser de experiências
- [X] T021 [US1] Criar testes unitários em cmd/import-linkedin/internal/parser/experience_test.go com cobertura >70%

**Checkpoint**: Neste ponto, a História de Usuário 1 deve estar totalmente funcional e testável independentemente

---

## Fase 4: História de Usuário 2 - Importar Educação (Prioridade: P2)

**Objetivo**: Permitir importação de formação acadêmica do CSV do LinkedIn

**Teste Independente**: Fornecer CSV de educação e verificar se institution, degree, field, datas e descrição são extraídos corretamente

### Implementação da História de Usuário 2

- [X] T022 [P] [US2] Implementar parse de Education.csv em cmd/import-linkedin/internal/parser/education.go mapeando: School Name→institution, Degree Name→degree, Field Of Study→field, Started On→start_date, Finished On→end_date, Description→description
- [X] T023 [US2] Implementar validação de educação (institution e degree obrigatórios, pelo menos uma data presente)
- [X] T024 [US2] Integrar conversão de datas no parser de educação
- [X] T025 [US2] Integrar split de descrições em bullets no parser de educação
- [X] T026 [US2] Criar testes unitários em cmd/import-linkedin/internal/parser/education_test.go

**Checkpoint**: Neste ponto, as Histórias de Usuário 1 E 2 devem ambas funcionar independentemente

---

## Fase 5: História de Usuário 3 - Importar Certificações (Prioridade: P3)

**Objetivo**: Permitir importação de certificações profissionais do CSV do LinkedIn

**Teste Independente**: Fornecer CSV de certificações e verificar se name, organization, issue_date, expiration_date são extraídos

### Implementação da História de Usuário 3

- [X] T027 [P] [US3] Implementar parse de Certifications.csv em cmd/import-linkedin/internal/parser/certification.go mapeando: Certification Name→name, Certification Authority→organization, Started On→issue_date, Finished On→expiration_date
- [X] T028 [US3] Implementar validação de certificações (name e organization obrigatórios, issue_date em formato válido)
- [X] T029 [US3] Integrar conversão de datas no parser de certificações
- [X] T030 [US3] Criar testes unitários em cmd/import-linkedin/internal/parser/certification_test.go

**Checkpoint**: Todas as histórias de importação (1-3) devem agora ser independentemente funcionais

---

## Fase 6: História de Usuário 4 - Comparar e Aplicar com Confirmação (Prioridade: P1)

**Objetivo**: Comparar dados importados com config.yaml atual, exibir diff visual, e permitir confirmação interativa das mudanças

**Teste Independente**: Ter config.yaml existente e CSV do LinkedIn, executar em modo dry-run e verificar se diff claro é exibido

### Implementação da História de Usuário 4

- [X] T031 [P] [US4] Implementar identificação única de entidades em cmd/import-linkedin/internal/comparator/id.go (Experience: company#role, Education: institution#degree#field, Certification: name#organization)
- [X] T032 [P] [US4] Implementar lógica de comparação em cmd/import-linkedin/internal/comparator/diff.go detectando: novas (added), modificadas (modified), removidas (removed)
- [X] T033 [US4] Implementar visualização de diff em cmd/import-linkedin/internal/ui/diff.go usando lipgloss (verde +, vermelho -, amarelo ~)
- [X] T034 [US4] Implementar prompts interativos em cmd/import-linkedin/internal/ui/confirm.go usando charmbracelet/huh (yes/no por mudança, aceitar tudo, rejeitar tudo, selecionar específicas)
- [X] T035 [US4] Implementar aplicação de mudanças em cmd/import-linkedin/internal/comparator/apply.go (atualiza config.yaml apenas com mudanças confirmadas)
- [X] T036 [US4] Criar testes unitários para diff e confirmação em cmd/import-linkedin/internal/comparator/diff_test.go e cmd/import-linkedin/internal/ui/confirm_test.go

**Checkpoint**: Neste ponto, todo o fluxo de importação com confirmação deve estar funcional

---

## Fase 7: Comandos da CLI

**Propósito**: Implementar comandos principais da CLI conforme contratos definidos

- [X] T037 [P] Implementar comando root em cmd/import-linkedin/commands/root.go com setup básico
- [X] T038 [P] Implementar comando import em cmd/import-linkedin/commands/import.go com flags: --experiences, --education, --certifications, --config, --dry-run, --yes, --backup
- [X] T039 [P] Implementar comando validate em cmd/import-linkedin/commands/validate.go validando CSVs sem importar
- [X] T040 [P] Implementar comando version em cmd/import-linkedin/commands/version.go
- [X] T041 [P] Implementar tratamento de códigos de saída (0=success, 1=erro genérico, 2=CSV parse error, 3=YAML error, 4=arquivo não encontrado, 5=permissão negada, 130=interrompido)
- [X] T042 Criar main.go em cmd/import-linkedin/main.go como entry point
- [X] T043 Criar testes de integração em cmd/import-linkedin/commands/import_test.go

**Checkpoint**: CLI completa e funcional com todos os comandos documentados

---

## Fase Final: Polimento & Preocupações Transversais

**Propósito**: Melhorias que afetam múltiplas histórias de usuário

- [X] T044 [P] Criar arquivos de teste de exemplo em cmd/import-linkedin/testdata/Experiences.csv, Education.csv, Certifications.csv
- [X] T045 Implementar tratamento de casos de borda: CSV vazio, datas mal formatadas, config.yaml inexistente, entradas duplicadas, descrição vazia, encoding incorreto, interrupção durante confirmação
- [X] T046 [P] Otimizar performance: streaming de CSV, caching de config.yaml, lazy loading de parsers
- [X] T047 Garantir cobertura de testes >70% conforme constituição do projeto
- [X] T048 [P] Adicionar logging apropriado para operações (sem expor dados sensíveis)
- [X] T049 Validar quickstart.md executando todos os comandos documentados

---

## Dependências & Ordem de Execução

### Dependências de Fase

| Fase | Dependências | Pode Começar Quando |
|------|--------------|---------------------|
| **Fase 1: Setup** | Nenhuma | Imediatamente |
| **Fase 2: Fundacional** | Fase 1 completa | Após T001-T006 |
| **Fase 3: US1 (P1)** | Fase 2 completa | Após T007-T016 |
| **Fase 4: US2 (P2)** | Fase 2 completa | Após T007-T016 (independente de US1) |
| **Fase 5: US3 (P3)** | Fase 2 completa | Após T007-T016 (independente de US1/US2) |
| **Fase 6: US4 (P1)** | Fases 3, 4, 5 completas | Requer parsers de todas entidades |
| **Fase 7: Comandos** | Fases 2-6 completas | Após lógica interna pronta |
| **Fase Final** | Todas as anteriores | Após CLI funcional |

### Dependências entre Tarefas

**Dentro de cada História:**
- Modelos antes de serviços (ex: T007 antes de T017)
- Parsers antes de testes (ex: T017 antes de T021)
- Lógica de diff antes de UI de confirmação (T032 antes de T034)

**Entre Histórias:**
- US1, US2, US3 são independentes entre si (podem rodar em paralelo após Fase 2)
- US4 depende de todas as entidades (US1+US2+US3) para comparação completa

### Oportunidades de Paralelismo

**Fase 2 (Fundacional) - Após Setup:**
```
T007 (Experience model)     ──┐
T008 (Education model)        ──┼──→ T011 (Config model) → T012-T016
T009 (Certification model)    ──┘
```

**Fases 3, 4, 5 (Histórias de Usuário):**
```
T017-T021 (US1) ──┐
T022-T026 (US2) ──┼──→ Independentes, podem rodar em paralelo
T027-T030 (US3) ──┘
```

**Fase 7 (Comandos):**
```
T037 (root) ──┐
T038 (import) ──┼──→ T042 (main.go)
T039 (validate)─┤
T040 (version)──┘
```

---

## Estratégia de Implementação

### MVP Primeiro (Apenas História de Usuário 1 + Setup + Fundacional)

1. Completar **Fase 1**: Setup (T001-T006)
2. Completar **Fase 2**: Fundacional (T007-T016)
3. Completar **Fase 3**: História de Usuário 1 (T017-T021)
   - Neste ponto: Parsing de experiências está pronto
4. Completar **Fase 6** parcial: Apenas lógica de diff/comparação para experiências
5. Completar **Fase 7** parcial: Comando import apenas para experiências
6. **PARE e VALIDE**: Teste a importação de experiências independentemente
7. Demo funcional: Usuário pode importar experiências do LinkedIn

**Após MVP Validado:**
8. Adicionar **Fase 4**: US2 (Educação) → Testar independentemente
9. Adicionar **Fase 5**: US3 (Certificações) → Testar independentemente
10. Expandir **Fase 6**: Diff completo com todas entidades
11. Expandir **Fase 7**: Comandos validate, version
12. **Fase Final**: Polimento e cobertura de testes

### Entrega Incremental

| Entrega | Tarefas | Funcionalidade | Valor para Usuário |
|---------|---------|----------------|-------------------|
| **MVP v0.1** | F1+F2+F3+parte F6+F7 | Importar experiências apenas | Economia de tempo maior (mais dados) |
| **v0.2** | +F4 | Adicionar educação | Currículo mais completo |
| **v0.3** | +F5 | Adicionar certificações | Perfil profissional completo |
| **v1.0** | +Fase Final | Dry-run, backup, validação | Segurança e controle total |

### Sequência Recomendada para Desenvolvedor Único

```
Semana 1: Fase 1 (Setup) + Fase 2 (Fundacional)
         ├─ T001-T006 (estrutura)
         └─ T007-T016 (models + parsers base)

Semana 2: Fase 3 (US1 - Experiências) + Início Fase 6
         ├─ T017-T021 (parser experiências)
         └─ T031-T033 (diff básico)

Semana 3: Fase 4 (US2 - Educação) + Fase 5 (US3 - Certificações)
         ├─ T022-T026 (parser educação)
         └─ T027-T030 (parser certificações)

Semana 4: Fase 6 (US4 - Confirmação) + Fase 7 (Comandos)
         ├─ T034-T036 (UI + aplicação)
         └─ T037-T043 (comandos CLI)

Semana 5: Fase Final (Polimento)
         └─ T044-T049 (testes, edge cases, docs)
```

---

## Resumo

| Métrica | Valor |
|---------|-------|
| **Total de Tarefas** | 49 |
| **Tarefas Setup (F1)** | 6 |
| **Tarefas Fundacionais (F2)** | 10 |
| **Tarefas US1 (F3)** | 5 |
| **Tarefas US2 (F4)** | 5 |
| **Tarefas US3 (F5)** | 4 |
| **Tarefas US4 (F6)** | 6 |
| **Tarefas Comandos (F7)** | 7 |
| **Tarefas Polimento (FF)** | 6 |
| **Histórias de Usuário** | 4 |
| **Ponto de Entrada** | cmd/import-linkedin/main.go |
| **Cobertura de Testes Meta** | >70% |

### Arquivos Principais Criados

```
cmd/import-linkedin/
├── main.go                              # Entry point (T042)
├── commands/
│   ├── root.go                          # Comando root (T037)
│   ├── import.go                        # Comando import (T038)
│   ├── validate.go                      # Comando validate (T039)
│   └── version.go                       # Comando version (T040)
├── internal/
│   ├── models/
│   │   ├── experience.go                  # T007
│   │   ├── education.go                 # T008
│   │   ├── certification.go             # T009
│   │   ├── change.go                    # T010
│   │   └── config.go                    # T011
│   ├── parser/
│   │   ├── csv.go                       # T012
│   │   ├── date.go                      # T013
│   │   ├── experience.go                # T017
│   │   ├── education.go                 # T022
│   │   └── certification.go             # T027
│   ├── transformer/
│   │   └── description.go               # T014
│   ├── comparator/
│   │   ├── id.go                        # T031
│   │   ├── diff.go                      # T032
│   │   └── apply.go                     # T035
│   ├── config/
│   │   ├── yaml.go                      # T015
│   │   └── backup.go                    # T016
│   └── ui/
│       ├── diff.go                      # T033
│       └── confirm.go                   # T034
└── testdata/
    ├── Experiences.csv                  # T044
    ├── Education.csv                    # T044
    └── Certifications.csv               # T044
```

### Próximos Passos Sugeridos

1. **Criar branch de feature**: `git checkout -b feat/003-linkedin-import`
2. **Iniciar Fase 1**: Executar T001-T006 para estrutura base
3. **Verificar cobertura**: Após cada fase, rodar `go test -cover ./...`
4. **Commit frequente**: Um commit por tarefa ou grupo lógico
5. **Validar com dados reais**: Testar com CSVs reais do LinkedIn após Fase 7
