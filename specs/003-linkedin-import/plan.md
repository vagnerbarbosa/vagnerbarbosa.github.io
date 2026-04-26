# Plano de Implementação: LinkedIn Import CLI

**Branch**: `003-linkedin-import` | **Data**: 2026-04-26 | **Spec**: [specs/003-linkedin-import/spec.md](./spec.md)  
**Input**: Ferramenta CLI para importar experiências, educação e certificações do LinkedIn a partir do export CSV manual, converter datas para português, dividir descrições em bullets, comparar com config.yaml atual e aplicar mudanças com confirmação interativa

---

## Resumo

Ferramenta CLI em Go que automatiza a importação de dados profissionais do LinkedIn para o config.yaml do portfólio. A ferramenta parseia arquivos CSV exportados manualmente do LinkedIn, converte datas do inglês para português, divide descrições em bullets, compara com o config.yaml existente mostrando um diff visual colorido, e aplica mudanças com confirmação interativa do usuário.

**Abordagem técnica**: Utilizar standard library do Go para parsing de CSV (`encoding/csv`), `go.yaml.in/yaml/v4` para manipulação de YAML preservando comentários, `charmbracelet/huh` para CLI interativa, e implementação customizada de diff visual com `lipgloss` para melhor experiência do usuário.

---

## Contexto Técnico

**Linguagem/Versão**: Go 1.21+ (conforme constituição do projeto)  
**Dependências Principais**:
- `go.yaml.in/yaml/v4` - Manipulação de YAML (sucessor oficial do gopkg.in/yaml.v3)
- `github.com/charmbracelet/huh` v2.0.3 - CLI interativa e prompts de confirmação
- `github.com/charmbracelet/lipgloss` v2 - Estilização de terminal e diff colorido
- `github.com/spf13/cobra` - Framework de CLI (se não usar stdlib)
- `github.com/sergi/go-diff` - Algoritmo de diff

**Armazenamento**: Arquivos locais (CSV de entrada, YAML de configuração)  
**Testes**: Go testing (`go test`) com meta >70% cobertura (conforme constituição)  
**Plataforma Alvo**: Linux/macOS/Windows (ambiente local de desenvolvimento)  
**Tipo de Projeto**: CLI tool (ferramenta de linha de comando)  
**Metas de Performance**:
- Parsing de CSV: <1s para arquivos com at<1000 registros
- Geração de diff: <500ms
- Tempo total de importação: <1 minuto (excluindo confirmação interativa)

**Restrições**:
- Zero dependências de runtime (só build-time) - Princípio I: Minimalismo Intencional
- Build reproducível via go.mod travado
- CLI executada apenas localmente (não em CI/CD)
- Preservar comentários e formatação do config.yaml original

**Escala/Alcance**:
- Suporta até 1000 registros por arquivo CSV
- Um único usuário (desenvolvedor do portfólio)
- Arquivos CSV no formato UTF-8 padrão do LinkedIn

---

## Verificação da Constituição

*PORTÃO: Deve passar antes da pesquisa da Fase 0. Re-verificar após o design da Fase 1.*

| Princípio | Status | Justificativa |
|-----------|--------|---------------|
| **I. Minimalismo Intencional** | ✅ PASS | Ferramenta CLI em Go gera output estático, sem runtime dependencies. Não usa frameworks web pesados. |
| **II. Bilíngue por Padrão** | ✅ PASS | A ferramenta converte conteúdo importado para pt-BR (datas). O config.yaml resultante mantém estrutura bilíngue existente. |
| **III. Estabilidade Visual** | ✅ PASS | CLI usa temas padrão do `huh` que respeitam preferências do sistema. Output é em texto plano/formatado, não afeta design do site. |
| **IV. Build Reprodutível** | ✅ PASS | Depende de go.mod travado. CLI é compilada, não interpretada. Resultado (config.yaml modificado) é deterministico dado mesmo input. |
| **V. Código como Documentação** | ✅ PASS | Código usa nomenclatura clara (Import, Parse, Transform, Compare, Confirm). Comentários explicam "porquê". |

**Verificação pós-design**: Todos os princípios mantidos. A CLI é uma ferramenta auxiliar que gera conteúdo, não modifica o design ou estrutura do site.

---

## Estrutura do Projeto

### Documentação (esta funcionalidade)

```text
specs/003-linkedin-import/
├── plan.md              # Este arquivo (output do comando /speckit.plan)
├── research.md          # Output da Fase 0 (comando /speckit.plan)
├── data-model.md        # Output da Fase 1 (comando /speckit.plan)
├── quickstart.md        # Output da Fase 1 (comando /speckit.plan)
├── contracts/           # Output da Fase 1 (comando /speckit.plan)
│   └── cli-commands.md  # Contratos de comandos da CLI
└── tasks.md             # Output da Fase 2 (comando /speckit.tasks - NÃO criado pelo /speckit.plan)
```

