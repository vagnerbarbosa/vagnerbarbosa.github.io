# Plano de Implementação: Cobertura de Testes Total do Projeto

**Branch**: `005-test-coverage-total` | **Data**: 2026-05-16 | **Spec**: [spec.md](./spec.md)
**Input**: Especificação da funcionalidade de `/specs/005-linkedin-import-test-coverage/spec.md`

## Resumo

O objetivo é elevar a cobertura de instruções (statement coverage) de todos os pacotes do projeto para 100%. Isso envolve a refatoração do `main()` do gerador para torná-lo testável, a criação de casos de teste para todos os caminhos de erro no parser e transformer do LinkedIn Import, e a validação completa de todos os comandos CLI.

## Contexto Técnico

**Linguagem/Versão**: Go 1.26
**Dependências Principais**: `testing`, `reflect`, `html/template`, `gopkg.in/yaml.v3`
**Armazenamento**: N/A (Processamento em memória e arquivos locais)
**Testes**: `go test` nativo com `-cover`
**Plataforma Alvo**: Linux / macOS / Windows
**Tipo de Projeto**: CLI Tool / Static Site Generator
**Metas de Performance**: Cobertura de instruções = 100.0% em todos os pacotes.
**Restrições**: Não introduzir dependências externas de teste (preferir `std lib`).
**Escala/Alcance**: Todo o código-fonte em `cmd/` e `internal/`.

## Verificação da Constituição

- [x] **Minimalismo Intencional**: A implementação de testes não adiciona dependências de runtime nem complexidade ao código de produção.
- [x] **Bilíngue por Padrão**: N/A (Código de teste interno).
- [x] **Estabilidade Visual**: N/A.
- [x] **Build Reprodutível**: Os testes rodam via `go test` e são integrados ao CI via GitHub Actions.
- [x] **Código como Documentação**: Aumentar a cobertura de testes serve como documentação viva do comportamento esperado do sistema.

**Veredito**: Totalmente alinhado com a constituição.

## Estrutura do Projeto

### Documentação (esta funcionalidade)

```text
specs/005-linkedin-import-test-coverage/
├── plan.md              # Este arquivo
├── research.md          # Análise de lacunas de cobertura
├── data-model.md        # N/A (não altera modelo de dados)
├── quickstart.md        # Como rodar os testes de cobertura
└── tasks.md             # Lista de tarefas de implementação
```

### Código Fonte (raiz do repositório)

```text
cmd/
├── generator/
│   ├── main.go           # Refatoração para testabilidade
│   └── main_test.go      # Novos casos de teste para main()
└── import-linkedin/
    ├── commands/         # Testes para todos os comandos CLI
    └── internal/
        ├── comparator/   # Testes de caminhos de erro e modificações
        ├── parser/        # Testes de datas e CSVs malformados
        └── transformer/   # Testes de regex e separadores de tech stack
internal/
└── config/
    └── config_test.go    # Validação de 100% de cobertura
```

**Decisão de Estrutura**: Projeto único (PADRÃO). O foco é a adição de casos de teste nos arquivos `*_test.go` existentes e a refatoração mínima de `main.go` para permitir a cobertura do entrypoint.
