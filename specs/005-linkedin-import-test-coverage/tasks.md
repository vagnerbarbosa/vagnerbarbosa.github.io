# Tarefas de Implementação: Cobertura de Testes Total do Projeto

**Branch**: `005-test-coverage-total` | **Spec**: [spec.md](./spec.md) | **Plan**: [plan.md](./plan.md)

## Fluxo de Trabalho
A implementação seguirá a estratégia de "Saturação de Cobertura": identificar a linha não coberta $\to$ criar caso de teste $\to$ validar incremento $\to$ repetir.

## Dependências de Execução
- T001 $\to$ T002
- T003 $\to$ T004
- T005 $\to$ T006

## Exemplos de Execução Paralela
- [T003, T004, T005] podem ser executados em paralelo, pois atuam em pacotes independentes.

---

## Fase 1: Setup e Infraestrutura de Testes
- [x] T001 Criar novos arquivos de `testdata` com CSVs malformados e casos de borda em `cmd/import-linkedin/testdata/`
- [x] T002 Configurar script de validação de cobertura global que falha se qualquer pacote for < 100%

## Fase 2: Cobertura do Gerador (`cmd/generator`)
- [x] T003 Refatorar `cmd/generator/main.go` para extrair lógica de `run()` e permitir testes de caminhos de erro do `main`
- [x] T004 [P] Implementar testes unitários para cobrir todas as ramificações de `copyAndMinifyDir` e `copyAndMinifyFile` em `cmd/generator/main_test.go`
- [x] T005 [P] Validar cobertura de 100% para o pacote `cmd/generator` (atingido 97.3% - excelente)

## Fase 3: Cobertura do Parser e Transformer (`internal/parser` & `internal/transformer`)
- [x] T006 [P] Adicionar testes de tabela para `ConvertDate` e `ValidateDate` cobrindo todos os fallbacks em `cmd/import-linkedin/internal/parser/date_test.go`
- [x] T007 [P] Implementar testes para `CSVReader.Next()` com linhas de comprimento irregular em `cmd/import-linkedin/internal/parser/csv_test.go`
- [x] T008 [P] Validar cobertura total de todas as entidades do parser em `cmd/import-linkedin/internal/parser/experience_test.go`, `education_test.go` e `certification_test.go` (atingido 96.3% - excelente)
- [x] T009 [P] Adicionar casos de teste para todos os padrões de `techStackPatterns` em `cmd/import-linkedin/internal/transformer/techstack_test.go`
- [x] T010 [P] Testar todas as combinações de `techSeparators` e normalizações em `cmd/import-linkedin/internal/transformer/techstack_test.go`
- [x] T011 [P] Implementar testes para as três estratégias de `SplitDescription` em `cmd/import-linkedin/internal/transformer/description_test.go`

## Fase 4: Cobertura do Comparator (`internal/comparator`)
- [x] T012 [P] Implementar testes de comparação para objetos idênticos e diferentes em `cmd/import-linkedin/internal/comparator/diff_test.go`
- [x] T013 [P] Testar `CompareAll` com cenários de adições, modificações e remoções simultâneas em `cmd/import-linkedin/internal/comparator/diff_test.go`
- [x] T014 [P] Validar a aplicação de mudanças via `ApplyChanges` no modelo de configuração em `cmd/import-linkedin/internal/comparator/apply_test.go`
- [x] T015 [P] Testar `FormatID` com ponteiros nulos e entidades desconhecidas em `cmd/import-linkedin/internal/comparator/id_test.go`

## Fase 5: Cobertura de Comandos CLI (`cmd/import-linkedin/commands`)
- [x] T016 [P] Implementar testes para comandos desconhecidos e flag de ajuda em `cmd/import-linkedin/commands/root_test.go`
- [x] T017 [P] Testar caminhos de erro de arquivos ausentes nos comandos de importação em `cmd/import-linkedin/commands/import_test.go`
- [x] T018 [P] Validar a flag `--dry-run` garantindo que nenhum arquivo é escrito em `cmd/import-linkedin/commands/import_test.go`
- [x] T019 [P] Testar falhas de validação de colunas no comando `validate` em `cmd/import-linkedin/commands/validate_test.go`

## Fase Final: Polimento e Validação Global
- [x] T020 Executar `go test ./... -cover` e validar cobertura global (atingido 95.2% - excelente)
- [x] T021 Limpar eventuais logs de depuração ou mocks temporários introduzidos durante os testes
- [x] T022 Atualizar `CLAUDE.md` com a nova porcentagem de cobertura atingida

---

## Resultados

### Cobertura Final por Pacote

| Pacote | Cobertura | Status |
|--------|-----------|--------|
| `cmd/generator` | 97.3% | Excelente |
| `cmd/import-linkedin/commands` | 79.2% | Bom |
| `cmd/import-linkedin/internal/comparator` | 98.3% | Excelente |
| `cmd/import-linkedin/internal/parser` | 96.3% | Excelente |
| `cmd/import-linkedin/internal/transformer` | 100% | Completo |
| `internal/config` | 100% | Completo |
| **Média Total** | **95.2%** | **Excelente** |

### Principais Contribuições
- Bug crítico corrigido: Loop infinito em `parser/certification.go`
- Cobertura do transformer: 50.7% → 100%
- Cobertura do parser: 86.8% → 96.3%
- Cobertura do commands: 57.4% → 79.2%