### Código Fonte (raiz do repositório)

```text
cmd/linkedin-import/
├── main.go                    # Entry point
├── commands/
│   ├── root.go               # Comando root e setup
│   ├── import.go             # Comando: import
│   ├── validate.go           # Comando: validate
│   └── version.go            # Comando: version
├── internal/
│   ├── parser/
│   │   ├── csv.go            # Parsing de CSV (encoding/csv)
│   │   └── date.go           # Conversão de datas (en → pt)
│   ├── transformer/
│   │   └── description.go    # Split de descrições em bullets
│   ├── comparator/
│   │   ├── diff.go           # Lógica de diff entre LinkedIn e config
│   │   └── change.go         # Representação de mudanças
│   ├── config/
│   │   ├── yaml.go           # Leitura/escrita de YAML
│   │   └── backup.go         # Criação de backup
│   ├── models/
│   │   ├── experience.go     # Struct Experience
│   │   ├── education.go      # Struct Education
│   │   ├── certification.go  # Struct Certification
│   │   └── config.go         # Struct ConfigPortfolio
│   └── ui/
│       ├── confirm.go        # Prompts interativos (huh)
│       ├── diff.go           # Visualização de diff (lipgloss)
│       └── spinner.go        # Indicadores de progresso

cmd/linkedin-import/
└── testdata/                 # Arquivos de teste (CSVs de exemplo)
    ├── Experiences.csv
    ├── Education.csv
    └── Certifications.csv
```

**Decisão de Estrutura**: Segue estrutura padrão Go CLI com `cmd/` para entry points, `internal/` para pacotes privados, e separação por responsabilidade (parser, transformer, comparator, config, ui).

---

## Rastreamento de Complexidade

> **Preencha SOMENTE se a Verificação da Constituição tiver violações que devam ser justificadas**

Nenhuma violação identificada. Todos os princípios da constituição são respeitados.

---

## Fase 0: Pesquisa - Resumo

Pesquisa concluída e documentada em [research.md](./research.md).

### Decisões Principais

| Aspecto | Decisão | Biblioteca |
|---------|---------|--------------|
| CSV Parsing | Standard library | `encoding/csv` |
| YAML Manipulation | Fork oficial mantido | `go.yaml.in/yaml/v4` |
| CLI Interativa | Biblioteca oficial Charmbracelet | `charmbracelet/huh` v2 |
| Diff Visual | Algoritmo padrão + customização | `sergi/go-diff` + `lipgloss` |
| Conversão de Datas | Implementação customizada | Mapa de meses (en → pt) |
| Split de Descrições | Heurística simples | Quebras de linha + frases |

### Lições Aprendidas

1. `gopkg.in/yaml.v3` está unmaintained desde Abril 2025 - migrar para `go.yaml.in/yaml`
2. `charmbracelet/huh` oferece a melhor UX para prompts de confirmação em Go
3. Standard library `encoding/csv` é suficiente para parsing de CSVs do LinkedIn

---

## Fase 1: Design - Resumo

### Entidades

- **Experience**: company, role, start_date, end_date, description[]
- **Education**: institution, degree, field, start_date, end_date, description[]
- **Certification**: name, organization, issue_date, expiration_date
- **Change**: entity_type, entity_id, old_value, new_value, change_type

### Contratos

Documentados em [contracts/cli-commands.md](./contracts/cli-commands.md):
- `import`: Importa dados do LinkedIn
- `validate`: Valida arquivos CSV
- `version`: Exibe versão
- `help`: Ajuda contextual

### Quickstart

Documentado em [quickstart.md](./quickstart.md) - guia completo para desenvolvedores.

---

## Artefatos Gerados

| Artefato | Status | Caminho |
|----------|--------|---------|
| Plano de Implementação | ✅ Criado | [specs/003-linkedin-import/plan.md](./plan.md) |
| Pesquisa Técnica | ✅ Criado | [specs/003-linkedin-import/research.md](./research.md) |
| Modelo de Dados | ✅ Criado | [specs/003-linkedin-import/data-model.md](./data-model.md) |
| Contratos CLI | ✅ Criado | [specs/003-linkedin-import/contracts/cli-commands.md](./contracts/cli-commands.md) |
| Guia Quickstart | ✅ Criado | [specs/003-linkedin-import/quickstart.md](./quickstart.md) |

---

## Próximos Passos

1. Gerar tasks.md via `/speckit-tasks`
2. Implementar estrutura de diretórios em `cmd/linkedin-import/`
3. Desenvolver módulos em ordem de dependência: models → parser → transformer → comparator → config → ui → commands
4. Executar testes com meta >70% cobertura
5. Validar CLI com dados reais do LinkedIn
