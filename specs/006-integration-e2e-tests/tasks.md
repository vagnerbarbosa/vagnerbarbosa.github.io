# Tarefas: Implementação de Suíte de Testes de Integração e E2E

**Feature**: `006-integration-e2e-tests`
**Branch**: `006-integration-e2e-tests`
**Spec**: [spec.md](spec.md) | **Plan**: [plan.md](plan.md)

## Resumo da Implementação

Esta funcionalidade implementa a camada de testes de alta ordem para garantir que o pipeline de dados do LinkedIn e a geração do site estático funcionem corretamente de ponta a ponta.

- **Total de Tarefas**: 16
- **Distribuição por História**:
    - US1 (Integração/Golden Files): 5 tarefas
    - US2 (E2E Flow): 6 tarefas
    - US3 (Assets/SEO Regression): 3 tarefas
- **Fase de Polimento**: 2 tarefas
- **MVP**: Conclusão da US1 e US2.

---

## Fase 1: Setup (Infraestrutura)

- [x] T001 Criar diretório de Golden Files em `cmd/import-linkedin/testdata/golden/`
- [x] T002 [P] Configurar diretório de inputs de teste em `cmd/import-linkedin/testdata/input_csv/`

## Fase 2: Fundamentais (Blocking Prerequisites)

- [x] T003 Implementar helper de comparação semântica de YAML em `cmd/import-linkedin/internal/integration/compare.go`
- [x] T004 Implementar lógica de atualização de Golden Files (flag `-update`) em `cmd/import-linkedin/internal/integration/compare.go`

## Fase 3: US1 - Validação do Pipeline via Golden Files (Prioridade: P1)

**Objetivo**: Validar que o pipeline de importação produz o YAML esperado bit-a-bit.
**Critério de Teste Independente**: Executar teste de integração e validar que o output coincide com o arquivo Golden.

- [x] T005 [US1] Implementar teste de integração de pipeline em `cmd/import-linkedin/internal/integration/pipeline_test.go`
- [x] T006 [US1] Criar Golden File de referência para cenário "Happy Path" em `cmd/import-linkedin/testdata/golden/happy_path.yaml`
- [x] T007 [US1] Criar Golden File de referência para cenário "Dados Parciais" em `cmd/import-linkedin/testdata/golden/partial_data.yaml`
- [x] T008 [US1] Criar Golden File de referência para cenário "Erro de Parsing" em `cmd/import-linkedin/testdata/golden/parsing_error.yaml`
- [x] T009 [US1] Validar que a importação de CSVs malformados gera o output esperado do Golden File

## Fase 4: US2 - Validação do Fluxo End-to-End (Prioridade: P1)

**Objetivo**: Garantir que a sequência `Import -> Generate -> Public` resulte em um HTML válido.
**Critério de Teste Independente**: Execução do runner E2E resultando em `public/index.html` com as entidades importadas presentes.

- [x] T010 [US2] Implementar `E2ETestRunner` com isolamento em `t.TempDir()` em `cmd/generator/e2e_test.go`
- [x] T011 [US2] Implementar lógica de orquestração (Import $\rightarrow$ Generator) em `cmd/generator/e2e_test.go`
- [x] T012 [US2] Implementar validação de presença de Nome do Usuário no HTML em `cmd/generator/e2e_test.go`
- [x] T013 [US2] Implementar validação de presença de 1 Experiência e 1 Educação no HTML em `cmd/generator/e2e.test.go`
- [x] T014 [US2] Implementar limpeza automática da pasta `public/` antes de cada teste E2E em `cmd/generator/e2e_test.go`
- [x] T015 [US2] Validar que o fluxo E2E completa em menos de 15 segundos

## Fase 5: US3 - Prevenção de Regressão em Assets e SEO (Prioridade: P2)

**Objetivo**: Impedir a perda de arquivos críticos de SEO/PWA após builds.
**Critério de Teste Independente**: Verificação da existência e conteúdo básico dos arquivos de SEO no output do gerador.

- [x] T016 [US3] Implementar teste de existência e validade de `sitemap.xml` em `cmd/generator/e2e_test.go`
- [x] T017 [US3] Implementar teste de existência e validade de `robots.txt` em `cmd/generator/e2e_test.go`
- [x] T018 [US3] Implementar validação de cores de tema e ícones no `site.webmanifest` em `cmd/generator/e2e_test.go`

## Fase Final: Polimento & Cross-Cutting

- [x] T019 Configurar GitHub Actions para executar `go test -tags=integration ./...` em cada PR em `.github/workflows/deploy.yml`
- [x] T020 Criar documentação de testes em `docs/TESTING.md` e atualizar a seção de desenvolvimento no `README.md`

---

## Dependências e Ordem de Execução

`T001, T002` $\rightarrow$ `T003, T004` $\rightarrow$ `US1` $\rightarrow$ `US2` $\rightarrow$ `US3` $\rightarrow$ `Polimento`

## Exemplos de Execução Paralela

- **Sessão 1 ( la- la-linkedin)**: T005, T006, T007, T008, T009 (Todos dependem apenas da infra de Golden Files)
- **Sessão 2 (generator)**: T016, T017, T018 (Validações de SEO são independentes da lógica de importação)
